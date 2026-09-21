// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBrowserInstanceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppInstanceGroupSetId(v string) *ListBrowserInstanceGroupRequest
	GetAppInstanceGroupSetId() *string
	SetBizRegionId(v string) *ListBrowserInstanceGroupRequest
	GetBizRegionId() *string
	SetBrowserInstanceGroupId(v string) *ListBrowserInstanceGroupRequest
	GetBrowserInstanceGroupId() *string
	SetBrowserInstanceGroupName(v string) *ListBrowserInstanceGroupRequest
	GetBrowserInstanceGroupName() *string
	SetCloudBrowserName(v string) *ListBrowserInstanceGroupRequest
	GetCloudBrowserName() *string
	SetExcludedUserGroupIds(v []*string) *ListBrowserInstanceGroupRequest
	GetExcludedUserGroupIds() []*string
	SetOfficeSiteId(v string) *ListBrowserInstanceGroupRequest
	GetOfficeSiteId() *string
	SetPageNumber(v int32) *ListBrowserInstanceGroupRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListBrowserInstanceGroupRequest
	GetPageSize() *int32
	SetStatus(v []*string) *ListBrowserInstanceGroupRequest
	GetStatus() []*string
	SetTag(v []*ListBrowserInstanceGroupRequestTag) *ListBrowserInstanceGroupRequest
	GetTag() []*ListBrowserInstanceGroupRequestTag
	SetTier(v string) *ListBrowserInstanceGroupRequest
	GetTier() *string
	SetUserGroupIds(v []*string) *ListBrowserInstanceGroupRequest
	GetUserGroupIds() []*string
}

type ListBrowserInstanceGroupRequest struct {
	// The browser group set ID for exact match queries of active members in the set.
	//
	// example:
	//
	// set-3jm9d0abc00example
	AppInstanceGroupSetId *string `json:"AppInstanceGroupSetId,omitempty" xml:"AppInstanceGroupSetId,omitempty"`
	// Filters browser groups by business region.
	//
	// example:
	//
	// cn-hangzhou
	BizRegionId *string `json:"BizRegionId,omitempty" xml:"BizRegionId,omitempty"`
	// The cloud browser group ID for exact match queries.
	//
	// example:
	//
	// big-0c7loey7fzjq****
	BrowserInstanceGroupId *string `json:"BrowserInstanceGroupId,omitempty" xml:"BrowserInstanceGroupId,omitempty"`
	// The browser group name. Fuzzy match is supported.
	//
	// example:
	//
	// Office
	BrowserInstanceGroupName *string `json:"BrowserInstanceGroupName,omitempty" xml:"BrowserInstanceGroupName,omitempty"`
	// Performs a contains match by browser group name or ID.
	//
	// example:
	//
	// Office
	CloudBrowserName *string `json:"CloudBrowserName,omitempty" xml:"CloudBrowserName,omitempty"`
	// Excludes browser groups that are authorized to the specified user groups.
	//
	// if can be null:
	// true
	ExcludedUserGroupIds []*string `json:"ExcludedUserGroupIds,omitempty" xml:"ExcludedUserGroupIds,omitempty" type:"Repeated"`
	// Filters browser groups by office network ID.
	//
	// example:
	//
	// cn-hangzhou+dir-843734****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The page number, starting from page 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Filters by browser group status.
	//
	// **Limit:*	- Only one status value can be specified at a time.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The tag filter parameters. This parameter is not supported in customer-facing scenarios. Do not specify this parameter.
	Tag []*ListBrowserInstanceGroupRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Filters by version of the browser.
	//
	// - `Basic`: Basic Edition.
	//
	// - `Pro`: Premium Edition.
	//
	// Use `Pro` to query MAU browser groups.
	//
	// example:
	//
	// Pro
	Tier *string `json:"Tier,omitempty" xml:"Tier,omitempty"`
	// Filters by authorized user group IDs.
	//
	// if can be null:
	// true
	UserGroupIds []*string `json:"UserGroupIds,omitempty" xml:"UserGroupIds,omitempty" type:"Repeated"`
}

