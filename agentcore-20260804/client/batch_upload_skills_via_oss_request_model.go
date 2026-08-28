// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUploadSkillsViaOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *BatchUploadSkillsViaOssRequestBody) *BatchUploadSkillsViaOssRequest
	GetBody() *BatchUploadSkillsViaOssRequestBody
}

type BatchUploadSkillsViaOssRequest struct {
	// The request body.
	Body *BatchUploadSkillsViaOssRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s BatchUploadSkillsViaOssRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUploadSkillsViaOssRequest) GoString() string {
	return s.String()
}

func (s *BatchUploadSkillsViaOssRequest) GetBody() *BatchUploadSkillsViaOssRequestBody {
	return s.Body
}

func (s *BatchUploadSkillsViaOssRequest) SetBody(v *BatchUploadSkillsViaOssRequestBody) *BatchUploadSkillsViaOssRequest {
	s.Body = v
	return s
}

func (s *BatchUploadSkillsViaOssRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchUploadSkillsViaOssRequestBody struct {
	// The OSS object name (path).
	//
	// This parameter is required.
	//
	// example:
	//
	// skill/import/user123/ns-123456/2026/04/20/uuid/skills.zip
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	// Specifies whether to overwrite an existing Skill. Default value: false.
	//
	// example:
	//
	// false
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
}

func (s BatchUploadSkillsViaOssRequestBody) String() string {
	return dara.Prettify(s)
}

func (s BatchUploadSkillsViaOssRequestBody) GoString() string {
	return s.String()
}

func (s *BatchUploadSkillsViaOssRequestBody) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *BatchUploadSkillsViaOssRequestBody) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *BatchUploadSkillsViaOssRequestBody) SetOssObjectName(v string) *BatchUploadSkillsViaOssRequestBody {
	s.OssObjectName = &v
	return s
}

func (s *BatchUploadSkillsViaOssRequestBody) SetOverwrite(v bool) *BatchUploadSkillsViaOssRequestBody {
	s.Overwrite = &v
	return s
}

func (s *BatchUploadSkillsViaOssRequestBody) Validate() error {
	return dara.Validate(s)
}
