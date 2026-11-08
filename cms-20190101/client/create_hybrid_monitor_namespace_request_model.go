// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridMonitorNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateHybridMonitorNamespaceRequest
	GetDescription() *string
	SetNamespace(v string) *CreateHybridMonitorNamespaceRequest
	GetNamespace() *string
	SetNamespaceRegion(v string) *CreateHybridMonitorNamespaceRequest
	GetNamespaceRegion() *string
	SetNamespaceType(v string) *CreateHybridMonitorNamespaceRequest
	GetNamespaceType() *string
	SetRegionId(v string) *CreateHybridMonitorNamespaceRequest
	GetRegionId() *string
	SetSpec(v string) *CreateHybridMonitorNamespaceRequest
	GetSpec() *string
}

type CreateHybridMonitorNamespaceRequest struct {
	// The description of the metric namespace.
	//
	// example:
	//
	// Alibaba Cloud product metric repository.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the metric namespace.
	//
	// Format: The name can contain only lowercase letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region where the monitoring data is stored.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceRegion *string `json:"NamespaceRegion,omitempty" xml:"NamespaceRegion,omitempty"`
	// The storage solution for monitoring data. Valid values:
	//
	// - aliyun_prometheus: Stores monitoring data in Prometheus.
	//
	// > For more information about storage solutions for monitoring data, see [Storage solutions for monitoring data in Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/2594921.html).
	//
	// example:
	//
	// aliyun_prometheus
	NamespaceType *string `json:"NamespaceType,omitempty" xml:"NamespaceType,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The data storage duration. Valid values:
	//
	// - cms.s1.large: Stores data for 15 days.
	//
	// - cms.s1.xlarge: Stores data for 32 days.
	//
	// - cms.s1.2xlarge: Stores data for 63 days.
	//
	// - cms.s1.3xlarge (default): Stores data for 93 days.
	//
	// - cms.s1.6xlarge: Stores data for 185 days.
	//
	// - cms.s1.12xlarge: Stores data for 376 days.
	//
	// For more information about the pricing of different storage durations, see the **Pricing*	- table in [Dashboard](https://help.aliyun.com/document_detail/223532.html).
	//
	// example:
	//
	// cms.s1.3xlarge
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s CreateHybridMonitorNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridMonitorNamespaceRequest) GoString() string {
	return s.String()
}

func (s *CreateHybridMonitorNamespaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateHybridMonitorNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *CreateHybridMonitorNamespaceRequest) GetNamespaceRegion() *string {
	return s.NamespaceRegion
}

func (s *CreateHybridMonitorNamespaceRequest) GetNamespaceType() *string {
	return s.NamespaceType
}

func (s *CreateHybridMonitorNamespaceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateHybridMonitorNamespaceRequest) GetSpec() *string {
	return s.Spec
}

func (s *CreateHybridMonitorNamespaceRequest) SetDescription(v string) *CreateHybridMonitorNamespaceRequest {
	s.Description = &v
	return s
}

func (s *CreateHybridMonitorNamespaceRequest) SetNamespace(v string) *CreateHybridMonitorNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *CreateHybridMonitorNamespaceRequest) SetNamespaceRegion(v string) *CreateHybridMonitorNamespaceRequest {
	s.NamespaceRegion = &v
	return s
}

func (s *CreateHybridMonitorNamespaceRequest) SetNamespaceType(v string) *CreateHybridMonitorNamespaceRequest {
	s.NamespaceType = &v
	return s
}

func (s *CreateHybridMonitorNamespaceRequest) SetRegionId(v string) *CreateHybridMonitorNamespaceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateHybridMonitorNamespaceRequest) SetSpec(v string) *CreateHybridMonitorNamespaceRequest {
	s.Spec = &v
	return s
}

func (s *CreateHybridMonitorNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
