// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupsForUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ListGroupsForUserResponseBodyData) *ListGroupsForUserResponseBody
	GetData() []*ListGroupsForUserResponseBodyData
	SetMaxResults(v int64) *ListGroupsForUserResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListGroupsForUserResponseBody
	GetNextToken() *string
	SetTotalCount(v int64) *ListGroupsForUserResponseBody
	GetTotalCount() *int64
}

type ListGroupsForUserResponseBody struct {
	Data []*ListGroupsForUserResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// NTxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1000
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListGroupsForUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBody) GetData() []*ListGroupsForUserResponseBodyData {
	return s.Data
}

func (s *ListGroupsForUserResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListGroupsForUserResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGroupsForUserResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListGroupsForUserResponseBody) SetData(v []*ListGroupsForUserResponseBodyData) *ListGroupsForUserResponseBody {
	s.Data = v
	return s
}

func (s *ListGroupsForUserResponseBody) SetMaxResults(v int64) *ListGroupsForUserResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGroupsForUserResponseBody) SetNextToken(v string) *ListGroupsForUserResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGroupsForUserResponseBody) SetTotalCount(v int64) *ListGroupsForUserResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListGroupsForUserResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListGroupsForUserResponseBodyData struct {
	// example:
	//
	// user_d6sbsuumeta4h66ec3il7yxxxx
	GroupId *string `json:"groupId,omitempty" xml:"groupId,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	GroupMemberRelationSourceId *string `json:"groupMemberRelationSourceId,omitempty" xml:"groupMemberRelationSourceId,omitempty"`
	// example:
	//
	// build_in
	GroupMemberRelationSourceType *string `json:"groupMemberRelationSourceType,omitempty" xml:"groupMemberRelationSourceType,omitempty"`
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s ListGroupsForUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGroupsForUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGroupsForUserResponseBodyData) GetGroupId() *string {
	return s.GroupId
}

func (s *ListGroupsForUserResponseBodyData) GetGroupMemberRelationSourceId() *string {
	return s.GroupMemberRelationSourceId
}

func (s *ListGroupsForUserResponseBodyData) GetGroupMemberRelationSourceType() *string {
	return s.GroupMemberRelationSourceType
}

func (s *ListGroupsForUserResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListGroupsForUserResponseBodyData) SetGroupId(v string) *ListGroupsForUserResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyData) SetGroupMemberRelationSourceId(v string) *ListGroupsForUserResponseBodyData {
	s.GroupMemberRelationSourceId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyData) SetGroupMemberRelationSourceType(v string) *ListGroupsForUserResponseBodyData {
	s.GroupMemberRelationSourceType = &v
	return s
}

func (s *ListGroupsForUserResponseBodyData) SetInstanceId(v string) *ListGroupsForUserResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListGroupsForUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
