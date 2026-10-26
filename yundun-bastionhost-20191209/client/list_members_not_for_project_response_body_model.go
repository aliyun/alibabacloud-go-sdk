// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMembersNotForProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListMembersNotForProjectResponseBody
	GetMaxResults() *int32
	SetMembers(v []*ListMembersNotForProjectResponseBodyMembers) *ListMembersNotForProjectResponseBody
	GetMembers() []*ListMembersNotForProjectResponseBodyMembers
	SetNextToken(v string) *ListMembersNotForProjectResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMembersNotForProjectResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListMembersNotForProjectResponseBody
	GetTotalCount() *int32
}

type ListMembersNotForProjectResponseBody struct {
	// The maximum number of entries to return. If a valid value is specified for the MaxResults request parameter, the value of the parameter is returned. Otherwise, the default value is returned.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The list of member accounts.
	Members []*ListMembersNotForProjectResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// The token for the next request, which is used for paging. Specify this token in the NextToken request parameter for the next paged query. If this field is empty in the response, no more data is available.
	//
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMembersNotForProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMembersNotForProjectResponseBody) GoString() string {
	return s.String()
}

func (s *ListMembersNotForProjectResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListMembersNotForProjectResponseBody) GetMembers() []*ListMembersNotForProjectResponseBodyMembers {
	return s.Members
}

func (s *ListMembersNotForProjectResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMembersNotForProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMembersNotForProjectResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListMembersNotForProjectResponseBody) SetMaxResults(v int32) *ListMembersNotForProjectResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListMembersNotForProjectResponseBody) SetMembers(v []*ListMembersNotForProjectResponseBodyMembers) *ListMembersNotForProjectResponseBody {
	s.Members = v
	return s
}

func (s *ListMembersNotForProjectResponseBody) SetNextToken(v string) *ListMembersNotForProjectResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMembersNotForProjectResponseBody) SetRequestId(v string) *ListMembersNotForProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMembersNotForProjectResponseBody) SetTotalCount(v int32) *ListMembersNotForProjectResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMembersNotForProjectResponseBody) Validate() error {
	if s.Members != nil {
		for _, item := range s.Members {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMembersNotForProjectResponseBodyMembers struct {
	// The UID of the member accounts.
	//
	// example:
	//
	// 122****3763
	MemberId *int64 `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s ListMembersNotForProjectResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s ListMembersNotForProjectResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListMembersNotForProjectResponseBodyMembers) GetMemberId() *int64 {
	return s.MemberId
}

func (s *ListMembersNotForProjectResponseBodyMembers) SetMemberId(v int64) *ListMembersNotForProjectResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *ListMembersNotForProjectResponseBodyMembers) Validate() error {
	return dara.Validate(s)
}
