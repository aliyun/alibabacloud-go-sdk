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
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The authorization chain settings. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// The ending page for document conversion. The default value is -1, which converts the file until the last page of the file.
	//
	// >
	//
	// 	- If the source is a spreadsheet file, specify the index number of the corresponding sheet instead.
	//
	// 	- If you convert a large number of pages within the document, we recommend that you split the pages into several document conversion tasks to prevent conversion timeouts.
	//
	// 	- This parameter takes effect only when you convert the file into an image. It does not take effect when you convert the file into a PDF or TXT file.
	//
	// example:
	//
	// -1
	EndPage *int64 `json:"EndPage,omitempty" xml:"EndPage,omitempty"`
	// Specifies whether to return only the first resulting image when you convert a spreadsheet document to images. The number of rows and the number of columns in the first image are determined by the automatic splitting process. Valid values:
	//
	// 	- false (default): does not return only the first resulting image. All the resulting images are returned.
	//
	// 	- true: returns only the first resulting image. A thumbnail is generated.
	//
	// >  This parameter takes effect only when the **LongPicture*	- parameter is set to `true`.
	//
	// example:
	//
	// false
	FirstPage *bool `json:"FirstPage,omitempty" xml:"FirstPage,omitempty"`
	// Specifies whether to convert all rows of a spreadsheet document to one single image or a single-page PDF document when you convert the table document to an image or a PDF document. Valid values:
	//
	// 	- false (default): converts all rows of the document to multiple images or a multi-page PDF document. This is the default value.
	//
	// 	- true: converts all rows of the document to one single image or a single-page PDF document.
	//
	// example:
	//
	// false
	FitToHeight *bool `json:"FitToHeight,omitempty" xml:"FitToHeight,omitempty"`
	// Specifies whether to convert all columns of a spreadsheet document to one single image or a single-page PDF document when you convert the spreadsheet file to an image or a PDF document. Valid values:
	//
	// 	- false (default): converts all columns of the document to multiple images or a multi-page PDF document.
	//
	// 	- true: converts all columns of the document to one single image or a single-page PDF document.
	//
	// example:
	//
	// false
	FitToWidth *bool `json:"FitToWidth,omitempty" xml:"FitToWidth,omitempty"`
	// Specifies whether to retain line feeds in the output file when a document is converted to a text file. Valid values:
	//
	// 	- false (default): does not retain the line feeds.
	//
	// 	- true: retains the line feeds.
	//
	// example:
	//
	// false
	HoldLineFeed *bool `json:"HoldLineFeed,omitempty" xml:"HoldLineFeed,omitempty"`
	// The dots per inch (DPI) of output images. Valid values: 96 to 600. Default value: 96.
	//
	// example:
	//
	// 96
	ImageDPI *int64 `json:"ImageDPI,omitempty" xml:"ImageDPI,omitempty"`
	// Specifies whether to convert the document to a long image. Valid values:
	//
	// 	- false (default): does not convert the document to a long image.
	//
	// 	- true: converts the document to a long image.
	//
	// >  You can convert up to 20 pages of a document into a long image. If you convert more than 20 pages to a long image, an error may occur.
	//
	// example:
	//
	// false
	LongPicture *bool `json:"LongPicture,omitempty" xml:"LongPicture,omitempty"`
	// Specifies whether to convert the document to a long text file. Valid values:
	//
	// 	- false (default): does not convert the document to a long text file. Each page of the document is converted to a text file.
	//
	// 	- true: converts the entire document to a long text file.
	//
	// example:
	//
	// false
	LongText *bool `json:"LongText,omitempty" xml:"LongText,omitempty"`
	// The maximum number of spreadsheet columns to be converted to an image. By default, all columns within the spreadsheet file are converted.
	//
	// >  This parameter takes effect only when the **LongPicture*	- parameter is set to `true`.
	//
	// example:
	//
	// 10
	MaxSheetColumn *int64 `json:"MaxSheetColumn,omitempty" xml:"MaxSheetColumn,omitempty"`
	// The maximum number of spreadsheet rows to be converted to an image. By default, all rows within the spreadsheet file are converted.
	//
	// >  This parameter takes effect only when the **LongPicture*	- parameter is set to `true`.
	//
	// example:
	//
	// 10
	MaxSheetRow *int64 `json:"MaxSheetRow,omitempty" xml:"MaxSheetRow,omitempty"`
	// The notification settings. For information about the asynchronous notification format, see [Asynchronous message examples](https://help.aliyun.com/document_detail/2743997.html).
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// The numbers of pages to be converted. This parameter takes precedence over the StartPage and EndPage parameters. The value of this parameter can be in different formats:
	//
	// 	- If you specify pages separately by page number, separate page numbers with commas (,). Example: 1,2
	//
	// 	- If you specify consecutive pages by using a page range, connect the starting and ending page numbers with a hyphen (-). Example: 1,2-4,7
	//
	// example:
	//
	// 1,2-4,7
	Pages *string `json:"Pages,omitempty" xml:"Pages,omitempty"`
	// Specifies whether to place sheets of paper horizontally for converting a spreadsheet document to images. Conversion to images is similar to printing the content on a sheet of paper. Valid values:
	//
	// 	- false (default): does not place sheets of paper horizontally. Paper sheets are placed vertically.
	//
	// 	- true: places sheets of paper horizontally.
	//
	// example:
	//
	// false
	PaperHorizontal *bool `json:"PaperHorizontal,omitempty" xml:"PaperHorizontal,omitempty"`
	// The paper size for converting a spreadsheet document to images. Conversion to images is similar to printing the content on a sheet of paper. Valid values:
	//
	// 	- A0
	//
	// 	- A2
	//
	// 	- A4 (default)
	//
	// >  This parameter takes effect only when the **FitToHeight*	- and **FitToWidth*	- parameters are specified.
	//
	// example:
	//
	// A4
	PaperSize *string `json:"PaperSize,omitempty" xml:"PaperSize,omitempty"`
	// The password that protects the source document. To convert a password-protected document, specify this parameter.
	//
	// example:
	//
	// ********
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The name of the project.[](~~478153~~)
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The quality of the output file. Valid values: 0 to 100. A smaller value indicates lower quality and better conversion performance. By default, the system specifies an appropriate value that provides an optimal balance between the quality and conversion performance based on the document content.
	//
	// example:
	//
	// 60
	Quality *int64 `json:"Quality,omitempty" xml:"Quality,omitempty"`
	// The percentage scale relative to the source document. Valid values: 20 to 200. The default value is 100, which indicates that the document is not scaled.
	//
	// >  A value that is less than 100 indicates a size reduction. A value that is greater than 100 indicates an enlargement.
	//
	// example:
	//
	// 100
	ScalePercentage *int64 `json:"ScalePercentage,omitempty" xml:"ScalePercentage,omitempty"`
	// The number of sheets to be converted to an image. By default, all sheets within the spreadsheet file are converted.
	//
	// example:
	//
	// 1
	SheetCount *int64 `json:"SheetCount,omitempty" xml:"SheetCount,omitempty"`
	// The index number of the sheet to be converted to an image. The value ranges from 1 to the index number of the last sheet. By default, the conversion starts from the first sheet.
	//
	// example:
	//
	// 1
	SheetIndex *int64 `json:"SheetIndex,omitempty" xml:"SheetIndex,omitempty"`
	// Specifies whether to display comments in resulting images when a text document is converted to images. Valid values:
	//
	// 	- false (default): does not display comments in resulting images.
	//
	// 	- true: displays comments in resulting images.
	//
	// example:
	//
	// false
	ShowComments *bool `json:"ShowComments,omitempty" xml:"ShowComments,omitempty"`
	// The name extension of the source file. By default, the type of the source file is determined based on the name extension of the source object in OSS. If the object in OSS does not have a name extension, you can specify this parameter. Valid values:
	//
	// 	- Text documents: doc, docx, wps, wpss, docm, dotm, dot, dotx, and html
	//
	// 	- Presentation documents: pptx, ppt, pot, potx, pps, ppsx, dps, dpt, pptm, potm, ppsm, and dpss
	//
	// 	- Spreadsheet documents: xls, xlt, et, ett, xlsx, xltx, csv, xlsb, xlsm, xltm, and ets
	//
	// 	- PDF documents: pdf
	//
	// example:
	//
	// doc
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// The URI of the source file.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://test-bucket/test-object
	SourceURI     *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	SourcesShrink *string `json:"Sources,omitempty" xml:"Sources,omitempty"`
	// The starting page for document conversion. Default value: 1.
	//
	// >
	//
	// 	- If the document is a spreadsheet file, specify the index number of the corresponding sheet instead.
	//
	// 	- This parameter takes effect only when you convert the file to an image format. It does not take effect when you convert the file into a PDF or TXT file.
	//
	// example:
	//
	// 1
	StartPage *int64 `json:"StartPage,omitempty" xml:"StartPage,omitempty"`
	// The custom tags in dictionary format. You can use the custom tags to search for the task.
	//
	// example:
	//
	// {"test":"val1"}
	TagsShrink *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The format of the output file. Valid values:
	//
	// 	- png: a PNG image.
	//
	// 	- jpg: a JPG image.
	//
	// 	- pdf: a PDF file.
	//
	// 	- txt: a TXT file. You can specify this value to extract the text content of the source document. Only presentation, text, or spreadsheet documents can be converted to a TXT file. If the source document is a spreadsheet, only one TXT is created and sheet-related parameters do not take effect.
	//
	// This parameter is required.
	//
	// example:
	//
	// png
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
	// The address template of the output file.
	//
	// Specify the value in the `oss://{bucket}/{tags.custom}/{dirname}/{barename}.{autoext}` format. For more information, see [TargetURI template](https://help.aliyun.com/document_detail/465762.html).
	//
	// >  Specify at least one of the TargetURI and TargetURIPrefix parameters.
	//
	// example:
	//
	// oss://{bucket}/{tags.custom}/{dirname}/{barename}.{autoext}
	TargetURI *string `json:"TargetURI,omitempty" xml:"TargetURI,omitempty"`
	// The prefix of the storage address of the output file.
	//
	// Specify the prefix in the `oss://${Bucket}/${Prefix}/` format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Prefix}` is the prefix of the output file.
	//
	// >  Specify at least one of the TargetURI and TargetURIPrefix parameters.
	//
	// example:
	//
	// oss://bucket1/
	TargetURIPrefix *string `json:"TargetURIPrefix,omitempty" xml:"TargetURIPrefix,omitempty"`
	// The trim policy for converting a spreadsheet file. Empty rows and columns may generate blank spaces in the output file if no appropriate trim policy is specified.
	TrimPolicyShrink *string `json:"TrimPolicy,omitempty" xml:"TrimPolicy,omitempty"`
	// The custom information, which is returned in an asynchronous notification and facilitates notification management. The maximum information length is 2,048 bytes.
	//
	// example:
	//
	// {"file_id": "abc"}
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
