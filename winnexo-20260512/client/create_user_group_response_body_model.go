// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateUserGroupResponseBody
	GetCode() *string
	SetMessage(v string) *CreateUserGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateUserGroupResponseBody
	GetRequestId() *string
	SetUserGroup(v interface{}) *CreateUserGroupResponseBody
	GetUserGroup() interface{}
}

type CreateUserGroupResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error description. This value is empty if the request is successful.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request trace ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The information about the newly created user group.
	//
	// example:
	//
	// {"userGroupId":"7ea8973f-7a5c-4e8a-956b-4fe0e7e2eb11","userGroupName":"华南销售","parentId":null,"level":1,"description":"华南销售组织","childGroupCount":0,"directMemberCount":0,"sourceType":"internal","externalSyncStatus":null,"gmtCreate":"2026-08-27T09:00:00Z","gmtModified":"2026-08-27T09:00:00Z"}
	UserGroup interface{} `json:"userGroup,omitempty" xml:"userGroup,omitempty"`
}

func (s CreateUserGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateUserGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserGroupResponseBody) GetUserGroup() interface{} {
	return s.UserGroup
}

func (s *CreateUserGroupResponseBody) SetCode(v string) *CreateUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetMessage(v string) *CreateUserGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetRequestId(v string) *CreateUserGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetUserGroup(v interface{}) *CreateUserGroupResponseBody {
	s.UserGroup = v
	return s
}

func (s *CreateUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
