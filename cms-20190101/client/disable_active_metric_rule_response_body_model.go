// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableActiveMetricRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableActiveMetricRuleResponseBody
	GetCode() *string
	SetMessage(v string) *DisableActiveMetricRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableActiveMetricRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DisableActiveMetricRuleResponseBody
	GetSuccess() *bool
}

type DisableActiveMetricRuleResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The Request is not authorization.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F82E6667-7811-4BA0-842F-5B2DC42BBAAD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableActiveMetricRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableActiveMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DisableActiveMetricRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableActiveMetricRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableActiveMetricRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableActiveMetricRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DisableActiveMetricRuleResponseBody) SetCode(v string) *DisableActiveMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DisableActiveMetricRuleResponseBody) SetMessage(v string) *DisableActiveMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DisableActiveMetricRuleResponseBody) SetRequestId(v string) *DisableActiveMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableActiveMetricRuleResponseBody) SetSuccess(v bool) *DisableActiveMetricRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DisableActiveMetricRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
