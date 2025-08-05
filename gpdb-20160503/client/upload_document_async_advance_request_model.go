// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iUploadDocumentAsyncAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChunkOverlap(v int32) *UploadDocumentAsyncAdvanceRequest
	GetChunkOverlap() *int32
	SetChunkSize(v int32) *UploadDocumentAsyncAdvanceRequest
	GetChunkSize() *int32
	SetCollection(v string) *UploadDocumentAsyncAdvanceRequest
	GetCollection() *string
	SetDBInstanceId(v string) *UploadDocumentAsyncAdvanceRequest
	GetDBInstanceId() *string
	SetDocumentLoaderName(v string) *UploadDocumentAsyncAdvanceRequest
	GetDocumentLoaderName() *string
	SetDryRun(v bool) *UploadDocumentAsyncAdvanceRequest
	GetDryRun() *bool
	SetFileName(v string) *UploadDocumentAsyncAdvanceRequest
	GetFileName() *string
	SetFileUrlObject(v io.Reader) *UploadDocumentAsyncAdvanceRequest
	GetFileUrlObject() io.Reader
	SetMetadata(v map[string]interface{}) *UploadDocumentAsyncAdvanceRequest
	GetMetadata() map[string]interface{}
	SetNamespace(v string) *UploadDocumentAsyncAdvanceRequest
	GetNamespace() *string
	SetNamespacePassword(v string) *UploadDocumentAsyncAdvanceRequest
	GetNamespacePassword() *string
	SetOwnerId(v int64) *UploadDocumentAsyncAdvanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UploadDocumentAsyncAdvanceRequest
	GetRegionId() *string
	SetSeparators(v []*string) *UploadDocumentAsyncAdvanceRequest
	GetSeparators() []*string
	SetSplitterModel(v string) *UploadDocumentAsyncAdvanceRequest
	GetSplitterModel() *string
	SetTextSplitterName(v string) *UploadDocumentAsyncAdvanceRequest
	GetTextSplitterName() *string
	SetVlEnhance(v bool) *UploadDocumentAsyncAdvanceRequest
	GetVlEnhance() *bool
	SetZhTitleEnhance(v bool) *UploadDocumentAsyncAdvanceRequest
	GetZhTitleEnhance() *bool
}

