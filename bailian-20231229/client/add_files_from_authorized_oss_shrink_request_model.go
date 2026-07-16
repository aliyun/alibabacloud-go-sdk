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
	SetOverWriteFileByOssKey(v bool) *AddFilesFromAuthorizedOssShrinkRequest
	GetOverWriteFileByOssKey() *bool
	SetTagsShrink(v string) *AddFilesFromAuthorizedOssShrinkRequest
	GetTagsShrink() *string
}

type AddFilesFromAuthorizedOssShrinkRequest struct {
	// The ID of the category to which the files are imported. This is the `CategoryId` returned by the AddCategory operation. You can also obtain the category ID by clicking the ID icon next to the category name on the <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) - Files tab<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) - Files tab. You can pass in `default` to use the system-created default category.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category type. Optional. Default value: UNSTRUCTURED. Valid values:
	//
	// - UNSTRUCTURED: category for building knowledge base scenarios.
	//
	// <props="china">
	//
	// > This operation does not support importing SESSION_FILE for agent application [conversation interaction](https://www.alibabacloud.com/help/en/model-studio/user-guide/file-interaction). Use the **AddFile*	- operation to upload SESSION_FILE from a local source.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The list of files to import. A maximum of 10 files can be uploaded at a time.
	//
	// > A maximum of 10 files can be uploaded at a time.
	//
	// >
	//
	// This parameter is required.
	FileDetailsShrink *string `json:"FileDetails,omitempty" xml:"FileDetails,omitempty"`
	// The name of the OSS bucket. For more information, see [Buckets](https://help.aliyun.com/document_detail/177682.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// bucketNamexxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The region ID of the OSS bucket. For more information, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
	// Specifies whether to overwrite files with the same OssKey in the category. Default value: false, which means files are not overwritten.
	//
	// example:
	//
	// false
	OverWriteFileByOssKey *bool `json:"OverWriteFileByOssKey,omitempty" xml:"OverWriteFileByOssKey,omitempty"`
	// The list of tags associated with the file. Default value: empty, which means the file is not associated with any tags. A maximum of 10 tags can be specified.
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
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

func (s *AddFilesFromAuthorizedOssShrinkRequest) GetOverWriteFileByOssKey() *bool {
	return s.OverWriteFileByOssKey
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

func (s *AddFilesFromAuthorizedOssShrinkRequest) SetOverWriteFileByOssKey(v bool) *AddFilesFromAuthorizedOssShrinkRequest {
	s.OverWriteFileByOssKey = &v
	return s
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) SetTagsShrink(v string) *AddFilesFromAuthorizedOssShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *AddFilesFromAuthorizedOssShrinkRequest) Validate() error {
	return dara.Validate(s)
}
