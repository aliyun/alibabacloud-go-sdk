// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListSkillsResponseBodyData) *ListSkillsResponseBody
	GetData() *ListSkillsResponseBodyData
	SetMaxResults(v int32) *ListSkillsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListSkillsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSkillsResponseBody
	GetRequestId() *string
}

type ListSkillsResponseBody struct {
	// The returned data.
	Data *ListSkillsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The maximum number of entries to return per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// The pagination token for the next page.
	//
	// example:
	//
	// next-page-token
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListSkillsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBody) GetData() *ListSkillsResponseBodyData {
	return s.Data
}

func (s *ListSkillsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListSkillsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSkillsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSkillsResponseBody) SetData(v *ListSkillsResponseBodyData) *ListSkillsResponseBody {
	s.Data = v
	return s
}

func (s *ListSkillsResponseBody) SetMaxResults(v int32) *ListSkillsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListSkillsResponseBody) SetNextToken(v string) *ListSkillsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSkillsResponseBody) SetRequestId(v string) *ListSkillsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSkillsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSkillsResponseBodyData struct {
	// The data on the current page.
	PageItems []*ListSkillsResponseBodyDataPageItems `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 5
	PagesAvailable *int32 `json:"pagesAvailable,omitempty" xml:"pagesAvailable,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListSkillsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodyData) GetPageItems() []*ListSkillsResponseBodyDataPageItems {
	return s.PageItems
}

func (s *ListSkillsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSkillsResponseBodyData) GetPagesAvailable() *int32 {
	return s.PagesAvailable
}

func (s *ListSkillsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSkillsResponseBodyData) SetPageItems(v []*ListSkillsResponseBodyDataPageItems) *ListSkillsResponseBodyData {
	s.PageItems = v
	return s
}

func (s *ListSkillsResponseBodyData) SetPageNumber(v int32) *ListSkillsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListSkillsResponseBodyData) SetPagesAvailable(v int32) *ListSkillsResponseBodyData {
	s.PagesAvailable = &v
	return s
}

func (s *ListSkillsResponseBodyData) SetTotalCount(v int32) *ListSkillsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListSkillsResponseBodyData) Validate() error {
	if s.PageItems != nil {
		for _, item := range s.PageItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSkillsResponseBodyDataPageItems struct {
	// The business tags as a JSON array string.
	//
	// example:
	//
	// Sample property value
	BizTags *string `json:"bizTags,omitempty" xml:"bizTags,omitempty"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The total number of downloads.
	//
	// example:
	//
	// 10
	DownloadCount *int64 `json:"downloadCount,omitempty" xml:"downloadCount,omitempty"`
	// The version that is being edited.
	//
	// example:
	//
	// 1.0.0
	EditingVersion *string `json:"editingVersion,omitempty" xml:"editingVersion,omitempty"`
	// Indicates whether the Skill is enabled.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The source tag.
	//
	// example:
	//
	// UPLOAD
	From *string `json:"from,omitempty" xml:"from,omitempty"`
	// The label mapping.
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The name.
	//
	// example:
	//
	// skill-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of online versions.
	//
	// example:
	//
	// 1
	OnlineCnt *int32 `json:"onlineCnt,omitempty" xml:"onlineCnt,omitempty"`
	// The resource owner.
	//
	// example:
	//
	// alice
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The version that is under review.
	//
	// example:
	//
	// 1.0.0
	ReviewingVersion *string `json:"reviewingVersion,omitempty" xml:"reviewingVersion,omitempty"`
	// The visibility scope.
	//
	// example:
	//
	// PRIVATE
	Scope *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The update time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1787671022000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// Indicates whether the current user has write permissions.
	Writeable *bool `json:"writeable,omitempty" xml:"writeable,omitempty"`
}

func (s ListSkillsResponseBodyDataPageItems) String() string {
	return dara.Prettify(s)
}

func (s ListSkillsResponseBodyDataPageItems) GoString() string {
	return s.String()
}

func (s *ListSkillsResponseBodyDataPageItems) GetBizTags() *string {
	return s.BizTags
}

func (s *ListSkillsResponseBodyDataPageItems) GetDescription() *string {
	return s.Description
}

func (s *ListSkillsResponseBodyDataPageItems) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *ListSkillsResponseBodyDataPageItems) GetEditingVersion() *string {
	return s.EditingVersion
}

func (s *ListSkillsResponseBodyDataPageItems) GetEnable() *bool {
	return s.Enable
}

func (s *ListSkillsResponseBodyDataPageItems) GetFrom() *string {
	return s.From
}

func (s *ListSkillsResponseBodyDataPageItems) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListSkillsResponseBodyDataPageItems) GetName() *string {
	return s.Name
}

func (s *ListSkillsResponseBodyDataPageItems) GetOnlineCnt() *int32 {
	return s.OnlineCnt
}

func (s *ListSkillsResponseBodyDataPageItems) GetOwner() *string {
	return s.Owner
}

func (s *ListSkillsResponseBodyDataPageItems) GetReviewingVersion() *string {
	return s.ReviewingVersion
}

func (s *ListSkillsResponseBodyDataPageItems) GetScope() *string {
	return s.Scope
}

func (s *ListSkillsResponseBodyDataPageItems) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListSkillsResponseBodyDataPageItems) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListSkillsResponseBodyDataPageItems) GetWriteable() *bool {
	return s.Writeable
}

func (s *ListSkillsResponseBodyDataPageItems) SetBizTags(v string) *ListSkillsResponseBodyDataPageItems {
	s.BizTags = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetDescription(v string) *ListSkillsResponseBodyDataPageItems {
	s.Description = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetDownloadCount(v int64) *ListSkillsResponseBodyDataPageItems {
	s.DownloadCount = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetEditingVersion(v string) *ListSkillsResponseBodyDataPageItems {
	s.EditingVersion = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetEnable(v bool) *ListSkillsResponseBodyDataPageItems {
	s.Enable = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetFrom(v string) *ListSkillsResponseBodyDataPageItems {
	s.From = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetLabels(v map[string]*string) *ListSkillsResponseBodyDataPageItems {
	s.Labels = v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetName(v string) *ListSkillsResponseBodyDataPageItems {
	s.Name = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetOnlineCnt(v int32) *ListSkillsResponseBodyDataPageItems {
	s.OnlineCnt = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetOwner(v string) *ListSkillsResponseBodyDataPageItems {
	s.Owner = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetReviewingVersion(v string) *ListSkillsResponseBodyDataPageItems {
	s.ReviewingVersion = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetScope(v string) *ListSkillsResponseBodyDataPageItems {
	s.Scope = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetUpdateTime(v int64) *ListSkillsResponseBodyDataPageItems {
	s.UpdateTime = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetWorkspaceId(v string) *ListSkillsResponseBodyDataPageItems {
	s.WorkspaceId = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) SetWriteable(v bool) *ListSkillsResponseBodyDataPageItems {
	s.Writeable = &v
	return s
}

func (s *ListSkillsResponseBodyDataPageItems) Validate() error {
	return dara.Validate(s)
}