func (s ListBrowserInstanceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupRequest) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupRequest) GetAppInstanceGroupSetId() *string {
	return s.AppInstanceGroupSetId
}

func (s *ListBrowserInstanceGroupRequest) GetBizRegionId() *string {
	return s.BizRegionId
}

func (s *ListBrowserInstanceGroupRequest) GetBrowserInstanceGroupId() *string {
	return s.BrowserInstanceGroupId
}

func (s *ListBrowserInstanceGroupRequest) GetBrowserInstanceGroupName() *string {
	return s.BrowserInstanceGroupName
}

func (s *ListBrowserInstanceGroupRequest) GetCloudBrowserName() *string {
	return s.CloudBrowserName
}

func (s *ListBrowserInstanceGroupRequest) GetExcludedUserGroupIds() []*string {
	return s.ExcludedUserGroupIds
}

func (s *ListBrowserInstanceGroupRequest) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *ListBrowserInstanceGroupRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListBrowserInstanceGroupRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListBrowserInstanceGroupRequest) GetStatus() []*string {
	return s.Status
}

func (s *ListBrowserInstanceGroupRequest) GetTag() []*ListBrowserInstanceGroupRequestTag {
	return s.Tag
}

func (s *ListBrowserInstanceGroupRequest) GetTier() *string {
	return s.Tier
}

func (s *ListBrowserInstanceGroupRequest) GetUserGroupIds() []*string {
	return s.UserGroupIds
}

func (s *ListBrowserInstanceGroupRequest) SetAppInstanceGroupSetId(v string) *ListBrowserInstanceGroupRequest {
	s.AppInstanceGroupSetId = &v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetBizRegionId(v string) *ListBrowserInstanceGroupRequest {
	s.BizRegionId = &v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetBrowserInstanceGroupId(v string) *ListBrowserInstanceGroupRequest {
	s.BrowserInstanceGroupId = &v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetBrowserInstanceGroupName(v string) *ListBrowserInstanceGroupRequest {
	s.BrowserInstanceGroupName = &v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetCloudBrowserName(v string) *ListBrowserInstanceGroupRequest {
	s.CloudBrowserName = &v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetExcludedUserGroupIds(v []*string) *ListBrowserInstanceGroupRequest {
	s.ExcludedUserGroupIds = v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetOfficeSiteId(v string) *ListBrowserInstanceGroupRequest {
	s.OfficeSiteId = &v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetPageNumber(v int32) *ListBrowserInstanceGroupRequest {
	s.PageNumber = &v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetPageSize(v int32) *ListBrowserInstanceGroupRequest {
	s.PageSize = &v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetStatus(v []*string) *ListBrowserInstanceGroupRequest {
	s.Status = v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetTag(v []*ListBrowserInstanceGroupRequestTag) *ListBrowserInstanceGroupRequest {
	s.Tag = v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetTier(v string) *ListBrowserInstanceGroupRequest {
	s.Tier = &v
	return s
}

func (s *ListBrowserInstanceGroupRequest) SetUserGroupIds(v []*string) *ListBrowserInstanceGroupRequest {
	s.UserGroupIds = v
	return s
}

func (s *ListBrowserInstanceGroupRequest) Validate() error {
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

type ListBrowserInstanceGroupRequestTag struct {
	// The tag key. Do not specify this parameter in customer-facing scenarios.
	//
	// example:
	//
	// -
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value. Do not specify this parameter in customer-facing scenarios.
	//
	// example:
	//
	// -
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListBrowserInstanceGroupRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListBrowserInstanceGroupRequestTag) GoString() string {
	return s.String()
}

func (s *ListBrowserInstanceGroupRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListBrowserInstanceGroupRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListBrowserInstanceGroupRequestTag) SetKey(v string) *ListBrowserInstanceGroupRequestTag {
	s.Key = &v
	return s
}

func (s *ListBrowserInstanceGroupRequestTag) SetValue(v string) *ListBrowserInstanceGroupRequestTag {
	s.Value = &v
	return s
}

func (s *ListBrowserInstanceGroupRequestTag) Validate() error {
	return dara.Validate(s)
}
