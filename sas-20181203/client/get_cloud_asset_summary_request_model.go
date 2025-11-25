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
	// 云产品的子类型。
	//
	// 资产的类型-子类型。取值：
	//
	// - **0**：云服务器 ECS
	//
	//     	- **1**：磁盘（存储）
	//
	//     	- **2**：安全组
	//
	//     	- **100**：实例
	//
	// - **1**：负载均衡
	//
	//     	- **0**：负载均衡
	//
	//     	- **1**：应用型负载均衡
	//
	// - **3**：云数据库 RDS
	//
	//     	- **0**：实例
	//
	// - **4**：云数据库 MongoDB 版
	//
	//     	- **0**：实例
	//
	// - **5**：云数据库 Tair（兼容 Redis）
	//
	//     	- **0**：实例
	//
	// - **6**：容器镜像服务
	//
	//     	- **1**：企业版
	//
	//     	- **2**：个人版
	//
	// - **8**：容器服务Kubernetes版
	//
	//     	- **0**：集群
	//
	// - **9**：专有网络VPC
	//
	//     	- **0**：NAT网关
	//
	//     	- **1**：EIP
	//
	//     	- **2**：VPN
	//
	//     	- **3**：FLOW_LOG
	//
	// - **11**：操作审计
	//
	//     	- **0**：跟踪
	//
	// - **12**：CDN
	//
	//     	- **0**：实例
	//
	// - **13**：数字证书管理服务（原SSL证书）
	//
	//     	- **0**：证书
	//
	// - **14**：云效
	//
	//     	- **0**：组织
	//
	// - **16**：DDoS防护
	//
	//     	- **0**：实例
	//
	// - **17**：Web应用防火墙
	//
	//     	- **0**：域名
	//
	// - **18**：对象存储
	//
	//     	- **0**：Bucket
	//
	// - **19**：云原生关系型数据库 PolarDB
	//
	//     	- **0**：集群
	//
	// - **20**：云数据库 PostgreSQL 版
	//
	//     	- **0**：实例
	//
	// - **21**：微服务引擎
	//
	//     	- **0**：集群
	//
	// - **22**：文件存储NAS
	//
	//     	- **0**：文件系统
	//
	// - **23**：数据安全中心
	//
	//     	- **0**：实例
	//
	// - **24**：弹性公网IP
	//
	//     	- **0**：任播弹性公网IP
	//
	// - **25**：云身份服务-EIAM
	//
	//     	- **0**：实例
	//
	// - **26**：PolarDB-X
	//
	//     	- **0**：实例
	//
	// - **27**：Elasticsearch
	//
	//     	- **0**：实例
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
