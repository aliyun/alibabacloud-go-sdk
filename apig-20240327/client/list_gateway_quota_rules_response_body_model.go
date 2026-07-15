// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGatewayQuotaRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGatewayQuotaRulesResponseBody
	GetCode() *string
	SetData(v *ListGatewayQuotaRulesResponseBodyData) *ListGatewayQuotaRulesResponseBody
	GetData() *ListGatewayQuotaRulesResponseBodyData
	SetMaxResults(v int32) *ListGatewayQuotaRulesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListGatewayQuotaRulesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListGatewayQuotaRulesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListGatewayQuotaRulesResponseBody
	GetRequestId() *string
}

type ListGatewayQuotaRulesResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	//
	// example:
	//
	// {"totalSize":100}
	Data *ListGatewayQuotaRulesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The maximum number of records to retrieve at a time. This parameter is not supported.
	//
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token. This parameter is not supported.
	//
	// example:
	//
	// 762b1fa4e2434fd3959b1f66481979cf
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGatewayQuotaRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayQuotaRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListGatewayQuotaRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGatewayQuotaRulesResponseBody) GetData() *ListGatewayQuotaRulesResponseBodyData {
	return s.Data
}

func (s *ListGatewayQuotaRulesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListGatewayQuotaRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGatewayQuotaRulesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListGatewayQuotaRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGatewayQuotaRulesResponseBody) SetCode(v string) *ListGatewayQuotaRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBody) SetData(v *ListGatewayQuotaRulesResponseBodyData) *ListGatewayQuotaRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListGatewayQuotaRulesResponseBody) SetMaxResults(v int32) *ListGatewayQuotaRulesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBody) SetMessage(v string) *ListGatewayQuotaRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBody) SetNextToken(v string) *ListGatewayQuotaRulesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBody) SetRequestId(v string) *ListGatewayQuotaRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListGatewayQuotaRulesResponseBodyData struct {
	// The list of rules.
	//
	// example:
	//
	// [{"ruleId":"rule-001"}]
	Items []*ListGatewayQuotaRulesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The current page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 100
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListGatewayQuotaRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayQuotaRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListGatewayQuotaRulesResponseBodyData) GetItems() []*ListGatewayQuotaRulesResponseBodyDataItems {
	return s.Items
}

func (s *ListGatewayQuotaRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListGatewayQuotaRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListGatewayQuotaRulesResponseBodyData) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *ListGatewayQuotaRulesResponseBodyData) SetItems(v []*ListGatewayQuotaRulesResponseBodyDataItems) *ListGatewayQuotaRulesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyData) SetPageNumber(v int32) *ListGatewayQuotaRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyData) SetPageSize(v int32) *ListGatewayQuotaRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyData) SetTotalSize(v int32) *ListGatewayQuotaRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyData) Validate() error {
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

type ListGatewayQuotaRulesResponseBodyDataItems struct {
	// The period type.
	//
	// example:
	//
	// week
	PeriodType *string `json:"periodType,omitempty" xml:"periodType,omitempty"`
	// The quota dimension.
	//
	// example:
	//
	// token
	QuotaDimension *string `json:"quotaDimension,omitempty" xml:"quotaDimension,omitempty"`
	// The quota limit.
	//
	// example:
	//
	// 1000
	QuotaLimit *int64 `json:"quotaLimit,omitempty" xml:"quotaLimit,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// qr-xxxxx
	RuleId *string `json:"ruleId,omitempty" xml:"ruleId,omitempty"`
	// The rule name.
	//
	// example:
	//
	// daily-token-limit
	RuleName *string `json:"ruleName,omitempty" xml:"ruleName,omitempty"`
	// The rule status.
	//
	// example:
	//
	// enabled
	RuleStatus *string `json:"ruleStatus,omitempty" xml:"ruleStatus,omitempty"`
	// The time zone for the natural period, in UTC+x format.
	//
	// example:
	//
	// GMT+8
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty"`
	// The reset period type.
	//
	// example:
	//
	// calendar
	WindowAlignment *string `json:"windowAlignment,omitempty" xml:"windowAlignment,omitempty"`
}

func (s ListGatewayQuotaRulesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListGatewayQuotaRulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) GetPeriodType() *string {
	return s.PeriodType
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) GetQuotaDimension() *string {
	return s.QuotaDimension
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) GetQuotaLimit() *int64 {
	return s.QuotaLimit
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) GetRuleId() *string {
	return s.RuleId
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) GetRuleName() *string {
	return s.RuleName
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) GetRuleStatus() *string {
	return s.RuleStatus
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) GetTimezone() *string {
	return s.Timezone
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) GetWindowAlignment() *string {
	return s.WindowAlignment
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) SetPeriodType(v string) *ListGatewayQuotaRulesResponseBodyDataItems {
	s.PeriodType = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) SetQuotaDimension(v string) *ListGatewayQuotaRulesResponseBodyDataItems {
	s.QuotaDimension = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) SetQuotaLimit(v int64) *ListGatewayQuotaRulesResponseBodyDataItems {
	s.QuotaLimit = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) SetRuleId(v string) *ListGatewayQuotaRulesResponseBodyDataItems {
	s.RuleId = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) SetRuleName(v string) *ListGatewayQuotaRulesResponseBodyDataItems {
	s.RuleName = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) SetRuleStatus(v string) *ListGatewayQuotaRulesResponseBodyDataItems {
	s.RuleStatus = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) SetTimezone(v string) *ListGatewayQuotaRulesResponseBodyDataItems {
	s.Timezone = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) SetWindowAlignment(v string) *ListGatewayQuotaRulesResponseBodyDataItems {
	s.WindowAlignment = &v
	return s
}

func (s *ListGatewayQuotaRulesResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
