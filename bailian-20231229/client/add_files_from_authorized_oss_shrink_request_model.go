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
	// The ID of the destination category for the imported files. This is the CategoryId returned by the AddCategory operation. You can also obtain the category ID by clicking the ID icon next to the category name on the <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center) > Files tab<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center) > Files tab. You can set this parameter to default to use the default category created by the system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cate_cdd11b1b79a74e8bbd675c356a91ee35xxxxxxxx
	CategoryId *string `json:"CategoryId,omitempty" xml:"CategoryId,omitempty"`
	// The category type. This parameter is optional. Default value: UNSTRUCTURED. Valid values:
	//
	// - UNSTRUCTURED: category, used for building knowledge bases.
	//
	// <props="china">
	//
	// > This operation does not support importing SESSION_FILE for [session interaction](https://www.alibabacloud.com/help/en/model-studio/user-guide/file-interaction) in agent applications. Use the **AddFile*	- operation to upload SESSION_FILE from your local machine.
	//
	// This parameter is required.
	//
	// example:
	//
	// UNSTRUCTURED
	CategoryType *string `json:"CategoryType,omitempty" xml:"CategoryType,omitempty"`
	// The list of files to import. You can upload up to 10 files at a time.
	//
	// > You can upload up to 10 files at a time.
	//
	// >
	//
	// This parameter is required.
	FileDetailsShrink *string `json:"FileDetails,omitempty" xml:"FileDetails,omitempty"`
	// The name of the OSS bucket. For more information, see [Buckets](https://help.aliyun.com/document_detail/177682.html). Before you import data from OSS for the first time, you must complete the authorization as prompted in the Model Studio console and add a tag with the key bailian-datahub-access and the value read to the destination bucket. You must also authorize the OSS service-linked role as prompted in the Model Studio console (service name: datahub.sfm.aliyuncs.com).
	//
	// This parameter is required.
	//
	// example:
	//
	// bucketNamexxxxx
	OssBucketName *string `json:"OssBucketName,omitempty" xml:"OssBucketName,omitempty"`
	// The region ID of the OSS bucket. For more information about how to obtain the ID, see [OSS regions and endpoints](https://help.aliyun.com/document_detail/31837.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	OssRegionId *string `json:"OssRegionId,omitempty" xml:"OssRegionId,omitempty"`
	// Specifies whether to overwrite identical files in the category based on the OssKey. Default value: false, which means identical files are not overwritten.
	//
	// example:
	//
	// false
	OverWriteFileByOssKey *bool `json:"OverWriteFileByOssKey,omitempty" xml:"OverWriteFileByOssKey,omitempty"`
	// The list of tags associated with the file. Default value: empty, which means the file is not associated with any tags. You can specify up to 10 tags.
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
