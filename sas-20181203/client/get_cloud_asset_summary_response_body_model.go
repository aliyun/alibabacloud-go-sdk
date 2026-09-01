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
	// The cloud asset summary information.
	GroupedFields *GetCloudAssetSummaryResponseBodyGroupedFields `json:"GroupedFields,omitempty" xml:"GroupedFields,omitempty" type:"Struct"`
	// The ID of the request. Alibaba Cloud generates a unique identifier for each request. You can use the request ID to troubleshoot issues.
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
	// The list of cloud service statistics information.
	CloudAssetSummaryMetas []*GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas `json:"CloudAssetSummaryMetas,omitempty" xml:"CloudAssetSummaryMetas,omitempty" type:"Repeated"`
	// The total number of cloud service instances.
	//
	// example:
	//
	// 919
	InstanceCountTotal *int32 `json:"InstanceCountTotal,omitempty" xml:"InstanceCountTotal,omitempty"`
	// The total number of at-risk cloud service instances.
	//
	// example:
	//
	// 544
	InstanceRiskCountTotal *int32 `json:"InstanceRiskCountTotal,omitempty" xml:"InstanceRiskCountTotal,omitempty"`
	// The total number of cloud services billed by instance.
	//
	// example:
	//
	// 10
	InstanceSaleCountTotal *int32 `json:"InstanceSaleCountTotal,omitempty" xml:"InstanceSaleCountTotal,omitempty"`
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

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) GetInstanceSaleCountTotal() *int32 {
	return s.InstanceSaleCountTotal
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

func (s *GetCloudAssetSummaryResponseBodyGroupedFields) SetInstanceSaleCountTotal(v int32) *GetCloudAssetSummaryResponseBodyGroupedFields {
	s.InstanceSaleCountTotal = &v
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
	// - **0**: Elastic Compute Service (ECS) server
	//
	// - **1**: load balancing
	//
	// - **3**: ApsaraDB RDS database
	//
	// - **4**: ApsaraDB for MongoDB database
	//
	// - **5**: Tair (Redis® OSS-Compatible) database
	//
	// - **6**: Container Registry
	//
	// - **8**: Container Service for Kubernetes (ACK)
	//
	// - **9**: Virtual Private Cloud (VPC)
	//
	// - **11**: ActionTrail
	//
	// - **12**: CDN
	//
	// - **13**: Certificate Management Service (formerly SSL Certificates Service)
	//
	// - **14**: Apsara Devops
	//
	// - **15**: Resource Access Management (RAM)
	//
	// - **16**: Anti-DDoS
	//
	// - **17**: Web Application Firewall (WAF)
	//
	// - **18**: Object Storage Service (OSS)
	//
	// - **19**: cloud-native relational database PolarDB
	//
	// - **20**: ApsaraDB RDS for PostgreSQL database
	//
	// - **21**: Microservices Engine (MSE)
	//
	// - **22**: Apsara File Storage NAS
	//
	// - **23**: Data Security Center (DSC)
	//
	// - **24**: Elastic IP Address (EIP)
	//
	// - **25**: Alibaba Cloud IDaaS EIAM
	//
	// - **26**: PolarDB-X
	//
	// - **27**: Elasticsearch
	//
	// example:
	//
	// 16
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The total number of instances of this cloud service type.
	//
	// example:
	//
	// 16
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The total number of at-risk instances of this cloud service type.
	//
	// example:
	//
	// 5
	InstanceRiskCount *int32 `json:"InstanceRiskCount,omitempty" xml:"InstanceRiskCount,omitempty"`
	// The number of assets billed by instance.
	//
	// example:
	//
	// 1
	InstanceSaleCount *int32 `json:"InstanceSaleCount,omitempty" xml:"InstanceSaleCount,omitempty"`
	// Indicates whether the asset is billed by instance. Valid values:
	//
	// - **true**: Billed by instance.
	//
	// - **false**: Not billed by instance.
	IsInstanceSale *bool `json:"IsInstanceSale,omitempty" xml:"IsInstanceSale,omitempty"`
	// The server vendor. Valid values:
	//
	// - **0**: Alibaba Cloud asset
	//
	// - **1**: non-cloud asset
	//
	// - **2**: IDC asset
	//
	// - **3**, **4**, **5**, **7**: third-party cloud asset
	//
	// - **8**: lightweight asset
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

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetInstanceSaleCount() *int32 {
	return s.InstanceSaleCount
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) GetIsInstanceSale() *bool {
	return s.IsInstanceSale
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

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetInstanceSaleCount(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.InstanceSaleCount = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetIsInstanceSale(v bool) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.IsInstanceSale = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) SetVendor(v int32) *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas {
	s.Vendor = &v
	return s
}

func (s *GetCloudAssetSummaryResponseBodyGroupedFieldsCloudAssetSummaryMetas) Validate() error {
	return dara.Validate(s)
}
