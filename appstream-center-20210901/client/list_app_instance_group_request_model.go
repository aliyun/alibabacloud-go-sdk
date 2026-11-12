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
	SetExcludedUserGroupIds(v []*string) *ListAppInstanceGroupRequest
	GetExcludedUserGroupIds() []*string
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
	SetUserGroupIds(v []*string) *ListAppInstanceGroupRequest
	GetUserGroupIds() []*string
}

type ListAppInstanceGroupRequest struct {
	// The ID of the app image. You can obtain the ID from the Image Management page in the WUYING Cloud App console.
	//
	// example:
	//
	// img-8z4nztpaqvay4****
	AppCenterImageId *string `json:"AppCenterImageId,omitempty" xml:"AppCenterImageId,omitempty"`
	// The ID of the app instance group.
	//
	// example:
	//
	// aig-9ciijz60n4xsv****
	AppInstanceGroupId *string `json:"AppInstanceGroupId,omitempty" xml:"AppInstanceGroupId,omitempty"`
	// The name of the app instance group. Fuzzy match is supported. For example, if you specify `Office App`, app instance groups whose names contain `Office App` are queried, such as `My Office App` and `Office App A`.
	//
	// example:
	//
	// Office App
	AppInstanceGroupName *string `json:"AppInstanceGroupName,omitempty" xml:"AppInstanceGroupName,omitempty"`
	// The ID of the region to which the app instance group belongs. For more information about supported regions, see [Limits](https://help.aliyun.com/document_detail/426036.html).
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The list of excluded authorization user group IDs.
	//
	// if can be null:
	// true
	ExcludedUserGroupIds []*string `json:"ExcludedUserGroupIds,omitempty" xml:"ExcludedUserGroupIds,omitempty" type:"Repeated"`
	// The ID of the instance type for the purchased resources. You can obtain the ID by calling the [ListNodeInstanceType](~~ListNodeInstanceType~~) operation.
	//
	// example:
	//
	// appstreaming.vgpu.4c8g.2g
	NodeInstanceType *string `json:"NodeInstanceType,omitempty" xml:"NodeInstanceType,omitempty"`
	// The ID of the office network.
	//
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
	// The number of entries to return on each page. The value cannot exceed `100`.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// CloudApp
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// Deprecated
	//
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The statuses of the app instance groups.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The list of tags.
	//
	// if can be null:
	// true
	Tag []*ListAppInstanceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The list of authorized user group IDs.
	//
	// if can be null:
	// true
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
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

func (s *ListAppInstanceGroupRequest) GetExcludedUserGroupIds() []*string {
	return s.ExcludedUserGroupIds
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

func (s *ListAppInstanceGroupRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
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

func (s *ListAppInstanceGroupRequest) SetExcludedUserGroupIds(v []*string) *ListAppInstanceGroupRequest {
	s.ExcludedUserGroupIds = v
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

func (s *ListAppInstanceGroupRequest) SetUserGroupIds(v []*string) *ListAppInstanceGroupRequest {
	s.UserGroupIds = v
	return s
}

func (s *ListAppInstanceGroupRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAppInstanceGroupRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// department
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// design
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
