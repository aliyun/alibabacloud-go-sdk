// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertId(v int64) *UpdateAlertRuleResponseBody
	GetAlertId() *int64
	SetData(v string) *UpdateAlertRuleResponseBody
	GetData() *string
	SetRequestId(v string) *UpdateAlertRuleResponseBody
	GetRequestId() *string
}

type UpdateAlertRuleResponseBody struct {
	// The ID of the alert rule.
	//
	// example:
	//
	// 1234567
	AlertId *int64 `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	// The struct returned.
	//
	// example:
	//
	// -
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 6A9AEA84-7186-4D8D-B498-4585C6A2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlertRuleResponseBody) GetAlertId() *int64 {
	return s.AlertId
}

func (s *UpdateAlertRuleResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlertRuleResponseBody) SetAlertId(v int64) *UpdateAlertRuleResponseBody {
	s.AlertId = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) SetData(v string) *UpdateAlertRuleResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) SetRequestId(v string) *UpdateAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
