// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAssetTypes(v []*GetCloudAssetSummaryRequestCloudAssetTypes) *GetCloudAssetSummaryRequest
	GetCloudAssetTypes() []*GetCloudAssetSummaryRequestCloudAssetTypes
	SetIsSaleData(v bool) *GetCloudAssetSummaryRequest
	GetIsSaleData() *bool
	SetVendors(v []*int32) *GetCloudAssetSummaryRequest
	GetVendors() []*int32
}

type GetCloudAssetSummaryRequest struct {
	// The list of asset type information of cloud assets.
	CloudAssetTypes []*GetCloudAssetSummaryRequestCloudAssetTypes `json:"CloudAssetTypes,omitempty" xml:"CloudAssetTypes,omitempty" type:"Repeated"`
	IsSaleData      *bool                                         `json:"IsSaleData,omitempty" xml:"IsSaleData,omitempty"`
	// The list of cloud vendors to query.
	Vendors []*int32 `json:"Vendors,omitempty" xml:"Vendors,omitempty" type:"Repeated"`
}

func (s GetCloudAssetSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetCloudAssetSummaryRequest) GetCloudAssetTypes() []*GetCloudAssetSummaryRequestCloudAssetTypes {
	return s.CloudAssetTypes
}

func (s *GetCloudAssetSummaryRequest) GetIsSaleData() *bool {
	return s.IsSaleData
}

func (s *GetCloudAssetSummaryRequest) GetVendors() []*int32 {
	return s.Vendors
}

func (s *GetCloudAssetSummaryRequest) SetCloudAssetTypes(v []*GetCloudAssetSummaryRequestCloudAssetTypes) *GetCloudAssetSummaryRequest {
	s.CloudAssetTypes = v
	return s
}

func (s *GetCloudAssetSummaryRequest) SetIsSaleData(v bool) *GetCloudAssetSummaryRequest {
	s.IsSaleData = &v
	return s
}

func (s *GetCloudAssetSummaryRequest) SetVendors(v []*int32) *GetCloudAssetSummaryRequest {
	s.Vendors = v
	return s
}

