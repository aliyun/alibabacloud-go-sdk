// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutCustomMetricRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PutCustomMetricRuleResponseBody
	GetCode() *string
	SetMessage(v string) *PutCustomMetricRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *PutCustomMetricRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PutCustomMetricRuleResponseBody
	GetSuccess() *bool
}

type PutCustomMetricRuleResponseBody struct {
	// 状态码。
	//
	// > 200表示成功。
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// 返回信息。接口调用成功时，返回为空；接口调用失败时，返回失败原因。
	//
	// example:
	//
	// ComparisonOperator is mandatory for this action.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// 请求ID。
	//
	// example:
	//
	// 65D50468-ECEF-48F1-A6E1-D952E89D9432
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 操作是否成功。取值：
	//
	// - true：成功。
	//
	// - false：失败。
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PutCustomMetricRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PutCustomMetricRuleResponseBody) GoString() string {
	return s.String()
}

func (s *PutCustomMetricRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *PutCustomMetricRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PutCustomMetricRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PutCustomMetricRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PutCustomMetricRuleResponseBody) SetCode(v string) *PutCustomMetricRuleResponseBody {
	s.Code = &v
	return s
}

func (s *PutCustomMetricRuleResponseBody) SetMessage(v string) *PutCustomMetricRuleResponseBody {
	s.Message = &v
	return s
}

func (s *PutCustomMetricRuleResponseBody) SetRequestId(v string) *PutCustomMetricRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *PutCustomMetricRuleResponseBody) SetSuccess(v bool) *PutCustomMetricRuleResponseBody {
	s.Success = &v
	return s
}

func (s *PutCustomMetricRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
