// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadMaterialFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *UploadMaterialFileRequest
	GetBizId() *string
	SetDirectoryId(v string) *UploadMaterialFileRequest
	GetDirectoryId() *string
	SetFileUrl(v string) *UploadMaterialFileRequest
	GetFileUrl() *string
	SetName(v string) *UploadMaterialFileRequest
	GetName() *string
}

type UploadMaterialFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WS20250801152639000005
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://docmind-api-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/4a7f_209934244261306272_14fd429b731245a79f291c64acf3ac77
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s UploadMaterialFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadMaterialFileRequest) GoString() string {
	return s.String()
}

func (s *UploadMaterialFileRequest) GetBizId() *string {
	return s.BizId
}

func (s *UploadMaterialFileRequest) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UploadMaterialFileRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadMaterialFileRequest) GetName() *string {
	return s.Name
}

func (s *UploadMaterialFileRequest) SetBizId(v string) *UploadMaterialFileRequest {
	s.BizId = &v
	return s
}

func (s *UploadMaterialFileRequest) SetDirectoryId(v string) *UploadMaterialFileRequest {
	s.DirectoryId = &v
	return s
}

func (s *UploadMaterialFileRequest) SetFileUrl(v string) *UploadMaterialFileRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadMaterialFileRequest) SetName(v string) *UploadMaterialFileRequest {
	s.Name = &v
	return s
}

func (s *UploadMaterialFileRequest) Validate() error {
	return dara.Validate(s)
}
