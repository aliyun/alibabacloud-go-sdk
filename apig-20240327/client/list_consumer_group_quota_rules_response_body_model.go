// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerGroupQuotaRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumerGroupQuotaRulesResponseBody
	GetCode() *string
	SetData(v *ListConsumerGroupQuotaRulesResponseBodyData) *ListConsumerGroupQuotaRulesResponseBody
	GetData() *ListConsumerGroupQuotaRulesResponseBodyData
	SetMessage(v string) *ListConsumerGroupQuotaRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumerGroupQuotaRulesResponseBody
	GetRequestId() *string
}

type ListConsumerGroupQuotaRulesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {"totalSize":100}
	Data *ListConsumerGroupQuotaRulesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListConsumerGroupQuotaRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupQuotaRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupQuotaRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumerGroupQuotaRulesResponseBody) GetData() *ListConsumerGroupQuotaRulesResponseBodyData {
	return s.Data
}

func (s *ListConsumerGroupQuotaRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumerGroupQuotaRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumerGroupQuotaRulesResponseBody) SetCode(v string) *ListConsumerGroupQuotaRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBody) SetData(v *ListConsumerGroupQuotaRulesResponseBodyData) *ListConsumerGroupQuotaRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBody) SetMessage(v string) *ListConsumerGroupQuotaRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBody) SetRequestId(v string) *ListConsumerGroupQuotaRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConsumerGroupQuotaRulesResponseBodyData struct {
	// example:
	//
	// [{"ruleId":"rule-001"}]
	Items []*ListConsumerGroupQuotaRulesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 100
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListConsumerGroupQuotaRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupQuotaRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupQuotaRulesResponseBodyData) GetItems() []*ListConsumerGroupQuotaRulesResponseBodyDataItems {
	return s.Items
}

func (s *ListConsumerGroupQuotaRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerGroupQuotaRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerGroupQuotaRulesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListConsumerGroupQuotaRulesResponseBodyData) SetItems(v []*ListConsumerGroupQuotaRulesResponseBodyDataItems) *ListConsumerGroupQuotaRulesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyData) SetPageNumber(v int32) *ListConsumerGroupQuotaRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyData) SetPageSize(v int32) *ListConsumerGroupQuotaRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyData) SetTotalSize(v int32) *ListConsumerGroupQuotaRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConsumerGroupQuotaRulesResponseBodyDataItems struct {
	// example:
	//
	// gw-123456
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// prod-gateway
	GatewayName *string `json:"gatewayName,omitempty" xml:"gatewayName,omitempty"`
	// example:
	//
	// 30
	PeriodMultiplier *int64 `json:"periodMultiplier,omitempty" xml:"periodMultiplier,omitempty"`
	// example:
	//
	// week
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
	// rule-001
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
	// UTC+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// example:
	//
	// calendar
	WindowAlignment *string `json:"windowAlignment,omitempty" xml:"windowAlignment,omitempty"`
}

func (s ListConsumerGroupQuotaRulesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerGroupQuotaRulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetGatewayName() *string {
	return s.GatewayName
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetPeriodMultiplier() *int64 {
	return s.PeriodMultiplier
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetPeriodType() *string {
	return s.PeriodType
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetQuotaDimension() *string {
	return s.QuotaDimension
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetQuotaLimit() *int64 {
	return s.QuotaLimit
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetRuleId() *string {
	return s.RuleId
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetRuleName() *string {
	return s.RuleName
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetTimezone() *string {
	return s.Timezone
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) GetWindowAlignment() *string {
	return s.WindowAlignment
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetGatewayId(v string) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetGatewayName(v string) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.GatewayName = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetPeriodMultiplier(v int64) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.PeriodMultiplier = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetPeriodType(v string) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.PeriodType = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetQuotaDimension(v string) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.QuotaDimension = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetQuotaLimit(v int64) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.QuotaLimit = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetRuleId(v string) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.RuleId = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetRuleName(v string) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.RuleName = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetRuleStatus(v string) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.RuleStatus = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetTimezone(v string) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.Timezone = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) SetWindowAlignment(v string) *ListConsumerGroupQuotaRulesResponseBodyDataItems {
	s.WindowAlignment = &v
	return s
}

func (s *ListConsumerGroupQuotaRulesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
