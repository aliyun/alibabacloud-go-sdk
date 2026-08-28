// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetSkillDetailResponseBodyData) *GetSkillDetailResponseBody
	GetData() *GetSkillDetailResponseBodyData
	SetRequestId(v string) *GetSkillDetailResponseBody
	GetRequestId() *string
}

type GetSkillDetailResponseBody struct {
	// The returned data.
	Data *GetSkillDetailResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetSkillDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillDetailResponseBody) GetData() *GetSkillDetailResponseBodyData {
	return s.Data
}

func (s *GetSkillDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillDetailResponseBody) SetData(v *GetSkillDetailResponseBodyData) *GetSkillDetailResponseBody {
	s.Data = v
	return s
}

func (s *GetSkillDetailResponseBody) SetRequestId(v string) *GetSkillDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillDetailResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSkillDetailResponseBodyData struct {
	// The business tag JSON array string.
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
	// The version currently being edited.
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
	// The update time. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1787671022000
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
	// The list of all version summaries.
	Versions []*GetSkillDetailResponseBodyDataVersions `json:"versions,omitempty" xml:"versions,omitempty" type:"Repeated"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
	// Indicates whether the current user has write permissions.
	Writeable *bool `json:"writeable,omitempty" xml:"writeable,omitempty"`
}

func (s GetSkillDetailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSkillDetailResponseBodyData) GetBizTags() *string {
	return s.BizTags
}

func (s *GetSkillDetailResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetSkillDetailResponseBodyData) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *GetSkillDetailResponseBodyData) GetEditingVersion() *string {
	return s.EditingVersion
}

func (s *GetSkillDetailResponseBodyData) GetEnable() *bool {
	return s.Enable
}

func (s *GetSkillDetailResponseBodyData) GetFrom() *string {
	return s.From
}

func (s *GetSkillDetailResponseBodyData) GetLabels() map[string]*string {
	return s.Labels
}

func (s *GetSkillDetailResponseBodyData) GetName() *string {
	return s.Name
}

func (s *GetSkillDetailResponseBodyData) GetOnlineCnt() *int32 {
	return s.OnlineCnt
}

func (s *GetSkillDetailResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *GetSkillDetailResponseBodyData) GetReviewingVersion() *string {
	return s.ReviewingVersion
}

func (s *GetSkillDetailResponseBodyData) GetScope() *string {
	return s.Scope
}

func (s *GetSkillDetailResponseBodyData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetSkillDetailResponseBodyData) GetVersions() []*GetSkillDetailResponseBodyDataVersions {
	return s.Versions
}

func (s *GetSkillDetailResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSkillDetailResponseBodyData) GetWriteable() *bool {
	return s.Writeable
}

func (s *GetSkillDetailResponseBodyData) SetBizTags(v string) *GetSkillDetailResponseBodyData {
	s.BizTags = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetDescription(v string) *GetSkillDetailResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetDownloadCount(v int64) *GetSkillDetailResponseBodyData {
	s.DownloadCount = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetEditingVersion(v string) *GetSkillDetailResponseBodyData {
	s.EditingVersion = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetEnable(v bool) *GetSkillDetailResponseBodyData {
	s.Enable = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetFrom(v string) *GetSkillDetailResponseBodyData {
	s.From = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetLabels(v map[string]*string) *GetSkillDetailResponseBodyData {
	s.Labels = v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetName(v string) *GetSkillDetailResponseBodyData {
	s.Name = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetOnlineCnt(v int32) *GetSkillDetailResponseBodyData {
	s.OnlineCnt = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetOwner(v string) *GetSkillDetailResponseBodyData {
	s.Owner = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetReviewingVersion(v string) *GetSkillDetailResponseBodyData {
	s.ReviewingVersion = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetScope(v string) *GetSkillDetailResponseBodyData {
	s.Scope = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetUpdateTime(v int64) *GetSkillDetailResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetVersions(v []*GetSkillDetailResponseBodyDataVersions) *GetSkillDetailResponseBodyData {
	s.Versions = v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetWorkspaceId(v string) *GetSkillDetailResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) SetWriteable(v bool) *GetSkillDetailResponseBodyData {
	s.Writeable = &v
	return s
}

func (s *GetSkillDetailResponseBodyData) Validate() error {
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

type GetSkillDetailResponseBodyDataVersions struct {
	// The version author.
	//
	// example:
	//
	// alice
	Author *string `json:"author,omitempty" xml:"author,omitempty"`
	// The commit message.
	//
	// example:
	//
	// Update documentation
	CommitMsg *string `json:"commitMsg,omitempty" xml:"commitMsg,omitempty"`
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
	// The download count.
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

func (s GetSkillDetailResponseBodyDataVersions) String() string {
	return dara.Prettify(s)
}

func (s GetSkillDetailResponseBodyDataVersions) GoString() string {
	return s.String()
}

func (s *GetSkillDetailResponseBodyDataVersions) GetAuthor() *string {
	return s.Author
}

func (s *GetSkillDetailResponseBodyDataVersions) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *GetSkillDetailResponseBodyDataVersions) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetSkillDetailResponseBodyDataVersions) GetDescription() *string {
	return s.Description
}

func (s *GetSkillDetailResponseBodyDataVersions) GetDownloadCount() *int64 {
	return s.DownloadCount
}

func (s *GetSkillDetailResponseBodyDataVersions) GetPublishPipelineInfo() *string {
	return s.PublishPipelineInfo
}

func (s *GetSkillDetailResponseBodyDataVersions) GetStatus() *string {
	return s.Status
}

func (s *GetSkillDetailResponseBodyDataVersions) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetSkillDetailResponseBodyDataVersions) GetVersion() *string {
	return s.Version
}

func (s *GetSkillDetailResponseBodyDataVersions) SetAuthor(v string) *GetSkillDetailResponseBodyDataVersions {
	s.Author = &v
	return s
}

func (s *GetSkillDetailResponseBodyDataVersions) SetCommitMsg(v string) *GetSkillDetailResponseBodyDataVersions {
	s.CommitMsg = &v
	return s
}

func (s *GetSkillDetailResponseBodyDataVersions) SetCreateTime(v int64) *GetSkillDetailResponseBodyDataVersions {
	s.CreateTime = &v
	return s
}

func (s *GetSkillDetailResponseBodyDataVersions) SetDescription(v string) *GetSkillDetailResponseBodyDataVersions {
	s.Description = &v
	return s
}

func (s *GetSkillDetailResponseBodyDataVersions) SetDownloadCount(v int64) *GetSkillDetailResponseBodyDataVersions {
	s.DownloadCount = &v
	return s
}

func (s *GetSkillDetailResponseBodyDataVersions) SetPublishPipelineInfo(v string) *GetSkillDetailResponseBodyDataVersions {
	s.PublishPipelineInfo = &v
	return s
}

func (s *GetSkillDetailResponseBodyDataVersions) SetStatus(v string) *GetSkillDetailResponseBodyDataVersions {
	s.Status = &v
	return s
}

func (s *GetSkillDetailResponseBodyDataVersions) SetUpdateTime(v int64) *GetSkillDetailResponseBodyDataVersions {
	s.UpdateTime = &v
	return s
}

func (s *GetSkillDetailResponseBodyDataVersions) SetVersion(v string) *GetSkillDetailResponseBodyDataVersions {
	s.Version = &v
	return s
}

func (s *GetSkillDetailResponseBodyDataVersions) Validate() error {
	return dara.Validate(s)
}
