/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/apache/submarine/submarine-cloud/pkg/apis/submarine/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SubmarineServerLister helps list SubmarineServers.
type SubmarineServerLister interface {
	// List lists all SubmarineServers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.SubmarineServer, err error)
	// SubmarineServers returns an object that can list and get SubmarineServers.
	SubmarineServers(namespace string) SubmarineServerNamespaceLister
	SubmarineServerListerExpansion
}

// submarineServerLister implements the SubmarineServerLister interface.
type submarineServerLister struct {
	indexer cache.Indexer
}

// NewSubmarineServerLister returns a new SubmarineServerLister.
func NewSubmarineServerLister(indexer cache.Indexer) SubmarineServerLister {
	return &submarineServerLister{indexer: indexer}
}

// List lists all SubmarineServers in the indexer.
func (s *submarineServerLister) List(selector labels.Selector) (ret []*v1alpha1.SubmarineServer, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubmarineServer))
	})
	return ret, err
}

// SubmarineServers returns an object that can list and get SubmarineServers.
func (s *submarineServerLister) SubmarineServers(namespace string) SubmarineServerNamespaceLister {
	return submarineServerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubmarineServerNamespaceLister helps list and get SubmarineServers.
type SubmarineServerNamespaceLister interface {
	// List lists all SubmarineServers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.SubmarineServer, err error)
	// Get retrieves the SubmarineServer from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.SubmarineServer, error)
	SubmarineServerNamespaceListerExpansion
}

// submarineServerNamespaceLister implements the SubmarineServerNamespaceLister
// interface.
type submarineServerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SubmarineServers in the indexer for a given namespace.
func (s submarineServerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SubmarineServer, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SubmarineServer))
	})
	return ret, err
}

// Get retrieves the SubmarineServer from the indexer for a given namespace and name.
func (s submarineServerNamespaceLister) Get(name string) (*v1alpha1.SubmarineServer, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("submarineserver"), name)
	}
	return obj.(*v1alpha1.SubmarineServer), nil
}
