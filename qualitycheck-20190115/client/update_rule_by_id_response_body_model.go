// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRuleByIdResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateRuleByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateRuleByIdResponseBody
	GetMessage() *string
	SetMessages(v *UpdateRuleByIdResponseBodyMessages) *UpdateRuleByIdResponseBody
	GetMessages() *UpdateRuleByIdResponseBodyMessages
	SetRequestId(v string) *UpdateRuleByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRuleByIdResponseBody
	GetSuccess() *bool
}

type UpdateRuleByIdResponseBody struct {
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
	Message  *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateRuleByIdResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 82C91484-B2D5-4D2A-A21F-A6D73F4D55C6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleByIdResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRuleByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateRuleByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRuleByIdResponseBody) GetMessages() *UpdateRuleByIdResponseBodyMessages {
	return s.Messages
}

func (s *UpdateRuleByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRuleByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRuleByIdResponseBody) SetCode(v string) *UpdateRuleByIdResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetHttpStatusCode(v int32) *UpdateRuleByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetMessage(v string) *UpdateRuleByIdResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetMessages(v *UpdateRuleByIdResponseBodyMessages) *UpdateRuleByIdResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetRequestId(v string) *UpdateRuleByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleByIdResponseBody) SetSuccess(v bool) *UpdateRuleByIdResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRuleByIdResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateRuleByIdResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateRuleByIdResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleByIdResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateRuleByIdResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *UpdateRuleByIdResponseBodyMessages) SetMessage(v []*string) *UpdateRuleByIdResponseBodyMessages {
	s.Message = v
	return s
}

func (s *UpdateRuleByIdResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
