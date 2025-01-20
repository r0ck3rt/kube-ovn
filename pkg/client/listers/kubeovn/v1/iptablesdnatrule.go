/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kubeovn/kube-ovn/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// IptablesDnatRuleLister helps list IptablesDnatRules.
// All objects returned here must be treated as read-only.
type IptablesDnatRuleLister interface {
	// List lists all IptablesDnatRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.IptablesDnatRule, err error)
	// Get retrieves the IptablesDnatRule from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.IptablesDnatRule, error)
	IptablesDnatRuleListerExpansion
}

// iptablesDnatRuleLister implements the IptablesDnatRuleLister interface.
type iptablesDnatRuleLister struct {
	listers.ResourceIndexer[*kubeovnv1.IptablesDnatRule]
}

// NewIptablesDnatRuleLister returns a new IptablesDnatRuleLister.
func NewIptablesDnatRuleLister(indexer cache.Indexer) IptablesDnatRuleLister {
	return &iptablesDnatRuleLister{listers.New[*kubeovnv1.IptablesDnatRule](indexer, kubeovnv1.Resource("iptablesdnatrule"))}
}
