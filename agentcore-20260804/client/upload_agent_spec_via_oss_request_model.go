// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAgentSpecViaOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *UploadAgentSpecViaOssRequestBody) *UploadAgentSpecViaOssRequest
	GetBody() *UploadAgentSpecViaOssRequestBody
}

type UploadAgentSpecViaOssRequest struct {
	// The request body.
	Body *UploadAgentSpecViaOssRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s UploadAgentSpecViaOssRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSpecViaOssRequest) GoString() string {
	return s.String()
}

func (s *UploadAgentSpecViaOssRequest) GetBody() *UploadAgentSpecViaOssRequestBody {
	return s.Body
}

func (s *UploadAgentSpecViaOssRequest) SetBody(v *UploadAgentSpecViaOssRequestBody) *UploadAgentSpecViaOssRequest {
	s.Body = v
	return s
}

func (s *UploadAgentSpecViaOssRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UploadAgentSpecViaOssRequestBody struct {
	// The OSS object name (full path).
	//
	// This parameter is required.
	//
	// example:
	//
	// agentspec/export/user1/ns1/2024-01-01/uuid/123456.zip
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
	// Specifies whether to overwrite existing drafts. Default value: false.
	//
	// example:
	//
	// false
	Overwrite *bool `json:"overwrite,omitempty" xml:"overwrite,omitempty"`
	// Specifies whether to publish immediately after upload. Default value: false.
	//
	// example:
	//
	// false
	Publish *bool `json:"publish,omitempty" xml:"publish,omitempty"`
}

func (s UploadAgentSpecViaOssRequestBody) String() string {
	return dara.Prettify(s)
}

func (s UploadAgentSpecViaOssRequestBody) GoString() string {
	return s.String()
}

func (s *UploadAgentSpecViaOssRequestBody) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *UploadAgentSpecViaOssRequestBody) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *UploadAgentSpecViaOssRequestBody) GetPublish() *bool {
	return s.Publish
}

func (s *UploadAgentSpecViaOssRequestBody) SetOssObjectName(v string) *UploadAgentSpecViaOssRequestBody {
	s.OssObjectName = &v
	return s
}

func (s *UploadAgentSpecViaOssRequestBody) SetOverwrite(v bool) *UploadAgentSpecViaOssRequestBody {
	s.Overwrite = &v
	return s
}

func (s *UploadAgentSpecViaOssRequestBody) SetPublish(v bool) *UploadAgentSpecViaOssRequestBody {
	s.Publish = &v
	return s
}

func (s *UploadAgentSpecViaOssRequestBody) Validate() error {
	return dara.Validate(s)
}
