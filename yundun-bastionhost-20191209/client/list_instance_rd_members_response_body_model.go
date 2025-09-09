// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRdMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListInstanceRdMembersResponseBody
	GetMaxResults() *int32
	SetMembers(v []*ListInstanceRdMembersResponseBodyMembers) *ListInstanceRdMembersResponseBody
	GetMembers() []*ListInstanceRdMembersResponseBodyMembers
	SetNextToken(v string) *ListInstanceRdMembersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListInstanceRdMembersResponseBody
	GetRequestId() *string
}

type ListInstanceRdMembersResponseBody struct {
	// example:
	//
	// 20
	MaxResults *int32                                      `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Members    []*ListInstanceRdMembersResponseBodyMembers `json:"Members,omitempty" xml:"Members,omitempty" type:"Repeated"`
	// example:
	//
	// 4ieSWJCwxvW3dk3wF.BqkrZmP72nWu5zJ5NWydMqyEs****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// EC9BF0F4-8983-491A-BC8C-1B4DD94976DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListInstanceRdMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRdMembersResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceRdMembersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListInstanceRdMembersResponseBody) GetMembers() []*ListInstanceRdMembersResponseBodyMembers {
	return s.Members
}

func (s *ListInstanceRdMembersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListInstanceRdMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceRdMembersResponseBody) SetMaxResults(v int32) *ListInstanceRdMembersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListInstanceRdMembersResponseBody) SetMembers(v []*ListInstanceRdMembersResponseBodyMembers) *ListInstanceRdMembersResponseBody {
	s.Members = v
	return s
}

func (s *ListInstanceRdMembersResponseBody) SetNextToken(v string) *ListInstanceRdMembersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListInstanceRdMembersResponseBody) SetRequestId(v string) *ListInstanceRdMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceRdMembersResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListInstanceRdMembersResponseBodyMembers struct {
	// example:
	//
	// 1197234496852779
	MemberId *string `json:"MemberId,omitempty" xml:"MemberId,omitempty"`
}

func (s ListInstanceRdMembersResponseBodyMembers) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRdMembersResponseBodyMembers) GoString() string {
	return s.String()
}

func (s *ListInstanceRdMembersResponseBodyMembers) GetMemberId() *string {
	return s.MemberId
}

func (s *ListInstanceRdMembersResponseBodyMembers) SetMemberId(v string) *ListInstanceRdMembersResponseBodyMembers {
	s.MemberId = &v
	return s
}

func (s *ListInstanceRdMembersResponseBodyMembers) Validate() error {
	return dara.Validate(s)
}
