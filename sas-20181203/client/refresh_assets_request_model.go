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
	SetResourceDirectoryAccountId(v int64) *RefreshAssetsRequest
	GetResourceDirectoryAccountId() *int64
	SetVendor(v int32) *RefreshAssetsRequest
	GetVendor() *int32
}

type RefreshAssetsRequest struct {
	// The type of asset to synchronize. Default value: **ecs**. Valid values:
	//
	// - **cloud_product**: cloud product
	//
	// - **ecs**: server
	//
	// - **container_image**: container image
	//
	// example:
	//
	// cloud_product
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The subtype of the cloud product.
	//
	// > Refer to the following list for valid values.
	//
	// example:
	//
	// 0
	CloudAssetSubType *int32 `json:"CloudAssetSubType,omitempty" xml:"CloudAssetSubType,omitempty"`
	// The type of cloud product. Valid values:
	//
	// - **0**: cloud server
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
	// - **8**: container service for Kubernetes
	//
	// - **9**: VPC
	//
	// - **11**: ActionTrail
	//
	// - **12**: CDN
	//
	// - **13**: Certificate Management Service (formerly SSL Certificates Service)
	//
	// - **14**: Apsara Devops
	//
	// - **15**: access control
	//
	// - **16**: Anti-DDoS
	//
	// - **17**: Web application firewall
	//
	// - **18**: OSS
	//
	// - **19**: cloud-native relational database PolarDB
	//
	// - **20**: ApsaraDB RDS for PostgreSQL database
	//
	// - **21**: microservices engine
	//
	// - **22**: file storage NAS
	//
	// - **23**: Data Security Center
	//
	// - **24**: EIP
	//
	// example:
	//
	// 0
	CloudAssetType *int32 `json:"CloudAssetType,omitempty" xml:"CloudAssetType,omitempty"`
	// The ID of the Alibaba Cloud account of the member accounts in the resource directory.
	//
	// > Call the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	//
	// example:
	//
	// 127608589417****
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
	// The server vendor. Valid values:
	//
	// - **0**: Alibaba Cloud asset
	//
	// - **1**: Non-cloud asset
	//
	// - **2**: IDC asset
	//
	// - **3**, **4**, **5**, **7**: Third-party cloud asset
	//
	// - **8**: Lightweight asset
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

func (s *RefreshAssetsRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
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

func (s *RefreshAssetsRequest) SetResourceDirectoryAccountId(v int64) *RefreshAssetsRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *RefreshAssetsRequest) SetVendor(v int32) *RefreshAssetsRequest {
	s.Vendor = &v
	return s
}

func (s *RefreshAssetsRequest) Validate() error {
	return dara.Validate(s)
}
