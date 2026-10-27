// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateAgentCodePackageFromOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ValidateAgentCodePackageFromOssRequestBody) *ValidateAgentCodePackageFromOssRequest
	GetBody() *ValidateAgentCodePackageFromOssRequestBody
}

type ValidateAgentCodePackageFromOssRequest struct {
	// The request body.
	//
	// This parameter is required.
	Body *ValidateAgentCodePackageFromOssRequestBody `json:"body,omitempty" xml:"body,omitempty" type:"Struct"`
}

func (s ValidateAgentCodePackageFromOssRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateAgentCodePackageFromOssRequest) GoString() string {
	return s.String()
}

func (s *ValidateAgentCodePackageFromOssRequest) GetBody() *ValidateAgentCodePackageFromOssRequestBody {
	return s.Body
}

func (s *ValidateAgentCodePackageFromOssRequest) SetBody(v *ValidateAgentCodePackageFromOssRequestBody) *ValidateAgentCodePackageFromOssRequest {
	s.Body = v
	return s
}

func (s *ValidateAgentCodePackageFromOssRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ValidateAgentCodePackageFromOssRequestBody struct {
	// The name of the source OSS bucket.
	//
	// This parameter is required.
	//
	// example:
	//
	// my-agent-bucket
	OssBucketName *string `json:"ossBucketName,omitempty" xml:"ossBucketName,omitempty"`
	// The full object key of the source code package.
	//
	// This parameter is required.
	//
	// example:
	//
	// packages/main.zip
	OssObjectName *string `json:"ossObjectName,omitempty" xml:"ossObjectName,omitempty"`
}

func (s ValidateAgentCodePackageFromOssRequestBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateAgentCodePackageFromOssRequestBody) GoString() string {
	return s.String()
}

func (s *ValidateAgentCodePackageFromOssRequestBody) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *ValidateAgentCodePackageFromOssRequestBody) GetOssObjectName() *string {
	return s.OssObjectName
}

func (s *ValidateAgentCodePackageFromOssRequestBody) SetOssBucketName(v string) *ValidateAgentCodePackageFromOssRequestBody {
	s.OssBucketName = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssRequestBody) SetOssObjectName(v string) *ValidateAgentCodePackageFromOssRequestBody {
	s.OssObjectName = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssRequestBody) Validate() error {
	return dara.Validate(s)
}
