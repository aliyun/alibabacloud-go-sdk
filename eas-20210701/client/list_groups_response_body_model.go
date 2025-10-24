// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*Group) *ListGroupsResponseBody
	GetGroups() []*Group
	SetPageNumber(v int64) *ListGroupsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListGroupsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListGroupsResponseBody
	GetTotalCount() *int64
}

type ListGroupsResponseBody struct {
	// The service groups.
	Groups []*Group `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
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
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsResponseBody) GetGroups() []*Group {
	return s.Groups
}

func (s *ListGroupsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListGroupsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGroupsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupsResponseBody) SetGroups(v []*Group) *ListGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListGroupsResponseBody) SetPageNumber(v int64) *ListGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListGroupsResponseBody) SetPageSize(v int64) *ListGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListGroupsResponseBody) SetRequestId(v string) *ListGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGroupsResponseBody) SetTotalCount(v int64) *ListGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsResponseBody) Validate() error {
	if s.Groups != nil {
		for _, item := range s.Groups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
