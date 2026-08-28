// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadSkillViaOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UploadSkillViaOssRequestBody) *UploadSkillViaOssRequest
	GetBody() *UploadSkillViaOssRequestBody
}

type UploadSkillViaOssRequest struct {
	// The request body.
	Body *UploadSkillViaOssRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UploadSkillViaOssRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadSkillViaOssRequest) GoString() string {
	return s.String()
}

func (s *UploadSkillViaOssRequest) GetBody() *UploadSkillViaOssRequestBody {
	return s.Body
}

func (s *UploadSkillViaOssRequest) SetBody(v *UploadSkillViaOssRequestBody) *UploadSkillViaOssRequest {
	s.Body = v
	return s
}

func (s *UploadSkillViaOssRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadSkillViaOssRequestBody struct {
	// The commit message. This parameter is optional.
	//
	// example:
	//
	// Update documentation
	CommitMsg *string `json:"commitMsg,omitempty" xml:"commitMsg,omitempty"`
	// The OSS object name (path).
	//
	// This parameter is required.
	//
	// example:
	//
	// imports/example.zip
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	// Specifies whether to overwrite an existing Skill. Default value: false.
	//
	// example:
	//
	// false
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// The upload version number. This parameter is optional and used as a fallback when the ZIP package contains no version information.
	//
	// example:
	//
	// 1.0.0
	TargetVersion *string `json:"targetVersion,omitempty" xml:"targetVersion,omitempty"`
}

func (s UploadSkillViaOssRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UploadSkillViaOssRequestBody) GoString() string {
	return s.String()
}

func (s *UploadSkillViaOssRequestBody) GetCommitMsg() *string {
	return s.CommitMsg
}

func (s *UploadSkillViaOssRequestBody) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *UploadSkillViaOssRequestBody) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *UploadSkillViaOssRequestBody) GetTargetVersion() *string {
	return s.TargetVersion
}

func (s *UploadSkillViaOssRequestBody) SetCommitMsg(v string) *UploadSkillViaOssRequestBody {
	s.CommitMsg = &v
	return s
}

func (s *UploadSkillViaOssRequestBody) SetOssObjectName(v string) *UploadSkillViaOssRequestBody {
	s.OssObjectName = &v
	return s
}

func (s *UploadSkillViaOssRequestBody) SetOverwrite(v bool) *UploadSkillViaOssRequestBody {
	s.Overwrite = &v
	return s
}

func (s *UploadSkillViaOssRequestBody) SetTargetVersion(v string) *UploadSkillViaOssRequestBody {
	s.TargetVersion = &v
	return s
}

func (s *UploadSkillViaOssRequestBody) Validate() error {
	return dara.Validate(s)
}
