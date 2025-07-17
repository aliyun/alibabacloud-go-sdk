// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactGroupName(v string) *DescribeContactGroupsRequest
	GetContactGroupName() *string
	SetGroupIds(v string) *DescribeContactGroupsRequest
	GetGroupIds() *string
	SetIsDetail(v bool) *DescribeContactGroupsRequest
	GetIsDetail() *bool
	SetPage(v int64) *DescribeContactGroupsRequest
	GetPage() *int64
	SetRegionId(v string) *DescribeContactGroupsRequest
	GetRegionId() *string
	SetSize(v int64) *DescribeContactGroupsRequest
	GetSize() *int64
}

type DescribeContactGroupsRequest struct {
	// The name of the alert contact group.
	//
	// example:
	//
	// TestGroup
	ContactGroupName *string `json:"ContactGroupName,omitempty" xml:"ContactGroupName,omitempty"`
	// The ID of the alert contact group.
	//
	// example:
	//
	// 12345
	GroupIds *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// Specifies whether to return all the alert contacts in the queried alert contact group. Valid values:
	//
	// 	- `false`
	//
	// 	- `true`
	//
	// example:
	//
	// true
	IsDetail *bool `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of alert contact groups displayed on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeContactGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsRequest) GetContactGroupName() *string {
	return s.ContactGroupName
}

func (s *DescribeContactGroupsRequest) GetGroupIds() *string {
	return s.GroupIds
}

func (s *DescribeContactGroupsRequest) GetIsDetail() *bool {
	return s.IsDetail
}

func (s *DescribeContactGroupsRequest) GetPage() *int64 {
	return s.Page
}

func (s *DescribeContactGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeContactGroupsRequest) GetSize() *int64 {
	return s.Size
}

func (s *DescribeContactGroupsRequest) SetContactGroupName(v string) *DescribeContactGroupsRequest {
	s.ContactGroupName = &v
	return s
}

func (s *DescribeContactGroupsRequest) SetGroupIds(v string) *DescribeContactGroupsRequest {
	s.GroupIds = &v
	return s
}

func (s *DescribeContactGroupsRequest) SetIsDetail(v bool) *DescribeContactGroupsRequest {
	s.IsDetail = &v
	return s
}

func (s *DescribeContactGroupsRequest) SetPage(v int64) *DescribeContactGroupsRequest {
	s.Page = &v
	return s
}

func (s *DescribeContactGroupsRequest) SetRegionId(v string) *DescribeContactGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeContactGroupsRequest) SetSize(v int64) *DescribeContactGroupsRequest {
	s.Size = &v
	return s
}

func (s *DescribeContactGroupsRequest) Validate() error {
	return dara.Validate(s)
}
