// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAssetSchemasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAssetSchemas(v []*ListCloudAssetSchemasResponseBodyCloudAssetSchemas) *ListCloudAssetSchemasResponseBody
	GetCloudAssetSchemas() []*ListCloudAssetSchemasResponseBodyCloudAssetSchemas
	SetPageInfo(v *ListCloudAssetSchemasResponseBodyPageInfo) *ListCloudAssetSchemasResponseBody
	GetPageInfo() *ListCloudAssetSchemasResponseBodyPageInfo
	SetRequestId(v string) *ListCloudAssetSchemasResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListCloudAssetSchemasResponseBody
	GetSuccess() *bool
}

type ListCloudAssetSchemasResponseBody struct {
	// List of asset structure definitions
	CloudAssetSchemas []*ListCloudAssetSchemasResponseBodyCloudAssetSchemas `json:"CloudAssetSchemas,omitempty" xml:"CloudAssetSchemas,omitempty" type:"Repeated"`
	// The pagination information.
	PageInfo *ListCloudAssetSchemasResponseBodyPageInfo `json:"PageInfo,omitempty" xml:"PageInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// F8B6F758-BCD4-597A-8A2C-DA5A552C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListCloudAssetSchemasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetSchemasResponseBody) GoString() string {
	return s.String()
}

func (s *ListCloudAssetSchemasResponseBody) GetCloudAssetSchemas() []*ListCloudAssetSchemasResponseBodyCloudAssetSchemas {
	return s.CloudAssetSchemas
}

func (s *ListCloudAssetSchemasResponseBody) GetPageInfo() *ListCloudAssetSchemasResponseBodyPageInfo {
	return s.PageInfo
}

func (s *ListCloudAssetSchemasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListCloudAssetSchemasResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListCloudAssetSchemasResponseBody) SetCloudAssetSchemas(v []*ListCloudAssetSchemasResponseBodyCloudAssetSchemas) *ListCloudAssetSchemasResponseBody {
	s.CloudAssetSchemas = v
	return s
}

func (s *ListCloudAssetSchemasResponseBody) SetPageInfo(v *ListCloudAssetSchemasResponseBodyPageInfo) *ListCloudAssetSchemasResponseBody {
	s.PageInfo = v
	return s
}

func (s *ListCloudAssetSchemasResponseBody) SetRequestId(v string) *ListCloudAssetSchemasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBody) SetSuccess(v bool) *ListCloudAssetSchemasResponseBody {
	s.Success = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBody) Validate() error {
	if s.CloudAssetSchemas != nil {
		for _, item := range s.CloudAssetSchemas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.PageInfo != nil {
		if err := s.PageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListCloudAssetSchemasResponseBodyCloudAssetSchemas struct {
	// Subtype of the cloud product.
	//
	// example:
	//
	// 1
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
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Asset structure definition name
	//
	// example:
	//
	// ACS_ECS_Disk
	DataName *string `json:"DataName,omitempty" xml:"DataName,omitempty"`
	// Current asset structure definition text.
	//
	// example:
	//
	// [{\\"associatedData\\":[{\\"assetSubType\\":100,\\"assetType\\":0,\\"dataName\\":\\"ACS_ECS_Instance\\",\\"properties\\":[{\\"name\\":\\"InstanceId\\",\\"path\\":\\"InstanceId\\"}],\\"vendor\\":0}],\\"description\\":\\"The ID of the instance to which the disk is attached.\\",\\"example\\":\\"i-bp67acfmxazb4q****\\",\\"name\\":\\"InstanceId\\",\\"type\\":\\"STRING\\",\\"withAssociatedData\\":true},{\\"description\\":\\"Disk name\\",\\"example\\":\\"testDiskName\\",\\"name\\":\\"DiskName\\",\\"type\\":\\"STRING\\"},{\\"description\\":\\"Only encrypted disks\\",\\"example\\":\\"false\\",\\"name\\":\\"Encrypted\\",\\"type\\":\\"BOOLEAN\\"},{\\"description\\":\\"Disk status\\",\\"example\\":\\"In_use\\",\\"name\\":\\"Status\\",\\"type\\":\\"STRING\\"},{\\"description\\":\\"Disk category\\",\\"example\\":\\"cloud_ssd\\",\\"name\\":\\"Category\\",\\"type\\":\\"STRING\\"},{\\"description\\":\\"Disk type\\",\\"example\\":\\"system\\",\\"name\\":\\"Type\\",\\"type\\":\\"STRING\\"},{\\"description\\":\\"Specifies whether to set an automatic snapshot policy for the disk.\\",\\"example\\":\\"false\\",\\"name\\":\\"EnableAutomatedSnapshotPolicy\\",\\"type\\":\\"BOOLEAN\\"},{\\"description\\":\\"The ID of the automatic snapshot policy.\\",\\"example\\":\\"sp-bp67acfmxazb4p****\\",\\"name\\":\\"AutoSnapshotPolicyId\\",\\"type\\":\\"STRING\\"},{\\"description\\":\\"Disk, local disk, or elastic ephemeral disk ID\\",\\"example\\":\\"d-bp18um4r4f2fve24****\\",\\"name\\":\\"DiskId\\",\\"type\\":\\"STRING\\"}]
	Properties *string `json:"Properties,omitempty" xml:"Properties,omitempty"`
	// The source of the server. Valid values:
	//
	// 	- **0**: an asset provided by Alibaba Cloud.
	//
	// 	- **1**: a third-party cloud server
	//
	// 	- **2**: a server in a data center
	//
	// 	- **3**, **4**, **5**, and **7**: other cloud asset
	//
	// 	- **8**: a lightweight asset
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListCloudAssetSchemasResponseBodyCloudAssetSchemas) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetSchemasResponseBodyCloudAssetSchemas) GoString() string {
	return s.String()
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) GetDataName() *string {
	return s.DataName
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) GetProperties() *string {
	return s.Properties
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) SetAssetSubType(v int32) *ListCloudAssetSchemasResponseBodyCloudAssetSchemas {
	s.AssetSubType = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) SetAssetType(v int32) *ListCloudAssetSchemasResponseBodyCloudAssetSchemas {
	s.AssetType = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) SetDataName(v string) *ListCloudAssetSchemasResponseBodyCloudAssetSchemas {
	s.DataName = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) SetProperties(v string) *ListCloudAssetSchemasResponseBodyCloudAssetSchemas {
	s.Properties = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) SetVendor(v int32) *ListCloudAssetSchemasResponseBodyCloudAssetSchemas {
	s.Vendor = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBodyCloudAssetSchemas) Validate() error {
	return dara.Validate(s)
}

type ListCloudAssetSchemasResponseBodyPageInfo struct {
	// The number of entries returned on the current page.
	//
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 54
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListCloudAssetSchemasResponseBodyPageInfo) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetSchemasResponseBodyPageInfo) GoString() string {
	return s.String()
}

func (s *ListCloudAssetSchemasResponseBodyPageInfo) GetCount() *int32 {
	return s.Count
}

func (s *ListCloudAssetSchemasResponseBodyPageInfo) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudAssetSchemasResponseBodyPageInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudAssetSchemasResponseBodyPageInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListCloudAssetSchemasResponseBodyPageInfo) SetCount(v int32) *ListCloudAssetSchemasResponseBodyPageInfo {
	s.Count = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBodyPageInfo) SetCurrentPage(v int32) *ListCloudAssetSchemasResponseBodyPageInfo {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBodyPageInfo) SetPageSize(v int32) *ListCloudAssetSchemasResponseBodyPageInfo {
	s.PageSize = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBodyPageInfo) SetTotalCount(v int32) *ListCloudAssetSchemasResponseBodyPageInfo {
	s.TotalCount = &v
	return s
}

func (s *ListCloudAssetSchemasResponseBodyPageInfo) Validate() error {
	return dara.Validate(s)
}
