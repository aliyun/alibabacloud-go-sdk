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
	// The list of category IDs.
	CategoryIds []*string `json:"CategoryIds,omitempty" xml:"CategoryIds,omitempty" type:"Repeated"`
	// <props="china">
	//
	// Enables custom chunking (applies only to files appended in this request). For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html). Valid values (only one value can be specified at a time):
	//
	// - **length**: chunk by length. Strictly chunks according to the specified `ChunkSize` and `OverlapSize`. If these two parameters are not specified, the system uses default values (`ChunkSize` of 500 and `OverlapSize` of 100). Chunking by length does not support `Separator` (even if specified, it does not take effect).
	//
	// - **page**: chunk by page. If `ChunkSize` is specified, it is also considered during chunking (if not specified, the default value of 500 is used). Chunking by page does not support `OverlapSize` or `Separator` (even if specified, they do not take effect).
	//
	// - **h1**~**h5**: chunk by headings at the corresponding level (`h1` is the first-level heading, and so on, with support up to `h5` fifth-level heading). If `ChunkSize` is specified, it is also considered during chunking (if not specified, the default value of 500 is used). Chunking by heading does not support `OverlapSize` or `Separator` (even if specified, they do not take effect).
	//
	// - **regex**: chunk by regular expression. The `Separator` parameter must be specified. If `ChunkSize` is specified, it is also considered during chunking (if not specified, the default value of 500 is used). Chunking by regex does not support `OverlapSize` (even if specified, it does not take effect).
	//
	// Default value: empty, which uses intelligent chunking.
	//
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify this parameter.
	//
	// example:
	//
	// length
	ChunkMode *string `json:"ChunkMode,omitempty" xml:"ChunkMode,omitempty"`
	// <props="china">
	//
	// The chunk length, which is the maximum number of characters per text chunk (applies only to files appended in this request). When this length is exceeded:
	//
	// - **Intelligent chunking*	- (without specifying `chunkMode`): the text is likely to be truncated.
	//
	// - **Custom chunking*	- (with `chunkMode` specified): the text is forcibly split.
	//
	// Valid values: 1 to 6000. If this parameter is not specified, the default value of 500 is used.
	//
	// For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	//
	// > If you specify `ChunkSize` with a value less than 100, you must also specify `OverlapSize`. You can also leave both parameters unspecified (the system uses default values).
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify this parameter.
	//
	// example:
	//
	// 128
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// The list of file IDs.
	DocumentIds []*string `json:"DocumentIds,omitempty" xml:"DocumentIds,omitempty" type:"Repeated"`
	// Specifies whether to enable header assembly for Excel files. When enabled, the knowledge base treats the first row of all xlsx and xls files as headers and automatically appends them to each text chunk (data row), preventing the large language model from treating headers as regular data rows.
	//
	//
	// > Enable this feature only when all imported files are in xlsx or xls format and contain headers. Otherwise, leave it disabled.
	//
	// >
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
	// The knowledge base ID, which is the `Data.Id` returned by the **CreateIndex*	- operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 79c0alxxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// <props="china">
	//
	// The chunk overlap length (applies only to files appended in this request). It indicates the number of overlapping characters between the current text chunk and the previous text chunk. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html). Valid values: 0 to 1024.
	//
	// If this parameter is not specified, the default value of 100 is used.
	//
	// > The value of `OverlapSize` must be less than the value of `ChunkSize`. Otherwise, chunking exceptions may occur.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify this parameter.
	//
	// example:
	//
	// 16
	OverlapSize *int32 `json:"OverlapSize,omitempty" xml:"OverlapSize,omitempty"`
	// <props="china">
	//
	// The sentence separator, which takes effect only when `chunkMode` is set to **regex*	- (otherwise, it does not take effect even if specified). You can specify a regular expression (only one is supported) to split the file into small text chunks. For more information, see [Knowledge base](https://help.aliyun.com/document_detail/2807740.html).
	//
	// When using intelligent chunking (without specifying `chunkMode`), keep the default empty value.
	//
	//
	//
	// <props="intl">
	//
	// > This parameter is not yet available. Do not specify this parameter.
	//
	// example:
	//
	// (?<=。)
	Separator *string `json:"Separator,omitempty" xml:"Separator,omitempty"`
	// The data source type. Valid values:
	//
	// - DATA_CENTER_CATEGORY: category type. Imports all documents under specified categories in <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). Multiple categories are supported.
	//
	// - DATA_CENTER_FILE: document type. Imports specified files from <props="china">[Application Data](https://bailian.console.aliyun.com/?tab=app#/data-center)<props="intl">[Application Data](https://modelstudio.console.alibabacloud.com/?tab=app#/data-center). Multiple files are supported.
	//
	// > If this parameter is set to DATA_CENTER_CATEGORY, you must specify the `CategoryIds` parameter. If this parameter is set to DATA_CENTER_FILE, you must specify the `DocumentIds` parameter.
	//
	// >
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
