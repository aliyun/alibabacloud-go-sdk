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
	SetRequestId(v string) *CreateUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUserResponseBody
	GetSuccess() *bool
}

type CreateUserResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *CreateUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *CreateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUserResponseBody) GetSuccess() *bool {
	return s.Success
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

func (s *CreateUserResponseBody) SetRequestId(v string) *CreateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUserResponseBody) SetSuccess(v bool) *CreateUserResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateUserResponseBodyData struct {
	AuthMethod      *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	DisplayName     *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email           *string `json:"Email,omitempty" xml:"Email,omitempty"`
	InitialPassword *string `json:"InitialPassword,omitempty" xml:"InitialPassword,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note            *string `json:"Note,omitempty" xml:"Note,omitempty"`
}

func (s CreateUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBodyData) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *CreateUserResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *CreateUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *CreateUserResponseBodyData) GetInitialPassword() *string {
	return s.InitialPassword
}

func (s *CreateUserResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *CreateUserResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *CreateUserResponseBodyData) SetAuthMethod(v string) *CreateUserResponseBodyData {
	s.AuthMethod = &v
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

func (s *CreateUserResponseBodyData) SetInitialPassword(v string) *CreateUserResponseBodyData {
	s.InitialPassword = &v
	return s
}

func (s *CreateUserResponseBodyData) SetInstanceId(v string) *CreateUserResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *CreateUserResponseBodyData) SetName(v string) *CreateUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *CreateUserResponseBodyData) SetNote(v string) *CreateUserResponseBodyData {
	s.Note = &v
	return s
}

func (s *CreateUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
