// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudAssetInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCloudAssetQueryData(v []*ListCloudAssetInstancesRequestCloudAssetQueryData) *ListCloudAssetInstancesRequest
	GetCloudAssetQueryData() []*ListCloudAssetInstancesRequestCloudAssetQueryData
	SetCloudAssetTypes(v []*ListCloudAssetInstancesRequestCloudAssetTypes) *ListCloudAssetInstancesRequest
	GetCloudAssetTypes() []*ListCloudAssetInstancesRequestCloudAssetTypes
	SetCriteria(v string) *ListCloudAssetInstancesRequest
	GetCriteria() *string
	SetCurrentPage(v int32) *ListCloudAssetInstancesRequest
	GetCurrentPage() *int32
	SetIsSaleData(v bool) *ListCloudAssetInstancesRequest
	GetIsSaleData() *bool
	SetLogicalExp(v string) *ListCloudAssetInstancesRequest
	GetLogicalExp() *string
	SetPageSize(v int32) *ListCloudAssetInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *ListCloudAssetInstancesRequest
	GetRegionId() *string
	SetResourceDirectoryAccountId(v int64) *ListCloudAssetInstancesRequest
	GetResourceDirectoryAccountId() *int64
}

type ListCloudAssetInstancesRequest struct {
	// The data list queried by keyword.
	CloudAssetQueryData []*ListCloudAssetInstancesRequestCloudAssetQueryData `json:"CloudAssetQueryData,omitempty" xml:"CloudAssetQueryData,omitempty" type:"Repeated"`
	// The list of cloud asset instance types.
	CloudAssetTypes []*ListCloudAssetInstancesRequestCloudAssetTypes `json:"CloudAssetTypes,omitempty" xml:"CloudAssetTypes,omitempty" type:"Repeated"`
	// The search conditions for assets. This parameter is in JSON format and contains the following fields:
	//
	// - **name**: The search item.
	//
	// - **value**: The value of the search item.
	//
	// - **logicalExp**: The logical relationship between multiple search item values. Valid values:
	//
	//     - **OR**: The search item values are evaluated using the OR operator.
	//
	//     - **AND**: The search item values are evaluated using the AND operator.
	//
	// > You can call the [GetCloudAssetCriteria](~~GetCloudAssetCriteria~~) operation to query the supported search conditions.
	//
	// example:
	//
	// [{\\"name\\":\\"internetIp\\",\\"value\\":\\"192.168\\",\\"logicalExp\\":\\"OR\\"}]
	Criteria *string `json:"Criteria,omitempty" xml:"Criteria,omitempty"`
	// The page number of the current page in a paging query.
	//
	// example:
	//
	// 2
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// Specifies whether to return sale-related data. Valid values:
	//
	// - **true**: Returns sale-related data.
	//
	// - **false**: Does not return sale-related data.
	IsSaleData *bool `json:"IsSaleData,omitempty" xml:"IsSaleData,omitempty"`
	// The logical relationship between multiple search conditions. Valid values:
	//
	// - **OR**: The search conditions are evaluated using the OR operator.
	//
	// - **AND**: The search conditions are evaluated using the AND operator.
	//
	// example:
	//
	// OR
	LogicalExp *string `json:"LogicalExp,omitempty" xml:"LogicalExp,omitempty"`
	// The maximum number of entries per page. Maximum value: 100. Default value: 20.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the Alibaba Cloud account of the resource folder member accounts.
	//
	// > You can invoke the [DescribeMonitorAccounts](~~DescribeMonitorAccounts~~) operation to obtain this parameter.
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s ListCloudAssetInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListCloudAssetInstancesRequest) GetCloudAssetQueryData() []*ListCloudAssetInstancesRequestCloudAssetQueryData {
	return s.CloudAssetQueryData
}

func (s *ListCloudAssetInstancesRequest) GetCloudAssetTypes() []*ListCloudAssetInstancesRequestCloudAssetTypes {
	return s.CloudAssetTypes
}

func (s *ListCloudAssetInstancesRequest) GetCriteria() *string {
	return s.Criteria
}

func (s *ListCloudAssetInstancesRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListCloudAssetInstancesRequest) GetIsSaleData() *bool {
	return s.IsSaleData
}

func (s *ListCloudAssetInstancesRequest) GetLogicalExp() *string {
	return s.LogicalExp
}

func (s *ListCloudAssetInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListCloudAssetInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListCloudAssetInstancesRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *ListCloudAssetInstancesRequest) SetCloudAssetQueryData(v []*ListCloudAssetInstancesRequestCloudAssetQueryData) *ListCloudAssetInstancesRequest {
	s.CloudAssetQueryData = v
	return s
}

