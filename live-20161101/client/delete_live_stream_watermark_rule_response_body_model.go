// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamWatermarkRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLiveStreamWatermarkRuleResponseBody
	GetRequestId() *string
}

type DeleteLiveStreamWatermarkRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 5c6a2a0df228-4a64- af62-20e91******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveStreamWatermarkRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamWatermarkRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamWatermarkRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveStreamWatermarkRuleResponseBody) SetRequestId(v string) *DeleteLiveStreamWatermarkRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
