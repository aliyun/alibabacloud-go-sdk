// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppCenterImageId(v string) *ListAppInstanceGroupRequest
	GetAppCenterImageId() *string
	SetAppInstanceGroupId(v string) *ListAppInstanceGroupRequest
	GetAppInstanceGroupId() *string
	SetAppInstanceGroupName(v string) *ListAppInstanceGroupRequest
	GetAppInstanceGroupName() *string
	SetBizRegionId(v string) *ListAppInstanceGroupRequest
	GetBizRegionId() *string
	SetNodeInstanceType(v string) *ListAppInstanceGroupRequest
	GetNodeInstanceType() *string
	SetOfficeSiteId(v string) *ListAppInstanceGroupRequest
	GetOfficeSiteId() *string
	SetPageNumber(v int32) *ListAppInstanceGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAppInstanceGroupRequest
	GetPageSize() *int32
	SetProductType(v string) *ListAppInstanceGroupRequest
	GetProductType() *string
	SetRegionId(v string) *ListAppInstanceGroupRequest
	GetRegionId() *string
	SetStatus(v []*string) *ListAppInstanceGroupRequest
	GetStatus() []*string
	SetTag(v []*ListAppInstanceGroupRequestTag) *ListAppInstanceGroupRequest
	GetTag() []*ListAppInstanceGroupRequestTag
}

type ListAppInstanceGroupRequest struct {
	// The image ID of the app. You can obtain the ID from the Images page in the App Streaming console.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The ID of the delivery group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the delivery groups to query. Fuzzy match is used for queries. For example, if you set this parameter to `Office App`, all delivery groups whose names contain `Office App` are queried, such as `My Office Apps` and `Office App A`.
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The ID of the region where the delivery group resides. For information about the supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// Valid values:
	//
	// 	- cn-shanghai: China (Shanghai)
	//
	// 	- cn-hangzhou: China (Hangzhou)
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The ID of the resource specification that you purchase. You can call the [ListNodeInstanceType](~~ListNodeInstanceType~~) operation to obtain the ID.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// example:
	//
	// cn-hongkong+dir-643067****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. The value cannot be greater than `100`.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// Valid value:
	//
	// 	- CloudApp: App Streaming
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Deprecated
	//
	// The region ID
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the delivery groups.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// if can be null:
	// true
	Tag []*ListAppInstanceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListAppInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupRequest) GetAppCenterImageId() *string {
	return s.AppCenterImageId
}

func (s *ListAppInstanceGroupRequest) GetAppInstanceGroupId() *string {
	return s.AppInstanceGroupId
}

func (s *ListAppInstanceGroupRequest) GetAppInstanceGroupName() *string {
	return s.AppInstanceGroupName
}

func (s *ListAppInstanceGroupRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListAppInstanceGroupRequest) GetNodeInstanceType() *string {
	return s.NodeInstanceType
}

func (s *ListAppInstanceGroupRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListAppInstanceGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAppInstanceGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAppInstanceGroupRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ListAppInstanceGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAppInstanceGroupRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListAppInstanceGroupRequest) GetTag() []*ListAppInstanceGroupRequestTag {
	return s.Tag
}

func (s *ListAppInstanceGroupRequest) SetAppCenterImageId(v string) *ListAppInstanceGroupRequest {
	s.AppCenterImageId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetAppInstanceGroupId(v string) *ListAppInstanceGroupRequest {
	s.AppInstanceGroupId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetAppInstanceGroupName(v string) *ListAppInstanceGroupRequest {
	s.AppInstanceGroupName = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetBizRegionId(v string) *ListAppInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetNodeInstanceType(v string) *ListAppInstanceGroupRequest {
	s.NodeInstanceType = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetOfficeSiteId(v string) *ListAppInstanceGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetPageNumber(v int32) *ListAppInstanceGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetPageSize(v int32) *ListAppInstanceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetProductType(v string) *ListAppInstanceGroupRequest {
	s.ProductType = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetRegionId(v string) *ListAppInstanceGroupRequest {
	s.RegionId = &v
	return s
}

func (s *ListAppInstanceGroupRequest) SetStatus(v []*string) *ListAppInstanceGroupRequest {
	s.Status = v
	return s
}

func (s *ListAppInstanceGroupRequest) SetTag(v []*ListAppInstanceGroupRequestTag) *ListAppInstanceGroupRequest {
	s.Tag = v
	return s
}

func (s *ListAppInstanceGroupRequest) Validate() error {
	return dara.Validate(s)
}

type ListAppInstanceGroupRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAppInstanceGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListAppInstanceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *ListAppInstanceGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListAppInstanceGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListAppInstanceGroupRequestTag) SetKey(v string) *ListAppInstanceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *ListAppInstanceGroupRequestTag) SetValue(v string) *ListAppInstanceGroupRequestTag {
	s.Value = &v
	return s
}

func (s *ListAppInstanceGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
