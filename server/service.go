package main

import (
	"fmt"
	"sync"
)

var (
	Services sync.Map
)

type List struct {
	ServiceProviders sync.Map
	flushInterval    int64
}

// WeedOut wipe out those dead service
func (l *List) WeedOut() {
	flushFunc := func(k, v interface{}) bool {
		if !v.(bool) {
			l.ServiceProviders.Delete(k)
		}
		return true
	}

	l.ServiceProviders.Range(flushFunc)
}

func (l *List) Activate(provider string) error {
	_, isExist := l.ServiceProviders.Load(provider)
	if !isExist {
		return fmt.Errorf("service provider: %s not exist", provider)
	}
	l.ServiceProviders.Store(provider, true)
	return nil

}

func (l *List) insert(provider string) error {
	_, isExist := l.ServiceProviders.Load(provider)
	if isExist {
		return fmt.Errorf("service provider: %s already exist", provider)
	}
	l.ServiceProviders.Store(provider, true)
	return nil
}

func RegisterService(serviceName, serviceProvider string, flushInterval int64) error {

	_, isExist := Services.Load(serviceName)
	if isExist {
		return fmt.Errorf("service name : %s already exist", serviceName)
	}

	newList := new(List)
	newList.insert(serviceProvider)
	newList.flushInterval = flushInterval

	Services.Store(serviceProvider, serviceName)

	return nil
}

func StartService(serviceName, serviceProvider string) {

}
