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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	submarinev1alpha1 "github.com/apache/submarine/submarine-cloud/pkg/apis/submarine/v1alpha1"
	versioned "github.com/apache/submarine/submarine-cloud/pkg/client/clientset/versioned"
	internalinterfaces "github.com/apache/submarine/submarine-cloud/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/apache/submarine/submarine-cloud/pkg/client/listers/submarine/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SubmarineServerInformer provides access to a shared informer and lister for
// SubmarineServers.
type SubmarineServerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SubmarineServerLister
}

type submarineServerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSubmarineServerInformer constructs a new informer for SubmarineServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSubmarineServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSubmarineServerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSubmarineServerInformer constructs a new informer for SubmarineServer type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSubmarineServerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SubmarineV1alpha1().SubmarineServers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SubmarineV1alpha1().SubmarineServers(namespace).Watch(options)
			},
		},
		&submarinev1alpha1.SubmarineServer{},
		resyncPeriod,
		indexers,
	)
}

func (f *submarineServerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSubmarineServerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *submarineServerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&submarinev1alpha1.SubmarineServer{}, f.defaultInformer)
}

func (f *submarineServerInformer) Lister() v1alpha1.SubmarineServerLister {
	return v1alpha1.NewSubmarineServerLister(f.Informer().GetIndexer())
}
