// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApplicationScalingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteApplicationScalingRuleResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteApplicationScalingRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApplicationScalingRuleResponseBody
	GetRequestId() *string
}

type DeleteApplicationScalingRuleResponseBody struct {
	Code    *int32  `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 5d6fa0bc-cc3**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteApplicationScalingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApplicationScalingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApplicationScalingRuleResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteApplicationScalingRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApplicationScalingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApplicationScalingRuleResponseBody) SetCode(v int32) *DeleteApplicationScalingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) SetMessage(v string) *DeleteApplicationScalingRuleResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) SetRequestId(v string) *DeleteApplicationScalingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApplicationScalingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
