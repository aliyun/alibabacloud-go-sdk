// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelLimitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelLimitsResponseBody
	GetCode() *string
	SetErrorMessage(v string) *ListModelLimitsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v int64) *ListModelLimitsResponseBody
	GetHttpStatusCode() *int64
	SetLimits(v []*ListModelLimitsResponseBodyLimits) *ListModelLimitsResponseBody
	GetLimits() []*ListModelLimitsResponseBodyLimits
	SetMaxResults(v int64) *ListModelLimitsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListModelLimitsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListModelLimitsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListModelLimitsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListModelLimitsResponseBody
	GetTotalCount() *int64
}

type ListModelLimitsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameter is invalid.
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The model throttling information.
	Limits []*ListModelLimitsResponseBodyLimits `json:"limits,omitempty" xml:"limits,omitempty" type:"Repeated"`
	// The maximum number of records returned in a single request.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The token for the next request.
	//
	// example:
	//
	// lwytFRtLdNk=
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 36045E0A-551D-592D-B1BC-4C56596CE59E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the API call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListModelLimitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelLimitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelLimitsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelLimitsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListModelLimitsResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *ListModelLimitsResponseBody) GetLimits() []*ListModelLimitsResponseBodyLimits {
	return s.Limits
}

func (s *ListModelLimitsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListModelLimitsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelLimitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelLimitsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListModelLimitsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListModelLimitsResponseBody) SetCode(v string) *ListModelLimitsResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelLimitsResponseBody) SetErrorMessage(v string) *ListModelLimitsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListModelLimitsResponseBody) SetHttpStatusCode(v int64) *ListModelLimitsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListModelLimitsResponseBody) SetLimits(v []*ListModelLimitsResponseBodyLimits) *ListModelLimitsResponseBody {
	s.Limits = v
	return s
}

func (s *ListModelLimitsResponseBody) SetMaxResults(v int64) *ListModelLimitsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListModelLimitsResponseBody) SetNextToken(v string) *ListModelLimitsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListModelLimitsResponseBody) SetRequestId(v string) *ListModelLimitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelLimitsResponseBody) SetSuccess(v bool) *ListModelLimitsResponseBody {
	s.Success = &v
	return s
}

