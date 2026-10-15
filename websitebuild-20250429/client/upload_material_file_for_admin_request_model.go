// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMaterialFileForAdminRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBelongId(v string) *UploadMaterialFileForAdminRequest
	GetBelongId() *string
	SetBizId(v string) *UploadMaterialFileForAdminRequest
	GetBizId() *string
	SetFileUrl(v string) *UploadMaterialFileForAdminRequest
	GetFileUrl() *string
	SetName(v string) *UploadMaterialFileForAdminRequest
	GetName() *string
}

type UploadMaterialFileForAdminRequest struct {
	// The ID of the owner.
	//
	// example:
	//
	// 123456
	BelongId *string `json:"BelongId,omitempty" xml:"BelongId,omitempty"`
	// The business instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// WS20250814102215000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The file URL.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://docmind-api-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/4a7f_209934244261306272_14fd429b731245a79f291c64acf3ac77
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The file name.
	//
	// example:
	//
	// Video
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UploadMaterialFileForAdminRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadMaterialFileForAdminRequest) GoString() string {
	return s.String()
}

func (s *UploadMaterialFileForAdminRequest) GetBelongId() *string {
	return s.BelongId
}

func (s *UploadMaterialFileForAdminRequest) GetBizId() *string {
	return s.BizId
}

func (s *UploadMaterialFileForAdminRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadMaterialFileForAdminRequest) GetName() *string {
	return s.Name
}

func (s *UploadMaterialFileForAdminRequest) SetBelongId(v string) *UploadMaterialFileForAdminRequest {
	s.BelongId = &v
	return s
}

func (s *UploadMaterialFileForAdminRequest) SetBizId(v string) *UploadMaterialFileForAdminRequest {
	s.BizId = &v
	return s
}

func (s *UploadMaterialFileForAdminRequest) SetFileUrl(v string) *UploadMaterialFileForAdminRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadMaterialFileForAdminRequest) SetName(v string) *UploadMaterialFileForAdminRequest {
	s.Name = &v
	return s
}

func (s *UploadMaterialFileForAdminRequest) Validate() error {
	return dara.Validate(s)
}
