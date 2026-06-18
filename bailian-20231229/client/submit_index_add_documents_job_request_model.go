// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIndexAddDocumentsJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategoryIds(v []*string) *SubmitIndexAddDocumentsJobRequest
	GetCategoryIds() []*string
	SetChunkMode(v string) *SubmitIndexAddDocumentsJobRequest
	GetChunkMode() *string
	SetChunkSize(v int32) *SubmitIndexAddDocumentsJobRequest
	GetChunkSize() *int32
	SetDocumentIds(v []*string) *SubmitIndexAddDocumentsJobRequest
	GetDocumentIds() []*string
	SetEnableHeaders(v bool) *SubmitIndexAddDocumentsJobRequest
	GetEnableHeaders() *bool
	SetExtra(v *SubmitIndexAddDocumentsJobRequestExtra) *SubmitIndexAddDocumentsJobRequest
	GetExtra() *SubmitIndexAddDocumentsJobRequestExtra
	SetIndexId(v string) *SubmitIndexAddDocumentsJobRequest
	GetIndexId() *string
	SetOverlapSize(v int32) *SubmitIndexAddDocumentsJobRequest
	GetOverlapSize() *int32
	SetSeparator(v string) *SubmitIndexAddDocumentsJobRequest
	GetSeparator() *string
	SetSourceType(v string) *SubmitIndexAddDocumentsJobRequest
	GetSourceType() *string
}

type SubmitIndexAddDocumentsJobRequest struct {
	// A list of category IDs.
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// <props="china">
	//
	// The custom chunking mode. This setting applies only to the documents added in the current job. For more information, see [knowledge base](https://help.aliyun.com/document_detail/2807740.html). Valid values (you can specify only one value):
	//
	// - **length**: Splits the text by a fixed length. The chunking strictly follows the specified `ChunkSize` and `OverlapSize`. If you do not specify these parameters, the system uses the default values: a `ChunkSize` of 500 and an `OverlapSize` of 100. This mode ignores the `Separator` parameter.
	//
	// - **page**: Splits the text by page. If `ChunkSize` is specified, its value is also applied during chunking. If `ChunkSize` is not set, a default value of 500 is used. This mode ignores the `OverlapSize` and `Separator` parameters.
	//
	// - **h1**: Splits the text by level-1 headings. If `ChunkSize` is specified, its value is also applied during chunking. If `ChunkSize` is not set, a default value of 500 is used. This mode ignores the `OverlapSize` and `Separator` parameters.
	//
	// - **h2**: Splits the text by level-2 headings. If `ChunkSize` is specified, its value is also applied during chunking. If `ChunkSize` is not set, a default value of 500 is used. This mode ignores the `OverlapSize` and `Separator` parameters.
	//
	// - **regex**: Splits the text by using a regular expression. The `Separator` parameter is required for this mode. If `ChunkSize` is specified, its value is also applied during chunking. If `ChunkSize` is not set, a default value of 500 is used. This mode ignores the `OverlapSize` parameter.
	//
	// If this parameter is not set, intelligent chunking is used by default.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not specify it.
	//
	// example:
	//
	// length
	ChunkMode *string `json:"ChunkMode,omitempty" xml:"ChunkMode,omitempty"`
	// <props="china">
	//
	// The chunk size. Specifies the maximum number of characters for each text chunk. This setting applies only to the documents added in the current job. If a text segment exceeds this size, the behavior depends on the chunking mode:
	//
	// - **Intelligent chunking*	- (if `ChunkMode` is not set): The text may be truncated.
	//
	// - **Custom chunking*	- (if `ChunkMode` is set): The text is forcibly split.
	//
	// The value must be in the range of [1, 6000]. Defaults to 500 if not specified.
	//
	// For more information, see [knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	//
	// > If you specify a `ChunkSize` less than 100, you must also specify the `OverlapSize` parameter. You can also omit both parameters to use the system defaults.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not specify it.
	//
	// example:
	//
	// 128
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// A list of file IDs.
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// Specifies whether to include Excel file headers. If set to `true`, the knowledge base treats the first row of all .xlsx and .xls files as the header and automatically prepends it to each text chunk (data row). This prevents the large language model (LLM) from misinterpreting the header as a regular data row.
	//
	// > Enable this parameter only if all imported documents are Excel files that contain a header.
	//
	// Valid values:
	//
	// - true: Enabled.
	//
	// - false: Disabled.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableHeaders *bool                                   `json:"EnableHeaders,omitempty" xml:"EnableHeaders,omitempty"`
	Extra         *SubmitIndexAddDocumentsJobRequestExtra `json:"Extra,omitempty" xml:"Extra,omitempty" type:"Struct"`
	// The knowledge base ID. This is the `Data.Id` returned by the **CreateIndex*	- API.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// <props="china">
	//
	// Specifies the number of overlapping characters between adjacent text chunks. This setting applies only to the documents added in the current job. For more information, see [knowledge base](https://help.aliyun.com/document_detail/2807740.html). The value must be in the range of [0, 1024].
	//
	// Defaults to 100 if not specified.
	//
	// > The value of `OverlapSize` must be less than the value of `ChunkSize`. Otherwise, the chunking process may fail.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not specify it.
	//
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// <props="china">
	//
	// The separator for chunking. This parameter is used only when `ChunkMode` is set to **regex**. You can specify a single regular expression (multiple expressions are not supported) to split the file into smaller text chunks. For more information, see [knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	//
	// When you use intelligent chunking (when `ChunkMode` is not specified), leave this parameter empty.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not available. Do not specify it.
	//
	// example:
	//
	// (?<=。)
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The type of the data source. Valid values:
	//
	// - DATA_CENTER_CATEGORY: Imports all documents from specified categories in <props="china">[application data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[application data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). You can import documents from multiple categories.
	//
	// - DATA_CENTER_FILE: Imports specified files from <props="china">[application data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[application data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). You can import multiple files.
	//
	// > If you set this parameter to `DATA_CENTER_CATEGORY`, you must specify the `CategoryIds` parameter. If you set this parameter to `DATA_CENTER_FILE`, you must specify the `DocumentIds` parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// DATA_CENTER_FILE
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
}

func (s SubmitIndexAddDocumentsJobRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobRequest) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobRequest) GetCategoryIds() []*string {
	return s.CategoryIds
}

func (s *SubmitIndexAddDocumentsJobRequest) GetChunkMode() *string {
	return s.ChunkMode
}

func (s *SubmitIndexAddDocumentsJobRequest) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *SubmitIndexAddDocumentsJobRequest) GetDocumentIds() []*string {
	return s.DocumentIds
}

func (s *SubmitIndexAddDocumentsJobRequest) GetEnableHeaders() *bool {
	return s.EnableHeaders
}

func (s *SubmitIndexAddDocumentsJobRequest) GetExtra() *SubmitIndexAddDocumentsJobRequestExtra {
	return s.Extra
}

func (s *SubmitIndexAddDocumentsJobRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *SubmitIndexAddDocumentsJobRequest) GetOverlapSize() *int32 {
	return s.OverlapSize
}

func (s *SubmitIndexAddDocumentsJobRequest) GetSeparator() *string {
	return s.Separator
}

func (s *SubmitIndexAddDocumentsJobRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *SubmitIndexAddDocumentsJobRequest) SetCategoryIds(v []*string) *SubmitIndexAddDocumentsJobRequest {
	s.CategoryIds = v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetChunkMode(v string) *SubmitIndexAddDocumentsJobRequest {
	s.ChunkMode = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetChunkSize(v int32) *SubmitIndexAddDocumentsJobRequest {
	s.ChunkSize = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetDocumentIds(v []*string) *SubmitIndexAddDocumentsJobRequest {
	s.DocumentIds = v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetEnableHeaders(v bool) *SubmitIndexAddDocumentsJobRequest {
	s.EnableHeaders = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetExtra(v *SubmitIndexAddDocumentsJobRequestExtra) *SubmitIndexAddDocumentsJobRequest {
	s.Extra = v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetIndexId(v string) *SubmitIndexAddDocumentsJobRequest {
	s.IndexId = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetOverlapSize(v int32) *SubmitIndexAddDocumentsJobRequest {
	s.OverlapSize = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetSeparator(v string) *SubmitIndexAddDocumentsJobRequest {
	s.Separator = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) SetSourceType(v string) *SubmitIndexAddDocumentsJobRequest {
	s.SourceType = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequest) Validate() error {
	if s.Extra != nil {
		if err := s.Extra.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitIndexAddDocumentsJobRequestExtra struct {
	UniqueId *string `json:"uniqueId,omitempty" xml:"uniqueId,omitempty"`
}

func (s SubmitIndexAddDocumentsJobRequestExtra) String() string {
	return dara.Prettify(s)
}

func (s SubmitIndexAddDocumentsJobRequestExtra) GoString() string {
	return s.String()
}

func (s *SubmitIndexAddDocumentsJobRequestExtra) GetUniqueId() *string {
	return s.UniqueId
}

func (s *SubmitIndexAddDocumentsJobRequestExtra) SetUniqueId(v string) *SubmitIndexAddDocumentsJobRequestExtra {
	s.UniqueId = &v
	return s
}

func (s *SubmitIndexAddDocumentsJobRequestExtra) Validate() error {
	return dara.Validate(s)
}
