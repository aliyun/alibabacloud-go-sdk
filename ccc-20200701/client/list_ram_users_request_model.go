// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRamUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListRamUsersRequest
	GetInstanceId() *string
	SetPageNumber(v int32) *ListRamUsersRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRamUsersRequest
	GetPageSize() *int32
	SetSearchPattern(v string) *ListRamUsersRequest
	GetSearchPattern() *string
}

type ListRamUsersRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the results to return. Valid values: 1 to 1,000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Valid values: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The keyword to use for a fuzzy search based on the RAM user\\"s logon name or display name. This parameter is optional. If you leave this parameter empty, no filtering is applied.
	//
	// example:
	//
	// agent
	SearchPattern *string `json:"SearchPattern,omitempty" xml:"SearchPattern,omitempty"`
}

func (s ListRamUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListRamUsersRequest) GoString() string {
	return s.String()
}

func (s *ListRamUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListRamUsersRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRamUsersRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRamUsersRequest) GetSearchPattern() *string {
	return s.SearchPattern
}

func (s *ListRamUsersRequest) SetInstanceId(v string) *ListRamUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *ListRamUsersRequest) SetPageNumber(v int32) *ListRamUsersRequest {
	s.PageNumber = &v
	return s
}

func (s *ListRamUsersRequest) SetPageSize(v int32) *ListRamUsersRequest {
	s.PageSize = &v
	return s
}

func (s *ListRamUsersRequest) SetSearchPattern(v string) *ListRamUsersRequest {
	s.SearchPattern = &v
	return s
}

func (s *ListRamUsersRequest) Validate() error {
	return dara.Validate(s)
}
