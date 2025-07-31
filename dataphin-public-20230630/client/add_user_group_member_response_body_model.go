// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddUserGroupMemberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddUserGroupMemberResponseBody
	GetCode() *string
	SetData(v bool) *AddUserGroupMemberResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *AddUserGroupMemberResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddUserGroupMemberResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddUserGroupMemberResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddUserGroupMemberResponseBody
	GetSuccess() *bool
}

type AddUserGroupMemberResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
}

func (s AddUserGroupMemberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddUserGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *AddUserGroupMemberResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddUserGroupMemberResponseBody) GetData() *bool {
	return s.Data
}

func (s *AddUserGroupMemberResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddUserGroupMemberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddUserGroupMemberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddUserGroupMemberResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddUserGroupMemberResponseBody) SetCode(v string) *AddUserGroupMemberResponseBody {
	s.Code = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetData(v bool) *AddUserGroupMemberResponseBody {
	s.Data = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetHttpStatusCode(v int32) *AddUserGroupMemberResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetMessage(v string) *AddUserGroupMemberResponseBody {
	s.Message = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetRequestId(v string) *AddUserGroupMemberResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) SetSuccess(v bool) *AddUserGroupMemberResponseBody {
	s.Success = &v
	return s
}

func (s *AddUserGroupMemberResponseBody) Validate() error {
	return dara.Validate(s)
}
