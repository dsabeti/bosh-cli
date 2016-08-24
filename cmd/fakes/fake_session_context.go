// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-init/cmd"
	cmdconf "github.com/cloudfoundry/bosh-init/cmd/config"
)

type FakeSessionContext struct {
	EnvironmentStub        func() string
	environmentMutex       sync.RWMutex
	environmentArgsForCall []struct{}
	environmentReturns     struct {
		result1 string
	}
	CACertStub        func() string
	cACertMutex       sync.RWMutex
	cACertArgsForCall []struct{}
	cACertReturns     struct {
		result1 string
	}
	CredentialsStub        func() cmdconf.Creds
	credentialsMutex       sync.RWMutex
	credentialsArgsForCall []struct{}
	credentialsReturns     struct {
		result1 cmdconf.Creds
	}
	DeploymentStub        func() string
	deploymentMutex       sync.RWMutex
	deploymentArgsForCall []struct{}
	deploymentReturns     struct {
		result1 string
	}
}

func (fake *FakeSessionContext) Environment() string {
	fake.environmentMutex.Lock()
	fake.environmentArgsForCall = append(fake.environmentArgsForCall, struct{}{})
	fake.environmentMutex.Unlock()
	if fake.EnvironmentStub != nil {
		return fake.EnvironmentStub()
	} else {
		return fake.environmentReturns.result1
	}
}

func (fake *FakeSessionContext) EnvironmentCallCount() int {
	fake.environmentMutex.RLock()
	defer fake.environmentMutex.RUnlock()
	return len(fake.environmentArgsForCall)
}

func (fake *FakeSessionContext) EnvironmentReturns(result1 string) {
	fake.EnvironmentStub = nil
	fake.environmentReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSessionContext) CACert() string {
	fake.cACertMutex.Lock()
	fake.cACertArgsForCall = append(fake.cACertArgsForCall, struct{}{})
	fake.cACertMutex.Unlock()
	if fake.CACertStub != nil {
		return fake.CACertStub()
	} else {
		return fake.cACertReturns.result1
	}
}

func (fake *FakeSessionContext) CACertCallCount() int {
	fake.cACertMutex.RLock()
	defer fake.cACertMutex.RUnlock()
	return len(fake.cACertArgsForCall)
}

func (fake *FakeSessionContext) CACertReturns(result1 string) {
	fake.CACertStub = nil
	fake.cACertReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeSessionContext) Credentials() cmdconf.Creds {
	fake.credentialsMutex.Lock()
	fake.credentialsArgsForCall = append(fake.credentialsArgsForCall, struct{}{})
	fake.credentialsMutex.Unlock()
	if fake.CredentialsStub != nil {
		return fake.CredentialsStub()
	} else {
		return fake.credentialsReturns.result1
	}
}

func (fake *FakeSessionContext) CredentialsCallCount() int {
	fake.credentialsMutex.RLock()
	defer fake.credentialsMutex.RUnlock()
	return len(fake.credentialsArgsForCall)
}

func (fake *FakeSessionContext) CredentialsReturns(result1 cmdconf.Creds) {
	fake.CredentialsStub = nil
	fake.credentialsReturns = struct {
		result1 cmdconf.Creds
	}{result1}
}

func (fake *FakeSessionContext) Deployment() string {
	fake.deploymentMutex.Lock()
	fake.deploymentArgsForCall = append(fake.deploymentArgsForCall, struct{}{})
	fake.deploymentMutex.Unlock()
	if fake.DeploymentStub != nil {
		return fake.DeploymentStub()
	} else {
		return fake.deploymentReturns.result1
	}
}

func (fake *FakeSessionContext) DeploymentCallCount() int {
	fake.deploymentMutex.RLock()
	defer fake.deploymentMutex.RUnlock()
	return len(fake.deploymentArgsForCall)
}

func (fake *FakeSessionContext) DeploymentReturns(result1 string) {
	fake.DeploymentStub = nil
	fake.deploymentReturns = struct {
		result1 string
	}{result1}
}

var _ cmd.SessionContext = new(FakeSessionContext)