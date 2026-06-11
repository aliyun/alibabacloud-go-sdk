// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayQuotaRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayQuotaRuleResponseBody
	GetCode() *string
	SetData(v *GetGatewayQuotaRuleResponseBodyData) *GetGatewayQuotaRuleResponseBody
	GetData() *GetGatewayQuotaRuleResponseBodyData
	SetMessage(v string) *GetGatewayQuotaRuleResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayQuotaRuleResponseBody
	GetRequestId() *string
}

type GetGatewayQuotaRuleResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {"ruleId":1001}
	Data *GetGatewayQuotaRuleResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayQuotaRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayQuotaRuleResponseBody) GetData() *GetGatewayQuotaRuleResponseBodyData {
	return s.Data
}

func (s *GetGatewayQuotaRuleResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayQuotaRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayQuotaRuleResponseBody) SetCode(v string) *GetGatewayQuotaRuleResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBody) SetData(v *GetGatewayQuotaRuleResponseBodyData) *GetGatewayQuotaRuleResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayQuotaRuleResponseBody) SetMessage(v string) *GetGatewayQuotaRuleResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBody) SetRequestId(v string) *GetGatewayQuotaRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayQuotaRuleResponseBodyData struct {
	// example:
	//
	// 1745846400000
	BaseTimestamp *int64 `json:"baseTimestamp,omitempty" xml:"baseTimestamp,omitempty"`
	// example:
	//
	// 20
	ConsumerCount *int64 `json:"consumerCount,omitempty" xml:"consumerCount,omitempty"`
	// example:
	//
	// day
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// example:
	//
	// token
	QuotaDimension *string `json:"quotaDimension,omitempty" xml:"quotaDimension,omitempty"`
	// example:
	//
	// 1000
	QuotaLimit *int64 `json:"quotaLimit,omitempty" xml:"quotaLimit,omitempty"`
	// example:
	//
	// qr-d8j7fpmm1hks65xxxxxx
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// example:
	//
	// daily-token-limit
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// example:
	//
	// enabled
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// example:
	//
	// GMT+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// example:
	//
	// calendar
	WindowAlignment *string `json:"windowAlignment,omitempty" xml:"windowAlignment,omitempty"`
}

func (s GetGatewayQuotaRuleResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetBaseTimestamp() *int64 {
	return s.BaseTimestamp
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetConsumerCount() *int64 {
	return s.ConsumerCount
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetPeriodType() *string {
	return s.PeriodType
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetQuotaDimension() *string {
	return s.QuotaDimension
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetQuotaLimit() *int64 {
	return s.QuotaLimit
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetRuleId() *string {
	return s.RuleId
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetRuleName() *string {
	return s.RuleName
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetTimezone() *string {
	return s.Timezone
}

func (s *GetGatewayQuotaRuleResponseBodyData) GetWindowAlignment() *string {
	return s.WindowAlignment
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetBaseTimestamp(v int64) *GetGatewayQuotaRuleResponseBodyData {
	s.BaseTimestamp = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetConsumerCount(v int64) *GetGatewayQuotaRuleResponseBodyData {
	s.ConsumerCount = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetPeriodType(v string) *GetGatewayQuotaRuleResponseBodyData {
	s.PeriodType = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetQuotaDimension(v string) *GetGatewayQuotaRuleResponseBodyData {
	s.QuotaDimension = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetQuotaLimit(v int64) *GetGatewayQuotaRuleResponseBodyData {
	s.QuotaLimit = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetRuleId(v string) *GetGatewayQuotaRuleResponseBodyData {
	s.RuleId = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetRuleName(v string) *GetGatewayQuotaRuleResponseBodyData {
	s.RuleName = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetRuleStatus(v string) *GetGatewayQuotaRuleResponseBodyData {
	s.RuleStatus = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetTimezone(v string) *GetGatewayQuotaRuleResponseBodyData {
	s.Timezone = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) SetWindowAlignment(v string) *GetGatewayQuotaRuleResponseBodyData {
	s.WindowAlignment = &v
	return s
}

func (s *GetGatewayQuotaRuleResponseBodyData) Validate() error {
	return dara.Validate(s)
}
