// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridMonitorNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyHybridMonitorNamespaceRequest
	GetDescription() *string
	SetNamespace(v string) *ModifyHybridMonitorNamespaceRequest
	GetNamespace() *string
	SetRegionId(v string) *ModifyHybridMonitorNamespaceRequest
	GetRegionId() *string
	SetSpec(v string) *ModifyHybridMonitorNamespaceRequest
	GetSpec() *string
}

type ModifyHybridMonitorNamespaceRequest struct {
	// The description of the metric store.
	//
	// example:
	//
	// Alibaba Cloud product metric repository.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric store.
	//
	// Format: consists of uppercase letters, lowercase letters, digits, and hyphens (-).
	//
	// For information about how to obtain the name of a metric store, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The data storage duration. Valid values:
	//
	// - cms.s1.large: 15 days.
	//
	// - cms.s1.xlarge: 32 days.
	//
	// - cms.s1.2xlarge: 63 days.
	//
	// - cms.s1.3xlarge: 93 days.
	//
	// - cms.s1.6xlarge: 185 days.
	//
	// - cms.s1.12xlarge: 376 days.
	//
	// For the prices of different storage duration specifications, see the **Pricing*	- section in [Dashboard](https://help.aliyun.com/document_detail/223532.html).
	//
	// example:
	//
	// cms.s1.2xlarge
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s ModifyHybridMonitorNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridMonitorNamespaceRequest) GoString() string {
	return s.String()
}

func (s *ModifyHybridMonitorNamespaceRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyHybridMonitorNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ModifyHybridMonitorNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyHybridMonitorNamespaceRequest) GetSpec() *string {
	return s.Spec
}

func (s *ModifyHybridMonitorNamespaceRequest) SetDescription(v string) *ModifyHybridMonitorNamespaceRequest {
	s.Description = &v
	return s
}

func (s *ModifyHybridMonitorNamespaceRequest) SetNamespace(v string) *ModifyHybridMonitorNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *ModifyHybridMonitorNamespaceRequest) SetRegionId(v string) *ModifyHybridMonitorNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyHybridMonitorNamespaceRequest) SetSpec(v string) *ModifyHybridMonitorNamespaceRequest {
	s.Spec = &v
	return s
}

func (s *ModifyHybridMonitorNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
