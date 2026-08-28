// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetAgentSpecResponseBodyData) *GetAgentSpecResponseBody
	GetData() *GetAgentSpecResponseBodyData
	SetRequestId(v string) *GetAgentSpecResponseBody
	GetRequestId() *string
}

type GetAgentSpecResponseBody struct {
	// The returned data.
	Data *GetAgentSpecResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetAgentSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentSpecResponseBody) GetData() *GetAgentSpecResponseBodyData {
	return s.Data
}

func (s *GetAgentSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentSpecResponseBody) SetData(v *GetAgentSpecResponseBodyData) *GetAgentSpecResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentSpecResponseBody) SetRequestId(v string) *GetAgentSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentSpecResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentSpecResponseBodyData struct {
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
	// The number of downloads.
	//
	// example:
	//
	// 10
	DownloadCount *int64 `json:"downloadCount,omitempty" xml:"downloadCount,omitempty"`
	// The version that is currently being edited.
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
	// The version that is currently under review.
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
	// The list of version summaries.
	Versions []*GetAgentSpecResponseBodyDataVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
}

func (s GetAgentSpecResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentSpecResponseBodyData) GetBizTags() *string {
	return s.BizTags
}

func (s *GetAgentSpecResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetAgentSpecResponseBodyData) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *GetAgentSpecResponseBodyData) GetEditingVersion() *string {
	return s.EditingVersion
}

func (s *GetAgentSpecResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetAgentSpecResponseBodyData) GetFrom() *string {
	return s.From
}

func (s *GetAgentSpecResponseBodyData) GetLabels() map[string]*string {
	return s.Labels
}

func (s *GetAgentSpecResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetAgentSpecResponseBodyData) GetOnlineCnt() *int32 {
	return s.OnlineCnt
}

func (s *GetAgentSpecResponseBodyData) GetReviewingVersion() *string {
	return s.ReviewingVersion
}

func (s *GetAgentSpecResponseBodyData) GetScope() *string {
	return s.Scope
}

func (s *GetAgentSpecResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetAgentSpecResponseBodyData) GetVersions() []*GetAgentSpecResponseBodyDataVersions {
	return s.Versions
}

func (s *GetAgentSpecResponseBodyData) SetBizTags(v string) *GetAgentSpecResponseBodyData {
	s.BizTags = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetDescription(v string) *GetAgentSpecResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetDownloadCount(v int64) *GetAgentSpecResponseBodyData {
	s.DownloadCount = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetEditingVersion(v string) *GetAgentSpecResponseBodyData {
	s.EditingVersion = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetEnable(v bool) *GetAgentSpecResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetFrom(v string) *GetAgentSpecResponseBodyData {
	s.From = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetLabels(v map[string]*string) *GetAgentSpecResponseBodyData {
	s.Labels = v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetName(v string) *GetAgentSpecResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetOnlineCnt(v int32) *GetAgentSpecResponseBodyData {
	s.OnlineCnt = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetReviewingVersion(v string) *GetAgentSpecResponseBodyData {
	s.ReviewingVersion = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetScope(v string) *GetAgentSpecResponseBodyData {
	s.Scope = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetUpdateTime(v int64) *GetAgentSpecResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetAgentSpecResponseBodyData) SetVersions(v []*GetAgentSpecResponseBodyDataVersions) *GetAgentSpecResponseBodyData {
	s.Versions = v
	return s
}

func (s *GetAgentSpecResponseBodyData) Validate() error {
	if s.Versions != nil {
		for _, item := range s.Versions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAgentSpecResponseBodyDataVersions struct {
	// The version author.
	//
	// example:
	//
	// alice
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// The creation time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1787671022000
	CreateTime *int64 `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// The description.
	//
	// example:
	//
	// A sample description that explains the purpose of the resource
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The number of downloads.
	//
	// example:
	//
	// 10
	DownloadCount *int64 `json:"downloadCount,omitempty" xml:"downloadCount,omitempty"`
	// The publish pipeline information.
	//
	// example:
	//
	// {"status":"SUCCESS"}
	PublishPipelineInfo *string `json:"publishPipelineInfo,omitempty" xml:"publishPipelineInfo,omitempty"`
	// The status.
	//
	// example:
	//
	// READY
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The update time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1787671022000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1.0.0
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s GetAgentSpecResponseBodyDataVersions) String() string {
	return dara.Prettify(s)
}

func (s GetAgentSpecResponseBodyDataVersions) GoString() string {
	return s.String()
}

func (s *GetAgentSpecResponseBodyDataVersions) GetAuthor() *string {
	return s.Author
}

func (s *GetAgentSpecResponseBodyDataVersions) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetAgentSpecResponseBodyDataVersions) GetDescription() *string {
	return s.Description
}

func (s *GetAgentSpecResponseBodyDataVersions) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *GetAgentSpecResponseBodyDataVersions) GetPublishPipelineInfo() *string {
	return s.PublishPipelineInfo
}

func (s *GetAgentSpecResponseBodyDataVersions) GetStatus() *string {
	return s.Status
}

func (s *GetAgentSpecResponseBodyDataVersions) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetAgentSpecResponseBodyDataVersions) GetVersion() *string {
	return s.Version
}

func (s *GetAgentSpecResponseBodyDataVersions) SetAuthor(v string) *GetAgentSpecResponseBodyDataVersions {
	s.Author = &v
	return s
}

func (s *GetAgentSpecResponseBodyDataVersions) SetCreateTime(v int64) *GetAgentSpecResponseBodyDataVersions {
	s.CreateTime = &v
	return s
}

func (s *GetAgentSpecResponseBodyDataVersions) SetDescription(v string) *GetAgentSpecResponseBodyDataVersions {
	s.Description = &v
	return s
}

func (s *GetAgentSpecResponseBodyDataVersions) SetDownloadCount(v int64) *GetAgentSpecResponseBodyDataVersions {
	s.DownloadCount = &v
	return s
}

func (s *GetAgentSpecResponseBodyDataVersions) SetPublishPipelineInfo(v string) *GetAgentSpecResponseBodyDataVersions {
	s.PublishPipelineInfo = &v
	return s
}

func (s *GetAgentSpecResponseBodyDataVersions) SetStatus(v string) *GetAgentSpecResponseBodyDataVersions {
	s.Status = &v
	return s
}

func (s *GetAgentSpecResponseBodyDataVersions) SetUpdateTime(v int64) *GetAgentSpecResponseBodyDataVersions {
	s.UpdateTime = &v
	return s
}

func (s *GetAgentSpecResponseBodyDataVersions) SetVersion(v string) *GetAgentSpecResponseBodyDataVersions {
	s.Version = &v
	return s
}

func (s *GetAgentSpecResponseBodyDataVersions) Validate() error {
	return dara.Validate(s)
}
