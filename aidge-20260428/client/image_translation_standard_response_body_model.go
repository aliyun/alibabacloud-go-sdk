// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageTranslationStandardResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ImageTranslationStandardResponseBody
	GetCode() *string
	SetData(v *ImageTranslationStandardResponseBodyData) *ImageTranslationStandardResponseBody
	GetData() *ImageTranslationStandardResponseBodyData
	SetMessage(v string) *ImageTranslationStandardResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImageTranslationStandardResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImageTranslationStandardResponseBody
	GetSuccess() *bool
}

type ImageTranslationStandardResponseBody struct {
	// The response code. A value of 200 indicates a successful call. For other response codes, refer to the error code information.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The translation result data, including the translated image URL and usage information.
	Data *ImageTranslationStandardResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. Returns "Success" for successful calls, and returns specific error information for failed calls.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which uniquely identifies the request.
	//
	// example:
	//
	// 1CEC4D94-905A-1ED1-A7B4-1BFEFFB3D850
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success, and a value of false indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ImageTranslationStandardResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBody) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBody) GetCode() *string {
	return s.Code
}

func (s *ImageTranslationStandardResponseBody) GetData() *ImageTranslationStandardResponseBodyData {
	return s.Data
}

func (s *ImageTranslationStandardResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImageTranslationStandardResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageTranslationStandardResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImageTranslationStandardResponseBody) SetCode(v string) *ImageTranslationStandardResponseBody {
	s.Code = &v
	return s
}

func (s *ImageTranslationStandardResponseBody) SetData(v *ImageTranslationStandardResponseBodyData) *ImageTranslationStandardResponseBody {
	s.Data = v
	return s
}

func (s *ImageTranslationStandardResponseBody) SetMessage(v string) *ImageTranslationStandardResponseBody {
	s.Message = &v
	return s
}

func (s *ImageTranslationStandardResponseBody) SetRequestId(v string) *ImageTranslationStandardResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageTranslationStandardResponseBody) SetSuccess(v bool) *ImageTranslationStandardResponseBody {
	s.Success = &v
	return s
}

