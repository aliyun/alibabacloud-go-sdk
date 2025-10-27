// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupedFields(v *GetCloudAssetSummaryResponseBodyGroupedFields) *GetCloudAssetSummaryResponseBody
	GetGroupedFields() *GetCloudAssetSummaryResponseBodyGroupedFields
	SetRequestId(v string) *GetCloudAssetSummaryResponseBody
	GetRequestId() *string
}

type GetCloudAssetSummaryResponseBody struct {
	// The summary of cloud services.
	GroupedFields *GetCloudAssetSummaryResponseBodyGroupedFields `json:"GroupedFields,omitempty" xml:"GroupedFields,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F5CF78A7-30AA-59DB-847F-13EE3AE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCloudAssetSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetCloudAssetSummaryResponseBody) GetGroupedFields() *GetCloudAssetSummaryResponseBodyGroupedFields {
	return s.GroupedFields
}

func (s *GetCloudAssetSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCloudAssetSummaryResponseBody) SetGroupedFields(v *GetCloudAssetSummaryResponseBodyGroupedFields) *GetCloudAssetSummaryResponseBody {
	s.GroupedFields = v
	return s
}

func (s *GetCloudAssetSummaryResponseBody) SetRequestId(v string) *GetCloudAssetSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBody) Validate() error {
	if s.GroupedFields != nil {
		if err := s.GroupedFields.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetCloudAssetSummaryResponseBodyGroupedFields struct {
	// The statistics of cloud services.
	CloudAssetSummaryMetas []*GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas `json:"CloudAssetSummaryMetas,omitempty" xml:"CloudAssetSummaryMetas,omitempty" type:"Repeated"`
	// The total number of cloud service instances.
	//
	// example:
	//
	// 919
	InstanceCountTotal *int32 `json:"InstanceCountTotal,omitempty" xml:"InstanceCountTotal,omitempty"`
	// The total number of cloud service instances that are at risk.
	//
	// example:
	//
	// 544
	InstanceRiskCountTotal *int32 `json:"InstanceRiskCountTotal,omitempty" xml:"InstanceRiskCountTotal,omitempty"`
}

func (s GetCloudAssetSummaryResponseBodyGroupedFields) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetSummaryResponseBodyGroupedFields) GoString() string {
	return s.String()
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) GetCloudAssetSummaryMetas() []*GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	return s.CloudAssetSummaryMetas
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) GetInstanceCountTotal() *int32 {
	return s.InstanceCountTotal
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) GetInstanceRiskCountTotal() *int32 {
	return s.InstanceRiskCountTotal
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) SetCloudAssetSummaryMetas(v []*GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) *GetCloudAssetSummaryResponseBodyGroupedFields {
	s.CloudAssetSummaryMetas = v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) SetInstanceCountTotal(v int32) *GetCloudAssetSummaryResponseBodyGroupedFields {
	s.InstanceCountTotal = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) SetInstanceRiskCountTotal(v int32) *GetCloudAssetSummaryResponseBodyGroupedFields {
	s.InstanceRiskCountTotal = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) Validate() error {
	if s.CloudAssetSummaryMetas != nil {
		for _, item := range s.CloudAssetSummaryMetas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas struct {
	// The subtype of the cloud service.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the cloud service. Valid values:
	//
	// 	- **0**: Elastic Compute Service (ECS)
	//
	// 	- **1**: Server Load Balancer (SLB)
	//
	// 	- **3**: ApsaraDB RDS
	//
	// 	- **4**: ApsaraDB for MongoDB (MongoDB)
	//
	// 	- **5**: ApsaraDB for Redis (Redis)
	//
	// 	- **6**: Container Registry
	//
	// 	- **8**: Container Service for Kubernetes (ACK)
	//
	// 	- **9**: Virtual Private Cloud (VPC)
	//
	// 	- **11**: ActionTrail
	//
	// 	- **12**: Alibaba Cloud CDN (CDN)
	//
	// 	- **13**: Certificate Management Service (formerly SSL Certificates Service)
	//
	// 	- **14**: Apsara Devops
	//
	// 	- **15**: Resource Access Management (RAM)
	//
	// 	- **16**: Anti-DDoS
	//
	// 	- **17**: Web Application Firewall (WAF)
	//
	// 	- **18**: Object Storage Service (OSS)
	//
	// 	- **19**: PolarDB
	//
	// 	- **20**: ApsaraDB RDS for PostgreSQL
	//
	// 	- **21**: Microservices Engine (MSE)
	//
	// 	- **22**: File Storage NAS (NAS)
	//
	// 	- **23**: Data Security Center (DSC)
	//
	// 	- **24**: Elastic IP Address (EIP)
	//
	// 	- **25**: Identity as a Service (IDaaS) - Enterprise Identity Access Management (EIAM)
	//
	// 	- **26**: PolarDB for Xscale (PolarDB-X)
	//
	// 	- **27**: Elasticsearch
	//
	// example:
	//
	// 16
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The total number of cloud service instances of this type.
	//
	// example:
	//
	// 16
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The total number of cloud service instances that are at risk of this type.
	//
	// example:
	//
	// 5
	InstanceRiskCount *int32 `json:"InstanceRiskCount,omitempty" xml:"InstanceRiskCount,omitempty"`
	// The server type. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud
	//
	// 	- **1**: an asset outside Alibaba Cloud
	//
	// 	- **2**: an asset in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: an asset provided by a third-party service provider
	//
	// 	- **8**: a lightweight asset
	//
	// example:
	//
	// 3
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GoString() string {
	return s.String()
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetAssetType() *int32 {
	return s.AssetType
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetInstanceRiskCount() *int32 {
	return s.InstanceRiskCount
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetVendor() *int32 {
	return s.Vendor
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetAssetSubType(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.AssetSubType = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetAssetType(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.AssetType = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetInstanceCount(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.InstanceCount = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetInstanceRiskCount(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.InstanceRiskCount = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetVendor(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.Vendor = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) Validate() error {
	return dara.Validate(s)
}
