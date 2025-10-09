// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataQualityAlertRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *CreateDataQualityAlertRuleResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDataQualityAlertRuleResponseBody
	GetRequestId() *string
}

type CreateDataQualityAlertRuleResponseBody struct {
	// The user-defined rule ID returned after the monitoring rule is successfully created.
	//
	// example:
	//
	// 1010543619
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc14115***159376359
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDataQualityAlertRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataQualityAlertRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataQualityAlertRuleResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDataQualityAlertRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataQualityAlertRuleResponseBody) SetId(v int64) *CreateDataQualityAlertRuleResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDataQualityAlertRuleResponseBody) SetRequestId(v string) *CreateDataQualityAlertRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataQualityAlertRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
