// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRuleResponseBody
	GetCode() *string
	SetData(v string) *UpdateRuleResponseBody
	GetData() *string
	SetMessage(v string) *UpdateRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRuleResponseBody
	GetSuccess() *bool
}

type UpdateRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// xxx
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRuleResponseBody) SetCode(v string) *UpdateRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRuleResponseBody) SetData(v string) *UpdateRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRuleResponseBody) SetMessage(v string) *UpdateRuleResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRuleResponseBody) SetRequestId(v string) *UpdateRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleResponseBody) SetSuccess(v bool) *UpdateRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
