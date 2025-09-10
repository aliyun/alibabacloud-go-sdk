// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilesFromAuthorizedOssRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *AddFilesFromAuthorizedOssRequest
	GetCategoryId() *string
	SetCategoryType(v string) *AddFilesFromAuthorizedOssRequest
	GetCategoryType() *string
	SetFileDetails(v []*AddFilesFromAuthorizedOssRequestFileDetails) *AddFilesFromAuthorizedOssRequest
	GetFileDetails() []*AddFilesFromAuthorizedOssRequestFileDetails
	SetOssBucketName(v string) *AddFilesFromAuthorizedOssRequest
	GetOssBucketName() *string
	SetOssRegionId(v string) *AddFilesFromAuthorizedOssRequest
	GetOssRegionId() *string
	SetOverWriteFileByOssKey(v bool) *AddFilesFromAuthorizedOssRequest
	GetOverWriteFileByOssKey() *bool
	SetTags(v []*string) *AddFilesFromAuthorizedOssRequest
	GetTags() []*string
}

type AddFilesFromAuthorizedOssRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// This parameter is required.
	FileDetails []*AddFilesFromAuthorizedOssRequestFileDetails `json:"FileDetails,omitempty" xml:"FileDetails,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// bucketNamexxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	OssRegionId           *string   `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
	OverWriteFileByOssKey *bool     `json:"OverWriteFileByOssKey,omitempty" xml:"OverWriteFileByOssKey,omitempty"`
	Tags                  []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s AddFilesFromAuthorizedOssRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFilesFromAuthorizedOssRequest) GoString() string {
	return s.String()
}

func (s *AddFilesFromAuthorizedOssRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *AddFilesFromAuthorizedOssRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *AddFilesFromAuthorizedOssRequest) GetFileDetails() []*AddFilesFromAuthorizedOssRequestFileDetails {
	return s.FileDetails
}

func (s *AddFilesFromAuthorizedOssRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *AddFilesFromAuthorizedOssRequest) GetOssRegionId() *string {
	return s.OssRegionId
}

func (s *AddFilesFromAuthorizedOssRequest) GetOverWriteFileByOssKey() *bool {
	return s.OverWriteFileByOssKey
}

func (s *AddFilesFromAuthorizedOssRequest) GetTags() []*string {
	return s.Tags
}

func (s *AddFilesFromAuthorizedOssRequest) SetCategoryId(v string) *AddFilesFromAuthorizedOssRequest {
	s.CategoryId = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequest) SetCategoryType(v string) *AddFilesFromAuthorizedOssRequest {
	s.CategoryType = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequest) SetFileDetails(v []*AddFilesFromAuthorizedOssRequestFileDetails) *AddFilesFromAuthorizedOssRequest {
	s.FileDetails = v
	return s
}

func (s *AddFilesFromAuthorizedOssRequest) SetOssBucketName(v string) *AddFilesFromAuthorizedOssRequest {
	s.OssBucketName = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequest) SetOssRegionId(v string) *AddFilesFromAuthorizedOssRequest {
	s.OssRegionId = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequest) SetOverWriteFileByOssKey(v bool) *AddFilesFromAuthorizedOssRequest {
	s.OverWriteFileByOssKey = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequest) SetTags(v []*string) *AddFilesFromAuthorizedOssRequest {
	s.Tags = v
	return s
}

func (s *AddFilesFromAuthorizedOssRequest) Validate() error {
	return dara.Validate(s)
}

type AddFilesFromAuthorizedOssRequestFileDetails struct {
	// This parameter is required.
	//
	// example:
	//
	// this_is_temp_xxxx.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// root/path/this_is_temp_xxxx.pdf
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
}

func (s AddFilesFromAuthorizedOssRequestFileDetails) String() string {
	return dara.Prettify(s)
}

func (s AddFilesFromAuthorizedOssRequestFileDetails) GoString() string {
	return s.String()
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) GetFileName() *string {
	return s.FileName
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) GetOssKey() *string {
	return s.OssKey
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) SetFileName(v string) *AddFilesFromAuthorizedOssRequestFileDetails {
	s.FileName = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) SetOssKey(v string) *AddFilesFromAuthorizedOssRequestFileDetails {
	s.OssKey = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) Validate() error {
	return dara.Validate(s)
}
