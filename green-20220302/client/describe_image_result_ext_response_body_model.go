// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageResultExtResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeImageResultExtResponseBody
	GetCode() *int32
	SetData(v *DescribeImageResultExtResponseBodyData) *DescribeImageResultExtResponseBody
	GetData() *DescribeImageResultExtResponseBodyData
	SetMsg(v string) *DescribeImageResultExtResponseBody
	GetMsg() *string
	SetRequestId(v string) *DescribeImageResultExtResponseBody
	GetRequestId() *string
}

type DescribeImageResultExtResponseBody struct {
	// The returned HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *DescribeImageResultExtResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message that is returned in response to the request.
	//
	// example:
	//
	// success
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6CF2815C-C8C7-4A01-B52E-FF6E24F53492
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeImageResultExtResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResultExtResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeImageResultExtResponseBody) GetData() *DescribeImageResultExtResponseBodyData {
	return s.Data
}

func (s *DescribeImageResultExtResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *DescribeImageResultExtResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageResultExtResponseBody) SetCode(v int32) *DescribeImageResultExtResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeImageResultExtResponseBody) SetData(v *DescribeImageResultExtResponseBodyData) *DescribeImageResultExtResponseBody {
	s.Data = v
	return s
}

func (s *DescribeImageResultExtResponseBody) SetMsg(v string) *DescribeImageResultExtResponseBody {
	s.Msg = &v
	return s
}

func (s *DescribeImageResultExtResponseBody) SetRequestId(v string) *DescribeImageResultExtResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageResultExtResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeImageResultExtResponseBodyData struct {
	// If a custom image library is hit, information about the hit custom image library is returned.
	CustomImage []*DescribeImageResultExtResponseBodyDataCustomImage `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	// Person information list.
	PublicFigure []*DescribeImageResultExtResponseBodyDataPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	// Returns the text information in the hit image.
	TextInImage *DescribeImageResultExtResponseBodyDataTextInImage `json:"TextInImage,omitempty" xml:"TextInImage,omitempty" type:"Struct"`
}

func (s DescribeImageResultExtResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyData) GetCustomImage() []*DescribeImageResultExtResponseBodyDataCustomImage {
	return s.CustomImage
}

func (s *DescribeImageResultExtResponseBodyData) GetPublicFigure() []*DescribeImageResultExtResponseBodyDataPublicFigure {
	return s.PublicFigure
}

func (s *DescribeImageResultExtResponseBodyData) GetTextInImage() *DescribeImageResultExtResponseBodyDataTextInImage {
	return s.TextInImage
}

func (s *DescribeImageResultExtResponseBodyData) SetCustomImage(v []*DescribeImageResultExtResponseBodyDataCustomImage) *DescribeImageResultExtResponseBodyData {
	s.CustomImage = v
	return s
}

func (s *DescribeImageResultExtResponseBodyData) SetPublicFigure(v []*DescribeImageResultExtResponseBodyDataPublicFigure) *DescribeImageResultExtResponseBodyData {
	s.PublicFigure = v
	return s
}

func (s *DescribeImageResultExtResponseBodyData) SetTextInImage(v *DescribeImageResultExtResponseBodyDataTextInImage) *DescribeImageResultExtResponseBodyData {
	s.TextInImage = v
	return s
}

func (s *DescribeImageResultExtResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeImageResultExtResponseBodyDataCustomImage struct {
	// The image ID.
	//
	// example:
	//
	// 123456
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image library ID.
	//
	// example:
	//
	// 123456
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The image library name.
	//
	// example:
	//
	// 图库123
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s DescribeImageResultExtResponseBodyDataCustomImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyDataCustomImage) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) GetLibId() *string {
	return s.LibId
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) GetLibName() *string {
	return s.LibName
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) SetImageId(v string) *DescribeImageResultExtResponseBodyDataCustomImage {
	s.ImageId = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) SetLibId(v string) *DescribeImageResultExtResponseBodyDataCustomImage {
	s.LibId = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) SetLibName(v string) *DescribeImageResultExtResponseBodyDataCustomImage {
	s.LibName = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataCustomImage) Validate() error {
	return dara.Validate(s)
}

type DescribeImageResultExtResponseBodyDataPublicFigure struct {
	// Identified person coding information.
	//
	// example:
	//
	// yzazhzou
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
}

func (s DescribeImageResultExtResponseBodyDataPublicFigure) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyDataPublicFigure) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyDataPublicFigure) GetFigureId() *string {
	return s.FigureId
}

func (s *DescribeImageResultExtResponseBodyDataPublicFigure) SetFigureId(v string) *DescribeImageResultExtResponseBodyDataPublicFigure {
	s.FigureId = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataPublicFigure) Validate() error {
	return dara.Validate(s)
}

type DescribeImageResultExtResponseBodyDataTextInImage struct {
	// When a custom text library is hit, the custom library ID, custom library name, and custom word are returned.
	CustomTexts []*DescribeImageResultExtResponseBodyDataTextInImageCustomTexts `json:"CustomTexts,omitempty" xml:"CustomTexts,omitempty" type:"Repeated"`
	// Returns the text information in the recognized image.
	OcrDatas []*string `json:"OcrDatas,omitempty" xml:"OcrDatas,omitempty" type:"Repeated"`
	// The risk words that are hit. Multiple words are separated by commas (,).
	RiskWords []*string `json:"RiskWords,omitempty" xml:"RiskWords,omitempty" type:"Repeated"`
}

func (s DescribeImageResultExtResponseBodyDataTextInImage) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyDataTextInImage) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) GetCustomTexts() []*DescribeImageResultExtResponseBodyDataTextInImageCustomTexts {
	return s.CustomTexts
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) GetOcrDatas() []*string {
	return s.OcrDatas
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) GetRiskWords() []*string {
	return s.RiskWords
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) SetCustomTexts(v []*DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) *DescribeImageResultExtResponseBodyDataTextInImage {
	s.CustomTexts = v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) SetOcrDatas(v []*string) *DescribeImageResultExtResponseBodyDataTextInImage {
	s.OcrDatas = v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) SetRiskWords(v []*string) *DescribeImageResultExtResponseBodyDataTextInImage {
	s.RiskWords = v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImage) Validate() error {
	return dara.Validate(s)
}

type DescribeImageResultExtResponseBodyDataTextInImageCustomTexts struct {
	// Custom words, multiple words separated by commas.
	//
	// example:
	//
	// aaa,bbb
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// Custom library ID.
	//
	// example:
	//
	// 123456
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// Custom library name.
	//
	// example:
	//
	// test
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) GoString() string {
	return s.String()
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) GetKeyWords() *string {
	return s.KeyWords
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) GetLibId() *string {
	return s.LibId
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) GetLibName() *string {
	return s.LibName
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) SetKeyWords(v string) *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts {
	s.KeyWords = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) SetLibId(v string) *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts {
	s.LibId = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) SetLibName(v string) *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts {
	s.LibName = &v
	return s
}

func (s *DescribeImageResultExtResponseBodyDataTextInImageCustomTexts) Validate() error {
	return dara.Validate(s)
}
