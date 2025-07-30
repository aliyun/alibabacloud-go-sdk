// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInvalidRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *InvalidRuleResponseBody
	GetCode() *string
	SetData(v bool) *InvalidRuleResponseBody
	GetData() *bool
	SetMessage(v string) *InvalidRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *InvalidRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *InvalidRuleResponseBody
	GetSuccess() *bool
}

type InvalidRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s InvalidRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InvalidRuleResponseBody) GoString() string {
	return s.String()
}

func (s *InvalidRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *InvalidRuleResponseBody) GetData() *bool {
	return s.Data
}

func (s *InvalidRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InvalidRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InvalidRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *InvalidRuleResponseBody) SetCode(v string) *InvalidRuleResponseBody {
	s.Code = &v
	return s
}

func (s *InvalidRuleResponseBody) SetData(v bool) *InvalidRuleResponseBody {
	s.Data = &v
	return s
}

func (s *InvalidRuleResponseBody) SetMessage(v string) *InvalidRuleResponseBody {
	s.Message = &v
	return s
}

func (s *InvalidRuleResponseBody) SetRequestId(v string) *InvalidRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *InvalidRuleResponseBody) SetSuccess(v bool) *InvalidRuleResponseBody {
	s.Success = &v
	return s
}

func (s *InvalidRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
