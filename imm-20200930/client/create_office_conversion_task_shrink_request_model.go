// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOfficeConversionTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCredentialConfigShrink(v string) *CreateOfficeConversionTaskShrinkRequest
	GetCredentialConfigShrink() *string
	SetEndPage(v int64) *CreateOfficeConversionTaskShrinkRequest
	GetEndPage() *int64
	SetFirstPage(v bool) *CreateOfficeConversionTaskShrinkRequest
	GetFirstPage() *bool
	SetFitToHeight(v bool) *CreateOfficeConversionTaskShrinkRequest
	GetFitToHeight() *bool
	SetFitToWidth(v bool) *CreateOfficeConversionTaskShrinkRequest
	GetFitToWidth() *bool
	SetHoldLineFeed(v bool) *CreateOfficeConversionTaskShrinkRequest
	GetHoldLineFeed() *bool
	SetImageDPI(v int64) *CreateOfficeConversionTaskShrinkRequest
	GetImageDPI() *int64
	SetLongPicture(v bool) *CreateOfficeConversionTaskShrinkRequest
	GetLongPicture() *bool
	SetLongText(v bool) *CreateOfficeConversionTaskShrinkRequest
	GetLongText() *bool
	SetMaxSheetColumn(v int64) *CreateOfficeConversionTaskShrinkRequest
	GetMaxSheetColumn() *int64
	SetMaxSheetRow(v int64) *CreateOfficeConversionTaskShrinkRequest
	GetMaxSheetRow() *int64
	SetNotificationShrink(v string) *CreateOfficeConversionTaskShrinkRequest
	GetNotificationShrink() *string
	SetPages(v string) *CreateOfficeConversionTaskShrinkRequest
	GetPages() *string
	SetPaperHorizontal(v bool) *CreateOfficeConversionTaskShrinkRequest
	GetPaperHorizontal() *bool
	SetPaperSize(v string) *CreateOfficeConversionTaskShrinkRequest
	GetPaperSize() *string
	SetPassword(v string) *CreateOfficeConversionTaskShrinkRequest
	GetPassword() *string
	SetProjectName(v string) *CreateOfficeConversionTaskShrinkRequest
	GetProjectName() *string
	SetQuality(v int64) *CreateOfficeConversionTaskShrinkRequest
	GetQuality() *int64
	SetScalePercentage(v int64) *CreateOfficeConversionTaskShrinkRequest
	GetScalePercentage() *int64
	SetSheetCount(v int64) *CreateOfficeConversionTaskShrinkRequest
	GetSheetCount() *int64
	SetSheetIndex(v int64) *CreateOfficeConversionTaskShrinkRequest
	GetSheetIndex() *int64
	SetShowComments(v bool) *CreateOfficeConversionTaskShrinkRequest
	GetShowComments() *bool
	SetSourceType(v string) *CreateOfficeConversionTaskShrinkRequest
	GetSourceType() *string
	SetSourceURI(v string) *CreateOfficeConversionTaskShrinkRequest
	GetSourceURI() *string
	SetSourcesShrink(v string) *CreateOfficeConversionTaskShrinkRequest
	GetSourcesShrink() *string
	SetStartPage(v int64) *CreateOfficeConversionTaskShrinkRequest
	GetStartPage() *int64
	SetTagsShrink(v string) *CreateOfficeConversionTaskShrinkRequest
	GetTagsShrink() *string
	SetTargetType(v string) *CreateOfficeConversionTaskShrinkRequest
	GetTargetType() *string
	SetTargetURI(v string) *CreateOfficeConversionTaskShrinkRequest
	GetTargetURI() *string
	SetTargetURIPrefix(v string) *CreateOfficeConversionTaskShrinkRequest
	GetTargetURIPrefix() *string
	SetTrimPolicyShrink(v string) *CreateOfficeConversionTaskShrinkRequest
	GetTrimPolicyShrink() *string
	SetUserData(v string) *CreateOfficeConversionTaskShrinkRequest
	GetUserData() *string
}

