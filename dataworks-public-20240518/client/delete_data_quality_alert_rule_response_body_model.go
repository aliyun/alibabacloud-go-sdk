// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataQualityAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataQualityAlertRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataQualityAlertRuleResponseBody
	GetSuccess() *bool
}

type DeleteDataQualityAlertRuleResponseBody struct {
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataQualityAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataQualityAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataQualityAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataQualityAlertRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataQualityAlertRuleResponseBody) SetRequestId(v string) *DeleteDataQualityAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataQualityAlertRuleResponseBody) SetSuccess(v bool) *DeleteDataQualityAlertRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataQualityAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
