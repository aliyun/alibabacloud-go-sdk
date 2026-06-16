// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationProResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageTranslationProResponseBody
	GetCode() *string
	SetData(v *ImageTranslationProResponseBodyData) *ImageTranslationProResponseBody
	GetData() *ImageTranslationProResponseBodyData
	SetMessage(v string) *ImageTranslationProResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageTranslationProResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageTranslationProResponseBody
	GetSuccess() *bool
}

type ImageTranslationProResponseBody struct {
	// The response code. A value of 200 indicates success. For other response codes, see the error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The translation result data. ResultList contains the URL of the translation result. GenFiles contains EditInfo with the recognized text information.
	Data *ImageTranslationProResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. "Success" is returned for successful calls. A specific error message is returned for failed calls, such as "The content contains sensitive data. Try other input.".
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which uniquely identifies the request.
	//
	// example:
	//
	// 61785C32-80C2-19A3-9E55-2C3702C84B40
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageTranslationProResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBody) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageTranslationProResponseBody) GetData() *ImageTranslationProResponseBodyData {
	return s.Data
}

func (s *ImageTranslationProResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageTranslationProResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageTranslationProResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageTranslationProResponseBody) SetCode(v string) *ImageTranslationProResponseBody {
	s.Code = &v
	return s
}

func (s *ImageTranslationProResponseBody) SetData(v *ImageTranslationProResponseBodyData) *ImageTranslationProResponseBody {
	s.Data = v
	return s
}

func (s *ImageTranslationProResponseBody) SetMessage(v string) *ImageTranslationProResponseBody {
	s.Message = &v
	return s
}

func (s *ImageTranslationProResponseBody) SetRequestId(v string) *ImageTranslationProResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageTranslationProResponseBody) SetSuccess(v bool) *ImageTranslationProResponseBody {
	s.Success = &v
	return s
}

