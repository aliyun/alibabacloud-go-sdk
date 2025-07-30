// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleV4ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRuleV4ResponseBody
	GetCode() *string
	SetData(v int64) *UpdateRuleV4ResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateRuleV4ResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateRuleV4ResponseBody
	GetMessage() *string
	SetMessages(v *UpdateRuleV4ResponseBodyMessages) *UpdateRuleV4ResponseBody
	GetMessages() *UpdateRuleV4ResponseBodyMessages
	SetRequestId(v string) *UpdateRuleV4ResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRuleV4ResponseBody
	GetSuccess() *bool
}

type UpdateRuleV4ResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message  *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Messages *UpdateRuleV4ResponseBodyMessages `json:"Messages,omitempty" xml:"Messages,omitempty" type:"Struct"`
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleV4ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleV4ResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleV4ResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRuleV4ResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateRuleV4ResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateRuleV4ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRuleV4ResponseBody) GetMessages() *UpdateRuleV4ResponseBodyMessages {
	return s.Messages
}

func (s *UpdateRuleV4ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRuleV4ResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRuleV4ResponseBody) SetCode(v string) *UpdateRuleV4ResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetData(v int64) *UpdateRuleV4ResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetHttpStatusCode(v int32) *UpdateRuleV4ResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetMessage(v string) *UpdateRuleV4ResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetMessages(v *UpdateRuleV4ResponseBodyMessages) *UpdateRuleV4ResponseBody {
	s.Messages = v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetRequestId(v string) *UpdateRuleV4ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) SetSuccess(v bool) *UpdateRuleV4ResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRuleV4ResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateRuleV4ResponseBodyMessages struct {
	Message []*string `json:"Message,omitempty" xml:"Message,omitempty" type:"Repeated"`
}

func (s UpdateRuleV4ResponseBodyMessages) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleV4ResponseBodyMessages) GoString() string {
	return s.String()
}

func (s *UpdateRuleV4ResponseBodyMessages) GetMessage() []*string {
	return s.Message
}

func (s *UpdateRuleV4ResponseBodyMessages) SetMessage(v []*string) *UpdateRuleV4ResponseBodyMessages {
	s.Message = v
	return s
}

func (s *UpdateRuleV4ResponseBodyMessages) Validate() error {
	return dara.Validate(s)
}
