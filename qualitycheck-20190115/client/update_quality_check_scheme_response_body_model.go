// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQualityCheckSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateQualityCheckSchemeResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateQualityCheckSchemeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateQualityCheckSchemeResponseBody
	GetMessage() *string
	SetMessages(v *UpdateQualityCheckSchemeResponseBodyMessages) *UpdateQualityCheckSchemeResponseBody
	GetMessages() *UpdateQualityCheckSchemeResponseBodyMessages
	SetRequestId(v string) *UpdateQualityCheckSchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateQualityCheckSchemeResponseBody
	GetSuccess() *bool
}

type UpdateQualityCheckSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateQualityCheckSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 96138D8D-8D26-4E41-BFF4-77AED1088BBD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateQualityCheckSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityCheckSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateQualityCheckSchemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateQualityCheckSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateQualityCheckSchemeResponseBody) GetMessages() *UpdateQualityCheckSchemeResponseBodyMessages {
	return s.Messages
}

func (s *UpdateQualityCheckSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQualityCheckSchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateQualityCheckSchemeResponseBody) SetCode(v string) *UpdateQualityCheckSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetHttpStatusCode(v int32) *UpdateQualityCheckSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetMessage(v string) *UpdateQualityCheckSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetMessages(v *UpdateQualityCheckSchemeResponseBodyMessages) *UpdateQualityCheckSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetRequestId(v string) *UpdateQualityCheckSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) SetSuccess(v bool) *UpdateQualityCheckSchemeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateQualityCheckSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateQualityCheckSchemeResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s UpdateQualityCheckSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateQualityCheckSchemeResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *UpdateQualityCheckSchemeResponseBodyMessages) SetMessage(v []*string) *UpdateQualityCheckSchemeResponseBodyMessages {
	s.Message = v
	return s
}

func (s *UpdateQualityCheckSchemeResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
