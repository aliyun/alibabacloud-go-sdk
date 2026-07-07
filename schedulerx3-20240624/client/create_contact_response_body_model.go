// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateContactResponseBody
	GetCode() *int32
	SetData(v *CreateContactResponseBodyData) *CreateContactResponseBody
	GetData() *CreateContactResponseBodyData
	SetMessage(v string) *CreateContactResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateContactResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateContactResponseBody
	GetSuccess() *bool
}

type CreateContactResponseBody struct {
	// example:
	//
	// 200
	Code *int32                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CreateContactResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter error: content is null.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 2C3E52FF-CBE9-5C0E-8252-37ACFF1F5EFF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateContactResponseBody) GoString() string {
	return s.String()
}

func (s *CreateContactResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateContactResponseBody) GetData() *CreateContactResponseBodyData {
	return s.Data
}

func (s *CreateContactResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateContactResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateContactResponseBody) SetCode(v int32) *CreateContactResponseBody {
	s.Code = &v
	return s
}

func (s *CreateContactResponseBody) SetData(v *CreateContactResponseBodyData) *CreateContactResponseBody {
	s.Data = v
	return s
}

func (s *CreateContactResponseBody) SetMessage(v string) *CreateContactResponseBody {
	s.Message = &v
	return s
}

func (s *CreateContactResponseBody) SetRequestId(v string) *CreateContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateContactResponseBody) SetSuccess(v bool) *CreateContactResponseBody {
	s.Success = &v
	return s
}

func (s *CreateContactResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateContactResponseBodyData struct {
	// 新建联系人的 id，后续 Update/Delete 时使用
	//
	// example:
	//
	// 5000
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
}

func (s CreateContactResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateContactResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateContactResponseBodyData) GetContactId() *int64 {
	return s.ContactId
}

func (s *CreateContactResponseBodyData) SetContactId(v int64) *CreateContactResponseBodyData {
	s.ContactId = &v
	return s
}

func (s *CreateContactResponseBodyData) Validate() error {
	return dara.Validate(s)
}
