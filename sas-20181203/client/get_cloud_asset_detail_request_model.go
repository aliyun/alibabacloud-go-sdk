// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCloudAssetDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssetSubType(v int32) *GetCloudAssetDetailRequest
	GetAssetSubType() *int32
	SetAssetType(v int32) *GetCloudAssetDetailRequest
	GetAssetType() *int32
	SetCloudAssetInstances(v []*GetCloudAssetDetailRequestCloudAssetInstances) *GetCloudAssetDetailRequest
	GetCloudAssetInstances() []*GetCloudAssetDetailRequestCloudAssetInstances
	SetVendor(v int32) *GetCloudAssetDetailRequest
	GetVendor() *int32
}

type GetCloudAssetDetailRequest struct {
	// The subtype of the cloud service.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the cloud asset. Valid values:
	//
	// 	- **0**: Elastic Compute Service (ECS).
	//
	// 	- **1**: Server Load Balancer (SLB).
	//
	// 	- **3**: ApsaraDB RDS.
	//
	// 	- **4**: ApsaraDB for MongoDB.
	//
	// 	- **5**: ApsaraDB for Redis.
	//
	// 	- **6**: Container Registry.
	//
	// 	- **8**: Container Service for Kubernetes.
	//
	// 	- **9**: Virtual Private Cloud (VPC).
	//
	// 	- **11**: ActionTrail.
	//
	// 	- **12**: Alibaba Cloud CDN (CDN).
	//
	// 	- **13**: Certificate Management Service.
	//
	// 	- **14**: Apsara Devops.
	//
	// 	- **15**: Resource Access Management (RAM).
	//
	// 	- **16**: Anti-DDoS.
	//
	// 	- **17**: Web Application Firewall (WAF).
	//
	// 	- **18**: Object Storage Service (OSS).
	//
	// 	- **19**: PolarDB.
	//
	// 	- **20**: ApsaraDB RDS for PostgreSQL.
	//
	// 	- **21**: Microservices Engine (MSE).
	//
	// 	- **22**: File Storage NAS (NAS).
	//
	// 	- **23**: Data Security Center (DSC).
	//
	// 	- **24**: Elastic IP Address (EIP).
	//
	// 	- **25**: Identity as a Service (IDaaS)-Employee Identity and Access Management (EIAM).
	//
	// 	- **26**: PolarDB-X.
	//
	// 	- **27**: Elasticsearch.
	//
	// example:
	//
	// 14
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The details of the assets.
	CloudAssetInstances []*GetCloudAssetDetailRequestCloudAssetInstances `json:"CloudAssetInstances,omitempty" xml:"CloudAssetInstances,omitempty" type:"Repeated"`
	// The service provider of the cloud asset. Valid values:
	//
	// 	- **0**: Alibaba Cloud.
	//
	// 	- **1**: service provider that is unrecognized.
	//
	// 	- **2**: data center.
	//
	// 	- **3**, **4**, **5**, and **7**: third-party service provider.
	//
	// 	- **8**: simple application server.
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s GetCloudAssetDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetDetailRequest) GoString() string {
	return s.String()
}

func (s *GetCloudAssetDetailRequest) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *GetCloudAssetDetailRequest) GetAssetType() *int32 {
	return s.AssetType
}

func (s *GetCloudAssetDetailRequest) GetCloudAssetInstances() []*GetCloudAssetDetailRequestCloudAssetInstances {
	return s.CloudAssetInstances
}

func (s *GetCloudAssetDetailRequest) GetVendor() *int32 {
	return s.Vendor
}

func (s *GetCloudAssetDetailRequest) SetAssetSubType(v int32) *GetCloudAssetDetailRequest {
	s.AssetSubType = &v
	return s
}

func (s *GetCloudAssetDetailRequest) SetAssetType(v int32) *GetCloudAssetDetailRequest {
	s.AssetType = &v
	return s
}

func (s *GetCloudAssetDetailRequest) SetCloudAssetInstances(v []*GetCloudAssetDetailRequestCloudAssetInstances) *GetCloudAssetDetailRequest {
	s.CloudAssetInstances = v
	return s
}

func (s *GetCloudAssetDetailRequest) SetVendor(v int32) *GetCloudAssetDetailRequest {
	s.Vendor = &v
	return s
}

func (s *GetCloudAssetDetailRequest) Validate() error {
	if s.CloudAssetInstances != nil {
		for _, item := range s.CloudAssetInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetCloudAssetDetailRequestCloudAssetInstances struct {
	// The instance ID of the cloud asset.
	//
	// example:
	//
	// sg-wz9hf86vbzbrrde7****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region in which the cloud asset resides.
	//
	// > For more information about the mapping between region IDs and region names, see [Regions and zones](https://help.aliyun.com/document_detail/40654.html).
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetCloudAssetDetailRequestCloudAssetInstances) String() string {
	return dara.Prettify(s)
}

func (s GetCloudAssetDetailRequestCloudAssetInstances) GoString() string {
	return s.String()
}

func (s *GetCloudAssetDetailRequestCloudAssetInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCloudAssetDetailRequestCloudAssetInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCloudAssetDetailRequestCloudAssetInstances) SetInstanceId(v string) *GetCloudAssetDetailRequestCloudAssetInstances {
	s.InstanceId = &v
	return s
}

func (s *GetCloudAssetDetailRequestCloudAssetInstances) SetRegionId(v string) *GetCloudAssetDetailRequestCloudAssetInstances {
	s.RegionId = &v
	return s
}

func (s *GetCloudAssetDetailRequestCloudAssetInstances) Validate() error {
	return dara.Validate(s)
}
