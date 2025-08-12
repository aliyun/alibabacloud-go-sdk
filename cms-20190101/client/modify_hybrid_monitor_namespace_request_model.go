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
	// The description of the namespace.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the namespace.
	//
	// The name can contain letters, digits, and hyphens (-).
	//
	// For information about how to obtain the name of a namespace, see [DescribeHybridMonitorNamespaceList](https://help.aliyun.com/document_detail/428880.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	RegionId  *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The data retention period. Valid values:
	//
	// 	- cms.s1.large: Data is stored for 15 days.
	//
	// 	- cms.s1.xlarge: Data is stored for 32 days.
	//
	// 	- cms.s1.2xlarge: Data is stored for 63 days.
	//
	// 	- cms.s1.3xlarge: Data is stored for 93 days.
	//
	// 	- cms.s1.6xlarge: Data is stored for 185 days.
	//
	// 	- cms.s1.12xlarge: Data is stored for 376 days.
	//
	// For information about the pricing for different retention periods, see the **Pricing*	- section in [Billing of the dashboard feature](https://help.aliyun.com/document_detail/223532.html).
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
