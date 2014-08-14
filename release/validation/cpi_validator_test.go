package validation_test

import (
	fakesys "github.com/cloudfoundry/bosh-agent/system/fakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	bmrelease "github.com/cloudfoundry/bosh-micro-cli/release"
	bmreljob "github.com/cloudfoundry/bosh-micro-cli/release/jobs"

	. "github.com/cloudfoundry/bosh-micro-cli/release/validation"
)

var _ = Describe("CpiValidator", func() {
	var fakeFs *fakesys.FakeFileSystem

	BeforeEach(func() {
		fakeFs = fakesys.NewFakeFileSystem()
	})

	It("validates a valid release without error", func() {
		release := bmrelease.Release{
			Jobs: []bmreljob.Job{
				{
					Name:        "cpi",
					Version:     "fake-job-1-version",
					Fingerprint: "fake-job-1-fingerprint",
					Sha1:        "fake-job-1-sha",
					Templates: map[string]string{
						"cpi.erb":               "bin/cpi",
						"micro_discover_ip.erb": "bin/micro_discover_ip",
					},
				},
			},
		}
		validator := NewCpiValidator()

		err := validator.Validate(release)
		Expect(err).NotTo(HaveOccurred())
	})

	Context("when the cpi job is not present", func() {
		var validator CpiValidator
		var release bmrelease.Release

		BeforeEach(func() {
			release = bmrelease.Release{
				Jobs: []bmreljob.Job{
					{
						Name:        "non-cpi-job",
						Version:     "fake-job-1-version",
						Fingerprint: "fake-job-1-fingerprint",
						Sha1:        "fake-job-1-sha",
						Templates: map[string]string{
							"cpi.erb": "bin/cpi",
						},
					},
				},
			}
			validator = NewCpiValidator()
		})

		It("returns an error that the cpi job is not present", func() {
			err := validator.Validate(release)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Job `cpi' is missing from release"))
		})
	})

	Context("when the templates are missing a bin/cpi target", func() {
		var validator CpiValidator
		var release bmrelease.Release

		BeforeEach(func() {
			release = bmrelease.Release{
				Jobs: []bmreljob.Job{
					{
						Name:        "cpi",
						Version:     "fake-job-1-version",
						Fingerprint: "fake-job-1-fingerprint",
						Sha1:        "fake-job-1-sha",
						Templates: map[string]string{
							"cpi.erb":               "nonsense",
							"micro_discover_ip.erb": "bin/micro_discover_ip",
						},
					},
				},
			}
			validator = NewCpiValidator()
		})

		It("returns an error that the bin/cpi template target is missing", func() {
			err := validator.Validate(release)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Job `cpi' is missing bin/cpi target"))
		})
	})

	Context("when the templates are missing a bin/micro_discover_ip target", func() {
		var validator CpiValidator
		var release bmrelease.Release

		BeforeEach(func() {
			release = bmrelease.Release{
				Jobs: []bmreljob.Job{
					{
						Name:        "cpi",
						Version:     "fake-job-1-version",
						Fingerprint: "fake-job-1-fingerprint",
						Sha1:        "fake-job-1-sha",
						Templates: map[string]string{
							"cpi.erb": "nonsense",
						},
					},
				},
			}
			validator = NewCpiValidator()
		})

		It("returns an error that the bin/micro_discover_ip template target is missing", func() {
			err := validator.Validate(release)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Job `cpi' is missing bin/micro_discover_ip target"))
		})
	})
})