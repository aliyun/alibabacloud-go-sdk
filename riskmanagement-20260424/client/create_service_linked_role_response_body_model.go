// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceLinkedRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateServiceLinkedRoleResponseBody
	GetCode() *string
	SetData(v *CreateServiceLinkedRoleResponseBodyData) *CreateServiceLinkedRoleResponseBody
	GetData() *CreateServiceLinkedRoleResponseBodyData
	SetMessage(v string) *CreateServiceLinkedRoleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceLinkedRoleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateServiceLinkedRoleResponseBody
	GetSuccess() *bool
}

type CreateServiceLinkedRoleResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateServiceLinkedRoleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1B4C9A14-94E6-5EEB-BF39-7DACCE9AC0D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateServiceLinkedRoleResponseBody) GetData() *CreateServiceLinkedRoleResponseBodyData {
	return s.Data
}

func (s *CreateServiceLinkedRoleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceLinkedRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceLinkedRoleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateServiceLinkedRoleResponseBody) SetCode(v string) *CreateServiceLinkedRoleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetData(v *CreateServiceLinkedRoleResponseBodyData) *CreateServiceLinkedRoleResponseBody {
	s.Data = v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetMessage(v string) *CreateServiceLinkedRoleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) SetSuccess(v bool) *CreateServiceLinkedRoleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceLinkedRoleResponseBodyData struct {
	Body *CreateServiceLinkedRoleResponseBodyDataBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
}

func (s CreateServiceLinkedRoleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBodyData) GetBody() *CreateServiceLinkedRoleResponseBodyDataBody {
	return s.Body
}

func (s *CreateServiceLinkedRoleResponseBodyData) SetBody(v *CreateServiceLinkedRoleResponseBodyDataBody) *CreateServiceLinkedRoleResponseBodyData {
	s.Body = v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyData) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateServiceLinkedRoleResponseBodyDataBody struct {
	// example:
	//
	// E00516EB-A56A-5381-ACFE-E618DBC3D0EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceLinkedRoleResponseBodyDataBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceLinkedRoleResponseBodyDataBody) GoString() string {
	return s.String()
}

func (s *CreateServiceLinkedRoleResponseBodyDataBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceLinkedRoleResponseBodyDataBody) SetRequestId(v string) *CreateServiceLinkedRoleResponseBodyDataBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceLinkedRoleResponseBodyDataBody) Validate() error {
	return dara.Validate(s)
}
