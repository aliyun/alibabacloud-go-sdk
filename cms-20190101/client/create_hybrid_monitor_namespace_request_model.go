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
	// The description of the namespace.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the namespace.
	//
	// The name can contain lowercase letters, digits, and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// aliyun
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The region where the metric data is stored.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceRegion *string `json:"NamespaceRegion,omitempty" xml:"NamespaceRegion,omitempty"`
	// The storage scheme of metric data. Valid values:
	//
	// 	- m_prom_user: The metric data is stored in Simple Log Service.
	//
	// 	- m_prom_pool: The metric data is stored in the private storage space provided by CloudMonitor.
	//
	// >  For more information about the storage schemes of metric data, see [Data storage schemes for Hybrid Cloud Monitoring](https://help.aliyun.com/document_detail/2594921.html).
	//
	// example:
	//
	// m_prometheus
	NamespaceType *string `json:"NamespaceType,omitempty" xml:"NamespaceType,omitempty"`
	RegionId      *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The data retention period. Valid values:
	//
	// 	- cms.s1.large (Retention Period 15 Days)
	//
	// 	- cms.s1.xlarge (Retention Period 32 Days)
	//
	// 	- cms.s1.2xlarge (Retention Period 63 Days)
	//
	// 	- cms.s1.3xlarge (Retention Period 93 Days) (default)
	//
	// 	- cms.s1.6xlarge (Retention Period 185 Days)
	//
	// 	- cms.s1.12xlarge (Retention Period 367 Days)
	//
	// For information about the pricing for different retention periods, see the **Pricing*	- section in [Billing of the dashboard feature](https://help.aliyun.com/document_detail/223532.html).
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
