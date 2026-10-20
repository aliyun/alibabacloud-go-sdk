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
	FileDetails []*AddFilesFromAuthorizedOssRequestFileDetails `json:"FileDetails,omitempty" xml:"FileDetails,omitempty" type:"Repeated"`
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
	Tags []*string `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	if s.FileDetails != nil {
		for _, item := range s.FileDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddFilesFromAuthorizedOssRequestFileDetails struct {
	// The name of the file to import. Make sure that the file name extension is included.
	//
	// - Supported formats: pdf, docx, doc, txt, md, pptx, ppt, xlsx, xls, html, png, jpg, jpeg, bmp, and gif.
	//
	// - The file name must be 4 to 128 characters in length.
	//
	// - For more information about the limits on file uploads, see [Knowledge base quotas and limits](https://help.aliyun.com/document_detail/2880605.html).
	//
	// 	Notice: If the name of the imported file is the same as the name of an existing file in the knowledge base, the operation still returns SUCCESS for Status, but the file is not actually imported into the knowledge base and the existing file with the same name remains unchanged. Make sure that the name of each imported file is unique.</notice>
	//
	// > To add a data table and upload data, use the Alibaba Cloud Model Studio console. This is not supported by the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// this_is_temp_xxxx.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The key of the imported file in the OSS bucket. For more information, see [Object naming](https://help.aliyun.com/document_detail/273129.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// root/path/this_is_temp_xxxx.pdf
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The parser type. Valid values:
	//
	// - DOCMIND (Intelligent document parsing)
	//
	// - DOCMIND_DIGITAL (Electronic document parsing)
	//
	// - DOCMIND_LLM_VERSION (Large model document parsing)
	//
	// - DASH_QWEN_VL_PARSER (Qwen VL parsing)
	//
	// - DOCMIND_LLM_VERSION_MEDIA (Audio and video parsing)
	//
	// - AUTO_SELECT (Automatic parser selection)
	//
	// <props="intl">
	//
	// <note>The uploaded file is parsed by using the specified parser. If you set this parameter to AUTO_SELECT, the parser configured for the category is used.</note>
	//
	//
	// <props="china">
	//
	// <note>If CategoryType is set to UNSTRUCTURED, the file is parsed based on the data parsing settings of the current category.</note>
	//
	// <note>If CategoryType is set to SESSION_FILE, the system parses the file content by using the default method, which cannot be changed.</note>
	//
	// example:
	//
	// AUTO_SELECT
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// The parser configuration. This parameter is required only when the parser type is set to Qwen VL parsing.
	ParserConfig *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig `json:"ParserConfig,omitempty" xml:"ParserConfig,omitempty" type:"Struct"`
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

func (s *AddFilesFromAuthorizedOssRequestFileDetails) GetParser() *string {
	return s.Parser
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) GetParserConfig() *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig {
	return s.ParserConfig
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) SetFileName(v string) *AddFilesFromAuthorizedOssRequestFileDetails {
	s.FileName = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) SetOssKey(v string) *AddFilesFromAuthorizedOssRequestFileDetails {
	s.OssKey = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) SetParser(v string) *AddFilesFromAuthorizedOssRequestFileDetails {
	s.Parser = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) SetParserConfig(v *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig) *AddFilesFromAuthorizedOssRequestFileDetails {
	s.ParserConfig = v
	return s
}

func (s *AddFilesFromAuthorizedOssRequestFileDetails) Validate() error {
	if s.ParserConfig != nil {
		if err := s.ParserConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AddFilesFromAuthorizedOssRequestFileDetailsParserConfig struct {
	// The model name.
	//
	// example:
	//
	// qwen-vl-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The prompt used when calling Qwen VL parsing.
	//
	// example:
	//
	// # Role
	//
	// You are a professional image content annotator, skilled at identifying and describing the content in images.
	//
	// # Task objective
	//
	// Based on the input image, describe the content in the image in detail
	ModelPrompt *string `json:"ModelPrompt,omitempty" xml:"ModelPrompt,omitempty"`
}

func (s AddFilesFromAuthorizedOssRequestFileDetailsParserConfig) String() string {
	return dara.Prettify(s)
}

func (s AddFilesFromAuthorizedOssRequestFileDetailsParserConfig) GoString() string {
	return s.String()
}

func (s *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig) GetModelName() *string {
	return s.ModelName
}

func (s *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig) GetModelPrompt() *string {
	return s.ModelPrompt
}

func (s *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig) SetModelName(v string) *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig {
	s.ModelName = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig) SetModelPrompt(v string) *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig {
	s.ModelPrompt = &v
	return s
}

func (s *AddFilesFromAuthorizedOssRequestFileDetailsParserConfig) Validate() error {
	return dara.Validate(s)
}
