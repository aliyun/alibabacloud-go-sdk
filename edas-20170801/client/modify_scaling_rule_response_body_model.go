// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ModifyScalingRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *ModifyScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyScalingRuleResponseBody
	GetRequestId() *string
}

type ModifyScalingRuleResponseBody struct {
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message that is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CB36E997-FE54-476C-8C0D-********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyScalingRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ModifyScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyScalingRuleResponseBody) SetCode(v int32) *ModifyScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetMessage(v string) *ModifyScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) SetRequestId(v string) *ModifyScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
