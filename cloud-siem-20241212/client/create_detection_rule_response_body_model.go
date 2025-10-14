// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDetectionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDetectionRuleId(v string) *CreateDetectionRuleResponseBody
	GetDetectionRuleId() *string
	SetRequestId(v string) *CreateDetectionRuleResponseBody
	GetRequestId() *string
}

type CreateDetectionRuleResponseBody struct {
	// example:
	//
	// dr-ha1i09ob3zmqrs85****
	DetectionRuleId *string `json:"DetectionRuleId,omitempty" xml:"DetectionRuleId,omitempty"`
	// example:
	//
	// 5CC09D0C-1CD7-54BD-A853-DCB2D945****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDetectionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDetectionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDetectionRuleResponseBody) GetDetectionRuleId() *string {
	return s.DetectionRuleId
}

func (s *CreateDetectionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDetectionRuleResponseBody) SetDetectionRuleId(v string) *CreateDetectionRuleResponseBody {
	s.DetectionRuleId = &v
	return s
}

func (s *CreateDetectionRuleResponseBody) SetRequestId(v string) *CreateDetectionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDetectionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
