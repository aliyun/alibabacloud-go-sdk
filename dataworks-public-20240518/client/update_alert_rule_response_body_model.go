// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAlertRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAlertRuleResponseBody
	GetSuccess() *bool
}

type UpdateAlertRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D85FEE2B-6174-5817-AF9E-FDD02FEDA5BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAlertRuleResponseBody) SetRequestId(v string) *UpdateAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) SetSuccess(v bool) *UpdateAlertRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
