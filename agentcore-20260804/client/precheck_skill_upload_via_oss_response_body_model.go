// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckSkillUploadViaOssResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PrecheckSkillUploadViaOssResponseBodyData) *PrecheckSkillUploadViaOssResponseBody
	GetData() []*PrecheckSkillUploadViaOssResponseBodyData
	SetRequestId(v string) *PrecheckSkillUploadViaOssResponseBody
	GetRequestId() *string
}

type PrecheckSkillUploadViaOssResponseBody struct {
	// The returned data.
	Data []*PrecheckSkillUploadViaOssResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PrecheckSkillUploadViaOssResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrecheckSkillUploadViaOssResponseBody) GoString() string {
	return s.String()
}

func (s *PrecheckSkillUploadViaOssResponseBody) GetData() []*PrecheckSkillUploadViaOssResponseBodyData {
	return s.Data
}

func (s *PrecheckSkillUploadViaOssResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrecheckSkillUploadViaOssResponseBody) SetData(v []*PrecheckSkillUploadViaOssResponseBodyData) *PrecheckSkillUploadViaOssResponseBody {
	s.Data = v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBody) SetRequestId(v string) *PrecheckSkillUploadViaOssResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PrecheckSkillUploadViaOssResponseBodyData struct {
	// The version currently being edited.
	//
	// example:
	//
	// 1.0.0
	EditingVersion *string `json:"editingVersion,omitempty" xml:"editingVersion,omitempty"`
	// The entry path of the Skill package.
	//
	// example:
	//
	// SKILL.md
	EntryPath *string `json:"entryPath,omitempty" xml:"entryPath,omitempty"`
	// Indicates whether the Skill already exists.
	Exists *bool `json:"exists,omitempty" xml:"exists,omitempty"`
	// The highest published version.
	//
	// example:
	//
	// 1.0.0
	MaxPublishedVersion *string `json:"maxPublishedVersion,omitempty" xml:"maxPublishedVersion,omitempty"`
	// The resource owner.
	//
	// example:
	//
	// alice
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The version parsed from the uploaded content.
	//
	// example:
	//
	// 1.0.0
	ParsedVersion *string `json:"parsedVersion,omitempty" xml:"parsedVersion,omitempty"`
	// The pre-check result code.
	//
	// example:
	//
	// VALIDATION_FAILED
	PrecheckCode *string `json:"precheckCode,omitempty" xml:"precheckCode,omitempty"`
	// The reason description.
	//
	// example:
	//
	// Resource processing completed
	Reason *string `json:"reason,omitempty" xml:"reason,omitempty"`
	// The version currently under review.
	//
	// example:
	//
	// 1.0.0
	ReviewingVersion *string `json:"reviewingVersion,omitempty" xml:"reviewingVersion,omitempty"`
	// The Skill name.
	//
	// example:
	//
	// skill-example
	SkillName *string `json:"skillName,omitempty" xml:"skillName,omitempty"`
	// The target version.
	//
	// example:
	//
	// 1.0.0
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// ws-1234567890abcdef
	WorkspaceId *string `json:"workspaceId,omitempty" xml:"workspaceId,omitempty"`
}

func (s PrecheckSkillUploadViaOssResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PrecheckSkillUploadViaOssResponseBodyData) GoString() string {
	return s.String()
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetEditingVersion() *string {
	return s.EditingVersion
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetEntryPath() *string {
	return s.EntryPath
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetExists() *bool {
	return s.Exists
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetMaxPublishedVersion() *string {
	return s.MaxPublishedVersion
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetParsedVersion() *string {
	return s.ParsedVersion
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetPrecheckCode() *string {
	return s.PrecheckCode
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetReason() *string {
	return s.Reason
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetReviewingVersion() *string {
	return s.ReviewingVersion
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetSkillName() *string {
	return s.SkillName
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetEditingVersion(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.EditingVersion = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetEntryPath(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.EntryPath = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetExists(v bool) *PrecheckSkillUploadViaOssResponseBodyData {
	s.Exists = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetMaxPublishedVersion(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.MaxPublishedVersion = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetOwner(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.Owner = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetParsedVersion(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.ParsedVersion = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetPrecheckCode(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.PrecheckCode = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetReason(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.Reason = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetReviewingVersion(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.ReviewingVersion = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetSkillName(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.SkillName = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetTargetVersion(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.TargetVersion = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) SetWorkspaceId(v string) *PrecheckSkillUploadViaOssResponseBodyData {
	s.WorkspaceId = &v
	return s
}

func (s *PrecheckSkillUploadViaOssResponseBodyData) Validate() error {
	return dara.Validate(s)
}