type CreateOfficeConversionTaskShrinkRequest struct {
	// **If you do not have special requirements, leave this parameter empty.**
	//
	// The chained authorization configuration. This parameter is not required. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The end page for the document conversion. The default value is -1, which indicates that all pages from the start page to the last page are converted.
	//
	// > - If the source file is a spreadsheet, you must specify the worksheet number (\\`SheetIndex\\`).
	//
	// >
	//
	// > - If the document has many pages, we recommend that you convert them in batches. Otherwise, the conversion may time out.
	//
	// >
	//
	// > - This parameter takes effect only when you convert the document to images. It does not take effect when you convert the document to a PDF file or a text file.
	//
	// example:
	//
	// -1
	EndPage *int64 `json:"EndPage,omitempty" xml:"EndPage,omitempty"`
	// When you convert a spreadsheet document to images, specifies whether to return only the first image of the conversion result. The number of rows and columns in the image is the result of automatic splitting. Valid values:
	//
	// - false (default): No. All images are returned.
	//
	// - true: Yes. Only the first image is returned. This is used to extract a thumbnail.
	//
	// > This parameter takes effect only if you set the **LongPicture*	- parameter to `true`.
	//
	// example:
	//
	// false
	FirstPage *bool `json:"FirstPage,omitempty" xml:"FirstPage,omitempty"`
	// When you convert a spreadsheet document to images or a PDF file, specifies whether to render all rows on a single image or PDF page. Valid values:
	//
	// - false (default): No. The content is rendered on multiple images or PDF pages.
	//
	// - true: Yes. The content is rendered on a single image or PDF page.
	//
	// example:
	//
	// false
	FitToHeight *bool `json:"FitToHeight,omitempty" xml:"FitToHeight,omitempty"`
	// When you convert a spreadsheet document to images or a PDF file, specifies whether to render all columns on a single image or PDF page. Valid values:
	//
	// - false (default): No. The content is rendered on multiple images or PDF pages.
	//
	// - true: Yes. The content is rendered on a single image or PDF page.
	//
	// example:
	//
	// false
	FitToWidth *bool `json:"FitToWidth,omitempty" xml:"FitToWidth,omitempty"`
	// When you convert a document to text, specifies whether to keep the line feeds in the document. Valid values:
	//
	// - false (default): No. Line feeds are not kept.
	//
	// - true: Yes. Line feeds are kept.
	//
	// example:
	//
	// false
	HoldLineFeed *bool `json:"HoldLineFeed,omitempty" xml:"HoldLineFeed,omitempty"`
	// The DPI of the output image. Valid values: 96 to 600. The default value is 96.
	//
	// example:
	//
	// 96
	ImageDPI *int64 `json:"ImageDPI,omitempty" xml:"ImageDPI,omitempty"`
	// When you convert a document to images, specifies whether to convert it into a long image. Valid values:
	//
	// - false (default): No. The document is converted into multiple images.
	//
	// - true: Yes. The document is converted into a long image.
	//
	// > You can combine a maximum of 20 pages into a long image. If the number of pages exceeds this limit, the conversion task may fail.
	//
	// example:
	//
	// false
	LongPicture *bool `json:"LongPicture,omitempty" xml:"LongPicture,omitempty"`
	// When you convert a document to text, specifies whether to convert it into a long text file. Valid values:
	//
	// - false (default): No. Each page of the document is converted into a separate text file.
	//
	// - true: Yes. All content is placed in a single text file.
	//
	// example:
	//
	// false
	LongText *bool `json:"LongText,omitempty" xml:"LongText,omitempty"`
	// The maximum number of columns to convert when you convert a spreadsheet document to images. By default, all columns are converted.
	//
	// > This parameter takes effect only when you set **LongPicture*	- to `true`.
	//
	// example:
	//
	// 10
	MaxSheetColumn *int64 `json:"MaxSheetColumn,omitempty" xml:"MaxSheetColumn,omitempty"`
	// The maximum number of rows to convert when you convert a spreadsheet document to images. By default, all rows are converted.
	//
	// > This parameter takes effect only when you set **LongPicture*	- to `true`.
	//
	// example:
	//
	// 10
	MaxSheetRow *int64 `json:"MaxSheetRow,omitempty" xml:"MaxSheetRow,omitempty"`
	// The message notification configuration. For more information, click Notification. For more information about the format of asynchronous notification messages, see [Asynchronous notification message format](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The page numbers to convert. This parameter has a higher priority than the \\`StartPage\\` and \\`EndPage\\` parameters. The format is as follows:
	//
	// - Separate multiple page numbers with commas (,), for example, 1,2.
	//
	// - Specify a range of consecutive pages with a hyphen (-), for example, 1,2-4,7.
	//
	// example:
	//
	// 1,2-4,7
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// When you convert a spreadsheet document to images, specifies whether to place the paper horizontally. The output image is similar to a printed page. Valid values:
	//
	// - false (default): No. The paper is placed vertically.
	//
	// - true: Yes. The paper is placed horizontally.
	//
	// example:
	//
	// false
	PaperHorizontal *bool `json:"PaperHorizontal,omitempty" xml:"PaperHorizontal,omitempty"`
	// The paper size for converting a spreadsheet document to images. The output image is similar to a printed page. Valid values:
	//
	// - A0
	//
	// - A2
	//
	// - A4 (default)
	//
	// > This parameter takes effect only when you use it with the **FitToHeight*	- and **FitToWidth*	- parameters.
	//
	// example:
	//
	// A4
	PaperSize *string `json:"PaperSize,omitempty" xml:"PaperSize,omitempty"`
	// The password to open the document. Set this parameter if you want to convert a password-protected document.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The project name. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The quality of the converted file. Valid values: 0 to 100. A value of 0 indicates the lowest quality and the best performance. A value of 100 indicates the highest quality and the poorest performance. By default, the system sets an appropriate value based on the document content to balance quality and performance.
	//
	// example:
	//
	// 60
	Quality *int64 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// The scaling ratio of the document. Valid values: 20 to 199. The default value is 100, which indicates that the document is not scaled.
	//
	// > A value less than 100 indicates that the document is scaled down. A value greater than 100 indicates that the document is scaled up.
	//
	// example:
	//
	// 100
	ScalePercentage *int64 `json:"ScalePercentage,omitempty" xml:"ScalePercentage,omitempty"`
	// The number of worksheets to convert to images in the spreadsheet document. By default, all worksheets are converted.
	//
	// example:
	//
	// 1
	SheetCount *int64 `json:"SheetCount,omitempty" xml:"SheetCount,omitempty"`
	// The number of the worksheet to convert to images in the spreadsheet document. Valid values: 1 to the number of the last worksheet. The default value is 1.
	//
	// example:
	//
	// 1
	SheetIndex *int64 `json:"SheetIndex,omitempty" xml:"SheetIndex,omitempty"`
	// When you convert a word processor document to images, specifies whether to show comments. Valid values:
	//
	// - false (default): No. Comments are not shown.
	//
	// - true: Yes. Comments are shown.
	//
	// example:
	//
	// false
	ShowComments *bool `json:"ShowComments,omitempty" xml:"ShowComments,omitempty"`
	// The extension type of the source data. By default, the type of the source data is determined by the extension of the OSS object. If the OSS object does not have an extension, you can set this parameter. Valid values:
	//
	// - Word processor documents (Word): doc, docx, wps, wpss, docm, dotm, dot, and dotx
	//
	// - Presentation documents (PowerPoint): pptx, ppt, pot, potx, pps, ppsx, dps, dpt, pptm, potm, ppsm, and dpss
	//
	// - Spreadsheet documents (Excel): xls, xlt, et, ett, xlsx, xltx, csv, xlsb, xlsm, xltm, and ets
	//
	// - PDF documents: pdf
	//
	// example:
	//
	// doc
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The storage address of the source data.
	//
	// The OSS address must be in the oss\\://${Bucket}/${Object} format. \\`${Bucket}\\` is the name of the OSS bucket that is in the same region as the current project. \\`${Object}\\` is the full path of the file, including the file name extension.
	//
	// example:
	//
	// oss://test-bucket/test-object
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// A list of input images. The images are converted in the order of their URIs in the list. (**This parameter is not yet published. Do not use it.**)
	//
	// example:
	//
	// oss://imm-test/test.pptx
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The start page for the document conversion. The default value is 1.
	//
	// > - If the source file is a spreadsheet, you must specify the worksheet number.
	//
	// >
	//
	// > - This parameter takes effect only when you convert the document to images. It does not take effect when you convert the document to a PDF file or a text file.
	//
	// example:
	//
	// 1
	StartPage *int64 `json:"StartPage,omitempty" xml:"StartPage,omitempty"`
	// The custom tags. The value is a dictionary. You can use tags to search for tasks.
	//
	// example:
	//
	// {
	//
	//       "key": "value"
	//
	// }
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The type of the output file. Valid values:
	//
	// - png: Converts the document to PNG images.
	//
	// - jpg: Converts the document to JPG images.
	//
	// - pdf: Converts the document to a PDF file.
	//
	// - txt: Converts the document to a text-only file. This is mainly used to extract text content from the file. This option is supported only for presentation documents, word processor documents, and spreadsheet documents. When you convert a spreadsheet document, a single txt file is generated, and settings for sheet-related variables do not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// png
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The template for the output address of the converted document.
	//
	// The address must be in the `oss://{bucket}/{tags.custom}/{dirname}/{barename}.{autoext}` format. For more information, see [TargetURI templates](https://help.aliyun.com/document_detail/465762.html).
	//
	// > Specify either this parameter or \\`TargetURIPrefix\\`.
	//
	// example:
	//
	// oss://examplebucket/outputDocument.pdf
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The prefix of the storage address for the output file after document conversion.
	//
	// The prefix must be in the `oss://${Bucket}/${Prefix}/` format. \\`${Bucket}\\` is the name of the OSS bucket that is in the same region as the current project. \\`${Prefix}\\` is the prefix of the storage address for the output file.
	//
	// > Specify either this parameter or \\`TargetURI\\`.
	//
	// example:
	//
	// oss://examplebucket/outputprefix/
	TargetURIPrefix *string `json:"TargetURIPrefix,omitempty" xml:"TargetURIPrefix,omitempty"`
	// The trimming policy for spreadsheet conversion. For example, if a spreadsheet contains many empty rows and columns, a large amount of white space may be generated if no trimming policy is specified.
	TrimPolicyShrink *string `json:"TrimPolicy,omitempty" xml:"TrimPolicy,omitempty"`
	// The custom information. This information is returned in the asynchronous notification message to help you associate the notification with your services. The value can be up to 2,048 bytes in length.
	//
	// example:
	//
	// {
	//
	//       "id": "test-id",
	//
	//       "name": "test-name"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateOfficeConversionTaskShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOfficeConversionTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetEndPage() *int64 {
	return s.EndPage
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetFirstPage() *bool {
	return s.FirstPage
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetFitToHeight() *bool {
	return s.FitToHeight
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetFitToWidth() *bool {
	return s.FitToWidth
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetHoldLineFeed() *bool {
	return s.HoldLineFeed
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetImageDPI() *int64 {
	return s.ImageDPI
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetLongPicture() *bool {
	return s.LongPicture
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetLongText() *bool {
	return s.LongText
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetMaxSheetColumn() *int64 {
	return s.MaxSheetColumn
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetMaxSheetRow() *int64 {
	return s.MaxSheetRow
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetPages() *string {
	return s.Pages
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetPaperHorizontal() *bool {
	return s.PaperHorizontal
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetPaperSize() *string {
	return s.PaperSize
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetQuality() *int64 {
	return s.Quality
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetScalePercentage() *int64 {
	return s.ScalePercentage
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetSheetCount() *int64 {
	return s.SheetCount
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetSheetIndex() *int64 {
	return s.SheetIndex
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetShowComments() *bool {
	return s.ShowComments
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetSourcesShrink() *string {
	return s.SourcesShrink
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetStartPage() *int64 {
	return s.StartPage
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetTagsShrink() *string {
	return s.TagsShrink
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetTargetType() *string {
	return s.TargetType
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetTargetURI() *string {
	return s.TargetURI
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetTargetURIPrefix() *string {
	return s.TargetURIPrefix
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetTrimPolicyShrink() *string {
	return s.TrimPolicyShrink
}

func (s *CreateOfficeConversionTaskShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetCredentialConfigShrink(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetEndPage(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.EndPage = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetFirstPage(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.FirstPage = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetFitToHeight(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.FitToHeight = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetFitToWidth(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.FitToWidth = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetHoldLineFeed(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.HoldLineFeed = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetImageDPI(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.ImageDPI = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetLongPicture(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.LongPicture = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetLongText(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.LongText = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetMaxSheetColumn(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.MaxSheetColumn = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetMaxSheetRow(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.MaxSheetRow = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetNotificationShrink(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetPages(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.Pages = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetPaperHorizontal(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.PaperHorizontal = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetPaperSize(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.PaperSize = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetPassword(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.Password = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetProjectName(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetQuality(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.Quality = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetScalePercentage(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.ScalePercentage = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetSheetCount(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.SheetCount = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetSheetIndex(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.SheetIndex = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetShowComments(v bool) *CreateOfficeConversionTaskShrinkRequest {
	s.ShowComments = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetSourceType(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetSourceURI(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetSourcesShrink(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.SourcesShrink = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetStartPage(v int64) *CreateOfficeConversionTaskShrinkRequest {
	s.StartPage = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTagsShrink(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TagsShrink = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTargetType(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TargetType = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTargetURI(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TargetURI = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTargetURIPrefix(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TargetURIPrefix = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetTrimPolicyShrink(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.TrimPolicyShrink = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) SetUserData(v string) *CreateOfficeConversionTaskShrinkRequest {
	s.UserData = &v
	return s
}

func (s *CreateOfficeConversionTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}
