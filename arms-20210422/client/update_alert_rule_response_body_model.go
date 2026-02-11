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
	AlertId   *int64  `json:"AlertId,omitempty" xml:"AlertId,omitempty"`
	Data      *string `json:"Data,omitempty" xml:"Data,omitempty"`
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
