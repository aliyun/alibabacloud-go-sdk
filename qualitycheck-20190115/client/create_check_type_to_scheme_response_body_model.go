// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCheckTypeToSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateCheckTypeToSchemeResponseBody
	GetCode() *string
	SetData(v int64) *CreateCheckTypeToSchemeResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateCheckTypeToSchemeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateCheckTypeToSchemeResponseBody
	GetMessage() *string
	SetMessages(v *CreateCheckTypeToSchemeResponseBodyMessages) *CreateCheckTypeToSchemeResponseBody
	GetMessages() *CreateCheckTypeToSchemeResponseBodyMessages
	SetRequestId(v string) *CreateCheckTypeToSchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateCheckTypeToSchemeResponseBody
	GetSuccess() *bool
}

type CreateCheckTypeToSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 5
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                      `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *CreateCheckTypeToSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24F4CE647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateCheckTypeToSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckTypeToSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCheckTypeToSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateCheckTypeToSchemeResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateCheckTypeToSchemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateCheckTypeToSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateCheckTypeToSchemeResponseBody) GetMessages() *CreateCheckTypeToSchemeResponseBodyMessages {
	return s.Messages
}

func (s *CreateCheckTypeToSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCheckTypeToSchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateCheckTypeToSchemeResponseBody) SetCode(v string) *CreateCheckTypeToSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetData(v int64) *CreateCheckTypeToSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetHttpStatusCode(v int32) *CreateCheckTypeToSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetMessage(v string) *CreateCheckTypeToSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetMessages(v *CreateCheckTypeToSchemeResponseBodyMessages) *CreateCheckTypeToSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetRequestId(v string) *CreateCheckTypeToSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) SetSuccess(v bool) *CreateCheckTypeToSchemeResponseBody {
	s.Success = &v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateCheckTypeToSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s CreateCheckTypeToSchemeResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s CreateCheckTypeToSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *CreateCheckTypeToSchemeResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *CreateCheckTypeToSchemeResponseBodyMessages) SetMessage(v []*string) *CreateCheckTypeToSchemeResponseBodyMessages {
	s.Message = v
	return s
}

func (s *CreateCheckTypeToSchemeResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
