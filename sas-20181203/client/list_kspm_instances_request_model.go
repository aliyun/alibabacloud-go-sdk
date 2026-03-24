// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKspmInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAssetTypes(v []*ListKspmInstancesRequestCloudAssetTypes) *ListKspmInstancesRequest
	GetCloudAssetTypes() []*ListKspmInstancesRequestCloudAssetTypes
	SetCriteria(v string) *ListKspmInstancesRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *ListKspmInstancesRequest
	GetCurrentPage() *int32
	SetLogicalExp(v string) *ListKspmInstancesRequest
	GetLogicalExp() *string
	SetPageSize(v int32) *ListKspmInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListKspmInstancesRequest
	GetRegionId() *string
}

type ListKspmInstancesRequest struct {
	// List of Kubernetes asset type information.
	CloudAssetTypes []*ListKspmInstancesRequestCloudAssetTypes `json:"CloudAssetTypes,omitempty" xml:"CloudAssetTypes,omitempty" type:"Repeated"`
	// Set the conditions for searching assets. This parameter is in JSON format, and case sensitivity should be noted when entering parameters. It includes the following fields:
	//
	// - **name**: Search item
	//
	// - **value**: Value of the search item
	//
	// - **logicalExp**: The relationship between search items when multiple conditions are used. Values:
	//
	//     - **OR**: Indicates an **or*	- relationship between multiple conditions.
	//
	//     - **AND**: Indicates an **and*	- relationship between multiple conditions.
	//
	// > Supports using region, instance name, instance ID, alarm status, risk status, and tag search conditions.
	//
	// example:
	//
	// [{"name":"vulStatus","value":"YES","logicalExp":"AND"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// Set which page to start displaying the query results from. The default value is 1, indicating that the display starts from the first page.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Set the logical relationship that takes effect between multiple search conditions. Values:
	//
	// - **OR**: Indicates an **or*	- relationship between multiple search conditions.
	//
	// - **AND**: Indicates an **and*	- relationship between multiple search conditions.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The maximum number of data entries displayed per page during pagination. The default value is 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the region where the instance is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListKspmInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListKspmInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListKspmInstancesRequest) GetCloudAssetTypes() []*ListKspmInstancesRequestCloudAssetTypes {
	return s.CloudAssetTypes
}

func (s *ListKspmInstancesRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *ListKspmInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListKspmInstancesRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *ListKspmInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKspmInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListKspmInstancesRequest) SetCloudAssetTypes(v []*ListKspmInstancesRequestCloudAssetTypes) *ListKspmInstancesRequest {
	s.CloudAssetTypes = v
	return s
}

func (s *ListKspmInstancesRequest) SetCriteria(v string) *ListKspmInstancesRequest {
	s.Criteria = &v
	return s
}

func (s *ListKspmInstancesRequest) SetCurrentPage(v int32) *ListKspmInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListKspmInstancesRequest) SetLogicalExp(v string) *ListKspmInstancesRequest {
	s.LogicalExp = &v
	return s
}

func (s *ListKspmInstancesRequest) SetPageSize(v int32) *ListKspmInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListKspmInstancesRequest) SetRegionId(v string) *ListKspmInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListKspmInstancesRequest) Validate() error {
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

type ListKspmInstancesRequestCloudAssetTypes struct {
	// Subtype of the product.
	//
	// Asset type-subtype. Values:
	//
	// - **0**: Workload
	//
	//     	- **0**: Pod
	//
	//     	- **1**: DaemonSet
	//
	//     	- **2**: StatefulSet
	//
	// - **1**: Service
	//
	//     	- **0**: Service
	//
	// - **2**: Namespace
	//
	//     	- **0**: Namespace
	//
	// - **3**: Authorization
	//
	//     	- **0**: Role
	//
	//     	- **1**: ClusterRole
	//
	//     	- **2**: ClusterRoleBinding
	//
	//     	- **3**: RoleBinding
	//
	//     	- **4**: ServiceAccount
	//
	// - **4**: Storage
	//
	//     	- **0**: PersistentVolume
	//
	//     	- **1**: PersistentVolumeClaim
	//
	//     	- **2**: StorageClass
	//
	// - **5**: Container
	//
	//     	- **0**: Image
	//
	// - **6**: Network
	//
	//     	- **0**: Route
	//
	//     	- **0**: Ingress
	//
	// - **7**: Configuration
	//
	//     	- **0**: ConfigMap
	//
	// - **8**: Policies
	//
	//     	- **0**: LimitRanges
	//
	//     	- **1**: ResourceQuota
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// Type of the asset. Values:
	//
	// - **0**: Workload
	//
	// - **1**: Service
	//
	// - **2**: Namespace
	//
	// - **3**: Authorization
	//
	// - **4**: Storage
	//
	// - **5**: Container
	//
	// - **6**: Network
	//
	// - **7**: Configuration
	//
	// - **8**: Policies
	//
	// example:
	//
	// 0
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// Vendor of the asset. The fixed value is **200**.
	//
	// example:
	//
	// 200
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListKspmInstancesRequestCloudAssetTypes) String() string {
	return dara.Prettify(s)
}

func (s ListKspmInstancesRequestCloudAssetTypes) GoString() string {
	return s.String()
}

func (s *ListKspmInstancesRequestCloudAssetTypes) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *ListKspmInstancesRequestCloudAssetTypes) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListKspmInstancesRequestCloudAssetTypes) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListKspmInstancesRequestCloudAssetTypes) SetAssetSubType(v int32) *ListKspmInstancesRequestCloudAssetTypes {
	s.AssetSubType = &v
	return s
}

func (s *ListKspmInstancesRequestCloudAssetTypes) SetAssetType(v int32) *ListKspmInstancesRequestCloudAssetTypes {
	s.AssetType = &v
	return s
}

func (s *ListKspmInstancesRequestCloudAssetTypes) SetVendor(v int32) *ListKspmInstancesRequestCloudAssetTypes {
	s.Vendor = &v
	return s
}

func (s *ListKspmInstancesRequestCloudAssetTypes) Validate() error {
	return dara.Validate(s)
}