func (s *ImageTranslationProResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageTranslationProResponseBodyData struct {
	// The editor protocol, which contains the translation result files and editing information.
	GenFiles []*ImageTranslationProResponseBodyDataGenFiles `json:"GenFiles,omitempty" xml:"GenFiles,omitempty" type:"Repeated"`
	// The list of image translation results.
	ResultList []*ImageTranslationProResponseBodyDataResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	// The asynchronous task ID. This is not returned for synchronous calls.
	//
	// example:
	//
	// abc123-task-id
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The usage information, including the number of processed images.
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s ImageTranslationProResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyData) GetGenFiles() []*ImageTranslationProResponseBodyDataGenFiles {
	return s.GenFiles
}

func (s *ImageTranslationProResponseBodyData) GetResultList() []*ImageTranslationProResponseBodyDataResultList {
	return s.ResultList
}

func (s *ImageTranslationProResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *ImageTranslationProResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *ImageTranslationProResponseBodyData) SetGenFiles(v []*ImageTranslationProResponseBodyDataGenFiles) *ImageTranslationProResponseBodyData {
	s.GenFiles = v
	return s
}

func (s *ImageTranslationProResponseBodyData) SetResultList(v []*ImageTranslationProResponseBodyDataResultList) *ImageTranslationProResponseBodyData {
	s.ResultList = v
	return s
}

func (s *ImageTranslationProResponseBodyData) SetTaskId(v string) *ImageTranslationProResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ImageTranslationProResponseBodyData) SetUsageMap(v map[string]*int64) *ImageTranslationProResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *ImageTranslationProResponseBodyData) Validate() error {
	if s.GenFiles != nil {
		for _, item := range s.GenFiles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ResultList != nil {
		for _, item := range s.ResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImageTranslationProResponseBodyDataGenFiles struct {
	// The editor information, which contains recognized information such as text areas, product areas, and fonts.
	EditInfo *ImageTranslationProResponseBodyDataGenFilesEditInfo `json:"EditInfo,omitempty" xml:"EditInfo,omitempty" type:"Struct"`
	// The collection of translation results.
	ResultList []*ImageTranslationProResponseBodyDataGenFilesResultList `json:"ResultList,omitempty" xml:"ResultList,omitempty" type:"Repeated"`
	// The URL of the original image.
	//
	// example:
	//
	// https://img.alicdn.com/imgextra/i3/O1CN01HTDhDi28Fd85ZYs7H_!!6000000007903-0-tps-800-800.jpg
	SrcImage *string `json:"SrcImage,omitempty" xml:"SrcImage,omitempty"`
}

func (s ImageTranslationProResponseBodyDataGenFiles) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyDataGenFiles) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyDataGenFiles) GetEditInfo() *ImageTranslationProResponseBodyDataGenFilesEditInfo {
	return s.EditInfo
}

func (s *ImageTranslationProResponseBodyDataGenFiles) GetResultList() []*ImageTranslationProResponseBodyDataGenFilesResultList {
	return s.ResultList
}

func (s *ImageTranslationProResponseBodyDataGenFiles) GetSrcImage() *string {
	return s.SrcImage
}

func (s *ImageTranslationProResponseBodyDataGenFiles) SetEditInfo(v *ImageTranslationProResponseBodyDataGenFilesEditInfo) *ImageTranslationProResponseBodyDataGenFiles {
	s.EditInfo = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFiles) SetResultList(v []*ImageTranslationProResponseBodyDataGenFilesResultList) *ImageTranslationProResponseBodyDataGenFiles {
	s.ResultList = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFiles) SetSrcImage(v string) *ImageTranslationProResponseBodyDataGenFiles {
	s.SrcImage = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFiles) Validate() error {
	if s.EditInfo != nil {
		if err := s.EditInfo.Validate(); err != nil {
			return err
		}
	}
	if s.ResultList != nil {
		for _, item := range s.ResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImageTranslationProResponseBodyDataGenFilesEditInfo struct {
	// The list of font types.
	Font []*string `json:"Font,omitempty" xml:"Font,omitempty" type:"Repeated"`
	// The coordinate information of the product bounding box area.
	GoodsRects *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects `json:"GoodsRects,omitempty" xml:"GoodsRects,omitempty" type:"Struct"`
	// The list of target languages for translation.
	Languages []*string `json:"Languages,omitempty" xml:"Languages,omitempty" type:"Repeated"`
	// The URL of the image with all text removed.
	//
	// example:
	//
	// http://dashscope-a717.oss-cn-beijing.aliyuncs.com/repaired.png
	RepairedUrl *string `json:"RepairedUrl,omitempty" xml:"RepairedUrl,omitempty"`
	// The collection of global IDs of translated images.
	ResultImageIds []*string `json:"ResultImageIds,omitempty" xml:"ResultImageIds,omitempty" type:"Repeated"`
	// The list of text boxes, which contains information about all recognized text areas.
	TextAreas []*ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas `json:"TextAreas,omitempty" xml:"TextAreas,omitempty" type:"Repeated"`
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfo) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfo) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) GetFont() []*string {
	return s.Font
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) GetGoodsRects() *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects {
	return s.GoodsRects
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) GetLanguages() []*string {
	return s.Languages
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) GetRepairedUrl() *string {
	return s.RepairedUrl
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) GetResultImageIds() []*string {
	return s.ResultImageIds
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) GetTextAreas() []*ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas {
	return s.TextAreas
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) SetFont(v []*string) *ImageTranslationProResponseBodyDataGenFilesEditInfo {
	s.Font = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) SetGoodsRects(v *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) *ImageTranslationProResponseBodyDataGenFilesEditInfo {
	s.GoodsRects = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) SetLanguages(v []*string) *ImageTranslationProResponseBodyDataGenFilesEditInfo {
	s.Languages = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) SetRepairedUrl(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfo {
	s.RepairedUrl = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) SetResultImageIds(v []*string) *ImageTranslationProResponseBodyDataGenFilesEditInfo {
	s.ResultImageIds = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) SetTextAreas(v []*ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) *ImageTranslationProResponseBodyDataGenFilesEditInfo {
	s.TextAreas = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfo) Validate() error {
	if s.GoodsRects != nil {
		if err := s.GoodsRects.Validate(); err != nil {
			return err
		}
	}
	if s.TextAreas != nil {
		for _, item := range s.TextAreas {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects struct {
	// The rotation angle of the text box in degrees. A value less than 1 indicates a horizontal text box.
	//
	// example:
	//
	// 0
	Degree *int32 `json:"Degree,omitempty" xml:"Degree,omitempty"`
	// The height.
	//
	// example:
	//
	// 0
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The left coordinate.
	//
	// example:
	//
	// 0
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top coordinate.
	//
	// example:
	//
	// 0
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width.
	//
	// example:
	//
	// 0
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) GetDegree() *int32 {
	return s.Degree
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) GetHeight() *int32 {
	return s.Height
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) GetLeft() *int32 {
	return s.Left
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) GetTop() *int32 {
	return s.Top
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) GetWidth() *int32 {
	return s.Width
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) SetDegree(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects {
	s.Degree = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) SetHeight(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects {
	s.Height = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) SetLeft(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects {
	s.Left = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) SetTop(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects {
	s.Top = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) SetWidth(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects {
	s.Width = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoGoodsRects) Validate() error {
	return dara.Validate(s)
}

type ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas struct {
	// The text color, such as #ffffff.
	//
	// example:
	//
	// #515151
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The original text before translation.
	//
	// example:
	//
	// 萌趣造型·清脆响铃
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The font size.
	//
	// example:
	//
	// 32
	Fontsize *int32 `json:"Fontsize,omitempty" xml:"Fontsize,omitempty"`
	// The horizontal layout. Valid values: center, left, right.
	//
	// example:
	//
	// left
	HorizontalLayout *string `json:"HorizontalLayout,omitempty" xml:"HorizontalLayout,omitempty"`
	// The number of lines in the text box.
	//
	// example:
	//
	// 1
	LineCount *int32 `json:"LineCount,omitempty" xml:"LineCount,omitempty"`
	// The list of translated texts. Each element corresponds to the translation result for a target language.
	Texts []*ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// The vertical layout. Valid values: center, top, down.
	//
	// example:
	//
	// center
	VerticalLayout *string `json:"VerticalLayout,omitempty" xml:"VerticalLayout,omitempty"`
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) GetColor() *string {
	return s.Color
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) GetContent() *string {
	return s.Content
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) GetFontsize() *int32 {
	return s.Fontsize
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) GetHorizontalLayout() *string {
	return s.HorizontalLayout
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) GetLineCount() *int32 {
	return s.LineCount
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) GetTexts() []*ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	return s.Texts
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) GetVerticalLayout() *string {
	return s.VerticalLayout
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) SetColor(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas {
	s.Color = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) SetContent(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas {
	s.Content = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) SetFontsize(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas {
	s.Fontsize = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) SetHorizontalLayout(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas {
	s.HorizontalLayout = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) SetLineCount(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas {
	s.LineCount = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) SetTexts(v []*ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas {
	s.Texts = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) SetVerticalLayout(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas {
	s.VerticalLayout = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreas) Validate() error {
	if s.Texts != nil {
		for _, item := range s.Texts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts struct {
	// The color of the translated text.
	//
	// example:
	//
	// #515151
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The font size of the translated text.
	//
	// example:
	//
	// 29
	Fontsize *int32 `json:"Fontsize,omitempty" xml:"Fontsize,omitempty"`
	// The horizontal layout. Valid values: center, left, right.
	//
	// example:
	//
	// center
	HorizontalLayout *string `json:"HorizontalLayout,omitempty" xml:"HorizontalLayout,omitempty"`
	// The coordinates of the image repair area.
	ImageRect *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect `json:"ImageRect,omitempty" xml:"ImageRect,omitempty" type:"Struct"`
	// The target language code for translation.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The number of lines in the text box.
	//
	// example:
	//
	// 1
	LineCount *int32 `json:"LineCount,omitempty" xml:"LineCount,omitempty"`
	// The Ovis model error message and execution time.
	//
	// example:
	//
	// | ovis time: 0.748
	OvisErrMsg *string `json:"OvisErrMsg,omitempty" xml:"OvisErrMsg,omitempty"`
	// The coordinates of the text box area.
	TextRect *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect `json:"TextRect,omitempty" xml:"TextRect,omitempty" type:"Struct"`
	// Indicates whether the TextItem is valid. The item is invalid if this value does not exist or is false.
	//
	// example:
	//
	// true
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
	// The translated text content.
	//
	// example:
	//
	// Adorable Design · Crisp Bell Sound
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The vertical layout. Valid values: center, top, down.
	//
	// example:
	//
	// center
	VerticalLayout *string `json:"VerticalLayout,omitempty" xml:"VerticalLayout,omitempty"`
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetColor() *string {
	return s.Color
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetFontsize() *int32 {
	return s.Fontsize
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetHorizontalLayout() *string {
	return s.HorizontalLayout
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetImageRect() *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect {
	return s.ImageRect
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetLanguage() *string {
	return s.Language
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetLineCount() *int32 {
	return s.LineCount
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetOvisErrMsg() *string {
	return s.OvisErrMsg
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetTextRect() *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect {
	return s.TextRect
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetValid() *bool {
	return s.Valid
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetValue() *string {
	return s.Value
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) GetVerticalLayout() *string {
	return s.VerticalLayout
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetColor(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.Color = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetFontsize(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.Fontsize = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetHorizontalLayout(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.HorizontalLayout = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetImageRect(v *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.ImageRect = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetLanguage(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.Language = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetLineCount(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.LineCount = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetOvisErrMsg(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.OvisErrMsg = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetTextRect(v *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.TextRect = v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetValid(v bool) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.Valid = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetValue(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.Value = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) SetVerticalLayout(v string) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts {
	s.VerticalLayout = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTexts) Validate() error {
	if s.ImageRect != nil {
		if err := s.ImageRect.Validate(); err != nil {
			return err
		}
	}
	if s.TextRect != nil {
		if err := s.TextRect.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect struct {
	// The rotation angle of the text box in degrees. A value less than 1 indicates a horizontal text box.
	//
	// example:
	//
	// 0
	Degree *int32 `json:"Degree,omitempty" xml:"Degree,omitempty"`
	// The height.
	//
	// example:
	//
	// 54
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The left coordinate.
	//
	// example:
	//
	// 43
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top coordinate.
	//
	// example:
	//
	// 83
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width.
	//
	// example:
	//
	// 418
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) GetDegree() *int32 {
	return s.Degree
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) GetHeight() *int32 {
	return s.Height
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) GetLeft() *int32 {
	return s.Left
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) GetTop() *int32 {
	return s.Top
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) GetWidth() *int32 {
	return s.Width
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) SetDegree(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect {
	s.Degree = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) SetHeight(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect {
	s.Height = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) SetLeft(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect {
	s.Left = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) SetTop(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect {
	s.Top = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) SetWidth(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect {
	s.Width = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsImageRect) Validate() error {
	return dara.Validate(s)
}

type ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect struct {
	// The rotation angle of the text box in degrees. A value less than 1 indicates a horizontal text box.
	//
	// example:
	//
	// 0
	Degree *int32 `json:"Degree,omitempty" xml:"Degree,omitempty"`
	// The height.
	//
	// example:
	//
	// 30
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The left coordinate.
	//
	// example:
	//
	// 8
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top coordinate.
	//
	// example:
	//
	// 95
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width.
	//
	// example:
	//
	// 488
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) GetDegree() *int32 {
	return s.Degree
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) GetHeight() *int32 {
	return s.Height
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) GetLeft() *int32 {
	return s.Left
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) GetTop() *int32 {
	return s.Top
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) GetWidth() *int32 {
	return s.Width
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) SetDegree(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect {
	s.Degree = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) SetHeight(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect {
	s.Height = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) SetLeft(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect {
	s.Left = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) SetTop(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect {
	s.Top = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) SetWidth(v int32) *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect {
	s.Width = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesEditInfoTextAreasTextsTextRect) Validate() error {
	return dara.Validate(s)
}

type ImageTranslationProResponseBodyDataGenFilesResultList struct {
	// The URL of the translated image file.
	//
	// example:
	//
	// http://dashscope-a717.oss-cn-beijing.aliyuncs.com/xxx.jpg
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The target language for translation.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ImageTranslationProResponseBodyDataGenFilesResultList) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyDataGenFilesResultList) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyDataGenFilesResultList) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ImageTranslationProResponseBodyDataGenFilesResultList) GetLanguage() *string {
	return s.Language
}

func (s *ImageTranslationProResponseBodyDataGenFilesResultList) SetFileUrl(v string) *ImageTranslationProResponseBodyDataGenFilesResultList {
	s.FileUrl = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesResultList) SetLanguage(v string) *ImageTranslationProResponseBodyDataGenFilesResultList {
	s.Language = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataGenFilesResultList) Validate() error {
	return dara.Validate(s)
}

type ImageTranslationProResponseBodyDataResultList struct {
	// The URL of the image translation result image.
	//
	// example:
	//
	// http://dashscope-a717.oss-cn-beijing.aliyuncs.com/xxx.jpg
	FileUrl *string `json:"FileUrl,omitempty" xml:"FileUrl,omitempty"`
	// The target language for image translation.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s ImageTranslationProResponseBodyDataResultList) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationProResponseBodyDataResultList) GoString() string {
	return s.String()
}

func (s *ImageTranslationProResponseBodyDataResultList) GetFileUrl() *string {
	return s.FileUrl
}

func (s *ImageTranslationProResponseBodyDataResultList) GetLanguage() *string {
	return s.Language
}

func (s *ImageTranslationProResponseBodyDataResultList) SetFileUrl(v string) *ImageTranslationProResponseBodyDataResultList {
	s.FileUrl = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataResultList) SetLanguage(v string) *ImageTranslationProResponseBodyDataResultList {
	s.Language = &v
	return s
}

func (s *ImageTranslationProResponseBodyDataResultList) Validate() error {
	return dara.Validate(s)
}
