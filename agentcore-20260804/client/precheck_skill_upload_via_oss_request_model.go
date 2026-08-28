// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckSkillUploadViaOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *PrecheckSkillUploadViaOssRequestBody) *PrecheckSkillUploadViaOssRequest
	GetBody() *PrecheckSkillUploadViaOssRequestBody
}

type PrecheckSkillUploadViaOssRequest struct {
	// The request body.
	Body *PrecheckSkillUploadViaOssRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s PrecheckSkillUploadViaOssRequest) String() string {
	return dara.Prettify(s)
}

func (s PrecheckSkillUploadViaOssRequest) GoString() string {
	return s.String()
}

func (s *PrecheckSkillUploadViaOssRequest) GetBody() *PrecheckSkillUploadViaOssRequestBody {
	return s.Body
}

func (s *PrecheckSkillUploadViaOssRequest) SetBody(v *PrecheckSkillUploadViaOssRequestBody) *PrecheckSkillUploadViaOssRequest {
	s.Body = v
	return s
}

func (s *PrecheckSkillUploadViaOssRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PrecheckSkillUploadViaOssRequestBody struct {
	// The OSS object name (path).
	//
	// This parameter is required.
	//
	// example:
	//
	// skill/import/user123/ns-123456/2026/04/20/uuid/skills.zip
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
}

func (s PrecheckSkillUploadViaOssRequestBody) String() string {
	return dara.Prettify(s)
}

func (s PrecheckSkillUploadViaOssRequestBody) GoString() string {
	return s.String()
}

func (s *PrecheckSkillUploadViaOssRequestBody) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *PrecheckSkillUploadViaOssRequestBody) SetOssObjectName(v string) *PrecheckSkillUploadViaOssRequestBody {
	s.OssObjectName = &v
	return s
}

func (s *PrecheckSkillUploadViaOssRequestBody) Validate() error {
	return dara.Validate(s)
}
