// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupExternalId(v string) *ListGroupsRequest
	GetGroupExternalId() *string
	SetGroupIds(v []*string) *ListGroupsRequest
	GetGroupIds() []*string
	SetGroupName(v string) *ListGroupsRequest
	GetGroupName() *string
	SetGroupNameStartsWith(v string) *ListGroupsRequest
	GetGroupNameStartsWith() *string
	SetInstanceId(v string) *ListGroupsRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListGroupsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListGroupsRequest
	GetPageSize() *int64
}

type ListGroupsRequest struct {
	// The external ID of the group.
	//
	// example:
	//
	// group_external_id
	GroupExternalId *string `json:"GroupExternalId,omitempty" xml:"GroupExternalId,omitempty"`
	// The group IDs.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The name of the group. If you specify this parameter, the query is based on an exact match.
	//
	// example:
	//
	// name_test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The prefix of the group name. If you specify this parameter, the query follows the leftmost matching principle.
	//
	// example:
	//
	// name
	GroupNameStartsWith *string `json:"GroupNameStartsWith,omitempty" xml:"GroupNameStartsWith,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of the page to return.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsRequest) GetGroupExternalId() *string {
	return s.GroupExternalId
}

func (s *ListGroupsRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *ListGroupsRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *ListGroupsRequest) GetGroupNameStartsWith() *string {
	return s.GroupNameStartsWith
}

func (s *ListGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListGroupsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListGroupsRequest) SetGroupExternalId(v string) *ListGroupsRequest {
	s.GroupExternalId = &v
	return s
}

func (s *ListGroupsRequest) SetGroupIds(v []*string) *ListGroupsRequest {
	s.GroupIds = v
	return s
}

func (s *ListGroupsRequest) SetGroupName(v string) *ListGroupsRequest {
	s.GroupName = &v
	return s
}

func (s *ListGroupsRequest) SetGroupNameStartsWith(v string) *ListGroupsRequest {
	s.GroupNameStartsWith = &v
	return s
}

func (s *ListGroupsRequest) SetInstanceId(v string) *ListGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsRequest) SetPageNumber(v int64) *ListGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsRequest) SetPageSize(v int64) *ListGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsRequest) Validate() error {
	return dara.Validate(s)
}
