// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateUserGroupResponseBody
	GetCode() *string
	SetMessage(v string) *UpdateUserGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUserGroupResponseBody
	GetRequestId() *string
	SetUserGroup(v interface{}) *UpdateUserGroupResponseBody
	GetUserGroup() interface{}
}

type UpdateUserGroupResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description. This value is empty when the request is successful.
	//
	// example:
	//
	// The current zone list is illegal.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The updated user group information.
	//
	// example:
	//
	// {"userGroupId":"7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11","userGroupName":"华南销售","parentId":null,"level":1,"description":"华南销售组织","childGroupCount":0,"directMemberCount":2,"sourceType":"internal","externalSyncStatus":null,"gmtCreate":"2026-08-27T09:00:00Z","gmtModified":"2026-08-27T10:00:00Z"}
	UserGroup interface{} `json:"userGroup,omitempty" xml:"userGroup,omitempty"`
}

func (s UpdateUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateUserGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserGroupResponseBody) GetUserGroup() interface{} {
	return s.UserGroup
}

func (s *UpdateUserGroupResponseBody) SetCode(v string) *UpdateUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetMessage(v string) *UpdateUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetRequestId(v string) *UpdateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserGroupResponseBody) SetUserGroup(v interface{}) *UpdateUserGroupResponseBody {
	s.UserGroup = v
	return s
}

func (s *UpdateUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
