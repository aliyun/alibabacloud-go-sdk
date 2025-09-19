// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAssetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetType(v string) *RefreshAssetsRequest
	GetAssetType() *string
	SetCloudAssetSubType(v int32) *RefreshAssetsRequest
	GetCloudAssetSubType() *int32
	SetCloudAssetType(v int32) *RefreshAssetsRequest
	GetCloudAssetType() *int32
	SetVendor(v int32) *RefreshAssetsRequest
	GetVendor() *int32
}

type RefreshAssetsRequest struct {
	// The type of the asset that you want to synchronize. Valid values:
	//
	// 	- **cloud_product**: Alibaba Cloud service
	//
	// 	- **ecs**: Elastic Compute Service (ECS) instance
	//
	// 	- **container_image**: container image
	//
	// example:
	//
	// cloud_product
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The subtype of the cloud service.
	//
	// >  The following list describes the subtypes of cloud services.
	//
	// example:
	//
	// 0
	CloudAssetSubType *int32 `json:"CloudAssetSubType,omitempty" xml:"CloudAssetSubType,omitempty"`
	// The type of the cloud service. Valid values:
	//
	// 	- **0**: ECS
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
	// example:
	//
	// 0
	CloudAssetType *int32 `json:"CloudAssetType,omitempty" xml:"CloudAssetType,omitempty"`
	// The type of the server. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud
	//
	// 	- **1**: a third-party cloud asset
	//
	// 	- **2**: an asset in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: an asset provided by another cloud
	//
	// 	- **8**: a lightweight asset
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s RefreshAssetsRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshAssetsRequest) GoString() string {
	return s.String()
}

func (s *RefreshAssetsRequest) GetAssetType() *string {
	return s.AssetType
}

func (s *RefreshAssetsRequest) GetCloudAssetSubType() *int32 {
	return s.CloudAssetSubType
}

func (s *RefreshAssetsRequest) GetCloudAssetType() *int32 {
	return s.CloudAssetType
}

func (s *RefreshAssetsRequest) GetVendor() *int32 {
	return s.Vendor
}

func (s *RefreshAssetsRequest) SetAssetType(v string) *RefreshAssetsRequest {
	s.AssetType = &v
	return s
}

func (s *RefreshAssetsRequest) SetCloudAssetSubType(v int32) *RefreshAssetsRequest {
	s.CloudAssetSubType = &v
	return s
}

func (s *RefreshAssetsRequest) SetCloudAssetType(v int32) *RefreshAssetsRequest {
	s.CloudAssetType = &v
	return s
}

func (s *RefreshAssetsRequest) SetVendor(v int32) *RefreshAssetsRequest {
	s.Vendor = &v
	return s
}

func (s *RefreshAssetsRequest) Validate() error {
	return dara.Validate(s)
}
