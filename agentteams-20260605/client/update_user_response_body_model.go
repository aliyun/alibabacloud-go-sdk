// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateUserResponseBody
	GetCode() *string
	SetData(v *UpdateUserResponseBodyData) *UpdateUserResponseBody
	GetData() *UpdateUserResponseBodyData
	SetHttpStatusCode(v int32) *UpdateUserResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateUserResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateUserResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateUserResponseBody
	GetSuccess() *bool
}

type UpdateUserResponseBody struct {
	Code           *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *UpdateUserResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                      `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                       `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateUserResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateUserResponseBody) GetData() *UpdateUserResponseBodyData {
	return s.Data
}

func (s *UpdateUserResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateUserResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateUserResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateUserResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateUserResponseBody) SetCode(v string) *UpdateUserResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateUserResponseBody) SetData(v *UpdateUserResponseBodyData) *UpdateUserResponseBody {
	s.Data = v
	return s
}

func (s *UpdateUserResponseBody) SetHttpStatusCode(v int32) *UpdateUserResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateUserResponseBody) SetMessage(v string) *UpdateUserResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateUserResponseBody) SetRequestId(v string) *UpdateUserResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateUserResponseBody) SetSuccess(v bool) *UpdateUserResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateUserResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateUserResponseBodyData struct {
	AuthMethod  *string `json:"AuthMethod,omitempty" xml:"AuthMethod,omitempty"`
	CreatedAt   *string `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	Email       *string `json:"Email,omitempty" xml:"Email,omitempty"`
	InstanceId  *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Note        *string `json:"Note,omitempty" xml:"Note,omitempty"`
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateUserResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateUserResponseBodyData) GetAuthMethod() *string {
	return s.AuthMethod
}

func (s *UpdateUserResponseBodyData) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *UpdateUserResponseBodyData) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateUserResponseBodyData) GetEmail() *string {
	return s.Email
}

func (s *UpdateUserResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateUserResponseBodyData) GetName() *string {
	return s.Name
}

func (s *UpdateUserResponseBodyData) GetNote() *string {
	return s.Note
}

func (s *UpdateUserResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *UpdateUserResponseBodyData) SetAuthMethod(v string) *UpdateUserResponseBodyData {
	s.AuthMethod = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetCreatedAt(v string) *UpdateUserResponseBodyData {
	s.CreatedAt = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetDisplayName(v string) *UpdateUserResponseBodyData {
	s.DisplayName = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetEmail(v string) *UpdateUserResponseBodyData {
	s.Email = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetInstanceId(v string) *UpdateUserResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetName(v string) *UpdateUserResponseBodyData {
	s.Name = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetNote(v string) *UpdateUserResponseBodyData {
	s.Note = &v
	return s
}

func (s *UpdateUserResponseBodyData) SetStatus(v string) *UpdateUserResponseBodyData {
	s.Status = &v
	return s
}

func (s *UpdateUserResponseBodyData) Validate() error {
	return dara.Validate(s)
}
