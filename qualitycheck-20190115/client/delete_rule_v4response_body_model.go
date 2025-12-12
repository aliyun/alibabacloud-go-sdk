// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRuleV4ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRuleV4ResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DeleteRuleV4ResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteRuleV4ResponseBody
	GetMessage() *string
	SetMessages(v *DeleteRuleV4ResponseBodyMessages) *DeleteRuleV4ResponseBody
	GetMessages() *DeleteRuleV4ResponseBodyMessages
	SetRequestId(v string) *DeleteRuleV4ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRuleV4ResponseBody
	GetSuccess() *bool
}

type DeleteRuleV4ResponseBody struct {
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
	Message  *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *DeleteRuleV4ResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// F190ADE9-619A-447D-84E3-7E241A5C428E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRuleV4ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRuleV4ResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRuleV4ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRuleV4ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRuleV4ResponseBody) GetMessages() *DeleteRuleV4ResponseBodyMessages {
	return s.Messages
}

func (s *DeleteRuleV4ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRuleV4ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRuleV4ResponseBody) SetCode(v string) *DeleteRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetHttpStatusCode(v int32) *DeleteRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetMessage(v string) *DeleteRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetMessages(v *DeleteRuleV4ResponseBodyMessages) *DeleteRuleV4ResponseBody {
	s.Messages = v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetRequestId(v string) *DeleteRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleV4ResponseBody) SetSuccess(v bool) *DeleteRuleV4ResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRuleV4ResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteRuleV4ResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s DeleteRuleV4ResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s DeleteRuleV4ResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *DeleteRuleV4ResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *DeleteRuleV4ResponseBodyMessages) SetMessage(v []*string) *DeleteRuleV4ResponseBodyMessages {
	s.Message = v
	return s
}

func (s *DeleteRuleV4ResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
