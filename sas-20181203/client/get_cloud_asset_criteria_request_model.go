// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAssetTypes(v []*GetCloudAssetCriteriaRequestCloudAssetTypes) *GetCloudAssetCriteriaRequest
	GetCloudAssetTypes() []*GetCloudAssetCriteriaRequestCloudAssetTypes
	SetValue(v string) *GetCloudAssetCriteriaRequest
	GetValue() *string
}

type GetCloudAssetCriteriaRequest struct {
	// The types of cloud assets.
	CloudAssetTypes []*GetCloudAssetCriteriaRequestCloudAssetTypes `json:"CloudAssetTypes,omitempty" xml:"CloudAssetTypes,omitempty" type:"Repeated"`
	// The keyword for fuzzy search when you query the asset.
	//
	// example:
	//
	// testwww
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetCloudAssetCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetCriteriaRequest) GoString() string {
	return s.String()
}

func (s *GetCloudAssetCriteriaRequest) GetCloudAssetTypes() []*GetCloudAssetCriteriaRequestCloudAssetTypes {
	return s.CloudAssetTypes
}

func (s *GetCloudAssetCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *GetCloudAssetCriteriaRequest) SetCloudAssetTypes(v []*GetCloudAssetCriteriaRequestCloudAssetTypes) *GetCloudAssetCriteriaRequest {
	s.CloudAssetTypes = v
	return s
}

func (s *GetCloudAssetCriteriaRequest) SetValue(v string) *GetCloudAssetCriteriaRequest {
	s.Value = &v
	return s
}

func (s *GetCloudAssetCriteriaRequest) Validate() error {
	return dara.Validate(s)
}

type GetCloudAssetCriteriaRequestCloudAssetTypes struct {
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
	// 2
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
	// 9
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
}

func (s GetCloudAssetCriteriaRequestCloudAssetTypes) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetCriteriaRequestCloudAssetTypes) GoString() string {
	return s.String()
}

func (s *GetCloudAssetCriteriaRequestCloudAssetTypes) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *GetCloudAssetCriteriaRequestCloudAssetTypes) GetAssetType() *int32 {
	return s.AssetType
}

func (s *GetCloudAssetCriteriaRequestCloudAssetTypes) SetAssetSubType(v int32) *GetCloudAssetCriteriaRequestCloudAssetTypes {
	s.AssetSubType = &v
	return s
}

func (s *GetCloudAssetCriteriaRequestCloudAssetTypes) SetAssetType(v int32) *GetCloudAssetCriteriaRequestCloudAssetTypes {
	s.AssetType = &v
	return s
}

func (s *GetCloudAssetCriteriaRequestCloudAssetTypes) Validate() error {
	return dara.Validate(s)
}
