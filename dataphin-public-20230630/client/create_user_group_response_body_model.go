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
	SetHttpStatusCode(v int32) *CreateUserGroupResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateUserGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateUserGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUserGroupResponseBody
	GetSuccess() *bool
	SetUserGroupId(v string) *CreateUserGroupResponseBody
	GetUserGroupId() *string
}

type CreateUserGroupResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 2313131
	UserGroupId *string `json:"UserGroupId,omitempty" xml:"UserGroupId,omitempty"`
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

func (s *CreateUserGroupResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateUserGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUserGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUserGroupResponseBody) GetUserGroupId() *string {
	return s.UserGroupId
}

func (s *CreateUserGroupResponseBody) SetCode(v string) *CreateUserGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetHttpStatusCode(v int32) *CreateUserGroupResponseBody {
	s.HttpStatusCode = &v
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

func (s *CreateUserGroupResponseBody) SetSuccess(v bool) *CreateUserGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserGroupResponseBody) SetUserGroupId(v string) *CreateUserGroupResponseBody {
	s.UserGroupId = &v
	return s
}

func (s *CreateUserGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
