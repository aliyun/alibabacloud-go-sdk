// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMetricRuleBlackListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyMetricRuleBlackListResponseBody
	GetCode() *string
	SetCount(v string) *ModifyMetricRuleBlackListResponseBody
	GetCount() *string
	SetMessage(v string) *ModifyMetricRuleBlackListResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyMetricRuleBlackListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModifyMetricRuleBlackListResponseBody
	GetSuccess() *bool
}

type ModifyMetricRuleBlackListResponseBody struct {
	// The error code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The number of blacklist policies that are modified.
	//
	// example:
	//
	// 1
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
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
	// 008773AE-1D86-3231-90F9-1AF7F808F9CE
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

func (s ModifyMetricRuleBlackListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMetricRuleBlackListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMetricRuleBlackListResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyMetricRuleBlackListResponseBody) GetCount() *string {
	return s.Count
}

func (s *ModifyMetricRuleBlackListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMetricRuleBlackListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMetricRuleBlackListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModifyMetricRuleBlackListResponseBody) SetCode(v string) *ModifyMetricRuleBlackListResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMetricRuleBlackListResponseBody) SetCount(v string) *ModifyMetricRuleBlackListResponseBody {
	s.Count = &v
	return s
}

func (s *ModifyMetricRuleBlackListResponseBody) SetMessage(v string) *ModifyMetricRuleBlackListResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMetricRuleBlackListResponseBody) SetRequestId(v string) *ModifyMetricRuleBlackListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMetricRuleBlackListResponseBody) SetSuccess(v bool) *ModifyMetricRuleBlackListResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyMetricRuleBlackListResponseBody) Validate() error {
	return dara.Validate(s)
}
