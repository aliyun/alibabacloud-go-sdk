// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadDocumentAsyncShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkOverlap(v int32) *UploadDocumentAsyncShrinkRequest
	GetChunkOverlap() *int32
	SetChunkSize(v int32) *UploadDocumentAsyncShrinkRequest
	GetChunkSize() *int32
	SetCollection(v string) *UploadDocumentAsyncShrinkRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UploadDocumentAsyncShrinkRequest
	GetDBInstanceId() *string
	SetDocumentLoaderName(v string) *UploadDocumentAsyncShrinkRequest
	GetDocumentLoaderName() *string
	SetDryRun(v bool) *UploadDocumentAsyncShrinkRequest
	GetDryRun() *bool
	SetFileName(v string) *UploadDocumentAsyncShrinkRequest
	GetFileName() *string
	SetFileUrl(v string) *UploadDocumentAsyncShrinkRequest
	GetFileUrl() *string
	SetMetadataShrink(v string) *UploadDocumentAsyncShrinkRequest
	GetMetadataShrink() *string
	SetNamespace(v string) *UploadDocumentAsyncShrinkRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UploadDocumentAsyncShrinkRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UploadDocumentAsyncShrinkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UploadDocumentAsyncShrinkRequest
	GetRegionId() *string
	SetSeparatorsShrink(v string) *UploadDocumentAsyncShrinkRequest
	GetSeparatorsShrink() *string
	SetSplitterModel(v string) *UploadDocumentAsyncShrinkRequest
	GetSplitterModel() *string
	SetTextSplitterName(v string) *UploadDocumentAsyncShrinkRequest
	GetTextSplitterName() *string
	SetVlEnhance(v bool) *UploadDocumentAsyncShrinkRequest
	GetVlEnhance() *bool
	SetZhTitleEnhance(v bool) *UploadDocumentAsyncShrinkRequest
	GetZhTitleEnhance() *bool
}