func (s *GetCloudAssetSummaryRequest) Validate() error {
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

type GetCloudAssetSummaryRequestCloudAssetTypes struct {
	// The subtype of the cloud service.
	//
	// The asset type-subtype. Valid values:
	//
	// - **0**: Elastic Compute Service (ECS)
	//
	//     	- **1**: Disk (Storage)
	//
	//     	- **2**: Security Group
	//
	//     	- **100**: Instance
	//
	// - **1**: Server Load Balancer
	//
	//     	- **0**: Server Load Balancer (SLB)
	//
	//     	- **1**: Application Load Balancer (ALB)
	//
	// - **3**: ApsaraDB RDS
	//
	//     	- **0**: Instance
	//
	// - **4**: ApsaraDB for MongoDB
	//
	//     	- **0**: Instance
	//
	// - **5**: ApsaraDB for Tair (compatible with Redis)
	//
	//     	- **0**: Instance
	//
	// - **6**: Container Registry
	//
	//     	- **1**: Enterprise Edition
	//
	//     	- **2**: Personal Edition
	//
	// - **8**: Container Service for Kubernetes (ACK)
	//
	//     	- **0**: Cluster
	//
	// - **9**: Virtual Private Cloud (VPC)
	//
	//     	- **0**: NAT Gateway
	//
	//     	- **1**: EIP
	//
	//     	- **2**: VPN
	//
	//     	- **3**: FLOW_LOG
	//
	// - **11**: ActionTrail
	//
	//     	- **0**: Trail
	//
	// - **12**: Alibaba Cloud CDN
	//
	//     	- **0**: Instance
	//
	// - **13**: Certificate Management Service (formerly SSL Certificates Service)
	//
	//     	- **0**: Certificate
	//
	// - **14**: Apsara Devops
	//
	//     	- **0**: Organization
	//
	// - **16**: Anti-DDoS
	//
	//     	- **0**: Instance
	//
	// - **17**: Web Application Firewall (WAF)
	//
	//     	- **0**: Domain name
	//
	// - **18**: Object Storage Service (OSS)
	//
	//     	- **0**: Bucket
	//
	// - **19**: PolarDB
	//
	//     	- **0**: Cluster
	//
	// - **20**: ApsaraDB RDS for PostgreSQL
	//
	//     	- **0**: Instance
	//
	// - **21**: Microservices Engine (MSE)
	//
	//     	- **0**: Cluster
	//
	// - **22**: Apsara File Storage NAS
	//
	//     	- **0**: File system
	//
	// - **23**: Data Security Center (DSC)
	//
	//     	- **0**: Instance
	//
	// - **24**: Elastic IP Address (EIP)
	//
	//     	- **0**: Anycast EIP
	//
	// - **25**: Identity as a Service - EIAM
	//
	//     	- **0**: Instance
	//
	// - **26**: PolarDB-X
	//
	//     	- **0**: Instance
	//
	// - **27**: Elasticsearch
	//
	//     	- **0**: Instance
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of asset. Valid values:
	//
	// - **0**: Elastic Compute Service (ECS)
	//
	// - **1**: Server Load Balancer (SLB)
	//
	// - **3**: ApsaraDB RDS
	//
	// - **4**: ApsaraDB for MongoDB
	//
	// - **5**: ApsaraDB for Tair (compatible with Redis)
	//
	// - **6**: Container Registry
	//
	// - **8**: Container Service for Kubernetes (ACK)
	//
	// - **9**: Virtual Private Cloud (VPC)
	//
	// - **11**: ActionTrail
	//
	// - **12**: Alibaba Cloud CDN
	//
	// - **13**: Certificate Management Service (formerly SSL Certificates Service)
	//
	// - **14**: Apsara Devops
	//
	// - **16**: Anti-DDoS
	//
	// - **17**: Web Application Firewall (WAF)
	//
	// - **18**: Object Storage Service (OSS)
	//
	// - **19**: PolarDB
	//
	// - **20**: ApsaraDB RDS for PostgreSQL
	//
	// - **21**: Microservices Engine (MSE)
	//
	// - **22**: Apsara File Storage NAS
	//
	// - **23**: Data Security Center (DSC)
	//
	// - **24**: Elastic IP Address (EIP)
	//
	// - **25**: Identity as a Service - EIAM
	//
	// - **26**: PolarDB-X
	//
	// - **27**: Elasticsearch
	//
	// example:
	//
	// 4
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Server vendor. Valid values:
	//
	// - **0**: Alibaba Cloud assets
	//
	// - **1**: Off-cloud assets
	//
	// - **2**: IDC assets
	//
	// - **3**, **4**, **5**, **7**: Other cloud assets
	//
	// - **8**: Lightweight assets
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetCloudAssetSummaryRequestCloudAssetTypes) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetSummaryRequestCloudAssetTypes) GoString() string {
	return s.String()
}

func (s *GetCloudAssetSummaryRequestCloudAssetTypes) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *GetCloudAssetSummaryRequestCloudAssetTypes) GetAssetType() *int32 {
	return s.AssetType
}

func (s *GetCloudAssetSummaryRequestCloudAssetTypes) GetVendor() *int32 {
	return s.Vendor
}

func (s *GetCloudAssetSummaryRequestCloudAssetTypes) SetAssetSubType(v int32) *GetCloudAssetSummaryRequestCloudAssetTypes {
	s.AssetSubType = &v
	return s
}

func (s *GetCloudAssetSummaryRequestCloudAssetTypes) SetAssetType(v int32) *GetCloudAssetSummaryRequestCloudAssetTypes {
	s.AssetType = &v
	return s
}

func (s *GetCloudAssetSummaryRequestCloudAssetTypes) SetVendor(v int32) *GetCloudAssetSummaryRequestCloudAssetTypes {
	s.Vendor = &v
	return s
}

func (s *GetCloudAssetSummaryRequestCloudAssetTypes) Validate() error {
	return dara.Validate(s)
}
