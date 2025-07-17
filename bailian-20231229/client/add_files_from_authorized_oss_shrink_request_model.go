// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFilesFromAuthorizedOssShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryId(v string) *AddFilesFromAuthorizedOssShrinkRequest
	GetCategoryId() *string
	SetCategoryType(v string) *AddFilesFromAuthorizedOssShrinkRequest
	GetCategoryType() *string
	SetFileDetailsShrink(v string) *AddFilesFromAuthorizedOssShrinkRequest
	GetFileDetailsShrink() *string
	SetOssBucketName(v string) *AddFilesFromAuthorizedOssShrinkRequest
	GetOssBucketName() *string
	SetOssRegionId(v string) *AddFilesFromAuthorizedOssShrinkRequest
	GetOssRegionId() *string
	SetTagsShrink(v string) *AddFilesFromAuthorizedOssShrinkRequest
	GetTagsShrink() *string
}

type AddFilesFromAuthorizedOssShrinkRequest struct {
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
	FileDetailsShrink *string `json:"FileDetails,omitempty" xml:"FileDetails,omitempty"`
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
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
	TagsShrink  *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
}

func (s AddFilesFromAuthorizedOssShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddFilesFromAuthorizedOssShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) GetCategoryId() *string {
	return s.CategoryId
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) GetCategoryType() *string {
	return s.CategoryType
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) GetFileDetailsShrink() *string {
	return s.FileDetailsShrink
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) GetOssBucketName() *string {
	return s.OssBucketName
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) GetOssRegionId() *string {
	return s.OssRegionId
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) SetCategoryId(v string) *AddFilesFromAuthorizedOssShrinkRequest {
	s.CategoryId = &v
	return s
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) SetCategoryType(v string) *AddFilesFromAuthorizedOssShrinkRequest {
	s.CategoryType = &v
	return s
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) SetFileDetailsShrink(v string) *AddFilesFromAuthorizedOssShrinkRequest {
	s.FileDetailsShrink = &v
	return s
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) SetOssBucketName(v string) *AddFilesFromAuthorizedOssShrinkRequest {
	s.OssBucketName = &v
	return s
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) SetOssRegionId(v string) *AddFilesFromAuthorizedOssShrinkRequest {
	s.OssRegionId = &v
	return s
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) SetTagsShrink(v string) *AddFilesFromAuthorizedOssShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) Validate() error {
	return dara.Validate(s)
}
