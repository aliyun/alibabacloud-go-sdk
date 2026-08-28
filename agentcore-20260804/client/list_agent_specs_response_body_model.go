// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentSpecsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListAgentSpecsResponseBodyData) *ListAgentSpecsResponseBody
	GetData() *ListAgentSpecsResponseBodyData
	SetRequestId(v string) *ListAgentSpecsResponseBody
	GetRequestId() *string
}

type ListAgentSpecsResponseBody struct {
	// The returned data.
	Data *ListAgentSpecsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListAgentSpecsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpecsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentSpecsResponseBody) GetData() *ListAgentSpecsResponseBodyData {
	return s.Data
}

func (s *ListAgentSpecsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentSpecsResponseBody) SetData(v *ListAgentSpecsResponseBodyData) *ListAgentSpecsResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentSpecsResponseBody) SetRequestId(v string) *ListAgentSpecsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentSpecsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentSpecsResponseBodyData struct {
	// The data on the current page.
	PageItems []*ListAgentSpecsResponseBodyDataPageItems `json:"pageItems,omitempty" xml:"pageItems,omitempty" type:"Repeated"`
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

func (s ListAgentSpecsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpecsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentSpecsResponseBodyData) GetPageItems() []*ListAgentSpecsResponseBodyDataPageItems {
	return s.PageItems
}

func (s *ListAgentSpecsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAgentSpecsResponseBodyData) GetPagesAvailable() *int32 {
	return s.PagesAvailable
}

func (s *ListAgentSpecsResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAgentSpecsResponseBodyData) SetPageItems(v []*ListAgentSpecsResponseBodyDataPageItems) *ListAgentSpecsResponseBodyData {
	s.PageItems = v
	return s
}

func (s *ListAgentSpecsResponseBodyData) SetPageNumber(v int32) *ListAgentSpecsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListAgentSpecsResponseBodyData) SetPagesAvailable(v int32) *ListAgentSpecsResponseBodyData {
	s.PagesAvailable = &v
	return s
}