func (s *ListModelLimitsResponseBody) SetTotalCount(v int64) *ListModelLimitsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListModelLimitsResponseBody) Validate() error {
	if s.Limits != nil {
		for _, item := range s.Limits {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelLimitsResponseBodyLimits struct {
	// The model.
	//
	// example:
	//
	// qwen-plus
	Model *string `json:"model,omitempty" xml:"model,omitempty"`
	// The model throttling configuration for the current user account.
	ModelLimit *ListModelLimitsResponseBodyLimitsModelLimit `json:"modelLimit,omitempty" xml:"modelLimit,omitempty" type:"Struct"`
	// The model name.
	//
	// example:
	//
	// qwen-plus
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The custom model throttling configuration for the current workspace.
	WorkspaceLimit *ListModelLimitsResponseBodyLimitsWorkspaceLimit `json:"workspaceLimit,omitempty" xml:"workspaceLimit,omitempty" type:"Struct"`
}

func (s ListModelLimitsResponseBodyLimits) String() string {
	return dara.Prettify(s)
}

func (s ListModelLimitsResponseBodyLimits) GoString() string {
	return s.String()
}

func (s *ListModelLimitsResponseBodyLimits) GetModel() *string {
	return s.Model
}

func (s *ListModelLimitsResponseBodyLimits) GetModelLimit() *ListModelLimitsResponseBodyLimitsModelLimit {
	return s.ModelLimit
}

func (s *ListModelLimitsResponseBodyLimits) GetName() *string {
	return s.Name
}

func (s *ListModelLimitsResponseBodyLimits) GetWorkspaceLimit() *ListModelLimitsResponseBodyLimitsWorkspaceLimit {
	return s.WorkspaceLimit
}

func (s *ListModelLimitsResponseBodyLimits) SetModel(v string) *ListModelLimitsResponseBodyLimits {
	s.Model = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimits) SetModelLimit(v *ListModelLimitsResponseBodyLimitsModelLimit) *ListModelLimitsResponseBodyLimits {
	s.ModelLimit = v
	return s
}

func (s *ListModelLimitsResponseBodyLimits) SetName(v string) *ListModelLimitsResponseBodyLimits {
	s.Name = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimits) SetWorkspaceLimit(v *ListModelLimitsResponseBodyLimitsWorkspaceLimit) *ListModelLimitsResponseBodyLimits {
	s.WorkspaceLimit = v
	return s
}

func (s *ListModelLimitsResponseBodyLimits) Validate() error {
	if s.ModelLimit != nil {
		if err := s.ModelLimit.Validate(); err != nil {
			return err
		}
	}
	if s.WorkspaceLimit != nil {
		if err := s.WorkspaceLimit.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListModelLimitsResponseBodyLimitsModelLimit struct {
	// The maximum concurrency.
	//
	// example:
	//
	// 10
	AsyncUserConcurrencyLimit *int64 `json:"asyncUserConcurrencyLimit,omitempty" xml:"asyncUserConcurrencyLimit,omitempty"`
	// The queue size.
	//
	// example:
	//
	// 10
	AsyncUserQueueLimit *int64 `json:"asyncUserQueueLimit,omitempty" xml:"asyncUserQueueLimit,omitempty"`
	// The request throttling value.
	//
	// example:
	//
	// 10
	RequestLimit *int64 `json:"requestLimit,omitempty" xml:"requestLimit,omitempty"`
	// The time period for request throttling, in seconds.
	//
	// example:
	//
	// 1
	RequestLimitPeriod *int32 `json:"requestLimitPeriod,omitempty" xml:"requestLimitPeriod,omitempty"`
	// The usage throttling value.
	//
	// example:
	//
	// 10
	UsageLimit *int64 `json:"usageLimit,omitempty" xml:"usageLimit,omitempty"`
	// The usage throttling unit.
	//
	// example:
	//
	// 10
	UsageLimitField *string `json:"usageLimitField,omitempty" xml:"usageLimitField,omitempty"`
	// The time period for usage throttling, in seconds.
	//
	// example:
	//
	// 1
	UsageLimitPeriod *int32 `json:"usageLimitPeriod,omitempty" xml:"usageLimitPeriod,omitempty"`
}

func (s ListModelLimitsResponseBodyLimitsModelLimit) String() string {
	return dara.Prettify(s)
}

func (s ListModelLimitsResponseBodyLimitsModelLimit) GoString() string {
	return s.String()
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) GetAsyncUserConcurrencyLimit() *int64 {
	return s.AsyncUserConcurrencyLimit
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) GetAsyncUserQueueLimit() *int64 {
	return s.AsyncUserQueueLimit
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) GetRequestLimit() *int64 {
	return s.RequestLimit
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) GetRequestLimitPeriod() *int32 {
	return s.RequestLimitPeriod
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) GetUsageLimitField() *string {
	return s.UsageLimitField
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) GetUsageLimitPeriod() *int32 {
	return s.UsageLimitPeriod
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) SetAsyncUserConcurrencyLimit(v int64) *ListModelLimitsResponseBodyLimitsModelLimit {
	s.AsyncUserConcurrencyLimit = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) SetAsyncUserQueueLimit(v int64) *ListModelLimitsResponseBodyLimitsModelLimit {
	s.AsyncUserQueueLimit = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) SetRequestLimit(v int64) *ListModelLimitsResponseBodyLimitsModelLimit {
	s.RequestLimit = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) SetRequestLimitPeriod(v int32) *ListModelLimitsResponseBodyLimitsModelLimit {
	s.RequestLimitPeriod = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) SetUsageLimit(v int64) *ListModelLimitsResponseBodyLimitsModelLimit {
	s.UsageLimit = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) SetUsageLimitField(v string) *ListModelLimitsResponseBodyLimitsModelLimit {
	s.UsageLimitField = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) SetUsageLimitPeriod(v int32) *ListModelLimitsResponseBodyLimitsModelLimit {
	s.UsageLimitPeriod = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsModelLimit) Validate() error {
	return dara.Validate(s)
}

type ListModelLimitsResponseBodyLimitsWorkspaceLimit struct {
	// The maximum concurrency.
	//
	// example:
	//
	// 10
	AsyncUserConcurrencyLimit *int64 `json:"asyncUserConcurrencyLimit,omitempty" xml:"asyncUserConcurrencyLimit,omitempty"`
	// The queue size.
	//
	// example:
	//
	// 10
	AsyncUserQueueLimit *int64 `json:"asyncUserQueueLimit,omitempty" xml:"asyncUserQueueLimit,omitempty"`
	// The request throttling value.
	//
	// example:
	//
	// 10
	RequestLimit *int64 `json:"requestLimit,omitempty" xml:"requestLimit,omitempty"`
	// The time period for request throttling, in seconds.
	//
	// example:
	//
	// 1
	RequestLimitPeriod *int32 `json:"requestLimitPeriod,omitempty" xml:"requestLimitPeriod,omitempty"`
	// The usage throttling value.
	//
	// example:
	//
	// 10
	UsageLimit *int64 `json:"usageLimit,omitempty" xml:"usageLimit,omitempty"`
	// The usage throttling unit.
	//
	// example:
	//
	// token
	UsageLimitField *string `json:"usageLimitField,omitempty" xml:"usageLimitField,omitempty"`
	// The time period for usage throttling, in seconds.
	//
	// example:
	//
	// 1
	UsageLimitPeriod *int32 `json:"usageLimitPeriod,omitempty" xml:"usageLimitPeriod,omitempty"`
}

func (s ListModelLimitsResponseBodyLimitsWorkspaceLimit) String() string {
	return dara.Prettify(s)
}

func (s ListModelLimitsResponseBodyLimitsWorkspaceLimit) GoString() string {
	return s.String()
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) GetAsyncUserConcurrencyLimit() *int64 {
	return s.AsyncUserConcurrencyLimit
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) GetAsyncUserQueueLimit() *int64 {
	return s.AsyncUserQueueLimit
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) GetRequestLimit() *int64 {
	return s.RequestLimit
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) GetRequestLimitPeriod() *int32 {
	return s.RequestLimitPeriod
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) GetUsageLimit() *int64 {
	return s.UsageLimit
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) GetUsageLimitField() *string {
	return s.UsageLimitField
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) GetUsageLimitPeriod() *int32 {
	return s.UsageLimitPeriod
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) SetAsyncUserConcurrencyLimit(v int64) *ListModelLimitsResponseBodyLimitsWorkspaceLimit {
	s.AsyncUserConcurrencyLimit = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) SetAsyncUserQueueLimit(v int64) *ListModelLimitsResponseBodyLimitsWorkspaceLimit {
	s.AsyncUserQueueLimit = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) SetRequestLimit(v int64) *ListModelLimitsResponseBodyLimitsWorkspaceLimit {
	s.RequestLimit = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) SetRequestLimitPeriod(v int32) *ListModelLimitsResponseBodyLimitsWorkspaceLimit {
	s.RequestLimitPeriod = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) SetUsageLimit(v int64) *ListModelLimitsResponseBodyLimitsWorkspaceLimit {
	s.UsageLimit = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) SetUsageLimitField(v string) *ListModelLimitsResponseBodyLimitsWorkspaceLimit {
	s.UsageLimitField = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) SetUsageLimitPeriod(v int32) *ListModelLimitsResponseBodyLimitsWorkspaceLimit {
	s.UsageLimitPeriod = &v
	return s
}

func (s *ListModelLimitsResponseBodyLimitsWorkspaceLimit) Validate() error {
	return dara.Validate(s)
}
