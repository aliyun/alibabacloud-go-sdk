// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMonitorGroupDynamicRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteMonitorGroupDynamicRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteMonitorGroupDynamicRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMonitorGroupDynamicRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMonitorGroupDynamicRuleResponseBody
	GetSuccess() *bool
}

type DeleteMonitorGroupDynamicRuleResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 56B4516A-EB44-4C66-8854-0393B35F636B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMonitorGroupDynamicRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMonitorGroupDynamicRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) SetCode(v int32) *DeleteMonitorGroupDynamicRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) SetMessage(v string) *DeleteMonitorGroupDynamicRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) SetRequestId(v string) *DeleteMonitorGroupDynamicRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) SetSuccess(v bool) *DeleteMonitorGroupDynamicRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMonitorGroupDynamicRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
