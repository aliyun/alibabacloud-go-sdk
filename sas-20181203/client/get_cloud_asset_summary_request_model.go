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
	SetVendors(v []*int32) *GetCloudAssetSummaryRequest
	GetVendors() []*int32
}

type GetCloudAssetSummaryRequest struct {
	// List of asset type information for cloud assets
	CloudAssetTypes []*GetCloudAssetSummaryRequestCloudAssetTypes `json:"CloudAssetTypes,omitempty" xml:"CloudAssetTypes,omitempty" type:"Repeated"`
	// List of cloud vendors to be queried.
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

func (s *GetCloudAssetSummaryRequest) GetVendors() []*int32 {
	return s.Vendors
}

func (s *GetCloudAssetSummaryRequest) SetCloudAssetTypes(v []*GetCloudAssetSummaryRequestCloudAssetTypes) *GetCloudAssetSummaryRequest {
	s.CloudAssetTypes = v
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
	// Subtypes of cloud products. Asset type-subtype. Values:
	//
	// - **0**: ECS (Elastic Compute Service)
	//
	//   	- **1**: Disk (Storage)
	//
	//   	- **2**: Security Group
	//
	//   	- **100**: Instance
	//
	// - **1**: Load Balancer
	//
	//   	- **0**: Load Balancer
	//
	//   	- **1**: Application Load Balancer
	//
	// - **3**: ApsaraDB RDS
	//
	//   	- **0**: Instance
	//
	// - **4**: ApsaraDB for MongoDB
	//
	//   	- **0**: Instance
	//
	// - **5**: ApsaraDB Tair (Redis Compatible)
	//
	//   	- **0**: Instance
	//
	// - **6**: Container Registry
	//
	//   	- **1**: Enterprise Edition
	//
	//   	- **2**: Personal Edition
	//
	// - **8**: Container Service for Kubernetes
	//
	//   	- **0**: Cluster
	//
	// - **9**: Virtual Private Cloud (VPC)
	//
	//   	- **0**: NAT Gateway
	//
	//   	- **1**: EIP (Elastic IP)
	//
	//   	- **2**: VPN
	//
	//   	- **3**: FLOW_LOG
	//
	// - **11**: ActionTrail
	//
	//   	- **0**: Trail
	//
	// - **12**: CDN
	//
	//   	- **0**: Instance
	//
	// - **13**: Digital Certificate Management Service (formerly SSL Certificates)
	//
	//   	- **0**: Certificate
	//
	// - **14**: DevOps
	//
	//   	- **0**: Organization
	//
	// - **16**: DDoS Protection
	//
	//   	- **0**: Instance
	//
	// - **17**: Web Application Firewall
	//
	//   	- **0**: Domain
	//
	// - **18**: Object Storage
	//
	//   	- **0**: Bucket
	//
	// - **19**: PolarDB (Cloud-Native Relational Database)
	//
	//   	- **0**: Cluster
	//
	// - **20**: ApsaraDB for PostgreSQL
	//
	//   	- **0**: Instance
	//
	// - **21**: Microservices Engine
	//
	//   	- **0**: Cluster
	//
	// - **22**: File Storage NAS
	//
	//   	- **0**: File System
	//
	// - **23**: Data Security Center
	//
	//   	- **0**: Instance
	//
	// - **24**: Elastic Public IP
	//
	//   	- **0**: Anycast Elastic Public IP
	//
	// - **25**: Cloud Identity Service - EIAM
	//
	//   	- **0**: Instance
	//
	// - **26**: PolarDB-X
	//
	//   	- **0**: Instance
	//
	// - **27**: Elasticsearch
	//
	//   	- **0**: Instance
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of asset. Values:
	//
	// - **0**: Elastic Compute Service (ECS)
	//
	// - **1**: Load Balancer
	//
	// - **3**: ApsaraDB for RDS
	//
	// - **4**: ApsaraDB for MongoDB
	//
	// - **5**: ApsaraDB for Tair (Redis compatible)
	//
	// - **6**: Container Registry
	//
	// - **8**: Container Service for Kubernetes
	//
	// - **9**: Virtual Private Cloud (VPC)
	//
	// - **11**: ActionTrail
	//
	// - **12**: Content Delivery Network (CDN)
	//
	// - **13**: SSL Certificates (now known as Certificate Management Service)
	//
	// - **14**: DevOps
	//
	// - **16**: DDoS Protection
	//
	// - **17**: Web Application Firewall
	//
	// - **18**: Object Storage Service (OSS)
	//
	// - **19**: PolarDB
	//
	// - **20**: ApsaraDB for PostgreSQL
	//
	// - **21**: Microservices Engine
	//
	// - **22**: File Storage NAS
	//
	// - **23**: Data Security Center
	//
	// - **24**: Elastic IP Address
	//
	// - **25**: Cloud Identity Service - EIAM
	//
	// - **26**: PolarDB-X
	//
	// - **27**: Elasticsearch
	//
	// example:
	//
	// 4
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Server vendor. Values:
	//
	// - **0**: Alibaba Cloud Asset
	//
	// - **1**: Non-cloud Asset
	//
	// - **2**: IDC Asset
	//
	// - **3**, **4**, **5**, **7**: Other Cloud Assets
	//
	// - **8**: Lightweight Asset
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
