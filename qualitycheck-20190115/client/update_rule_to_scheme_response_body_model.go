// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleToSchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRuleToSchemeResponseBody
	GetCode() *string
	SetData(v int64) *UpdateRuleToSchemeResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateRuleToSchemeResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateRuleToSchemeResponseBody
	GetMessage() *string
	SetMessages(v *UpdateRuleToSchemeResponseBodyMessages) *UpdateRuleToSchemeResponseBody
	GetMessages() *UpdateRuleToSchemeResponseBodyMessages
	SetRequestId(v string) *UpdateRuleToSchemeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRuleToSchemeResponseBody
	GetSuccess() *bool
}

type UpdateRuleToSchemeResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 30
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateRuleToSchemeResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleToSchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleToSchemeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleToSchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRuleToSchemeResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateRuleToSchemeResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateRuleToSchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRuleToSchemeResponseBody) GetMessages() *UpdateRuleToSchemeResponseBodyMessages {
	return s.Messages
}

func (s *UpdateRuleToSchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRuleToSchemeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRuleToSchemeResponseBody) SetCode(v string) *UpdateRuleToSchemeResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetData(v int64) *UpdateRuleToSchemeResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetHttpStatusCode(v int32) *UpdateRuleToSchemeResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetMessage(v string) *UpdateRuleToSchemeResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetMessages(v *UpdateRuleToSchemeResponseBodyMessages) *UpdateRuleToSchemeResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetRequestId(v string) *UpdateRuleToSchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) SetSuccess(v bool) *UpdateRuleToSchemeResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRuleToSchemeResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleToSchemeResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateRuleToSchemeResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleToSchemeResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateRuleToSchemeResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *UpdateRuleToSchemeResponseBodyMessages) SetMessage(v []*string) *UpdateRuleToSchemeResponseBodyMessages {
	s.Message = v
	return s
}

func (s *UpdateRuleToSchemeResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
