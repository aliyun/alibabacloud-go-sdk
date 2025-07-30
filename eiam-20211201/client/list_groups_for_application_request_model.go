// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForApplicationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListGroupsForApplicationRequest
	GetApplicationId() *string
	SetGroupIds(v []*string) *ListGroupsForApplicationRequest
	GetGroupIds() []*string
	SetInstanceId(v string) *ListGroupsForApplicationRequest
	GetInstanceId() *string
	SetPageNumber(v int64) *ListGroupsForApplicationRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListGroupsForApplicationRequest
	GetPageSize() *int64
}

type ListGroupsForApplicationRequest struct {
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// app_mkv7rgt4d7i4u7zqtzev2mxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The group IDs. You can specify up to 100 group IDs at a time.
	//
	// example:
	//
	// group_miu8e4t4d7i4u7uwezgr54xxxx
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListGroupsForApplicationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForApplicationRequest) GoString() string {
	return s.String()
}

func (s *ListGroupsForApplicationRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListGroupsForApplicationRequest) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *ListGroupsForApplicationRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsForApplicationRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListGroupsForApplicationRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListGroupsForApplicationRequest) SetApplicationId(v string) *ListGroupsForApplicationRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListGroupsForApplicationRequest) SetGroupIds(v []*string) *ListGroupsForApplicationRequest {
	s.GroupIds = v
	return s
}

func (s *ListGroupsForApplicationRequest) SetInstanceId(v string) *ListGroupsForApplicationRequest {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForApplicationRequest) SetPageNumber(v int64) *ListGroupsForApplicationRequest {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsForApplicationRequest) SetPageSize(v int64) *ListGroupsForApplicationRequest {
	s.PageSize = &v
	return s
}

func (s *ListGroupsForApplicationRequest) Validate() error {
	return dara.Validate(s)
}
