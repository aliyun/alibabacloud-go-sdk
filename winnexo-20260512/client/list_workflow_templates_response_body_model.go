// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWorkflowTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListWorkflowTemplatesResponseBody
	GetCode() *string
	SetHasMore(v bool) *ListWorkflowTemplatesResponseBody
	GetHasMore() *bool
	SetItems(v []*ListWorkflowTemplatesResponseBodyItems) *ListWorkflowTemplatesResponseBody
	GetItems() []*ListWorkflowTemplatesResponseBodyItems
	SetMaxResults(v int64) *ListWorkflowTemplatesResponseBody
	GetMaxResults() *int64
	SetMessage(v string) *ListWorkflowTemplatesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListWorkflowTemplatesResponseBody
	GetNextToken() *string
	SetPage(v int64) *ListWorkflowTemplatesResponseBody
	GetPage() *int64
	SetPageSize(v int64) *ListWorkflowTemplatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListWorkflowTemplatesResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListWorkflowTemplatesResponseBody
	GetTotal() *int64
}

type ListWorkflowTemplatesResponseBody struct {
	// The business status code. A value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Indicates whether more data is available.
	//
	// example:
	//
	// false
	HasMore *bool `json:"hasMore,omitempty" xml:"hasMore,omitempty"`
	// The list of template summaries on the current page.
	//
	// example:
	//
	// example
	Items []*ListWorkflowTemplatesResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The maximum number of entries returned per page in the current request.
	//
	// example:
	//
	// 1
	MaxResults *int64 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The error message.
	//
	// example:
	//
	// The requested resource does not exist
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The pagination token for the next page. An empty string indicates that no more data is available.
	//
	// example:
	//
	// example
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The current page number.
	//
	// example:
	//
	// 1
	Page *int64 `json:"page,omitempty" xml:"page,omitempty"`
	// The number of entries per page in the current request.
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The total number of visible entries that match the filter conditions.
	//
	// example:
	//
	// 1
	Total *int64 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ListWorkflowTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplatesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListWorkflowTemplatesResponseBody) GetHasMore() *bool {
	return s.HasMore
}

func (s *ListWorkflowTemplatesResponseBody) GetItems() []*ListWorkflowTemplatesResponseBodyItems {
	return s.Items
}

func (s *ListWorkflowTemplatesResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListWorkflowTemplatesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListWorkflowTemplatesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListWorkflowTemplatesResponseBody) GetPage() *int64 {
	return s.Page
}

func (s *ListWorkflowTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListWorkflowTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListWorkflowTemplatesResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListWorkflowTemplatesResponseBody) SetCode(v string) *ListWorkflowTemplatesResponseBody {
	s.Code = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetHasMore(v bool) *ListWorkflowTemplatesResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetItems(v []*ListWorkflowTemplatesResponseBodyItems) *ListWorkflowTemplatesResponseBody {
	s.Items = v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetMaxResults(v int64) *ListWorkflowTemplatesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetMessage(v string) *ListWorkflowTemplatesResponseBody {
	s.Message = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetNextToken(v string) *ListWorkflowTemplatesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetPage(v int64) *ListWorkflowTemplatesResponseBody {
	s.Page = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetPageSize(v int64) *ListWorkflowTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetRequestId(v string) *ListWorkflowTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) SetTotal(v int64) *ListWorkflowTemplatesResponseBody {
	s.Total = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBody) Validate() error {
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

type ListWorkflowTemplatesResponseBodyItems struct {
	// The description of the template.
	//
	// example:
	//
	// example
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The current effective version number. An empty value indicates that only a draft exists and the template cannot be used for work items.
	//
	// example:
	//
	// v1
	EffectiveVersion *string `json:"effectiveVersion,omitempty" xml:"effectiveVersion,omitempty"`
	// The time when the template was last modified. The time is in the ISO 8601 standard.
	//
	// example:
	//
	// example
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The latest version number.
	//
	// example:
	//
	// v2
	LatestVersion *string `json:"latestVersion,omitempty" xml:"latestVersion,omitempty"`
	// The status of the latest version. Valid values:
	//
	// - DRAFT: Draft. The version contains uncommitted edits.
	//
	// - SUBMITTED: Submitted.
	//
	// example:
	//
	// DRAFT
	LatestVersionStatus *string `json:"latestVersionStatus,omitempty" xml:"latestVersionStatus,omitempty"`
	// The name of the template.
	//
	// example:
	//
	// example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of unique owners.
	//
	// example:
	//
	// 1
	OwnerCount *int64 `json:"ownerCount,omitempty" xml:"ownerCount,omitempty"`
	// The number of steps configured with reviewers.
	//
	// example:
	//
	// 1
	ReviewStepCount *int64 `json:"reviewStepCount,omitempty" xml:"reviewStepCount,omitempty"`
	// The number of steps.
	//
	// example:
	//
	// 1
	StepCount *int64 `json:"stepCount,omitempty" xml:"stepCount,omitempty"`
	// The ID of the template.
	//
	// example:
	//
	// example
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s ListWorkflowTemplatesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListWorkflowTemplatesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetEffectiveVersion() *string {
	return s.EffectiveVersion
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetLatestVersion() *string {
	return s.LatestVersion
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetLatestVersionStatus() *string {
	return s.LatestVersionStatus
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetName() *string {
	return s.Name
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetOwnerCount() *int64 {
	return s.OwnerCount
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetReviewStepCount() *int64 {
	return s.ReviewStepCount
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetStepCount() *int64 {
	return s.StepCount
}

func (s *ListWorkflowTemplatesResponseBodyItems) GetTemplateId() *string {
	return s.TemplateId
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetDescription(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetEffectiveVersion(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.EffectiveVersion = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetGmtModified(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetLatestVersion(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.LatestVersion = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetLatestVersionStatus(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.LatestVersionStatus = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetName(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.Name = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetOwnerCount(v int64) *ListWorkflowTemplatesResponseBodyItems {
	s.OwnerCount = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetReviewStepCount(v int64) *ListWorkflowTemplatesResponseBodyItems {
	s.ReviewStepCount = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetStepCount(v int64) *ListWorkflowTemplatesResponseBodyItems {
	s.StepCount = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) SetTemplateId(v string) *ListWorkflowTemplatesResponseBodyItems {
	s.TemplateId = &v
	return s
}

func (s *ListWorkflowTemplatesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
