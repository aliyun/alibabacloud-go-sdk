// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateUserResponseBody
	GetCode() *string
	SetData(v *CreateUserResponseBodyData) *CreateUserResponseBody
	GetData() *CreateUserResponseBodyData
	SetHttpStatusCode(v int32) *CreateUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateUserResponseBody
	GetMessage() *string
	SetParams(v []*string) *CreateUserResponseBody
	GetParams() []*string
	SetRequestId(v string) *CreateUserResponseBody
	GetRequestId() *string
}

type CreateUserResponseBody struct {
	// example:
	//
	// OK
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params         []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// BA03159C-E808-4FF1-B27E-A61B6E888D7F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateUserResponseBody) GetData() *CreateUserResponseBodyData {
	return s.Data
}

func (s *CreateUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateUserResponseBody) GetParams() []*string {
	return s.Params
}

func (s *CreateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserResponseBody) SetCode(v string) *CreateUserResponseBody {
	s.Code = &v
	return s
}

func (s *CreateUserResponseBody) SetData(v *CreateUserResponseBodyData) *CreateUserResponseBody {
	s.Data = v
	return s
}

func (s *CreateUserResponseBody) SetHttpStatusCode(v int32) *CreateUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateUserResponseBody) SetMessage(v string) *CreateUserResponseBody {
	s.Message = &v
	return s
}

func (s *CreateUserResponseBody) SetParams(v []*string) *CreateUserResponseBody {
	s.Params = v
	return s
}

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateUserResponseBodyData struct {
	AvatarUrl   *string `json:"AvatarUrl,omitempty" xml:"AvatarUrl,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// username@example.com
	Email *string `json:"Email,omitempty" xml:"Email,omitempty"`
	// example:
	//
	// 8032****
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// example:
	//
	// agent
	LoginName *string `json:"LoginName,omitempty" xml:"LoginName,omitempty"`
	// example:
	//
	// 1382114****
	Mobile   *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
	Nickname *string `json:"Nickname,omitempty" xml:"Nickname,omitempty"`
	// example:
	//
	// agent@ccc-test
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// example:
	//
	// ON_SITE
	WorkMode *string `json:"WorkMode,omitempty" xml:"WorkMode,omitempty"`
}

func (s CreateUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyData) GetAvatarUrl() *string {
	return s.AvatarUrl
}

func (s *CreateUserResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *CreateUserResponseBodyData) GetExtension() *string {
	return s.Extension
}

func (s *CreateUserResponseBodyData) GetLoginName() *string {
	return s.LoginName
}

func (s *CreateUserResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *CreateUserResponseBodyData) GetNickname() *string {
	return s.Nickname
}

func (s *CreateUserResponseBodyData) GetUserId() *string {
	return s.UserId
}

func (s *CreateUserResponseBodyData) GetWorkMode() *string {
	return s.WorkMode
}

func (s *CreateUserResponseBodyData) SetAvatarUrl(v string) *CreateUserResponseBodyData {
	s.AvatarUrl = &v
	return s
}

func (s *CreateUserResponseBodyData) SetDisplayName(v string) *CreateUserResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *CreateUserResponseBodyData) SetEmail(v string) *CreateUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *CreateUserResponseBodyData) SetExtension(v string) *CreateUserResponseBodyData {
	s.Extension = &v
	return s
}

func (s *CreateUserResponseBodyData) SetLoginName(v string) *CreateUserResponseBodyData {
	s.LoginName = &v
	return s
}

func (s *CreateUserResponseBodyData) SetMobile(v string) *CreateUserResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *CreateUserResponseBodyData) SetNickname(v string) *CreateUserResponseBodyData {
	s.Nickname = &v
	return s
}

func (s *CreateUserResponseBodyData) SetUserId(v string) *CreateUserResponseBodyData {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseBodyData) SetWorkMode(v string) *CreateUserResponseBodyData {
	s.WorkMode = &v
	return s
}

func (s *CreateUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
