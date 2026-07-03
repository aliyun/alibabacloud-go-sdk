// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayQuotaRuleSubjectUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayQuotaRuleSubjectUsageResponseBody
	GetCode() *string
	SetData(v *GetGatewayQuotaRuleSubjectUsageResponseBodyData) *GetGatewayQuotaRuleSubjectUsageResponseBody
	GetData() *GetGatewayQuotaRuleSubjectUsageResponseBodyData
	SetMessage(v string) *GetGatewayQuotaRuleSubjectUsageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayQuotaRuleSubjectUsageResponseBody
	GetRequestId() *string
}

type GetGatewayQuotaRuleSubjectUsageResponseBody struct {
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
	// {"usedAmount":500}
	Data *GetGatewayQuotaRuleSubjectUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1234567890
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBody) GetData() *GetGatewayQuotaRuleSubjectUsageResponseBodyData {
	return s.Data
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBody) SetCode(v string) *GetGatewayQuotaRuleSubjectUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBody) SetData(v *GetGatewayQuotaRuleSubjectUsageResponseBodyData) *GetGatewayQuotaRuleSubjectUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBody) SetMessage(v string) *GetGatewayQuotaRuleSubjectUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBody) SetRequestId(v string) *GetGatewayQuotaRuleSubjectUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayQuotaRuleSubjectUsageResponseBodyData struct {
	// The total cached token consumption.
	//
	// example:
	//
	// 20
	CachedAmount *int64 `json:"cachedAmount,omitempty" xml:"cachedAmount,omitempty"`
	// The paginated consumption details.
	//
	// example:
	//
	// {"totalSize":100}
	Details *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails `json:"details,omitempty" xml:"details,omitempty" type:"Struct"`
	// The total input token consumption.
	//
	// example:
	//
	// 300
	InputAmount *int64 `json:"inputAmount,omitempty" xml:"inputAmount,omitempty"`
	// The total output token consumption.
	//
	// example:
	//
	// 180
	OutputAmount *int64 `json:"outputAmount,omitempty" xml:"outputAmount,omitempty"`
	// Indicates whether the quota limit is exceeded.
	//
	// example:
	//
	// false
	OverLimit *bool `json:"overLimit,omitempty" xml:"overLimit,omitempty"`
	// The total quota of the subject.
	//
	// example:
	//
	// 1000
	TotalQuota *int64 `json:"totalQuota,omitempty" xml:"totalQuota,omitempty"`
	// The total used amount of the subject.
	//
	// example:
	//
	// 500
	UsedAmount *int64 `json:"usedAmount,omitempty" xml:"usedAmount,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) GetCachedAmount() *int64 {
	return s.CachedAmount
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) GetDetails() *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails {
	return s.Details
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) GetInputAmount() *int64 {
	return s.InputAmount
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) GetOutputAmount() *int64 {
	return s.OutputAmount
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) GetOverLimit() *bool {
	return s.OverLimit
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) GetUsedAmount() *int64 {
	return s.UsedAmount
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) SetCachedAmount(v int64) *GetGatewayQuotaRuleSubjectUsageResponseBodyData {
	s.CachedAmount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) SetDetails(v *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) *GetGatewayQuotaRuleSubjectUsageResponseBodyData {
	s.Details = v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) SetInputAmount(v int64) *GetGatewayQuotaRuleSubjectUsageResponseBodyData {
	s.InputAmount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) SetOutputAmount(v int64) *GetGatewayQuotaRuleSubjectUsageResponseBodyData {
	s.OutputAmount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) SetOverLimit(v bool) *GetGatewayQuotaRuleSubjectUsageResponseBodyData {
	s.OverLimit = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) SetTotalQuota(v int64) *GetGatewayQuotaRuleSubjectUsageResponseBodyData {
	s.TotalQuota = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) SetUsedAmount(v int64) *GetGatewayQuotaRuleSubjectUsageResponseBodyData {
	s.UsedAmount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyData) Validate() error {
	if s.Details != nil {
		if err := s.Details.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails struct {
	// The list of usage details.
	//
	// example:
	//
	// [{"model":"qwen-plus",\\"usedAmount\\":210}]
	Items []*GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalSize *int32 `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) GetItems() []*GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems {
	return s.Items
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) GetTotalSize() *int32 {
	return s.TotalSize
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) SetItems(v []*GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails {
	s.Items = v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) SetPageNumber(v int32) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails {
	s.PageNumber = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) SetPageSize(v int32) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails {
	s.PageSize = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) SetTotalSize(v int32) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails {
	s.TotalSize = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetails) Validate() error {
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

type GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems struct {
	// The cached token consumption.
	//
	// example:
	//
	// 10
	CachedAmount *int64 `json:"cachedAmount,omitempty" xml:"cachedAmount,omitempty"`
	// The input token consumption.
	//
	// example:
	//
	// 120
	InputAmount *int64 `json:"inputAmount,omitempty" xml:"inputAmount,omitempty"`
	// The model name.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The output token consumption.
	//
	// example:
	//
	// 80
	OutputAmount *int64 `json:"outputAmount,omitempty" xml:"outputAmount,omitempty"`
	// The consumption (request) time.
	//
	// example:
	//
	// 2026-06-05 13:16:31
	StartTime *string `json:"startTime,omitempty" xml:"startTime,omitempty"`
	// The total consumption.
	//
	// example:
	//
	// 210
	UsedAmount *int64 `json:"usedAmount,omitempty" xml:"usedAmount,omitempty"`
}

func (s GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) GoString() string {
	return s.String()
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) GetCachedAmount() *int64 {
	return s.CachedAmount
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) GetInputAmount() *int64 {
	return s.InputAmount
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) GetModel() *string {
	return s.Model
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) GetOutputAmount() *int64 {
	return s.OutputAmount
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) GetStartTime() *string {
	return s.StartTime
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) GetUsedAmount() *int64 {
	return s.UsedAmount
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) SetCachedAmount(v int64) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems {
	s.CachedAmount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) SetInputAmount(v int64) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems {
	s.InputAmount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) SetModel(v string) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems {
	s.Model = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) SetOutputAmount(v int64) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems {
	s.OutputAmount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) SetStartTime(v string) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems {
	s.StartTime = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) SetUsedAmount(v int64) *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems {
	s.UsedAmount = &v
	return s
}

func (s *GetGatewayQuotaRuleSubjectUsageResponseBodyDataDetailsItems) Validate() error {
	return dara.Validate(s)
}