func (s *ListCloudAssetInstancesRequest) SetCloudAssetTypes(v []*ListCloudAssetInstancesRequestCloudAssetTypes) *ListCloudAssetInstancesRequest {
	s.CloudAssetTypes = v
	return s
}

func (s *ListCloudAssetInstancesRequest) SetCriteria(v string) *ListCloudAssetInstancesRequest {
	s.Criteria = &v
	return s
}

func (s *ListCloudAssetInstancesRequest) SetCurrentPage(v int32) *ListCloudAssetInstancesRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListCloudAssetInstancesRequest) SetIsSaleData(v bool) *ListCloudAssetInstancesRequest {
	s.IsSaleData = &v
	return s
}

func (s *ListCloudAssetInstancesRequest) SetLogicalExp(v string) *ListCloudAssetInstancesRequest {
	s.LogicalExp = &v
	return s
}

func (s *ListCloudAssetInstancesRequest) SetPageSize(v int32) *ListCloudAssetInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *ListCloudAssetInstancesRequest) SetRegionId(v string) *ListCloudAssetInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ListCloudAssetInstancesRequest) SetResourceDirectoryAccountId(v int64) *ListCloudAssetInstancesRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *ListCloudAssetInstancesRequest) Validate() error {
	if s.CloudAssetQueryData != nil {
		for _, item := range s.CloudAssetQueryData {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type ListCloudAssetInstancesRequestCloudAssetQueryData struct {
	// The query content.
	//
	// example:
	//
	// 163.8.8.9
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The query operator. Currently, only INCLUDE is supported.
	//
	// example:
	//
	// INCLUDE
	Operator *string `json:"Operator,omitempty" xml:"Operator,omitempty"`
}

func (s ListCloudAssetInstancesRequestCloudAssetQueryData) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetInstancesRequestCloudAssetQueryData) GoString() string {
	return s.String()
}

func (s *ListCloudAssetInstancesRequestCloudAssetQueryData) GetData() *string {
	return s.Data
}

func (s *ListCloudAssetInstancesRequestCloudAssetQueryData) GetOperator() *string {
	return s.Operator
}

func (s *ListCloudAssetInstancesRequestCloudAssetQueryData) SetData(v string) *ListCloudAssetInstancesRequestCloudAssetQueryData {
	s.Data = &v
	return s
}

func (s *ListCloudAssetInstancesRequestCloudAssetQueryData) SetOperator(v string) *ListCloudAssetInstancesRequestCloudAssetQueryData {
	s.Operator = &v
	return s
}

func (s *ListCloudAssetInstancesRequestCloudAssetQueryData) Validate() error {
	return dara.Validate(s)
}

type ListCloudAssetInstancesRequestCloudAssetTypes struct {
	// The subtype of the cloud service.
	//
	// > For specific meanings, refer to the AssetSubType parameter in the [GetCloudAssetCriteria](~~GetCloudAssetCriteria~~) operation.
	//
	// example:
	//
	// 0
	AssetSubType *int32 `json:"AssetSubType,omitempty" xml:"AssetSubType,omitempty"`
	// The type of the cloud asset.
	//
	// > For specific meanings, refer to the AssetType parameter in the [GetCloudAssetCriteria](~~GetCloudAssetCriteria~~) operation.
	//
	// example:
	//
	// 18
	AssetType *int32 `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
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

func (s ListCloudAssetInstancesRequestCloudAssetTypes) String() string {
	return dara.Prettify(s)
}

func (s ListCloudAssetInstancesRequestCloudAssetTypes) GoString() string {
	return s.String()
}

func (s *ListCloudAssetInstancesRequestCloudAssetTypes) GetAssetSubType() *int32 {
	return s.AssetSubType
}

func (s *ListCloudAssetInstancesRequestCloudAssetTypes) GetAssetType() *int32 {
	return s.AssetType
}

func (s *ListCloudAssetInstancesRequestCloudAssetTypes) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListCloudAssetInstancesRequestCloudAssetTypes) SetAssetSubType(v int32) *ListCloudAssetInstancesRequestCloudAssetTypes {
	s.AssetSubType = &v
	return s
}

func (s *ListCloudAssetInstancesRequestCloudAssetTypes) SetAssetType(v int32) *ListCloudAssetInstancesRequestCloudAssetTypes {
	s.AssetType = &v
	return s
}

func (s *ListCloudAssetInstancesRequestCloudAssetTypes) SetVendor(v int32) *ListCloudAssetInstancesRequestCloudAssetTypes {
	s.Vendor = &v
	return s
}

func (s *ListCloudAssetInstancesRequestCloudAssetTypes) Validate() error {
	return dara.Validate(s)
}
