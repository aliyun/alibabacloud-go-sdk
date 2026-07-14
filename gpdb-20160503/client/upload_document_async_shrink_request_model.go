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
	// The size of overlapping data between consecutive chunks. The maximum value of this parameter cannot be greater than the value of the ChunkSize parameter.
	//
	// >  This parameter prevents context loss caused by data truncation. For example, when you upload long text, you can retain specific overlapping text content between consecutive chunks for better context understanding.
	//
	// example:
	//
	// 50
	ChunkOverlap *int32 `json:"ChunkOverlap,omitempty" xml:"ChunkOverlap,omitempty"`
	// The strategy for processing large data: the size of each chunk when data is split into smaller parts. Maximum value: 2048.
	//
	// example:
	//
	// 250
	ChunkSize *int32 `json:"ChunkSize,omitempty" xml:"ChunkSize,omitempty"`
	// The name of the document collection.
	//
	// >Created by the [CreateDocumentCollection](https://help.aliyun.com/document_detail/2618448.html) operation. You can call the [ListDocumentCollections](https://help.aliyun.com/document_detail/2618452.html) operation to query the created document collections.
	//
	// This parameter is required.
	//
	// example:
	//
	// document
	Collection *string `json:"Collection,omitempty" xml:"Collection,omitempty"`
	// The ID of the instance that has vector engine optimization enabled. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the details of all AnalyticDB for PostgreSQL instances in the target region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The name of the document loader. If you do not specify this parameter, the system automatically selects the corresponding document loader based on the file name extension in the following order. Valid values:
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
	// 	- ADBPGLoader (paid, first 3,000 pages free): .pdf, .doc, .docx, .ppt, .pptx, .xls, .xlsx, .xlsm, .csv, .txt, .jpg, .jpeg, .png, .bmp, .gif, .md, .html, .epub, .mobi, and .rtf
	//
	// example:
	//
	// PyMuPDFLoader
	DocumentLoaderName *string `json:"DocumentLoaderName,omitempty" xml:"DocumentLoaderName,omitempty"`
	// Specifies whether to perform only document understanding and chunking without vectorization and storage. Default value: false.
	//
	// >  You can set this parameter to true to check the chunking results and then optimize as needed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file name of the document.
	//
	// >	- The file name must include file name extension, such as .json, .md, or .pdf.
	//
	// >	- Supported image file extensions include .bmp, .jpg, .jpeg, .png, and .tiff.
	//
	// >	- You can upload images by using an archive. The archive file name must include file name extension. Supported archive extensions include .tar, .gz, and .zip.
	//
	// This parameter is required.
	//
	// example:
	//
	// mydoc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The publicly accessible URL of the document.
	//
	// > Use the SDK to call this operation. The SDK provides a method named UploadDocumentAsyncAdvance that allows you to directly upload local files.
	//
	// If the URL points to an image archive, the number of images in the archive cannot exceed 100.
	//
	// 	Notice:
	//
	// The maximum size of an image uploaded by using multimodal-embedding-v1 is 3 MB.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://xx/mydoc.txt
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The metadata. The value of this parameter must be the same as the Metadata parameter specified when you call the CreateDocumentCollection operation.
	MetadataShrink *string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// The namespace. Default value: public. You can call the CreateNamespace operation to create a namespace and call the ListNamespaces operation to query the list of namespaces.
	//
	// example:
	//
	// mynamespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The password of the namespace. The value is specified by the CreateNamespace operation.
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
	// The separators used to split large data.
	//
	// > 	- This is an important parameter that determines the effectiveness of data chunking. This parameter is related to the splitter specified by the TextSplitterName parameter.
	//
	// >	- In most cases, you do not need to specify this parameter. The server assigns separators based on the value of the TextSplitterName parameter.
	SeparatorsShrink *string `json:"Separators,omitempty" xml:"Separators,omitempty"`
	// The splitting model to use when DocumentLoaderName is set to ADBPGLoader and TextSplitterName is set to LLMSplitter. Default value: qwen3-8b.
	//
	// >
	//
	// > Currently supported splitting models:
	//
	// > qwq-plus, qwq-plus-latest,
	//
	// > qwen-max, qwen-max-latest,
	//
	// > qwen-plus, qwen-plus-latest,
	//
	// > qwen-turbo, qwen-turbo-latest,
	//
	// > qwen3-235b-a22b, qwen3-32b, qwen3-30b-a3b,
	//
	// > qwen3-14b, qwen3-8b, qwen3-4b, qwen3-1.7b, qwen3-0.6b,
	//
	// > qwq-32b
	//
	// > qwen2.5-14b-instruct-1m, qwen2.5-7b-instruct-1m
	//
	// > qwen2.5-72b-instruct, qwen2.5-32b-instruct,
	//
	// > qwen2.5-14b-instruct, qwen2.5-7b-instruct,
	//
	// > qwen2.5-3b-instruct, qwen2.5-1.5b-instruct, qwen2.5-0.5b-instruct
	//
	// example:
	//
	// qwen3-8b
	SplitterModel *string `json:"SplitterModel,omitempty" xml:"SplitterModel,omitempty"`
	// The name of the text splitter. Valid values:
	//
	// 	- **ChineseRecursiveTextSplitter**: inherits from RecursiveCharacterTextSplitter and uses `["
	//
	// ","
	//
	// ", "。|!|?", "\\.\\s|\\!\\s|\\?\\s", ";|;\\s", ",|,\\s"]` as the default separators with regular expression matching.
	//
	// 	- **RecursiveCharacterTextSplitter**: uses `["
	//
	// ", "
	//
	// ", " ", ""]` as the default separators. This splitter supports splitting code in languages such as C++, Go, Java, JS, PHP, Proto, Python, RST, Ruby, Rust, Scala, Swift, Markdown, LaTeX, HTML, Sol, and C Sharp.
	//
	// 	- **SpacyTextSplitter**: uses `
	//
	// ` as the default separator and the spaCy en_core_web_sm model. This splitter provides better splitting results.
	//
	// 	- **MarkdownHeaderTextSplitter**: splits text in the format of [("#", "head1"), ("##", "head2"), ("###", "head3"), ("####", "head4")]. This splitter is suitable for Markdown text.
	//
	// 	- **LLMSplitter**: uses an LLM to split text. The default model is qwen3-8b. This splitter takes effect only when ADBPGLoader is selected as the document loader.
	//
	// example:
	//
	// ChineseRecursiveTextSplitter
	TextSplitterName *string `json:"TextSplitterName,omitempty" xml:"TextSplitterName,omitempty"`
	// Specifies whether to enable VL-enhanced content recognition for complex documents. Default value: false.
	//
	// >
	//
	// > - For complex documents with disorganized layouts and formats, enable VL-enhanced content recognition.
	//
	// > - After VL-enhanced content recognition is enabled, document processing takes longer.
	//
	// > - After VL-enhanced content recognition is enabled, images in the document cannot be stored or recalled.
	//
	// example:
	//
	// false
	VlEnhance *bool `json:"VlEnhance,omitempty" xml:"VlEnhance,omitempty"`
	// Specifies whether to enable title enhancement.
	//
	// >You can identify the title text, mark the text in the metadata, and then combine the text with the upper-level title for text enhancement.
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
