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
	FileDetails []*AddFilesFromAuthorizedOssRequestFileDetails `json:"FileDetails,omitempty" xml:"FileDetails,omitempty" type:"Repeated"`
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
	// The name of the file to import. The file name must include the file format extension.
	//
	// - Supported formats: pdf, docx, doc, txt, md, pptx, ppt, xlsx, xls, html, png, jpg, jpeg, bmp, and gif.
	//
	// - The file name must be 4 to 128 characters in length.
	//
	// - For file upload requirements and limits, see [Knowledge base quotas and limits](https://help.aliyun.com/document_detail/2880605.html).
	//
	// 	Notice: If the name of the imported file is the same as an existing file in the knowledge base, the operation still returns a `Status` of `SUCCESS`, but the file is not actually imported. The existing file with the same name remains unchanged. Make sure that each imported file name is unique.
	//
	// > To create a data table and upload data, use the Model Studio console. This is not supported through the API.
	//
	// This parameter is required.
	//
	// example:
	//
	// this_is_temp_xxxx.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The key of the file in the OSS bucket. For more information, see [Object naming conventions](https://help.aliyun.com/document_detail/273129.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// root/path/this_is_temp_xxxx.pdf
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// The parser type. Valid values:
	//
	// - DOCMIND: intelligent document parsing.
	//
	// - DOCMIND_DIGITAL: electronic document parsing.
	//
	// - DOCMIND_LLM_VERSION: LLM-based document parsing.
	//
	// - DASH_QWEN_VL_PARSER: Qwen VL parsing.
	//
	// - DOCMIND_LLM_VERSION_MEDIA: audio and video parsing.
	//
	// - AUTO_SELECT: automatic parser selection.
	//
	// <props="intl">
	//
	// <note>The uploaded file is parsed by using the specified parser. If you set this parameter to AUTO_SELECT, the parser configured for the category is used.</note>
	//
	//
	// <props="china">
	//
	// <note>When CategoryType is UNSTRUCTURED, the parser parses the uploaded file based on the data parsing settings of the current category.</note>
	//
	// <note>When CategoryType is SESSION_FILE, the system uses the default method (which cannot be changed) to parse the file content.</note>
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
	// #角色
	//
	// 你是一个专业的图片内容标注人员，擅长识别并描述出图片中的内容。
	//
	// # 任务目标
	//
	// 请结合输入图片，详细描述图片中的内容。
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
