// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRuleV4ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddRuleV4ResponseBody
	GetCode() *string
	SetData(v int64) *AddRuleV4ResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *AddRuleV4ResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddRuleV4ResponseBody
	GetMessage() *string
	SetMessages(v *AddRuleV4ResponseBodyMessages) *AddRuleV4ResponseBody
	GetMessages() *AddRuleV4ResponseBodyMessages
	SetRequestId(v string) *AddRuleV4ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddRuleV4ResponseBody
	GetSuccess() *bool
}

type AddRuleV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *int64  `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *AddRuleV4ResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddRuleV4ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *AddRuleV4ResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddRuleV4ResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddRuleV4ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddRuleV4ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddRuleV4ResponseBody) GetMessages() *AddRuleV4ResponseBodyMessages {
	return s.Messages
}

func (s *AddRuleV4ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddRuleV4ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddRuleV4ResponseBody) SetCode(v string) *AddRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetData(v int64) *AddRuleV4ResponseBody {
	s.Data = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetHttpStatusCode(v int32) *AddRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetMessage(v string) *AddRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetMessages(v *AddRuleV4ResponseBodyMessages) *AddRuleV4ResponseBody {
	s.Messages = v
	return s
}

func (s *AddRuleV4ResponseBody) SetRequestId(v string) *AddRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddRuleV4ResponseBody) SetSuccess(v bool) *AddRuleV4ResponseBody {
	s.Success = &v
	return s
}

func (s *AddRuleV4ResponseBody) Validate() error {
	if s.Messages != nil {
		if err := s.Messages.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddRuleV4ResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s AddRuleV4ResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s AddRuleV4ResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *AddRuleV4ResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *AddRuleV4ResponseBodyMessages) SetMessage(v []*string) *AddRuleV4ResponseBodyMessages {
	s.Message = v
	return s
}

func (s *AddRuleV4ResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
