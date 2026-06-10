// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadAppSiteValidationFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *UploadAppSiteValidationFileRequest
	GetBizId() *string
	SetDomain(v string) *UploadAppSiteValidationFileRequest
	GetDomain() *string
	SetFile(v string) *UploadAppSiteValidationFileRequest
	GetFile() *string
	SetFileContent(v string) *UploadAppSiteValidationFileRequest
	GetFileContent() *string
	SetFileType(v string) *UploadAppSiteValidationFileRequest
	GetFileType() *string
	SetSiteHost(v string) *UploadAppSiteValidationFileRequest
	GetSiteHost() *string
}

type UploadAppSiteValidationFileRequest struct {
	// business ID
	//
	// example:
	//
	// WD20250703155602000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// domain name
	//
	// example:
	//
	// yjdw.bpu.edu.cn-waf
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// file name
	//
	// example:
	//
	// {\\"URI\\": \\"oss://imm-test-co-cn-chengdu/zqh/input/Image/indexImage/nn1.jpg\\", \\"LatLong\\":\\"+39.998800,+116.480900\\"}
	File *string `json:"File,omitempty" xml:"File,omitempty"`
	// file content
	//
	// example:
	//
	// abc
	FileContent *string `json:"FileContent,omitempty" xml:"FileContent,omitempty"`
	// file type
	//
	// example:
	//
	// APP_MAPPING
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// site host
	//
	// example:
	//
	// www.aliyun.com
	SiteHost *string `json:"SiteHost,omitempty" xml:"SiteHost,omitempty"`
}

func (s UploadAppSiteValidationFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadAppSiteValidationFileRequest) GoString() string {
	return s.String()
}

func (s *UploadAppSiteValidationFileRequest) GetBizId() *string {
	return s.BizId
}

func (s *UploadAppSiteValidationFileRequest) GetDomain() *string {
	return s.Domain
}

func (s *UploadAppSiteValidationFileRequest) GetFile() *string {
	return s.File
}

func (s *UploadAppSiteValidationFileRequest) GetFileContent() *string {
	return s.FileContent
}

func (s *UploadAppSiteValidationFileRequest) GetFileType() *string {
	return s.FileType
}

func (s *UploadAppSiteValidationFileRequest) GetSiteHost() *string {
	return s.SiteHost
}

func (s *UploadAppSiteValidationFileRequest) SetBizId(v string) *UploadAppSiteValidationFileRequest {
	s.BizId = &v
	return s
}

func (s *UploadAppSiteValidationFileRequest) SetDomain(v string) *UploadAppSiteValidationFileRequest {
	s.Domain = &v
	return s
}

func (s *UploadAppSiteValidationFileRequest) SetFile(v string) *UploadAppSiteValidationFileRequest {
	s.File = &v
	return s
}

func (s *UploadAppSiteValidationFileRequest) SetFileContent(v string) *UploadAppSiteValidationFileRequest {
	s.FileContent = &v
	return s
}

func (s *UploadAppSiteValidationFileRequest) SetFileType(v string) *UploadAppSiteValidationFileRequest {
	s.FileType = &v
	return s
}

func (s *UploadAppSiteValidationFileRequest) SetSiteHost(v string) *UploadAppSiteValidationFileRequest {
	s.SiteHost = &v
	return s
}

func (s *UploadAppSiteValidationFileRequest) Validate() error {
	return dara.Validate(s)
}
