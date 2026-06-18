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
	FileDetails []*AddFilesFromAuthorizedOssRequestFileDetails `json:"FileDetails,omitempty" xml:"FileDetails,omitempty" type:"Repeated"`
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
	// The name of the file to import. Note that the suffix must include the file format type.
	//
	// - Supported formats: pdf, docx, doc, txt, md, pptx, ppt, xlsx, xls, html, png, jpg, jpeg, bmp, gif.
	//
	// - The file name length is limited to 4-128 characters.
	//
	// - For file upload requirements and limits, see [Knowledge base quotas and limits](https://help.aliyun.com/document_detail/2880605.html).
	//
	// 	Notice: When the imported file name duplicates an existing file name in the knowledge base, the operation still returns `Status` as `SUCCESS`, but the file will not actually be imported into the knowledge base, and the existing file with the same name remains unchanged. Please ensure that each imported file name is unique.
	//
	// > To add a new data table and upload data, please use the Alibaba Cloud Model Studio console; the API does not support this.
	//
	// This parameter is required.
	//
	// example:
	//
	// this_is_temp_xxxx.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The key name (Key) of the imported file in the OSS Bucket. For details, see [Object naming](https://help.aliyun.com/document_detail/273129.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// root/path/this_is_temp_xxxx.pdf
	OssKey *string `json:"OssKey,omitempty" xml:"OssKey,omitempty"`
	// Parser type. Possible values include:
	//
	// - DOCMIND (Intelligent document parsing)
	//
	// - DOCMIND_DIGITAL (Digital document parsing)
	//
	// - DOCMIND_LLM_VERSION (LLM-based document parsing)
	//
	// - DASH_QWEN_VL_PARSER (Qwen VL parsing)
	//
	// - DOCMIND_LLM_VERSION_MEDIA (Audio/video parsing)
	//
	// - AUTO_SELECT (Automatically select parser)
	//
	// <props="intl">
	//
	// <note>The currently configured parser will be used to parse your uploaded files. If AUTO_SELECT is entered, the parser configured for the corresponding category will be used.</note>
	//
	//
	// <props="china">
	//
	// <note>When CategoryType is UNSTRUCTURED, the parser parses your uploaded files according to the data parsing settings of the current category.</note>
	//
	// <note>When CategoryType is SESSION_FILE, the system uses the default method (not changeable) to parse file content.</note>
	//
	// example:
	//
	// AUTO_SELECT
	Parser *string `json:"Parser,omitempty" xml:"Parser,omitempty"`
	// Parser configuration. Required only when the parser type is set to Qwen VL parsing.
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
	// Model name.
	//
	// example:
	//
	// qwen-vl-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// The prompt used when invoking Qwen VL parsing.
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
