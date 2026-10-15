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
	SetOssKey(v string) *UploadMaterialFileRequest
	GetOssKey() *string
}

type UploadMaterialFileRequest struct {
	// The business ID of the application instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// WS20250801152639000005
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The ID of the parent folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The file path.
	//
	// example:
	//
	// http://docmind-api-cn-hangzhou.oss-cn-hangzhou.aliyuncs.com/4a7f_209934244261306272_14fd429b731245a79f291c64acf3ac77
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The file name.
	//
	// example:
	//
	// TestFile
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The OSS object key.
	//
	// example:
	//
	// backend/detection/objects/r-0008ujvfksltf5j45n81/task-000hckiuwyau0gwn17vq.jpg
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
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

func (s *UploadMaterialFileRequest) GetOssKey() *string {
	return s.OssKey
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

func (s *UploadMaterialFileRequest) SetOssKey(v string) *UploadMaterialFileRequest {
	s.OssKey = &v
	return s
}

func (s *UploadMaterialFileRequest) Validate() error {
	return dara.Validate(s)
}
