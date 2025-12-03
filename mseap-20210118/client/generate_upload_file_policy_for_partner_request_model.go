// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateUploadFilePolicyForPartnerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *GenerateUploadFilePolicyForPartnerRequest
	GetBizType() *string
	SetFileName(v string) *GenerateUploadFilePolicyForPartnerRequest
	GetFileName() *string
	SetFileType(v string) *GenerateUploadFilePolicyForPartnerRequest
	GetFileType() *string
}

type GenerateUploadFilePolicyForPartnerRequest struct {
	// example:
	//
	// esp.website
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 2024/06/25/website_partner_third_party_leads_02.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// example:
	//
	// website_partner_third_party_leads
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
}

func (s GenerateUploadFilePolicyForPartnerRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateUploadFilePolicyForPartnerRequest) GoString() string {
	return s.String()
}

func (s *GenerateUploadFilePolicyForPartnerRequest) GetBizType() *string {
	return s.BizType
}

func (s *GenerateUploadFilePolicyForPartnerRequest) GetFileName() *string {
	return s.FileName
}

func (s *GenerateUploadFilePolicyForPartnerRequest) GetFileType() *string {
	return s.FileType
}

func (s *GenerateUploadFilePolicyForPartnerRequest) SetBizType(v string) *GenerateUploadFilePolicyForPartnerRequest {
	s.BizType = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerRequest) SetFileName(v string) *GenerateUploadFilePolicyForPartnerRequest {
	s.FileName = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerRequest) SetFileType(v string) *GenerateUploadFilePolicyForPartnerRequest {
	s.FileType = &v
	return s
}

func (s *GenerateUploadFilePolicyForPartnerRequest) Validate() error {
	return dara.Validate(s)
}
