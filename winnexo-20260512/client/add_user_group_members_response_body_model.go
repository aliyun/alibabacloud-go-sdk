// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMembersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAffectedCount(v int64) *AddUserGroupMembersResponseBody
	GetAffectedCount() *int64
	SetCode(v string) *AddUserGroupMembersResponseBody
	GetCode() *string
	SetMessage(v string) *AddUserGroupMembersResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddUserGroupMembersResponseBody
	GetRequestId() *string
	SetRequestedCount(v int64) *AddUserGroupMembersResponseBody
	GetRequestedCount() *int64
	SetUserGroupId(v string) *AddUserGroupMembersResponseBody
	GetUserGroupId() *string
}

type AddUserGroupMembersResponseBody struct {
	// The number of user group member relationships that were actually added.
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
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
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

func (s AddUserGroupMembersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMembersResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserGroupMembersResponseBody) GetAffectedCount() *int64 {
	return s.AffectedCount
}

func (s *AddUserGroupMembersResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddUserGroupMembersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddUserGroupMembersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserGroupMembersResponseBody) GetRequestedCount() *int64 {
	return s.RequestedCount
}

func (s *AddUserGroupMembersResponseBody) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *AddUserGroupMembersResponseBody) SetAffectedCount(v int64) *AddUserGroupMembersResponseBody {
	s.AffectedCount = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) SetCode(v string) *AddUserGroupMembersResponseBody {
	s.Code = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) SetMessage(v string) *AddUserGroupMembersResponseBody {
	s.Message = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) SetRequestId(v string) *AddUserGroupMembersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) SetRequestedCount(v int64) *AddUserGroupMembersResponseBody {
	s.RequestedCount = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) SetUserGroupId(v string) *AddUserGroupMembersResponseBody {
	s.UserGroupId = &v
	return s
}

func (s *AddUserGroupMembersResponseBody) Validate() error {
	return dara.Validate(s)
}