func (s *ImageTranslationStandardResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageTranslationStandardResponseBodyData struct {
	// The edit information.
	EditInfo *ImageTranslationStandardResponseBodyDataEditInfo `json:"EditInfo,omitempty" xml:"EditInfo,omitempty" type:"Struct"`
	// The URL of the image generated after image translation.
	//
	// example:
	//
	// http://dashscope-a717.oss-cn-beijing.aliyuncs.com/xxx.jpg
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	// The usage information, including the number of images processed.
	//
	// example:
	//
	// {"ProcessedImageCount":1}
	UsageMap map[string]*int64 `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s ImageTranslationStandardResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBodyData) GetEditInfo() *ImageTranslationStandardResponseBodyDataEditInfo {
	return s.EditInfo
}

func (s *ImageTranslationStandardResponseBodyData) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ImageTranslationStandardResponseBodyData) GetUsageMap() map[string]*int64 {
	return s.UsageMap
}

func (s *ImageTranslationStandardResponseBodyData) SetEditInfo(v *ImageTranslationStandardResponseBodyDataEditInfo) *ImageTranslationStandardResponseBodyData {
	s.EditInfo = v
	return s
}

func (s *ImageTranslationStandardResponseBodyData) SetImageUrl(v string) *ImageTranslationStandardResponseBodyData {
	s.ImageUrl = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyData) SetUsageMap(v map[string]*int64) *ImageTranslationStandardResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *ImageTranslationStandardResponseBodyData) Validate() error {
	if s.EditInfo != nil {
		if err := s.EditInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ImageTranslationStandardResponseBodyDataEditInfo struct {
	// The list of fonts used.
	//
	// example:
	//
	// ["NotoSansSC-Bold"]
	Font []*string `json:"Font,omitempty" xml:"Font,omitempty" type:"Repeated"`
	// The product area rectangles.
	GoodsRects *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects `json:"GoodsRects,omitempty" xml:"GoodsRects,omitempty" type:"Struct"`
	// The product image URL.
	//
	// example:
	//
	// https://xiuxiu-pro.meitudata.com/posters/34d78f9157e5560a4d612949ca6f6485.jpg
	GoodsUrl *string `json:"GoodsUrl,omitempty" xml:"GoodsUrl,omitempty"`
	// The list of target languages.
	//
	// example:
	//
	// ["zh"]
	Languages []*string `json:"Languages,omitempty" xml:"Languages,omitempty" type:"Repeated"`
	// The original image URL.
	//
	// example:
	//
	// https://xiuxiu-pro.meitudata.com/posters/34d78f9157e5560a4d612949ca6f6485.jpg
	PictUrl *string `json:"PictUrl,omitempty" xml:"PictUrl,omitempty"`
	// The repaired image URL.
	//
	// example:
	//
	// http://dashscope-a717.oss-cn-beijing.aliyuncs.com/xxx.png
	RepairedUrl *string `json:"RepairedUrl,omitempty" xml:"RepairedUrl,omitempty"`
	// The list of repaired image URLs.
	//
	// example:
	//
	// ["http://dashscope-a717.oss-cn-beijing.aliyuncs.com/xxx.png"]
	RepairedUrls []*string `json:"RepairedUrls,omitempty" xml:"RepairedUrls,omitempty" type:"Repeated"`
	// The list of result image IDs.
	//
	// example:
	//
	// ["c18ab570-81aa-11f1-a14a-5ee00dcfdc3c"]
	ResultImageIds []*string `json:"ResultImageIds,omitempty" xml:"ResultImageIds,omitempty" type:"Repeated"`
	// The list of result image URLs.
	//
	// example:
	//
	// ["http://dashscope-a717.oss-cn-beijing.aliyuncs.com/xxx.jpg"]
	ResultUrls []*string `json:"ResultUrls,omitempty" xml:"ResultUrls,omitempty" type:"Repeated"`
	// The list of text areas.
	TextAreas []*ImageTranslationStandardResponseBodyDataEditInfoTextAreas `json:"TextAreas,omitempty" xml:"TextAreas,omitempty" type:"Repeated"`
}

func (s ImageTranslationStandardResponseBodyDataEditInfo) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBodyDataEditInfo) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetFont() []*string {
	return s.Font
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetGoodsRects() *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects {
	return s.GoodsRects
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetGoodsUrl() *string {
	return s.GoodsUrl
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetLanguages() []*string {
	return s.Languages
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetPictUrl() *string {
	return s.PictUrl
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetRepairedUrl() *string {
	return s.RepairedUrl
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetRepairedUrls() []*string {
	return s.RepairedUrls
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetResultImageIds() []*string {
	return s.ResultImageIds
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetResultUrls() []*string {
	return s.ResultUrls
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) GetTextAreas() []*ImageTranslationStandardResponseBodyDataEditInfoTextAreas {
	return s.TextAreas
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetFont(v []*string) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.Font = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetGoodsRects(v *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.GoodsRects = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetGoodsUrl(v string) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.GoodsUrl = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetLanguages(v []*string) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.Languages = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetPictUrl(v string) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.PictUrl = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetRepairedUrl(v string) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.RepairedUrl = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetRepairedUrls(v []*string) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.RepairedUrls = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetResultImageIds(v []*string) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.ResultImageIds = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetResultUrls(v []*string) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.ResultUrls = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) SetTextAreas(v []*ImageTranslationStandardResponseBodyDataEditInfoTextAreas) *ImageTranslationStandardResponseBodyDataEditInfo {
	s.TextAreas = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfo) Validate() error {
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

type ImageTranslationStandardResponseBodyDataEditInfoGoodsRects struct {
	// The rotation angle.
	Degree *int32 `json:"Degree,omitempty" xml:"Degree,omitempty"`
	// The height.
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The left coordinate.
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top coordinate.
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width.
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) GetDegree() *int32 {
	return s.Degree
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) GetHeight() *int32 {
	return s.Height
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) GetLeft() *int32 {
	return s.Left
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) GetTop() *int32 {
	return s.Top
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) GetWidth() *int32 {
	return s.Width
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) SetDegree(v int32) *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects {
	s.Degree = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) SetHeight(v int32) *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects {
	s.Height = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) SetLeft(v int32) *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects {
	s.Left = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) SetTop(v int32) *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects {
	s.Top = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) SetWidth(v int32) *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects {
	s.Width = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoGoodsRects) Validate() error {
	return dara.Validate(s)
}

type ImageTranslationStandardResponseBodyDataEditInfoTextAreas struct {
	// The color.
	//
	// example:
	//
	// #000000
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The content.
	//
	// example:
	//
	// SOLIDWOOD WARDROBE
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The font size.
	//
	// example:
	//
	// 36
	Fontsize *int32 `json:"Fontsize,omitempty" xml:"Fontsize,omitempty"`
	// The horizontal layout.
	//
	// example:
	//
	// left
	HorizontalLayout *string `json:"HorizontalLayout,omitempty" xml:"HorizontalLayout,omitempty"`
	// The line count.
	//
	// example:
	//
	// 2
	LineCount *int32 `json:"LineCount,omitempty" xml:"LineCount,omitempty"`
	// The list of texts.
	Texts []*ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts `json:"Texts,omitempty" xml:"Texts,omitempty" type:"Repeated"`
	// The vertical layout.
	//
	// example:
	//
	// center
	VerticalLayout *string `json:"VerticalLayout,omitempty" xml:"VerticalLayout,omitempty"`
}

func (s ImageTranslationStandardResponseBodyDataEditInfoTextAreas) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBodyDataEditInfoTextAreas) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) GetColor() *string {
	return s.Color
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) GetContent() *string {
	return s.Content
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) GetFontsize() *int32 {
	return s.Fontsize
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) GetHorizontalLayout() *string {
	return s.HorizontalLayout
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) GetLineCount() *int32 {
	return s.LineCount
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) GetTexts() []*ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	return s.Texts
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) GetVerticalLayout() *string {
	return s.VerticalLayout
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) SetColor(v string) *ImageTranslationStandardResponseBodyDataEditInfoTextAreas {
	s.Color = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) SetContent(v string) *ImageTranslationStandardResponseBodyDataEditInfoTextAreas {
	s.Content = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) SetFontsize(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreas {
	s.Fontsize = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) SetHorizontalLayout(v string) *ImageTranslationStandardResponseBodyDataEditInfoTextAreas {
	s.HorizontalLayout = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) SetLineCount(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreas {
	s.LineCount = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) SetTexts(v []*ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) *ImageTranslationStandardResponseBodyDataEditInfoTextAreas {
	s.Texts = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) SetVerticalLayout(v string) *ImageTranslationStandardResponseBodyDataEditInfoTextAreas {
	s.VerticalLayout = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreas) Validate() error {
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

type ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts struct {
	// The color.
	//
	// example:
	//
	// #9d7746
	Color *string `json:"Color,omitempty" xml:"Color,omitempty"`
	// The font size.
	//
	// example:
	//
	// 42
	Fontsize *int32 `json:"Fontsize,omitempty" xml:"Fontsize,omitempty"`
	// The horizontal layout.
	//
	// example:
	//
	// left
	HorizontalLayout *string `json:"HorizontalLayout,omitempty" xml:"HorizontalLayout,omitempty"`
	// The image area.
	ImageRect *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect `json:"ImageRect,omitempty" xml:"ImageRect,omitempty" type:"Struct"`
	// The language.
	//
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The line count.
	//
	// example:
	//
	// 1
	LineCount *int32 `json:"LineCount,omitempty" xml:"LineCount,omitempty"`
	// The text area.
	TextRect *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect `json:"TextRect,omitempty" xml:"TextRect,omitempty" type:"Struct"`
	// Indicates whether the text is valid.
	//
	// example:
	//
	// true
	Valid *bool `json:"Valid,omitempty" xml:"Valid,omitempty"`
	// The text value.
	//
	// example:
	//
	// Solid Wood Wardrobe.
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The vertical layout.
	//
	// example:
	//
	// center
	VerticalLayout *string `json:"VerticalLayout,omitempty" xml:"VerticalLayout,omitempty"`
}

func (s ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetColor() *string {
	return s.Color
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetFontsize() *int32 {
	return s.Fontsize
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetHorizontalLayout() *string {
	return s.HorizontalLayout
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetImageRect() *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect {
	return s.ImageRect
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetLanguage() *string {
	return s.Language
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetLineCount() *int32 {
	return s.LineCount
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetTextRect() *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect {
	return s.TextRect
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetValid() *bool {
	return s.Valid
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetValue() *string {
	return s.Value
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) GetVerticalLayout() *string {
	return s.VerticalLayout
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetColor(v string) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.Color = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetFontsize(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.Fontsize = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetHorizontalLayout(v string) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.HorizontalLayout = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetImageRect(v *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.ImageRect = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetLanguage(v string) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.Language = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetLineCount(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.LineCount = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetTextRect(v *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.TextRect = v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetValid(v bool) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.Valid = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetValue(v string) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.Value = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) SetVerticalLayout(v string) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts {
	s.VerticalLayout = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTexts) Validate() error {
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

type ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect struct {
	// The rotation angle.
	Degree *int32 `json:"Degree,omitempty" xml:"Degree,omitempty"`
	// The height.
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The left coordinate.
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top coordinate.
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width.
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) GetDegree() *int32 {
	return s.Degree
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) GetHeight() *int32 {
	return s.Height
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) GetLeft() *int32 {
	return s.Left
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) GetTop() *int32 {
	return s.Top
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) GetWidth() *int32 {
	return s.Width
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) SetDegree(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect {
	s.Degree = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) SetHeight(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect {
	s.Height = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) SetLeft(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect {
	s.Left = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) SetTop(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect {
	s.Top = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) SetWidth(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect {
	s.Width = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsImageRect) Validate() error {
	return dara.Validate(s)
}

type ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect struct {
	// The rotation angle.
	Degree *int32 `json:"Degree,omitempty" xml:"Degree,omitempty"`
	// The height.
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// The left coordinate.
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// The top coordinate.
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// The width.
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) String() string {
	return dara.Prettify(s)
}

func (s ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) GoString() string {
	return s.String()
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) GetDegree() *int32 {
	return s.Degree
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) GetHeight() *int32 {
	return s.Height
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) GetLeft() *int32 {
	return s.Left
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) GetTop() *int32 {
	return s.Top
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) GetWidth() *int32 {
	return s.Width
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) SetDegree(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect {
	s.Degree = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) SetHeight(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect {
	s.Height = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) SetLeft(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect {
	s.Left = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) SetTop(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect {
	s.Top = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) SetWidth(v int32) *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect {
	s.Width = &v
	return s
}

func (s *ImageTranslationStandardResponseBodyDataEditInfoTextAreasTextsTextRect) Validate() error {
	return dara.Validate(s)
}
