// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMetricRulesResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMetricRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMetricRulesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMetricRulesResponseBody
	GetSuccess() *bool
}

type DeleteMetricRulesResponseBody struct {
	// The status code. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request. Alibaba Cloud generates a unique ID for each request. Use this ID to troubleshoot issues.
	//
	// example:
	//
	// 09B3A484-1A2B-5743-BB6B-D796416F2677
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation was successful. Valid values:
	//
	// - true: The operation was successful.
	//
	// - false: The operation failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetricRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRulesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMetricRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMetricRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetricRulesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetricRulesResponseBody) SetCode(v string) *DeleteMetricRulesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) SetMessage(v string) *DeleteMetricRulesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) SetRequestId(v string) *DeleteMetricRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) SetSuccess(v bool) *DeleteMetricRulesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetricRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
