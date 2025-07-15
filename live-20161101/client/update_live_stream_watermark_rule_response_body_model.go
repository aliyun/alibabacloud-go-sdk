// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamWatermarkRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLiveStreamWatermarkRuleResponseBody
	GetRequestId() *string
}

type UpdateLiveStreamWatermarkRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64-af62-20e91b9676b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLiveStreamWatermarkRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamWatermarkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamWatermarkRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLiveStreamWatermarkRuleResponseBody) SetRequestId(v string) *UpdateLiveStreamWatermarkRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLiveStreamWatermarkRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