type UploadDocumentAsyncAdvanceRequest struct {
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
	// The name of the document loader. You do not need to specify this parameter. A document loader is automatically specified based on the file extension. Valid values:
	//
	// 	- UnstructuredHTMLLoader: `.html`
	//
	// 	- UnstructuredMarkdownLoader: `.md`
	//
	// 	- PyMuPDFLoader: `.pdf`
	//
	// 	- PyPDFLoader: `.pdf`
	//
	// 	- RapidOCRPDFLoader: `.pdf`
	//
	// 	- PDFWithImageRefLoader: `.pdf` (with the text-image association feature)
	//
	// 	- JSONLoader: `.json`
	//
	// 	- CSVLoader: `.csv`
	//
	// 	- RapidOCRLoader: `.png`, `.jpg`, `.jpeg`, and `.bmp`
	//
	// 	- UnstructuredFileLoader: `.eml`, `.msg`, `.rst`, `.txt`, `.docx`, `.epub`, `.odt`, `.pptx`, and `.tsv`
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
	// The file name of the document.
	//
	// >
	//
	// 	- We recommend that you add an extension to the file name. Examples: `.json`, `.md`, and `.pdf`. If you do not add an extension, the default loader designed for unstructured data is used.
	//
	// 	- If an image file is involved, the file name must contain an extension. The following extensions are supported: `.bmp`, `.jpg`, `.jpeg`, `.png`, and `.tiff`.
	//
	// 	- You can use a compressed package to upload images. The package file name must contain an extension. Supported package file extensions: `.tar`, `.gz`, and `.zip`.
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
	FileUrlObject io.Reader `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The metadata. The value of this parameter must be the same as the Metadata parameter that is specified when you call the CreateDocumentCollection operation.
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
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
	Separators    []*string `json:"Separators,omitempty" xml:"Separators,omitempty" type:"Repeated"`
	SplitterModel *string   `json:"SplitterModel,omitempty" xml:"SplitterModel,omitempty"`
	// The name of the splitter. Valid values:
	//
	// 	- **ChineseRecursiveTextSplitter**: inherits from RecursiveCharacterTextSplitter, uses `["\\n\\n","\\n", "。|!|?", "\\.\\s|\\!\\s|\\?\\s", ";|;\\s", ",|,\\s"]` as separators by default, and uses regular expressions to match text.
	//
	// 	- **RecursiveCharacterTextSplitter**: uses `["\\n\\n", "\\n", " ", ""]` as separators by default. The splitter supports splitting code in languages such as `C++, Go, Java, JS, PHP, Proto, Python, RST, Ruby, Rust, Scala, Swift, Markdown, LaTeX, HTML, Sol, and C Sharp`.
	//
	// 	- **SpacyTextSplitter**: uses `\\n\\n` as separators by default and uses the en_core_web_sm model of spaCy. The splitter can obtain better splitting effect.
	//
	// 	- **MarkdownHeaderTextSplitter**: splits text in the `[("#", "head1"), ("##", "head2"), ("###", "head3"), ("####", "head4")]` format. The splitter is suitable for Markdown text.
	//
	// example:
	//
	// ChineseRecursiveTextSplitter
	TextSplitterName *string `json:"TextSplitterName,omitempty" xml:"TextSplitterName,omitempty"`
	VlEnhance        *bool   `json:"VlEnhance,omitempty" xml:"VlEnhance,omitempty"`
	// Specifies whether to enable title enhancement.
	//
	// >  You can determine the title text, mark the text in the metadata, and then combine the text with the upper-level title to implement text enhancement.
	//
	// example:
	//
	// false
	ZhTitleEnhance *bool `json:"ZhTitleEnhance,omitempty" xml:"ZhTitleEnhance,omitempty"`
}

func (s UploadDocumentAsyncAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadDocumentAsyncAdvanceRequest) GoString() string {
	return s.String()
}

func (s *UploadDocumentAsyncAdvanceRequest) GetChunkOverlap() *int32 {
	return s.ChunkOverlap
}

func (s *UploadDocumentAsyncAdvanceRequest) GetChunkSize() *int32 {
	return s.ChunkSize
}

func (s *UploadDocumentAsyncAdvanceRequest) GetCollection() *string {
	return s.Collection
}

func (s *UploadDocumentAsyncAdvanceRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UploadDocumentAsyncAdvanceRequest) GetDocumentLoaderName() *string {
	return s.DocumentLoaderName
}

func (s *UploadDocumentAsyncAdvanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UploadDocumentAsyncAdvanceRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadDocumentAsyncAdvanceRequest) GetFileUrlObject() io.Reader {
	return s.FileUrlObject
}

func (s *UploadDocumentAsyncAdvanceRequest) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *UploadDocumentAsyncAdvanceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *UploadDocumentAsyncAdvanceRequest) GetNamespacePassword() *string {
	return s.NamespacePassword
}

func (s *UploadDocumentAsyncAdvanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UploadDocumentAsyncAdvanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UploadDocumentAsyncAdvanceRequest) GetSeparators() []*string {
	return s.Separators
}

func (s *UploadDocumentAsyncAdvanceRequest) GetSplitterModel() *string {
	return s.SplitterModel
}

func (s *UploadDocumentAsyncAdvanceRequest) GetTextSplitterName() *string {
	return s.TextSplitterName
}

func (s *UploadDocumentAsyncAdvanceRequest) GetVlEnhance() *bool {
	return s.VlEnhance
}

func (s *UploadDocumentAsyncAdvanceRequest) GetZhTitleEnhance() *bool {
	return s.ZhTitleEnhance
}

func (s *UploadDocumentAsyncAdvanceRequest) SetChunkOverlap(v int32) *UploadDocumentAsyncAdvanceRequest {
	s.ChunkOverlap = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetChunkSize(v int32) *UploadDocumentAsyncAdvanceRequest {
	s.ChunkSize = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetCollection(v string) *UploadDocumentAsyncAdvanceRequest {
	s.Collection = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetDBInstanceId(v string) *UploadDocumentAsyncAdvanceRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetDocumentLoaderName(v string) *UploadDocumentAsyncAdvanceRequest {
	s.DocumentLoaderName = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetDryRun(v bool) *UploadDocumentAsyncAdvanceRequest {
	s.DryRun = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetFileName(v string) *UploadDocumentAsyncAdvanceRequest {
	s.FileName = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetFileUrlObject(v io.Reader) *UploadDocumentAsyncAdvanceRequest {
	s.FileUrlObject = v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetMetadata(v map[string]interface{}) *UploadDocumentAsyncAdvanceRequest {
	s.Metadata = v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetNamespace(v string) *UploadDocumentAsyncAdvanceRequest {
	s.Namespace = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetNamespacePassword(v string) *UploadDocumentAsyncAdvanceRequest {
	s.NamespacePassword = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetOwnerId(v int64) *UploadDocumentAsyncAdvanceRequest {
	s.OwnerId = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetRegionId(v string) *UploadDocumentAsyncAdvanceRequest {
	s.RegionId = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetSeparators(v []*string) *UploadDocumentAsyncAdvanceRequest {
	s.Separators = v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetSplitterModel(v string) *UploadDocumentAsyncAdvanceRequest {
	s.SplitterModel = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetTextSplitterName(v string) *UploadDocumentAsyncAdvanceRequest {
	s.TextSplitterName = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetVlEnhance(v bool) *UploadDocumentAsyncAdvanceRequest {
	s.VlEnhance = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) SetZhTitleEnhance(v bool) *UploadDocumentAsyncAdvanceRequest {
	s.ZhTitleEnhance = &v
	return s
}

func (s *UploadDocumentAsyncAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
