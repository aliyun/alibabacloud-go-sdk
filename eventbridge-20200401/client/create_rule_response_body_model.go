// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRuleResponseBody
	GetCode() *string
	SetData(v *CreateRuleResponseBodyData) *CreateRuleResponseBody
	GetData() *CreateRuleResponseBodyData
	SetMessage(v string) *CreateRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRuleResponseBody
	GetSuccess() *bool
}

type CreateRuleResponseBody struct {
	// The returned HTTP status code. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *CreateRuleResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned error message.
	//
	// example:
	//
	// Remote error. requestId: [xxxx], error code: [xxx], message: [The target in event rule is invalid! Endpoint is xxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1AD6D598-7506-5D2C-81EA-30E3241A903A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRuleResponseBody) GetData() *CreateRuleResponseBodyData {
	return s.Data
}

func (s *CreateRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRuleResponseBody) SetCode(v string) *CreateRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRuleResponseBody) SetData(v *CreateRuleResponseBodyData) *CreateRuleResponseBody {
	s.Data = v
	return s
}

func (s *CreateRuleResponseBody) SetMessage(v string) *CreateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRuleResponseBody) SetRequestId(v string) *CreateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponseBody) SetSuccess(v bool) *CreateRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateRuleResponseBodyData struct {
	// The ARN of the event rule. The ARN is used for authorization.
	//
	// example:
	//
	// acs:eventbridge:cn-hangzhou:123456789098****:eventbus/default/rule/MNSRule
	RuleARN *string `json:"RuleARN,omitempty" xml:"RuleARN,omitempty"`
}

func (s CreateRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateRuleResponseBodyData) GetRuleARN() *string {
	return s.RuleARN
}

func (s *CreateRuleResponseBodyData) SetRuleARN(v string) *CreateRuleResponseBodyData {
	s.RuleARN = &v
	return s
}

func (s *CreateRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
