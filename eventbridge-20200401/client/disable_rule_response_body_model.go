// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableRuleResponseBody
	GetCode() *string
	SetMessage(v string) *DisableRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableRuleResponseBody
	GetSuccess() *bool
}

type DisableRuleResponseBody struct {
	// The error code. The value Success indicates that the request is successful.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// EventRuleNotExisted
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C36345A1-75F3-5A1A-BFCF-33B8271971FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableRuleResponseBody) SetCode(v string) *DisableRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableRuleResponseBody) SetMessage(v string) *DisableRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableRuleResponseBody) SetRequestId(v string) *DisableRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableRuleResponseBody) SetSuccess(v bool) *DisableRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DisableRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
