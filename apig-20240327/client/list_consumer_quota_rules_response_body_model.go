// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerQuotaRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumerQuotaRulesResponseBody
	GetCode() *string
	SetData(v *ListConsumerQuotaRulesResponseBodyData) *ListConsumerQuotaRulesResponseBody
	GetData() *ListConsumerQuotaRulesResponseBodyData
	SetMessage(v string) *ListConsumerQuotaRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumerQuotaRulesResponseBody
	GetRequestId() *string
}

type ListConsumerQuotaRulesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// example:
	//
	// {"totalSize":100}
	Data *ListConsumerQuotaRulesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListConsumerQuotaRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerQuotaRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerQuotaRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumerQuotaRulesResponseBody) GetData() *ListConsumerQuotaRulesResponseBodyData {
	return s.Data
}

func (s *ListConsumerQuotaRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumerQuotaRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumerQuotaRulesResponseBody) SetCode(v string) *ListConsumerQuotaRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBody) SetData(v *ListConsumerQuotaRulesResponseBodyData) *ListConsumerQuotaRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerQuotaRulesResponseBody) SetMessage(v string) *ListConsumerQuotaRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBody) SetRequestId(v string) *ListConsumerQuotaRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConsumerQuotaRulesResponseBodyData struct {
	// example:
	//
	// [{"ruleId":"rule-001"}]
	Items []*ListConsumerQuotaRulesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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

func (s ListConsumerQuotaRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerQuotaRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerQuotaRulesResponseBodyData) GetItems() []*ListConsumerQuotaRulesResponseBodyDataItems {
	return s.Items
}

func (s *ListConsumerQuotaRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerQuotaRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerQuotaRulesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListConsumerQuotaRulesResponseBodyData) SetItems(v []*ListConsumerQuotaRulesResponseBodyDataItems) *ListConsumerQuotaRulesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyData) SetPageNumber(v int32) *ListConsumerQuotaRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyData) SetPageSize(v int32) *ListConsumerQuotaRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyData) SetTotalSize(v int32) *ListConsumerQuotaRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyData) Validate() error {
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

type ListConsumerQuotaRulesResponseBodyDataItems struct {
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
	PeriodMultiplier *string `json:"periodMultiplier,omitempty" xml:"periodMultiplier,omitempty"`
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
	// qr-d8j7fpmm1hksxxxxxx
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

func (s ListConsumerQuotaRulesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerQuotaRulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetGatewayName() *string {
	return s.GatewayName
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetPeriodMultiplier() *string {
	return s.PeriodMultiplier
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetPeriodType() *string {
	return s.PeriodType
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetQuotaDimension() *string {
	return s.QuotaDimension
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetQuotaLimit() *int64 {
	return s.QuotaLimit
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetRuleId() *string {
	return s.RuleId
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetRuleName() *string {
	return s.RuleName
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetTimezone() *string {
	return s.Timezone
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) GetWindowAlignment() *string {
	return s.WindowAlignment
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetGatewayId(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.GatewayId = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetGatewayName(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.GatewayName = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetPeriodMultiplier(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.PeriodMultiplier = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetPeriodType(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.PeriodType = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetQuotaDimension(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.QuotaDimension = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetQuotaLimit(v int64) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.QuotaLimit = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetRuleId(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.RuleId = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetRuleName(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.RuleName = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetRuleStatus(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.RuleStatus = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetTimezone(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.Timezone = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) SetWindowAlignment(v string) *ListConsumerQuotaRulesResponseBodyDataItems {
	s.WindowAlignment = &v
	return s
}

func (s *ListConsumerQuotaRulesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
