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
	// Specifies the target category for file import. This is the `CategoryId` returned by the AddCategory operation. You can also obtain the category ID from the <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) - Files tab<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) - Files tab by clicking the ID icon next to the category name. You can also pass in default, which uses the system-created "Default Category".
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// Category type. Optional. The default value is UNSTRUCTURED. Valid values:
	//
	// - UNSTRUCTURED: Category used for building knowledge base scenarios.
	//
	// <props="china">
	//
	// > This operation does not support importing SESSION_FILE used for agent application [session interaction](https://help.aliyun.com/zh/model-studio/user-guide/file-interaction). Please use the **AddFile*	- operation to upload SESSION_FILE from local.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The list of files to import. Up to 10 files can be uploaded at a time.
	//
	// > Up to 10 files can be uploaded at a time.
	//
	// >
	//
	// This parameter is required.
	FileDetailsShrink *string `json:"FileDetails,omitempty" xml:"FileDetails,omitempty"`
	// The OSS Bucket name. For details, see [Buckets](https://help.aliyun.com/document_detail/177682.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// bucketNamexxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The region ID of the OSS Bucket. For how to obtain it, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
	// Whether to overwrite the same file in the category by OssKey. The default value is false, meaning no overwrite.
	//
	// example:
	//
	// false
	OverWriteFileByOssKey *bool `json:"OverWriteFileByOssKey,omitempty" xml:"OverWriteFileByOssKey,omitempty"`
	// The list of tags associated with the file. The default is empty, meaning the file is not associated with any tags. Up to 10 tags can be passed in.
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
