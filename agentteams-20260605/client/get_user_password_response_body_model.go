// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserPasswordResponseBody
	GetCode() *string
	SetData(v *GetUserPasswordResponseBodyData) *GetUserPasswordResponseBody
	GetData() *GetUserPasswordResponseBodyData
	SetHttpStatusCode(v int32) *GetUserPasswordResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserPasswordResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUserPasswordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserPasswordResponseBody
	GetSuccess() *bool
}

type GetUserPasswordResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetUserPasswordResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUserPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserPasswordResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserPasswordResponseBody) GetData() *GetUserPasswordResponseBodyData {
	return s.Data
}

func (s *GetUserPasswordResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserPasswordResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserPasswordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserPasswordResponseBody) SetCode(v string) *GetUserPasswordResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserPasswordResponseBody) SetData(v *GetUserPasswordResponseBodyData) *GetUserPasswordResponseBody {
	s.Data = v
	return s
}

func (s *GetUserPasswordResponseBody) SetHttpStatusCode(v int32) *GetUserPasswordResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserPasswordResponseBody) SetMessage(v string) *GetUserPasswordResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserPasswordResponseBody) SetRequestId(v string) *GetUserPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserPasswordResponseBody) SetSuccess(v bool) *GetUserPasswordResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserPasswordResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetUserPasswordResponseBodyData struct {
	InitialPassword *string `json:"InitialPassword,omitempty" xml:"InitialPassword,omitempty"`
	InstanceId      *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	Name            *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetUserPasswordResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUserPasswordResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUserPasswordResponseBodyData) GetInitialPassword() *string {
	return s.InitialPassword
}

func (s *GetUserPasswordResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetUserPasswordResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetUserPasswordResponseBodyData) SetInitialPassword(v string) *GetUserPasswordResponseBodyData {
	s.InitialPassword = &v
	return s
}

func (s *GetUserPasswordResponseBodyData) SetInstanceId(v string) *GetUserPasswordResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *GetUserPasswordResponseBodyData) SetName(v string) *GetUserPasswordResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetUserPasswordResponseBodyData) Validate() error {
	return dara.Validate(s)
}
