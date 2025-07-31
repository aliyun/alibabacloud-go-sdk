// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataQualityAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataQualityAlertRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataQualityAlertRuleResponseBody
	GetSuccess() *bool
}

type UpdateDataQualityAlertRuleResponseBody struct {
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataQualityAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataQualityAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataQualityAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataQualityAlertRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataQualityAlertRuleResponseBody) SetRequestId(v string) *UpdateDataQualityAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataQualityAlertRuleResponseBody) SetSuccess(v bool) *UpdateDataQualityAlertRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataQualityAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
