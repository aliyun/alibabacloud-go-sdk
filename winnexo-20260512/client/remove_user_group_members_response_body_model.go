// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedCount(v int64) *RemoveUserGroupMembersResponseBody
	GetAffectedCount() *int64
	SetCode(v string) *RemoveUserGroupMembersResponseBody
	GetCode() *string
	SetMessage(v string) *RemoveUserGroupMembersResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveUserGroupMembersResponseBody
	GetRequestId() *string
	SetRequestedCount(v int64) *RemoveUserGroupMembersResponseBody
	GetRequestedCount() *int64
	SetUserGroupId(v string) *RemoveUserGroupMembersResponseBody
	GetUserGroupId() *string
}

type RemoveUserGroupMembersResponseBody struct {
	// The number of member relationships actually removed.
	//
	// example:
	//
	// 2
	AffectedCount *int64 `json:"affectedCount,omitempty" xml:"affectedCount,omitempty"`
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The number of requested members before deduplication.
	//
	// example:
	//
	// 2
	RequestedCount *int64 `json:"requestedCount,omitempty" xml:"requestedCount,omitempty"`
	// The ID of the target user group.
	//
	// example:
	//
	// 7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11
	UserGroupId *string `json:"userGroupId,omitempty" xml:"userGroupId,omitempty"`
}

func (s RemoveUserGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveUserGroupMembersResponseBody) GetAffectedCount() *int64 {
	return s.AffectedCount
}

func (s *RemoveUserGroupMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveUserGroupMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveUserGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveUserGroupMembersResponseBody) GetRequestedCount() *int64 {
	return s.RequestedCount
}

func (s *RemoveUserGroupMembersResponseBody) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *RemoveUserGroupMembersResponseBody) SetAffectedCount(v int64) *RemoveUserGroupMembersResponseBody {
	s.AffectedCount = &v
	return s
}

func (s *RemoveUserGroupMembersResponseBody) SetCode(v string) *RemoveUserGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveUserGroupMembersResponseBody) SetMessage(v string) *RemoveUserGroupMembersResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveUserGroupMembersResponseBody) SetRequestId(v string) *RemoveUserGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveUserGroupMembersResponseBody) SetRequestedCount(v int64) *RemoveUserGroupMembersResponseBody {
	s.RequestedCount = &v
	return s
}

func (s *RemoveUserGroupMembersResponseBody) SetUserGroupId(v string) *RemoveUserGroupMembersResponseBody {
	s.UserGroupId = &v
	return s
}

func (s *RemoveUserGroupMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
