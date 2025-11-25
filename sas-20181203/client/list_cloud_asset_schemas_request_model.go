// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAssetSchemasRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssociatedDataOnly(v bool) *ListCloudAssetSchemasRequest
	GetAssociatedDataOnly() *bool
	SetCloudAssetTypes(v []*ListCloudAssetSchemasRequestCloudAssetTypes) *ListCloudAssetSchemasRequest
	GetCloudAssetTypes() []*ListCloudAssetSchemasRequestCloudAssetTypes
	SetCurrentPage(v int32) *ListCloudAssetSchemasRequest
	GetCurrentPage() *int32
	SetDataNames(v []*string) *ListCloudAssetSchemasRequest
	GetDataNames() []*string
	SetLang(v string) *ListCloudAssetSchemasRequest
	GetLang() *string
	SetPageSize(v int32) *ListCloudAssetSchemasRequest
	GetPageSize() *int32
}

type ListCloudAssetSchemasRequest struct {
	// Whether to filter out attributes that can be associated with other assets.
	//
	// example:
	//
	// true
	AssociatedDataOnly *bool `json:"AssociatedDataOnly,omitempty" xml:"AssociatedDataOnly,omitempty"`
	// The types of cloud assets.
	CloudAssetTypes []*ListCloudAssetSchemasRequestCloudAssetTypes `json:"CloudAssetTypes,omitempty" xml:"CloudAssetTypes,omitempty" type:"Repeated"`
	// The number of the page to return. Default value: **1**.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// List of asset data names
	DataNames []*string `json:"DataNames,omitempty" xml:"DataNames,omitempty" type:"Repeated"`
	// The language type for requesting and receiving messages, with a default value of **zh**. The values can be: - **zh**: Chinese - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListCloudAssetSchemasRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetSchemasRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAssetSchemasRequest) GetAssociatedDataOnly() *bool {
	return s.AssociatedDataOnly
}

func (s *ListCloudAssetSchemasRequest) GetCloudAssetTypes() []*ListCloudAssetSchemasRequestCloudAssetTypes {
	return s.CloudAssetTypes
}

func (s *ListCloudAssetSchemasRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudAssetSchemasRequest) GetDataNames() []*string {
	return s.DataNames
}

func (s *ListCloudAssetSchemasRequest) GetLang() *string {
	return s.Lang
}

func (s *ListCloudAssetSchemasRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudAssetSchemasRequest) SetAssociatedDataOnly(v bool) *ListCloudAssetSchemasRequest {
	s.AssociatedDataOnly = &v
	return s
}

func (s *ListCloudAssetSchemasRequest) SetCloudAssetTypes(v []*ListCloudAssetSchemasRequestCloudAssetTypes) *ListCloudAssetSchemasRequest {
	s.CloudAssetTypes = v
	return s
}

func (s *ListCloudAssetSchemasRequest) SetCurrentPage(v int32) *ListCloudAssetSchemasRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudAssetSchemasRequest) SetDataNames(v []*string) *ListCloudAssetSchemasRequest {
	s.DataNames = v
	return s
}

func (s *ListCloudAssetSchemasRequest) SetLang(v string) *ListCloudAssetSchemasRequest {
	s.Lang = &v
	return s
}

func (s *ListCloudAssetSchemasRequest) SetPageSize(v int32) *ListCloudAssetSchemasRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudAssetSchemasRequest) Validate() error {
	if s.CloudAssetTypes != nil {
		for _, item := range s.CloudAssetTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListCloudAssetSchemasRequestCloudAssetTypes struct {
	// The subtype of the cloud service or asset. Valid values:
	//
	// 	- **0**: ECS
	//
	//     	- **0**: instance
	//
	//     	- **1**: disk (storage)
	//
	//     	- **2**: security group
	//
	// 	- **1**: SLB
	//
	//     	- **0**: SLB
	//
	//     	- **1**: Application Load Balancer (ALB)
	//
	// 	- **3**: ApsaraDB RDS
	//
	//     	- **0**: instance
	//
	// 	- **4**: MongoDB
	//
	//     	- **0**: instance
	//
	// 	- **5**: Redis
	//
	//     	- **0**: instance
	//
	// 	- **6**: Container Registry
	//
	//     	- **1**: Enterprise Edition
	//
	//     	- **2**: Personal Edition
	//
	// 	- **8**: ACK
	//
	//     	- **0**: cluster
	//
	// 	- **9**: VPC
	//
	//     	- **0**: NAT gateway
	//
	//     	- **1**: Elastic IP address (EIP)
	//
	//     	- **2**: VPN
	//
	//     	- **3**: VPC Flow Logs
	//
	// 	- **11**: ActionTrail
	//
	//     	- **0**: trail
	//
	// 	- **12**: CDN
	//
	//     	- **0**: instance
	//
	// 	- **13**: Certificate Management Service (formerly SSL Certificates Service)
	//
	//     	- **0**: certificate
	//
	// 	- **14**: Apsara Devops
	//
	//     	- **0**: organization
	//
	// 	- **16**: Anti-DDoS
	//
	//     	- **0**: instance
	//
	// 	- **17**: WAF
	//
	//     	- **0**: domain name
	//
	// 	- **18**: OSS
	//
	//     	- **0**: bucket
	//
	// 	- **19**: PolarDB
	//
	//     	- **0**: cluster
	//
	// 	- **20**: ApsaraDB RDS for PostgreSQL
	//
	//     	- **0**: instance
	//
	// 	- **21**: MSE
	//
	//     	- **0**: cluster
	//
	// 	- **22**: NAS
	//
	//     	- **0**: file system
	//
	// 	- **23**: DSC
	//
	//     	- **0**: instance
	//
	// 	- **24**: EIP
	//
	//     	- **0**: Anycast EIP
	//
	// 	- **25**: IDaaS EIAM
	//
	//     	- **0**: instance
	//
	// 	- **26**: PolarDB-X
	//
	//     	- **0**: instance
	//
	// 	- **27**: Elasticsearch
	//
	//     	- **0**: instance
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the asset. Valid values:
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
	// 	- **25**: IDaaS EIAM
	//
	// 	- **26**: PolarDB-X
	//
	// 	- **27**: Elasticsearch
	//
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The server type. Valid values:
	//
	// 	- **0**: a cloud asset provided by Alibaba Cloud
	//
	// 	- **1**: a cloud asset outside Alibaba Cloud
	//
	// 	- **2**: a cloud asset in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: a cloud asset provided by a third-party service provider
	//
	// 	- **8**: a lightweight cloud asset
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListCloudAssetSchemasRequestCloudAssetTypes) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetSchemasRequestCloudAssetTypes) GoString() string {
	return s.String()
}

func (s *ListCloudAssetSchemasRequestCloudAssetTypes) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *ListCloudAssetSchemasRequestCloudAssetTypes) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListCloudAssetSchemasRequestCloudAssetTypes) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListCloudAssetSchemasRequestCloudAssetTypes) SetAssetSubType(v int32) *ListCloudAssetSchemasRequestCloudAssetTypes {
	s.AssetSubType = &v
	return s
}

func (s *ListCloudAssetSchemasRequestCloudAssetTypes) SetAssetType(v int32) *ListCloudAssetSchemasRequestCloudAssetTypes {
	s.AssetType = &v
	return s
}

func (s *ListCloudAssetSchemasRequestCloudAssetTypes) SetVendor(v int32) *ListCloudAssetSchemasRequestCloudAssetTypes {
	s.Vendor = &v
	return s
}

func (s *ListCloudAssetSchemasRequestCloudAssetTypes) Validate() error {
	return dara.Validate(s)
}
