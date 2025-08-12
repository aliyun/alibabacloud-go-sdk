// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetricRuleBlackListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMetricRuleBlackListResponseBody
	GetCode() *string
	SetCount(v int32) *DeleteMetricRuleBlackListResponseBody
	GetCount() *int32
	SetMessage(v string) *DeleteMetricRuleBlackListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMetricRuleBlackListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMetricRuleBlackListResponseBody
	GetSuccess() *bool
}

type DeleteMetricRuleBlackListResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of blacklist policies that are deleted.
	//
	// example:
	//
	// 1
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B88D233C-A004-3AB8-AD9C-30CBDD4440C5
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

func (s DeleteMetricRuleBlackListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetricRuleBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetricRuleBlackListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMetricRuleBlackListResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DeleteMetricRuleBlackListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMetricRuleBlackListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetricRuleBlackListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetricRuleBlackListResponseBody) SetCode(v string) *DeleteMetricRuleBlackListResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMetricRuleBlackListResponseBody) SetCount(v int32) *DeleteMetricRuleBlackListResponseBody {
	s.Count = &v
	return s
}

func (s *DeleteMetricRuleBlackListResponseBody) SetMessage(v string) *DeleteMetricRuleBlackListResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMetricRuleBlackListResponseBody) SetRequestId(v string) *DeleteMetricRuleBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetricRuleBlackListResponseBody) SetSuccess(v bool) *DeleteMetricRuleBlackListResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetricRuleBlackListResponseBody) Validate() error {
	return dara.Validate(s)
}
