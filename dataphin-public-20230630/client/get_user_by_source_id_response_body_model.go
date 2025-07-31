// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserBySourceIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserBySourceIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetUserBySourceIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserBySourceIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserBySourceIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserBySourceIdResponseBody
	GetSuccess() *bool
	SetUser(v *GetUserBySourceIdResponseBodyUser) *GetUserBySourceIdResponseBody
	GetUser() *GetUserBySourceIdResponseBodyUser
}

type GetUserBySourceIdResponseBody struct {
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
	Success *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	User    *GetUserBySourceIdResponseBodyUser `json:"User,omitempty" xml:"User,omitempty" type:"Struct"`
}

func (s GetUserBySourceIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserBySourceIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserBySourceIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserBySourceIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserBySourceIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserBySourceIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserBySourceIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserBySourceIdResponseBody) GetUser() *GetUserBySourceIdResponseBodyUser {
	return s.User
}

func (s *GetUserBySourceIdResponseBody) SetCode(v string) *GetUserBySourceIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetHttpStatusCode(v int32) *GetUserBySourceIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetMessage(v string) *GetUserBySourceIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetRequestId(v string) *GetUserBySourceIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetSuccess(v bool) *GetUserBySourceIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserBySourceIdResponseBody) SetUser(v *GetUserBySourceIdResponseBodyUser) *GetUserBySourceIdResponseBody {
	s.User = v
	return s
}

func (s *GetUserBySourceIdResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUserBySourceIdResponseBodyUser struct {
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 23231231
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetUserBySourceIdResponseBodyUser) String() string {
	return dara.Prettify(s)
}

func (s GetUserBySourceIdResponseBodyUser) GoString() string {
	return s.String()
}

func (s *GetUserBySourceIdResponseBodyUser) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetUserBySourceIdResponseBodyUser) GetId() *string {
	return s.Id
}

func (s *GetUserBySourceIdResponseBodyUser) SetDisplayName(v string) *GetUserBySourceIdResponseBodyUser {
	s.DisplayName = &v
	return s
}

func (s *GetUserBySourceIdResponseBodyUser) SetId(v string) *GetUserBySourceIdResponseBodyUser {
	s.Id = &v
	return s
}

func (s *GetUserBySourceIdResponseBodyUser) Validate() error {
	return dara.Validate(s)
}
