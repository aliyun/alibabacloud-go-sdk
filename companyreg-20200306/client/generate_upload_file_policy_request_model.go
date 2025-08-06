// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadFilePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *GenerateUploadFilePolicyRequest
	GetBizType() *string
	SetFileName(v string) *GenerateUploadFilePolicyRequest
	GetFileName() *string
	SetFileType(v string) *GenerateUploadFilePolicyRequest
	GetFileType() *string
}

type GenerateUploadFilePolicyRequest struct {
	// example:
	//
	// esp.isp
	BizType  *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// company_apply_business_license
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GenerateUploadFilePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadFilePolicyRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyRequest) GetBizType() *string {
	return s.BizType
}

func (s *GenerateUploadFilePolicyRequest) GetFileName() *string {
	return s.FileName
}

func (s *GenerateUploadFilePolicyRequest) GetFileType() *string {
	return s.FileType
}

func (s *GenerateUploadFilePolicyRequest) SetBizType(v string) *GenerateUploadFilePolicyRequest {
	s.BizType = &v
	return s
}

func (s *GenerateUploadFilePolicyRequest) SetFileName(v string) *GenerateUploadFilePolicyRequest {
	s.FileName = &v
	return s
}

func (s *GenerateUploadFilePolicyRequest) SetFileType(v string) *GenerateUploadFilePolicyRequest {
	s.FileType = &v
	return s
}

func (s *GenerateUploadFilePolicyRequest) Validate() error {
	return dara.Validate(s)
}
