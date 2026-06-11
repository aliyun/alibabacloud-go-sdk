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
	Data      *GetSkillDetailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
	BizTags          *string                                   `json:"BizTags,omitempty" xml:"BizTags,omitempty"`
	Description      *string                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	DownloadCount    *int64                                    `json:"DownloadCount,omitempty" xml:"DownloadCount,omitempty"`
	EditingVersion   *string                                   `json:"EditingVersion,omitempty" xml:"EditingVersion,omitempty"`
	Enable           *bool                                     `json:"Enable,omitempty" xml:"Enable,omitempty"`
	From             *string                                   `json:"From,omitempty" xml:"From,omitempty"`
	Labels           map[string]*string                        `json:"Labels,omitempty" xml:"Labels,omitempty"`
	Name             *string                                   `json:"Name,omitempty" xml:"Name,omitempty"`
	NamespaceId      *string                                   `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	OnlineCnt        *int32                                    `json:"OnlineCnt,omitempty" xml:"OnlineCnt,omitempty"`
	Owner            *string                                   `json:"Owner,omitempty" xml:"Owner,omitempty"`
	ReviewingVersion *string                                   `json:"ReviewingVersion,omitempty" xml:"ReviewingVersion,omitempty"`
	Scope            *string                                   `json:"Scope,omitempty" xml:"Scope,omitempty"`
	UpdateTime       *int64                                    `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Versions         []*GetSkillDetailResponseBodyDataVersions `json:"Versions,omitempty" xml:"Versions,omitempty" type:"Repeated"`
	Writeable        *bool                                     `json:"Writeable,omitempty" xml:"Writeable,omitempty"`
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

func (s *GetSkillDetailResponseBodyData) GetNamespaceId() *string {
	return s.NamespaceId
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

func (s *GetSkillDetailResponseBodyData) SetNamespaceId(v string) *GetSkillDetailResponseBodyData {
	s.NamespaceId = &v
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
	Author              *string `json:"Author,omitempty" xml:"Author,omitempty"`
	CommitMsg           *string `json:"CommitMsg,omitempty" xml:"CommitMsg,omitempty"`
	CreateTime          *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description         *string `json:"Description,omitempty" xml:"Description,omitempty"`
	DownloadCount       *int64  `json:"DownloadCount,omitempty" xml:"DownloadCount,omitempty"`
	PublishPipelineInfo *string `json:"PublishPipelineInfo,omitempty" xml:"PublishPipelineInfo,omitempty"`
	Status              *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime          *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	Version             *string `json:"Version,omitempty" xml:"Version,omitempty"`
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
