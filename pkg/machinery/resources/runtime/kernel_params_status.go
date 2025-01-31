// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package runtime

import (
	"github.com/cosi-project/runtime/pkg/resource"
	"github.com/cosi-project/runtime/pkg/resource/meta"
	"github.com/cosi-project/runtime/pkg/resource/typed"
)

// KernelParamStatusType is type of KernelParam resource.
const KernelParamStatusType = resource.Type("KernelParamStatuses.runtime.talos.dev")

// KernelParamStatus resource holds defined sysctl flags status.
type KernelParamStatus = typed.Resource[KernelParamStatusSpec, KernelParamStatusRD]

// KernelParamStatusSpec describes status of the defined sysctls.
type KernelParamStatusSpec struct {
	Current     string `yaml:"current"`
	Default     string `yaml:"default"`
	Unsupported bool   `yaml:"unsupported"`
}

// NewKernelParamStatus initializes a KernelParamStatus resource.
func NewKernelParamStatus(namespace resource.Namespace, id resource.ID) *KernelParamStatus {
	return typed.NewResource[KernelParamStatusSpec, KernelParamStatusRD](
		resource.NewMetadata(namespace, KernelParamStatusType, id, resource.VersionUndefined),
		KernelParamStatusSpec{},
	)
}

// DeepCopy implements typed.DeepCopyable interface.
func (spec KernelParamStatusSpec) DeepCopy() KernelParamStatusSpec {
	return spec
}

// KernelParamStatusRD is auxiliary resource data for KernelParamStatus.
type KernelParamStatusRD struct{}

// ResourceDefinition implements meta.ResourceDefinitionProvider interface.
func (KernelParamStatusRD) ResourceDefinition(resource.Metadata, KernelParamStatusSpec) meta.ResourceDefinitionSpec {
	return meta.ResourceDefinitionSpec{
		Type:             KernelParamStatusType,
		Aliases:          []resource.Type{"sysctls", "kernelparameters", "kernelparams"},
		DefaultNamespace: NamespaceName,
		PrintColumns: []meta.PrintColumn{
			{
				Name:     "Current",
				JSONPath: `{.current}`,
			},
			{
				Name:     "Default",
				JSONPath: `{.default}`,
			},
			{
				Name:     "Unsupported",
				JSONPath: `{.unsupported}`,
			},
		},
	}
}
