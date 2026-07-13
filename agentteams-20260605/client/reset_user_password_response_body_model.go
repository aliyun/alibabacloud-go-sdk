// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ResetUserPasswordResponseBody
	GetCode() *string
	SetData(v *ResetUserPasswordResponseBodyData) *ResetUserPasswordResponseBody
	GetData() *ResetUserPasswordResponseBodyData
	SetHttpStatusCode(v int32) *ResetUserPasswordResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ResetUserPasswordResponseBody
	GetMessage() *string
	SetRequestId(v string) *ResetUserPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ResetUserPasswordResponseBody
	GetSuccess() *bool
}

type ResetUserPasswordResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *ResetUserPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBody) GetCode() *string {
	return s.Code
}

func (s *ResetUserPasswordResponseBody) GetData() *ResetUserPasswordResponseBodyData {
	return s.Data
}

func (s *ResetUserPasswordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ResetUserPasswordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ResetUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetUserPasswordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ResetUserPasswordResponseBody) SetCode(v string) *ResetUserPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetData(v *ResetUserPasswordResponseBodyData) *ResetUserPasswordResponseBody {
	s.Data = v
	return s
}

func (s *ResetUserPasswordResponseBody) SetHttpStatusCode(v int32) *ResetUserPasswordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetMessage(v string) *ResetUserPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetRequestId(v string) *ResetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetUserPasswordResponseBody) SetSuccess(v bool) *ResetUserPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *ResetUserPasswordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ResetUserPasswordResponseBodyData struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name       *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Password   *string `json:"Password,omitempty" xml:"Password,omitempty"`
	Subject    *string `json:"Subject,omitempty" xml:"Subject,omitempty"`
	UserPoolId *string `json:"UserPoolId,omitempty" xml:"UserPoolId,omitempty"`
}

func (s ResetUserPasswordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ResetUserPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *ResetUserPasswordResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ResetUserPasswordResponseBodyData) GetName() *string {
	return s.Name
}

func (s *ResetUserPasswordResponseBodyData) GetPassword() *string {
	return s.Password
}

func (s *ResetUserPasswordResponseBodyData) GetSubject() *string {
	return s.Subject
}

func (s *ResetUserPasswordResponseBodyData) GetUserPoolId() *string {
	return s.UserPoolId
}

func (s *ResetUserPasswordResponseBodyData) SetInstanceId(v string) *ResetUserPasswordResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ResetUserPasswordResponseBodyData) SetName(v string) *ResetUserPasswordResponseBodyData {
	s.Name = &v
	return s
}

func (s *ResetUserPasswordResponseBodyData) SetPassword(v string) *ResetUserPasswordResponseBodyData {
	s.Password = &v
	return s
}

func (s *ResetUserPasswordResponseBodyData) SetSubject(v string) *ResetUserPasswordResponseBodyData {
	s.Subject = &v
	return s
}

func (s *ResetUserPasswordResponseBodyData) SetUserPoolId(v string) *ResetUserPasswordResponseBodyData {
	s.UserPoolId = &v
	return s
}

func (s *ResetUserPasswordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
