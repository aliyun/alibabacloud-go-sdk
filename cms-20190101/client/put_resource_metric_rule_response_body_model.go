// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutResourceMetricRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutResourceMetricRuleResponseBody
	GetCode() *string
	SetMessage(v string) *PutResourceMetricRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutResourceMetricRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutResourceMetricRuleResponseBody
	GetSuccess() *bool
}

type PutResourceMetricRuleResponseBody struct {
	// The responses code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 65D50468-ECEF-48F1-A6E1-D952E89D9436
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

func (s PutResourceMetricRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutResourceMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutResourceMetricRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutResourceMetricRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutResourceMetricRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutResourceMetricRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutResourceMetricRuleResponseBody) SetCode(v string) *PutResourceMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) SetMessage(v string) *PutResourceMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) SetRequestId(v string) *PutResourceMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) SetSuccess(v bool) *PutResourceMetricRuleResponseBody {
	s.Success = &v
	return s
}

func (s *PutResourceMetricRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
