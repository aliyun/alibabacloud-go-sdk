// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSkillViaOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommitMsg(v string) *UploadSkillViaOssRequest
	GetCommitMsg() *string
	SetNamespaceId(v string) *UploadSkillViaOssRequest
	GetNamespaceId() *string
	SetOssObjectName(v string) *UploadSkillViaOssRequest
	GetOssObjectName() *string
	SetOverwrite(v bool) *UploadSkillViaOssRequest
	GetOverwrite() *bool
	SetTargetVersion(v string) *UploadSkillViaOssRequest
	GetTargetVersion() *string
}

type UploadSkillViaOssRequest struct {
	// The commit message. This parameter is optional.
	CommitMsg *string `json:"CommitMsg,omitempty" xml:"CommitMsg,omitempty"`
	// The workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The OSS object name (path).
	//
	// This parameter is required.
	OssObjectName *string `json:"OssObjectName,omitempty" xml:"OssObjectName,omitempty"`
	// Specifies whether to overwrite an existing skill. Default value: false.
	//
	// example:
	//
	// false
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The target upload version number. This parameter is optional and used as a fallback when the ZIP file contains no version information.
	//
	// example:
	//
	// 1.0.0
	TargetVersion *string `json:"TargetVersion,omitempty" xml:"TargetVersion,omitempty"`
}

func (s UploadSkillViaOssRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadSkillViaOssRequest) GoString() string {
	return s.String()
}

func (s *UploadSkillViaOssRequest) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *UploadSkillViaOssRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *UploadSkillViaOssRequest) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *UploadSkillViaOssRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *UploadSkillViaOssRequest) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *UploadSkillViaOssRequest) SetCommitMsg(v string) *UploadSkillViaOssRequest {
	s.CommitMsg = &v
	return s
}

func (s *UploadSkillViaOssRequest) SetNamespaceId(v string) *UploadSkillViaOssRequest {
	s.NamespaceId = &v
	return s
}

func (s *UploadSkillViaOssRequest) SetOssObjectName(v string) *UploadSkillViaOssRequest {
	s.OssObjectName = &v
	return s
}

func (s *UploadSkillViaOssRequest) SetOverwrite(v bool) *UploadSkillViaOssRequest {
	s.Overwrite = &v
	return s
}

func (s *UploadSkillViaOssRequest) SetTargetVersion(v string) *UploadSkillViaOssRequest {
	s.TargetVersion = &v
	return s
}

func (s *UploadSkillViaOssRequest) Validate() error {
	return dara.Validate(s)
}