type UploadDocumentAsyncShrinkRequest struct {
	// The size of data that is overlapped between consecutive chunks. The maximum value of this parameter cannot be greater than the value of the ChunkSize parameter.
	//
	// >  This parameter is used to prevent context missing that may occur due to data truncation. For example, when you upload a long text, you can retain specific overlapped text content between consecutive chunks to better understand the context.
	//
	// example:
	//
	// 50
	ChunkOverlap *int32 `json:"ChunkOverlap,omitempty" xml:"ChunkOverlap,omitempty"`
	// Strategy for processing large data: the size of each chunk when the data is split into smaller parts. Maximum value is 2048.
	//
	// example:
	//
	// 250
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// The name of the document library.
	//
	// > Created by the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) API. You can call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) API to view the document libraries that have already been created.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// Instance ID with vector engine optimization acceleration enabled. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) API to view details of all AnalyticDB PostgreSQL instances in the target region, including the instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// Specifies the document loader to use for processing the file. If this parameter is omitted, the system automatically selects a loader based on the file\\"s extension.Valid Values:[List of valid loader names would go here] Valid values:
	//
	// 	- UnstructuredHTMLLoader: .html
	//
	// 	- UnstructuredMarkdownLoader: .md
	//
	// 	- PyMuPDFLoader: .pdf
	//
	// 	- PyPDFLoader: .pdf
	//
	// 	- RapidOCRPDFLoader: .pdf
	//
	// 	- PDFWithImageRefLoader: .pdf (with the text-image association feature)
	//
	// 	- JSONLoader: .json
	//
	// 	- CSVLoader: .csv
	//
	// 	- RapidOCRLoader: .png, .jpg, .jpeg, and .bmp
	//
	// 	- UnstructuredFileLoader: .eml, .msg, .rst, .txt, .docx, .epub, .odt, .pptx, and .tsv
	//
	// 	- ADBPGLoader (free of charge for the first 3,000 pages): .pdf, .doc, .docx, .ppt, .pptx, .xls, .xlsx, .xlsm, .csv, .txt, .jpg, .jpeg, .png, .bmp, .gif, .md, .html, .epub, .mobi, and .rtf
	//
	// example:
	//
	// PyMuPDFLoader
	DocumentLoaderName *string `json:"DocumentLoaderName,omitempty" xml:"DocumentLoaderName,omitempty"`
	// Specifies whether to perform only document understanding and chunking, but not vectorization and storage. Default value: false.
	//
	// >  You can set this parameter to true, check the chunking effect, and then perform optimization if needed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The name of the file being uploaded.
	//
	// >
	//
	// 	- File name: .json, .md, and .pdf.
	//
	// 	- Images: .bmp,. jpg,. jpeg,. png, and. tiff.
	//
	// 	- Compressed packages. The package file name must contain an extension: .tar, .gz, and .zip.
	//
	// This parameter is required.
	//
	// example:
	//
	// mydoc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The URL of the publicly accessible document.
	//
	// >  > - It is recommended to call this interface using the SDK, which provides a method called UploadDocumentAsyncAdvance for directly uploading local files. > - If the URL points to an image archive, the number of images in the archive should not exceed 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xx/mydoc.txt
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The metadata. The value of this parameter must be the same as the Metadata parameter that is specified when you call the CreateDocumentCollection operation.
	MetadataShrink *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Namespace, default is public. You can create one through the CreateNamespace interface and view the list via the ListNamespaces interface.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password corresponding to the namespace.  > This value is specified by the CreateNamespace interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// testpassword
	NamespacePassword *string `json:"NamespacePassword,omitempty" xml:"NamespacePassword,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The separators that are used to split large amounts of data.
	//
	// >
	//
	// 	- This is an important parameter that determines the chunking effect. This parameter is related to the splitter that is specified by the TextSplitterName parameter.
	//
	// 	- In most cases, you do not need to specify this parameter. The server assigns separators based on the value of the TextSplitterName parameter.
	SeparatorsShrink *string `json:"Separators,omitempty" xml:"Separators,omitempty"`
	// When DocumentLoaderName is set to ADBPGLoader and TextSplitterName is set to LLMSplitter, you can specify the splitting model. Default Value: qwen3-8b.
	//
	// >  Supported splitting models: qwq-plus, qwq-plus-latest, qwen-max, qwen-max-latest, qwen-plus, qwen-plus-latest, qwen-turbo, qwen-turbo-latest, qwen3-235b-a22b, qwen3-32b,qwen3-30b-a3b, qwen3-14b, qwen3-8b, qwen3-4b, qwen3-1.7b, qwen3-0.6b, qwq-32b qwen2.5-14b-instruct-1m, qwen2.5-7b-instruct-1m, qwen2.5-72b-Instruct, qwen2.5-32b-Instruct, qwen2.5-14b-Instruct, qwen2.5-7b-Instruct, qwen2.5-3b-instruct, qwen2.5-1.5b-instruct, qwen2.5-0.5b-instruct.
	//
	// example:
	//
	// qwen3-8b
	SplitterModel *string `json:"SplitterModel,omitempty" xml:"SplitterModel,omitempty"`
	// The name of the separator. Valid values:
	//
	// 	- **ChineseRecursiveTextSplitter**: Inherits from RecursiveCharacterTextSplitter and, by default, uses the delimiters` ["\\n\\n","\\n", "。 |! |?", "\\.\\s|\\! \\s|\\?\\s", ";|;\\s", ",|,\\s"]  `, employing regular expressions to match text.
	//
	// 	- **RecursiveCharacterTextSplitter**: Uses the delimiters `["\\n\\n", "\\n", " ", ""]` by default. The splitter supports splitting code in languages such as C++, Go, Java, JS, PHP, Proto, Python, RST, Ruby, Rust, Scala, Swift, Markdown, LaTeX, HTML, Sol, and C Sharp.
	//
	// 	- **SpacyTextSplitter**: Uses the delimiters `\\n\\n` by default and leverages the spaCy en_core_web_sm model. The splitter can achieve better text splitting performance.
	//
	// 	- **MarkdownHeaderTextSplitter**: Splits text in the [("#", "head1"), ("##", "head2"), ("###", "head3"), ("####", "head4") format. This splitter works well with Markdown text.
	//
	// 	- **LLMSplitter**: Use LLM to split text. The default model is qwen3-8b. Currently, this splitter works only when ADBPGLoader is selected.
	//
	// example:
	//
	// ChineseRecursiveTextSplitter
	TextSplitterName *string `json:"TextSplitterName,omitempty" xml:"TextSplitterName,omitempty"`
	// Specifies whether to enable VL-enhanced content recognition for complex documents. Default value: false.
	//
	// >
	//
	// 	- For complex documents with confusing typesetting and formatting, we recommend that you enable VL-enhanced content recognition.
	//
	// 	- Document processing time is longer after VL-enhanced content recognition is enabled.
	//
	// 	- After VL-enhanced content recognition is enabled, images in documents cannot be stored or recalled.
	//
	// example:
	//
	// false
	VlEnhance *bool `json:"VlEnhance,omitempty" xml:"VlEnhance,omitempty"`
	// Specifies whether to enable title enhancement.
	//
	// >  You can determine the title text, mark the text in the metadata, and then combine the text with the upper-level title to implement text enhancement.
	//
	// example:
	//
	// false
	ZhTitleEnhance *bool `json:"ZhTitleEnhance,omitempty" xml:"ZhTitleEnhance,omitempty"`
}

func (s UploadDocumentAsyncShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentAsyncShrinkRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentAsyncShrinkRequest) GetChunkOverlap() *int32 {
	return s.ChunkOverlap
}

func (s *UploadDocumentAsyncShrinkRequest) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *UploadDocumentAsyncShrinkRequest) GetCollection() *string {
	return s.Collection
}

func (s *UploadDocumentAsyncShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UploadDocumentAsyncShrinkRequest) GetDocumentLoaderName() *string {
	return s.DocumentLoaderName
}

func (s *UploadDocumentAsyncShrinkRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UploadDocumentAsyncShrinkRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadDocumentAsyncShrinkRequest) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadDocumentAsyncShrinkRequest) GetMetadataShrink() *string {
	return s.MetadataShrink
}

func (s *UploadDocumentAsyncShrinkRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UploadDocumentAsyncShrinkRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UploadDocumentAsyncShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UploadDocumentAsyncShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadDocumentAsyncShrinkRequest) GetSeparatorsShrink() *string {
	return s.SeparatorsShrink
}

func (s *UploadDocumentAsyncShrinkRequest) GetSplitterModel() *string {
	return s.SplitterModel
}

func (s *UploadDocumentAsyncShrinkRequest) GetTextSplitterName() *string {
	return s.TextSplitterName
}

func (s *UploadDocumentAsyncShrinkRequest) GetVlEnhance() *bool {
	return s.VlEnhance
}

func (s *UploadDocumentAsyncShrinkRequest) GetZhTitleEnhance() *bool {
	return s.ZhTitleEnhance
}

func (s *UploadDocumentAsyncShrinkRequest) SetChunkOverlap(v int32) *UploadDocumentAsyncShrinkRequest {
	s.ChunkOverlap = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetChunkSize(v int32) *UploadDocumentAsyncShrinkRequest {
	s.ChunkSize = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetCollection(v string) *UploadDocumentAsyncShrinkRequest {
	s.Collection = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetDBInstanceId(v string) *UploadDocumentAsyncShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetDocumentLoaderName(v string) *UploadDocumentAsyncShrinkRequest {
	s.DocumentLoaderName = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetDryRun(v bool) *UploadDocumentAsyncShrinkRequest {
	s.DryRun = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetFileName(v string) *UploadDocumentAsyncShrinkRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetFileUrl(v string) *UploadDocumentAsyncShrinkRequest {
	s.FileUrl = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetMetadataShrink(v string) *UploadDocumentAsyncShrinkRequest {
	s.MetadataShrink = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetNamespace(v string) *UploadDocumentAsyncShrinkRequest {
	s.Namespace = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetNamespacePassword(v string) *UploadDocumentAsyncShrinkRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetOwnerId(v int64) *UploadDocumentAsyncShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetRegionId(v string) *UploadDocumentAsyncShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetSeparatorsShrink(v string) *UploadDocumentAsyncShrinkRequest {
	s.SeparatorsShrink = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetSplitterModel(v string) *UploadDocumentAsyncShrinkRequest {
	s.SplitterModel = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetTextSplitterName(v string) *UploadDocumentAsyncShrinkRequest {
	s.TextSplitterName = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetVlEnhance(v bool) *UploadDocumentAsyncShrinkRequest {
	s.VlEnhance = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) SetZhTitleEnhance(v bool) *UploadDocumentAsyncShrinkRequest {
	s.ZhTitleEnhance = &v
	return s
}

func (s *UploadDocumentAsyncShrinkRequest) Validate() error {
	return dara.Validate(s)
}
