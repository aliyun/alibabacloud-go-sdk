// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRuleResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteRuleResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteRuleResponseBody
	GetMessage() *string
	SetMessages(v *DeleteRuleResponseBodyMessages) *DeleteRuleResponseBody
	GetMessages() *DeleteRuleResponseBodyMessages
	SetRequestId(v string) *DeleteRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRuleResponseBody
	GetSuccess() *bool
}

type DeleteRuleResponseBody struct {
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
	Message  *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *DeleteRuleResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRuleResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRuleResponseBody) GetMessages() *DeleteRuleResponseBodyMessages {
	return s.Messages
}

func (s *DeleteRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRuleResponseBody) SetCode(v string) *DeleteRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRuleResponseBody) SetHttpStatusCode(v int32) *DeleteRuleResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRuleResponseBody) SetMessage(v string) *DeleteRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRuleResponseBody) SetMessages(v *DeleteRuleResponseBodyMessages) *DeleteRuleResponseBody {
	s.Messages = v
	return s
}

func (s *DeleteRuleResponseBody) SetRequestId(v string) *DeleteRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleResponseBody) SetSuccess(v bool) *DeleteRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRuleResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRuleResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s DeleteRuleResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *DeleteRuleResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *DeleteRuleResponseBodyMessages) SetMessage(v []*string) *DeleteRuleResponseBodyMessages {
	s.Message = v
	return s
}

func (s *DeleteRuleResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
