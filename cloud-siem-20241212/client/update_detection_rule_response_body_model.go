// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDetectionRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDetectionRuleResponseBody
	GetRequestId() *string
}

type UpdateDetectionRuleResponseBody struct {
	// example:
	//
	// B88A2D41-87B8-537E-A7D3-3416A39F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDetectionRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDetectionRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDetectionRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDetectionRuleResponseBody) SetRequestId(v string) *UpdateDetectionRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDetectionRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
