/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was automatically generated by lister-gen

package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterServicePlanLister helps list ClusterServicePlans.
type ClusterServicePlanLister interface {
	// List lists all ClusterServicePlans in the indexer.
	List(selector labels.Selector) (ret []*servicecatalog.ClusterServicePlan, err error)
	// Get retrieves the ClusterServicePlan from the index for a given name.
	Get(name string) (*servicecatalog.ClusterServicePlan, error)
	ClusterServicePlanListerExpansion
}

// clusterServicePlanLister implements the ClusterServicePlanLister interface.
type clusterServicePlanLister struct {
	indexer cache.Indexer
}

// NewClusterServicePlanLister returns a new ClusterServicePlanLister.
func NewClusterServicePlanLister(indexer cache.Indexer) ClusterServicePlanLister {
	return &clusterServicePlanLister{indexer: indexer}
}

// List lists all ClusterServicePlans in the indexer.
func (s *clusterServicePlanLister) List(selector labels.Selector) (ret []*servicecatalog.ClusterServicePlan, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ClusterServicePlan))
	})
	return ret, err
}

// Get retrieves the ClusterServicePlan from the index for a given name.
func (s *clusterServicePlanLister) Get(name string) (*servicecatalog.ClusterServicePlan, error) {
	key := &servicecatalog.ClusterServicePlan{ObjectMeta: v1.ObjectMeta{Name: name}}
	obj, exists, err := s.indexer.Get(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("clusterserviceplan"), name)
	}
	return obj.(*servicecatalog.ClusterServicePlan), nil
}
