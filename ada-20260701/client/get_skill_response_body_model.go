// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanDelete(v bool) *GetSkillResponseBody
	GetCanDelete() *bool
	SetCanModify(v bool) *GetSkillResponseBody
	GetCanModify() *bool
	SetCreatedAt(v int64) *GetSkillResponseBody
	GetCreatedAt() *int64
	SetDescription(v string) *GetSkillResponseBody
	GetDescription() *string
	SetDownloadUrl(v string) *GetSkillResponseBody
	GetDownloadUrl() *string
	SetDownloadUrlNetwork(v string) *GetSkillResponseBody
	GetDownloadUrlNetwork() *string
	SetIconUrl(v string) *GetSkillResponseBody
	GetIconUrl() *string
	SetMetadata(v interface{}) *GetSkillResponseBody
	GetMetadata() interface{}
	SetName(v string) *GetSkillResponseBody
	GetName() *string
	SetOfficial(v bool) *GetSkillResponseBody
	GetOfficial() *bool
	SetRequestId(v string) *GetSkillResponseBody
	GetRequestId() *string
	SetSkillId(v string) *GetSkillResponseBody
	GetSkillId() *string
	SetStatus(v string) *GetSkillResponseBody
	GetStatus() *string
	SetUpdatedAt(v int64) *GetSkillResponseBody
	GetUpdatedAt() *int64
	SetVisibility(v string) *GetSkillResponseBody
	GetVisibility() *string
}

type GetSkillResponseBody struct {
	// Indicates whether the current caller can delete the Skill.
	//
	// example:
	//
	// true
	CanDelete *bool `json:"CanDelete,omitempty" xml:"CanDelete,omitempty"`
	// Indicates whether the current caller can modify the Skill.
	//
	// example:
	//
	// true
	CanModify *bool `json:"CanModify,omitempty" xml:"CanModify,omitempty"`
	// The creation time of the Skill, in Unix millisecond timestamp.
	//
	// example:
	//
	// 1760000000000
	CreatedAt *int64 `json:"CreatedAt,omitempty" xml:"CreatedAt,omitempty"`
	// The description in the current Skill main record.
	//
	// example:
	//
	// A Skill for performing code review and risk alerts
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The bundle download URL. Returned when a network type is specified, an accessible Artifact exists, and pre-signing succeeds.
	//
	// example:
	//
	// https://example.com/artifacts/code-review-v2.zip?signature=example
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The network type used to generate the download URL.
	//
	// example:
	//
	// public
	DownloadUrlNetwork *string `json:"DownloadUrlNetwork,omitempty" xml:"DownloadUrlNetwork,omitempty"`
	// The Skill icon URL, sourced from the iconUrl in the metadata. This field may be empty if no icon is configured.
	//
	// example:
	//
	// https://example.com/icons/code-review.png
	IconUrl *string `json:"IconUrl,omitempty" xml:"IconUrl,omitempty"`
	// The Skill metadata, mapped to the metadata field in the backend response.
	//
	// example:
	//
	// {"skillMd":"# Code Review\\nCheck code quality.","artifactId":"artifact_example003"}
	Metadata interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The Skill name.
	//
	// example:
	//
	// code-review
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Indicates whether the Skill is an official Skill.
	//
	// example:
	//
	// false
	Official *bool `json:"Official,omitempty" xml:"Official,omitempty"`
	// The request ID, used for Tracing Analysis and troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Skill ID。
	//
	// example:
	//
	// skill_example123
	SkillId *string `json:"SkillId,omitempty" xml:"SkillId,omitempty"`
	// The current Skill status. Common values are DRAFT and PUBLISHED.
	//
	// example:
	//
	// PUBLISHED
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The update time of the Skill, in Unix millisecond timestamp.
	//
	// example:
	//
	// 1760000300000
	UpdatedAt *int64 `json:"UpdatedAt,omitempty" xml:"UpdatedAt,omitempty"`
	// The visibility of the current Skill. Common values are user and tenant.
	//
	// example:
	//
	// user
	Visibility *string `json:"Visibility,omitempty" xml:"Visibility,omitempty"`
}

func (s GetSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponseBody) GoString() string {
	return s.String()
}

func (s *GetSkillResponseBody) GetCanDelete() *bool {
	return s.CanDelete
}

func (s *GetSkillResponseBody) GetCanModify() *bool {
	return s.CanModify
}

func (s *GetSkillResponseBody) GetCreatedAt() *int64 {
	return s.CreatedAt
}

func (s *GetSkillResponseBody) GetDescription() *string {
	return s.Description
}

func (s *GetSkillResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetSkillResponseBody) GetDownloadUrlNetwork() *string {
	return s.DownloadUrlNetwork
}

func (s *GetSkillResponseBody) GetIconUrl() *string {
	return s.IconUrl
}

func (s *GetSkillResponseBody) GetMetadata() interface{} {
	return s.Metadata
}

func (s *GetSkillResponseBody) GetName() *string {
	return s.Name
}

func (s *GetSkillResponseBody) GetOfficial() *bool {
	return s.Official
}

func (s *GetSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSkillResponseBody) GetSkillId() *string {
	return s.SkillId
}

func (s *GetSkillResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSkillResponseBody) GetUpdatedAt() *int64 {
	return s.UpdatedAt
}

func (s *GetSkillResponseBody) GetVisibility() *string {
	return s.Visibility
}

func (s *GetSkillResponseBody) SetCanDelete(v bool) *GetSkillResponseBody {
	s.CanDelete = &v
	return s
}

func (s *GetSkillResponseBody) SetCanModify(v bool) *GetSkillResponseBody {
	s.CanModify = &v
	return s
}

func (s *GetSkillResponseBody) SetCreatedAt(v int64) *GetSkillResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetSkillResponseBody) SetDescription(v string) *GetSkillResponseBody {
	s.Description = &v
	return s
}

func (s *GetSkillResponseBody) SetDownloadUrl(v string) *GetSkillResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *GetSkillResponseBody) SetDownloadUrlNetwork(v string) *GetSkillResponseBody {
	s.DownloadUrlNetwork = &v
	return s
}

func (s *GetSkillResponseBody) SetIconUrl(v string) *GetSkillResponseBody {
	s.IconUrl = &v
	return s
}

func (s *GetSkillResponseBody) SetMetadata(v interface{}) *GetSkillResponseBody {
	s.Metadata = v
	return s
}

func (s *GetSkillResponseBody) SetName(v string) *GetSkillResponseBody {
	s.Name = &v
	return s
}

func (s *GetSkillResponseBody) SetOfficial(v bool) *GetSkillResponseBody {
	s.Official = &v
	return s
}

func (s *GetSkillResponseBody) SetRequestId(v string) *GetSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSkillResponseBody) SetSkillId(v string) *GetSkillResponseBody {
	s.SkillId = &v
	return s
}

func (s *GetSkillResponseBody) SetStatus(v string) *GetSkillResponseBody {
	s.Status = &v
	return s
}

func (s *GetSkillResponseBody) SetUpdatedAt(v int64) *GetSkillResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetSkillResponseBody) SetVisibility(v string) *GetSkillResponseBody {
	s.Visibility = &v
	return s
}

func (s *GetSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