func (s *ListAgentSpecsResponseBodyData) SetTotalCount(v int32) *ListAgentSpecsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListAgentSpecsResponseBodyData) Validate() error {
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

type ListAgentSpecsResponseBodyDataPageItems struct {
	// The business tags.
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
	// The download count.
	//
	// example:
	//
	// 10
	DownloadCount *int64 `json:"downloadCount,omitempty" xml:"downloadCount,omitempty"`
	// The version currently being edited.
	//
	// example:
	//
	// 1.0.0
	EditingVersion *string `json:"editingVersion,omitempty" xml:"editingVersion,omitempty"`
	// Indicates whether the AgentSpec is enabled.
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
	// The source.
	//
	// example:
	//
	// UPLOAD
	From *string `json:"from,omitempty" xml:"from,omitempty"`
	// The version labels.
	Labels map[string]*string `json:"labels,omitempty" xml:"labels,omitempty"`
	// The list of MCP server references.
	McpServers []*ListAgentSpecsResponseBodyDataPageItemsMcpServers `json:"mcpServers,omitempty" xml:"mcpServers,omitempty" type:"Repeated"`
	// The name.
	//
	// example:
	//
	// agentspec-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The number of online versions.
	//
	// example:
	//
	// 1
	OnlineCnt *int32 `json:"onlineCnt,omitempty" xml:"onlineCnt,omitempty"`
	// The version currently under review.
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
	// The list of Skill references.
	Skills []*ListAgentSpecsResponseBodyDataPageItemsSkills `json:"skills,omitempty" xml:"skills,omitempty" type:"Repeated"`
	// The update time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1787671022000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s ListAgentSpecsResponseBodyDataPageItems) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpecsResponseBodyDataPageItems) GoString() string {
	return s.String()
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetBizTags() *string {
	return s.BizTags
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetDescription() *string {
	return s.Description
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetEditingVersion() *string {
	return s.EditingVersion
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetEnable() *bool {
	return s.Enable
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetFrom() *string {
	return s.From
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetLabels() map[string]*string {
	return s.Labels
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetMcpServers() []*ListAgentSpecsResponseBodyDataPageItemsMcpServers {
	return s.McpServers
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetName() *string {
	return s.Name
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetOnlineCnt() *int32 {
	return s.OnlineCnt
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetReviewingVersion() *string {
	return s.ReviewingVersion
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetScope() *string {
	return s.Scope
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetSkills() []*ListAgentSpecsResponseBodyDataPageItemsSkills {
	return s.Skills
}

func (s *ListAgentSpecsResponseBodyDataPageItems) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetBizTags(v string) *ListAgentSpecsResponseBodyDataPageItems {
	s.BizTags = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetDescription(v string) *ListAgentSpecsResponseBodyDataPageItems {
	s.Description = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetDownloadCount(v int64) *ListAgentSpecsResponseBodyDataPageItems {
	s.DownloadCount = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetEditingVersion(v string) *ListAgentSpecsResponseBodyDataPageItems {
	s.EditingVersion = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetEnable(v bool) *ListAgentSpecsResponseBodyDataPageItems {
	s.Enable = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetFrom(v string) *ListAgentSpecsResponseBodyDataPageItems {
	s.From = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetLabels(v map[string]*string) *ListAgentSpecsResponseBodyDataPageItems {
	s.Labels = v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetMcpServers(v []*ListAgentSpecsResponseBodyDataPageItemsMcpServers) *ListAgentSpecsResponseBodyDataPageItems {
	s.McpServers = v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetName(v string) *ListAgentSpecsResponseBodyDataPageItems {
	s.Name = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetOnlineCnt(v int32) *ListAgentSpecsResponseBodyDataPageItems {
	s.OnlineCnt = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetReviewingVersion(v string) *ListAgentSpecsResponseBodyDataPageItems {
	s.ReviewingVersion = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetScope(v string) *ListAgentSpecsResponseBodyDataPageItems {
	s.Scope = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetSkills(v []*ListAgentSpecsResponseBodyDataPageItemsSkills) *ListAgentSpecsResponseBodyDataPageItems {
	s.Skills = v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) SetUpdateTime(v int64) *ListAgentSpecsResponseBodyDataPageItems {
	s.UpdateTime = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItems) Validate() error {
	if s.McpServers != nil {
		for _, item := range s.McpServers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Skills != nil {
		for _, item := range s.Skills {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentSpecsResponseBodyDataPageItemsMcpServers struct {
	// The name.
	//
	// example:
	//
	// agentspec-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAgentSpecsResponseBodyDataPageItemsMcpServers) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpecsResponseBodyDataPageItemsMcpServers) GoString() string {
	return s.String()
}

func (s *ListAgentSpecsResponseBodyDataPageItemsMcpServers) GetName() *string {
	return s.Name
}

func (s *ListAgentSpecsResponseBodyDataPageItemsMcpServers) SetName(v string) *ListAgentSpecsResponseBodyDataPageItemsMcpServers {
	s.Name = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItemsMcpServers) Validate() error {
	return dara.Validate(s)
}

type ListAgentSpecsResponseBodyDataPageItemsSkills struct {
	// The name.
	//
	// example:
	//
	// agentspec-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ListAgentSpecsResponseBodyDataPageItemsSkills) String() string {
	return dara.Prettify(s)
}

func (s ListAgentSpecsResponseBodyDataPageItemsSkills) GoString() string {
	return s.String()
}

func (s *ListAgentSpecsResponseBodyDataPageItemsSkills) GetName() *string {
	return s.Name
}

func (s *ListAgentSpecsResponseBodyDataPageItemsSkills) SetName(v string) *ListAgentSpecsResponseBodyDataPageItemsSkills {
	s.Name = &v
	return s
}

func (s *ListAgentSpecsResponseBodyDataPageItemsSkills) Validate() error {
	return dara.Validate(s)
}
