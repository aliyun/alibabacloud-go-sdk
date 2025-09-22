// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageQueueModerationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImageQueueModerationResponseBody
	GetCode() *int32
	SetData(v *ImageQueueModerationResponseBodyData) *ImageQueueModerationResponseBody
	GetData() *ImageQueueModerationResponseBodyData
	SetMsg(v string) *ImageQueueModerationResponseBody
	GetMsg() *string
	SetRequestId(v string) *ImageQueueModerationResponseBody
	GetRequestId() *string
}

type ImageQueueModerationResponseBody struct {
	// example:
	//
	// 200
	Code *int32                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ImageQueueModerationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ImageQueueModerationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBody) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImageQueueModerationResponseBody) GetData() *ImageQueueModerationResponseBodyData {
	return s.Data
}

func (s *ImageQueueModerationResponseBody) GetMsg() *string {
	return s.Msg
}

func (s *ImageQueueModerationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImageQueueModerationResponseBody) SetCode(v int32) *ImageQueueModerationResponseBody {
	s.Code = &v
	return s
}

func (s *ImageQueueModerationResponseBody) SetData(v *ImageQueueModerationResponseBodyData) *ImageQueueModerationResponseBody {
	s.Data = v
	return s
}

func (s *ImageQueueModerationResponseBody) SetMsg(v string) *ImageQueueModerationResponseBody {
	s.Msg = &v
	return s
}

func (s *ImageQueueModerationResponseBody) SetRequestId(v string) *ImageQueueModerationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImageQueueModerationResponseBody) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyData struct {
	// example:
	//
	// data1234
	DataId *string                                  `json:"DataId,omitempty" xml:"DataId,omitempty"`
	Ext    *ImageQueueModerationResponseBodyDataExt `json:"Ext,omitempty" xml:"Ext,omitempty" type:"Struct"`
	// example:
	//
	// xxxxx-xxxxx
	ManualTaskId *string                                       `json:"ManualTaskId,omitempty" xml:"ManualTaskId,omitempty"`
	Result       []*ImageQueueModerationResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageQueueModerationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyData) GetDataId() *string {
	return s.DataId
}

func (s *ImageQueueModerationResponseBodyData) GetExt() *ImageQueueModerationResponseBodyDataExt {
	return s.Ext
}

func (s *ImageQueueModerationResponseBodyData) GetManualTaskId() *string {
	return s.ManualTaskId
}

func (s *ImageQueueModerationResponseBodyData) GetResult() []*ImageQueueModerationResponseBodyDataResult {
	return s.Result
}

func (s *ImageQueueModerationResponseBodyData) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ImageQueueModerationResponseBodyData) SetDataId(v string) *ImageQueueModerationResponseBodyData {
	s.DataId = &v
	return s
}

func (s *ImageQueueModerationResponseBodyData) SetExt(v *ImageQueueModerationResponseBodyDataExt) *ImageQueueModerationResponseBodyData {
	s.Ext = v
	return s
}

func (s *ImageQueueModerationResponseBodyData) SetManualTaskId(v string) *ImageQueueModerationResponseBodyData {
	s.ManualTaskId = &v
	return s
}

func (s *ImageQueueModerationResponseBodyData) SetResult(v []*ImageQueueModerationResponseBodyDataResult) *ImageQueueModerationResponseBodyData {
	s.Result = v
	return s
}

func (s *ImageQueueModerationResponseBodyData) SetRiskLevel(v string) *ImageQueueModerationResponseBodyData {
	s.RiskLevel = &v
	return s
}

func (s *ImageQueueModerationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExt struct {
	CustomImage  []*ImageQueueModerationResponseBodyDataExtCustomImage  `json:"CustomImage,omitempty" xml:"CustomImage,omitempty" type:"Repeated"`
	FaceData     []*ImageQueueModerationResponseBodyDataExtFaceData     `json:"FaceData,omitempty" xml:"FaceData,omitempty" type:"Repeated"`
	LogoData     []*ImageQueueModerationResponseBodyDataExtLogoData     `json:"LogoData,omitempty" xml:"LogoData,omitempty" type:"Repeated"`
	OcrResult    []*ImageQueueModerationResponseBodyDataExtOcrResult    `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Repeated"`
	PublicFigure []*ImageQueueModerationResponseBodyDataExtPublicFigure `json:"PublicFigure,omitempty" xml:"PublicFigure,omitempty" type:"Repeated"`
	Recognition  []*ImageQueueModerationResponseBodyDataExtRecognition  `json:"Recognition,omitempty" xml:"Recognition,omitempty" type:"Repeated"`
	TextInImage  *ImageQueueModerationResponseBodyDataExtTextInImage    `json:"TextInImage,omitempty" xml:"TextInImage,omitempty" type:"Struct"`
	VlContent    *ImageQueueModerationResponseBodyDataExtVlContent      `json:"VlContent,omitempty" xml:"VlContent,omitempty" type:"Struct"`
}

func (s ImageQueueModerationResponseBodyDataExt) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExt) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExt) GetCustomImage() []*ImageQueueModerationResponseBodyDataExtCustomImage {
	return s.CustomImage
}

func (s *ImageQueueModerationResponseBodyDataExt) GetFaceData() []*ImageQueueModerationResponseBodyDataExtFaceData {
	return s.FaceData
}

func (s *ImageQueueModerationResponseBodyDataExt) GetLogoData() []*ImageQueueModerationResponseBodyDataExtLogoData {
	return s.LogoData
}

func (s *ImageQueueModerationResponseBodyDataExt) GetOcrResult() []*ImageQueueModerationResponseBodyDataExtOcrResult {
	return s.OcrResult
}

func (s *ImageQueueModerationResponseBodyDataExt) GetPublicFigure() []*ImageQueueModerationResponseBodyDataExtPublicFigure {
	return s.PublicFigure
}

func (s *ImageQueueModerationResponseBodyDataExt) GetRecognition() []*ImageQueueModerationResponseBodyDataExtRecognition {
	return s.Recognition
}

func (s *ImageQueueModerationResponseBodyDataExt) GetTextInImage() *ImageQueueModerationResponseBodyDataExtTextInImage {
	return s.TextInImage
}

func (s *ImageQueueModerationResponseBodyDataExt) GetVlContent() *ImageQueueModerationResponseBodyDataExtVlContent {
	return s.VlContent
}

func (s *ImageQueueModerationResponseBodyDataExt) SetCustomImage(v []*ImageQueueModerationResponseBodyDataExtCustomImage) *ImageQueueModerationResponseBodyDataExt {
	s.CustomImage = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExt) SetFaceData(v []*ImageQueueModerationResponseBodyDataExtFaceData) *ImageQueueModerationResponseBodyDataExt {
	s.FaceData = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExt) SetLogoData(v []*ImageQueueModerationResponseBodyDataExtLogoData) *ImageQueueModerationResponseBodyDataExt {
	s.LogoData = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExt) SetOcrResult(v []*ImageQueueModerationResponseBodyDataExtOcrResult) *ImageQueueModerationResponseBodyDataExt {
	s.OcrResult = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExt) SetPublicFigure(v []*ImageQueueModerationResponseBodyDataExtPublicFigure) *ImageQueueModerationResponseBodyDataExt {
	s.PublicFigure = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExt) SetRecognition(v []*ImageQueueModerationResponseBodyDataExtRecognition) *ImageQueueModerationResponseBodyDataExt {
	s.Recognition = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExt) SetTextInImage(v *ImageQueueModerationResponseBodyDataExtTextInImage) *ImageQueueModerationResponseBodyDataExt {
	s.TextInImage = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExt) SetVlContent(v *ImageQueueModerationResponseBodyDataExtVlContent) *ImageQueueModerationResponseBodyDataExt {
	s.VlContent = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExt) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtCustomImage struct {
	// example:
	//
	// 123456
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// example:
	//
	// 图库123
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 图库123
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtCustomImage) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtCustomImage) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtCustomImage) GetImageId() *string {
	return s.ImageId
}

func (s *ImageQueueModerationResponseBodyDataExtCustomImage) GetLibId() *string {
	return s.LibId
}

func (s *ImageQueueModerationResponseBodyDataExtCustomImage) GetLibName() *string {
	return s.LibName
}

func (s *ImageQueueModerationResponseBodyDataExtCustomImage) SetImageId(v string) *ImageQueueModerationResponseBodyDataExtCustomImage {
	s.ImageId = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtCustomImage) SetLibId(v string) *ImageQueueModerationResponseBodyDataExtCustomImage {
	s.LibId = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtCustomImage) SetLibName(v string) *ImageQueueModerationResponseBodyDataExtCustomImage {
	s.LibName = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtCustomImage) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtFaceData struct {
	// example:
	//
	// 18
	Age    *int32                                                 `json:"Age,omitempty" xml:"Age,omitempty"`
	Bang   *ImageQueueModerationResponseBodyDataExtFaceDataBang   `json:"Bang,omitempty" xml:"Bang,omitempty" type:"Struct"`
	Gender *ImageQueueModerationResponseBodyDataExtFaceDataGender `json:"Gender,omitempty" xml:"Gender,omitempty" type:"Struct"`
	// example:
	//
	// Common
	Glasses   *string                                                   `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	Hairstyle *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle `json:"Hairstyle,omitempty" xml:"Hairstyle,omitempty" type:"Struct"`
	Hat       *ImageQueueModerationResponseBodyDataExtFaceDataHat       `json:"Hat,omitempty" xml:"Hat,omitempty" type:"Struct"`
	Location  *ImageQueueModerationResponseBodyDataExtFaceDataLocation  `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	Mask      *ImageQueueModerationResponseBodyDataExtFaceDataMask      `json:"Mask,omitempty" xml:"Mask,omitempty" type:"Struct"`
	Mustache  *ImageQueueModerationResponseBodyDataExtFaceDataMustache  `json:"Mustache,omitempty" xml:"Mustache,omitempty" type:"Struct"`
	Quality   *ImageQueueModerationResponseBodyDataExtFaceDataQuality   `json:"Quality,omitempty" xml:"Quality,omitempty" type:"Struct"`
	// example:
	//
	// 85.88
	Smile *float32 `json:"Smile,omitempty" xml:"Smile,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtFaceData) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtFaceData) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetAge() *int32 {
	return s.Age
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetBang() *ImageQueueModerationResponseBodyDataExtFaceDataBang {
	return s.Bang
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetGender() *ImageQueueModerationResponseBodyDataExtFaceDataGender {
	return s.Gender
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetGlasses() *string {
	return s.Glasses
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetHairstyle() *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle {
	return s.Hairstyle
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetHat() *ImageQueueModerationResponseBodyDataExtFaceDataHat {
	return s.Hat
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetLocation() *ImageQueueModerationResponseBodyDataExtFaceDataLocation {
	return s.Location
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetMask() *ImageQueueModerationResponseBodyDataExtFaceDataMask {
	return s.Mask
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetMustache() *ImageQueueModerationResponseBodyDataExtFaceDataMustache {
	return s.Mustache
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetQuality() *ImageQueueModerationResponseBodyDataExtFaceDataQuality {
	return s.Quality
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) GetSmile() *float32 {
	return s.Smile
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetAge(v int32) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Age = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetBang(v *ImageQueueModerationResponseBodyDataExtFaceDataBang) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Bang = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetGender(v *ImageQueueModerationResponseBodyDataExtFaceDataGender) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Gender = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetGlasses(v string) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Glasses = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetHairstyle(v *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Hairstyle = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetHat(v *ImageQueueModerationResponseBodyDataExtFaceDataHat) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Hat = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetLocation(v *ImageQueueModerationResponseBodyDataExtFaceDataLocation) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Location = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetMask(v *ImageQueueModerationResponseBodyDataExtFaceDataMask) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Mask = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetMustache(v *ImageQueueModerationResponseBodyDataExtFaceDataMustache) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Mustache = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetQuality(v *ImageQueueModerationResponseBodyDataExtFaceDataQuality) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Quality = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) SetSmile(v float32) *ImageQueueModerationResponseBodyDataExtFaceData {
	s.Smile = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceData) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtFaceDataBang struct {
	// example:
	//
	// 81.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// Has
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataBang) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataBang) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataBang) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataBang) GetValue() *string {
	return s.Value
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataBang) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataBang {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataBang) SetValue(v string) *ImageQueueModerationResponseBodyDataExtFaceDataBang {
	s.Value = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataBang) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtFaceDataGender struct {
	// example:
	//
	// 81.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// FeMale
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataGender) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataGender) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataGender) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataGender) GetValue() *string {
	return s.Value
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataGender) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataGender {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataGender) SetValue(v string) *ImageQueueModerationResponseBodyDataExtFaceDataGender {
	s.Value = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataGender) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtFaceDataHairstyle struct {
	// example:
	//
	// 81.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// Short
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataHairstyle) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataHairstyle) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle) GetValue() *string {
	return s.Value
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle) SetValue(v string) *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle {
	s.Value = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHairstyle) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtFaceDataHat struct {
	// example:
	//
	// 88.88
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// Wear
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataHat) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataHat) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHat) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHat) GetValue() *string {
	return s.Value
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHat) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataHat {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHat) SetValue(v string) *ImageQueueModerationResponseBodyDataExtFaceDataHat {
	s.Value = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataHat) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtFaceDataLocation struct {
	// example:
	//
	// 26
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 83
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 41
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 84
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataLocation) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataLocation) GetH() *int32 {
	return s.H
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataLocation) GetW() *int32 {
	return s.W
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataLocation) GetX() *int32 {
	return s.X
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataLocation) SetH(v int32) *ImageQueueModerationResponseBodyDataExtFaceDataLocation {
	s.H = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataLocation) SetW(v int32) *ImageQueueModerationResponseBodyDataExtFaceDataLocation {
	s.W = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataLocation) SetX(v int32) *ImageQueueModerationResponseBodyDataExtFaceDataLocation {
	s.X = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataLocation) SetY(v int32) *ImageQueueModerationResponseBodyDataExtFaceDataLocation {
	s.Y = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataLocation) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtFaceDataMask struct {
	// example:
	//
	// 99.99
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// Wear
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataMask) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataMask) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMask) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMask) GetValue() *string {
	return s.Value
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMask) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataMask {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMask) SetValue(v string) *ImageQueueModerationResponseBodyDataExtFaceDataMask {
	s.Value = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMask) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtFaceDataMustache struct {
	// example:
	//
	// 99.99
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// Has
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataMustache) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataMustache) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMustache) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMustache) GetValue() *string {
	return s.Value
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMustache) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataMustache {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMustache) SetValue(v string) *ImageQueueModerationResponseBodyDataExtFaceDataMustache {
	s.Value = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataMustache) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtFaceDataQuality struct {
	// example:
	//
	// 5.88
	Blur *float32 `json:"Blur,omitempty" xml:"Blur,omitempty"`
	// example:
	//
	// 100.0
	Integrity *float32 `json:"Integrity,omitempty" xml:"Integrity,omitempty"`
	// example:
	//
	// 5.88
	Pitch *float32 `json:"Pitch,omitempty" xml:"Pitch,omitempty"`
	// example:
	//
	// 5.18
	Roll *float32 `json:"Roll,omitempty" xml:"Roll,omitempty"`
	// example:
	//
	// 5.18
	Yaw *float32 `json:"Yaw,omitempty" xml:"Yaw,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataQuality) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtFaceDataQuality) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) GetBlur() *float32 {
	return s.Blur
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) GetIntegrity() *float32 {
	return s.Integrity
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) GetPitch() *float32 {
	return s.Pitch
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) GetRoll() *float32 {
	return s.Roll
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) GetYaw() *float32 {
	return s.Yaw
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) SetBlur(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataQuality {
	s.Blur = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) SetIntegrity(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataQuality {
	s.Integrity = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) SetPitch(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataQuality {
	s.Pitch = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) SetRoll(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataQuality {
	s.Roll = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) SetYaw(v float32) *ImageQueueModerationResponseBodyDataExtFaceDataQuality {
	s.Yaw = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtFaceDataQuality) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtLogoData struct {
	Location *ImageQueueModerationResponseBodyDataExtLogoDataLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	Logo     []*ImageQueueModerationResponseBodyDataExtLogoDataLogo   `json:"Logo,omitempty" xml:"Logo,omitempty" type:"Repeated"`
}

func (s ImageQueueModerationResponseBodyDataExtLogoData) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtLogoData) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtLogoData) GetLocation() *ImageQueueModerationResponseBodyDataExtLogoDataLocation {
	return s.Location
}

func (s *ImageQueueModerationResponseBodyDataExtLogoData) GetLogo() []*ImageQueueModerationResponseBodyDataExtLogoDataLogo {
	return s.Logo
}

func (s *ImageQueueModerationResponseBodyDataExtLogoData) SetLocation(v *ImageQueueModerationResponseBodyDataExtLogoDataLocation) *ImageQueueModerationResponseBodyDataExtLogoData {
	s.Location = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtLogoData) SetLogo(v []*ImageQueueModerationResponseBodyDataExtLogoDataLogo) *ImageQueueModerationResponseBodyDataExtLogoData {
	s.Logo = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtLogoData) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtLogoDataLocation struct {
	// example:
	//
	// 440
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 330
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtLogoDataLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtLogoDataLocation) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLocation) GetH() *int32 {
	return s.H
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLocation) GetW() *int32 {
	return s.W
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLocation) GetX() *int32 {
	return s.X
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLocation) SetH(v int32) *ImageQueueModerationResponseBodyDataExtLogoDataLocation {
	s.H = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLocation) SetW(v int32) *ImageQueueModerationResponseBodyDataExtLogoDataLocation {
	s.W = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLocation) SetX(v int32) *ImageQueueModerationResponseBodyDataExtLogoDataLocation {
	s.X = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLocation) SetY(v int32) *ImageQueueModerationResponseBodyDataExtLogoDataLocation {
	s.Y = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLocation) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtLogoDataLogo struct {
	// example:
	//
	// 99.1
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// logo_sns
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 钉钉
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtLogoDataLogo) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtLogoDataLogo) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLogo) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLogo) GetLabel() *string {
	return s.Label
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLogo) GetName() *string {
	return s.Name
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLogo) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataExtLogoDataLogo {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLogo) SetLabel(v string) *ImageQueueModerationResponseBodyDataExtLogoDataLogo {
	s.Label = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLogo) SetName(v string) *ImageQueueModerationResponseBodyDataExtLogoDataLogo {
	s.Name = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtLogoDataLogo) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtOcrResult struct {
	Location *ImageQueueModerationResponseBodyDataExtOcrResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// example:
	//
	// abcd
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtOcrResult) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtOcrResult) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResult) GetLocation() *ImageQueueModerationResponseBodyDataExtOcrResultLocation {
	return s.Location
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResult) GetText() *string {
	return s.Text
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResult) SetLocation(v *ImageQueueModerationResponseBodyDataExtOcrResultLocation) *ImageQueueModerationResponseBodyDataExtOcrResult {
	s.Location = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResult) SetText(v string) *ImageQueueModerationResponseBodyDataExtOcrResult {
	s.Text = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResult) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtOcrResultLocation struct {
	// example:
	//
	// 44
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 33
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtOcrResultLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtOcrResultLocation) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResultLocation) GetH() *int32 {
	return s.H
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResultLocation) GetW() *int32 {
	return s.W
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResultLocation) GetX() *int32 {
	return s.X
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResultLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResultLocation) SetH(v int32) *ImageQueueModerationResponseBodyDataExtOcrResultLocation {
	s.H = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResultLocation) SetW(v int32) *ImageQueueModerationResponseBodyDataExtOcrResultLocation {
	s.W = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResultLocation) SetX(v int32) *ImageQueueModerationResponseBodyDataExtOcrResultLocation {
	s.X = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResultLocation) SetY(v int32) *ImageQueueModerationResponseBodyDataExtOcrResultLocation {
	s.Y = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtOcrResultLocation) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtPublicFigure struct {
	// example:
	//
	// xxx001
	FigureId *string `json:"FigureId,omitempty" xml:"FigureId,omitempty"`
	// example:
	//
	// yzazhzou
	FigureName *string                                                        `json:"FigureName,omitempty" xml:"FigureName,omitempty"`
	Location   []*ImageQueueModerationResponseBodyDataExtPublicFigureLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Repeated"`
}

func (s ImageQueueModerationResponseBodyDataExtPublicFigure) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtPublicFigure) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigure) GetFigureId() *string {
	return s.FigureId
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigure) GetFigureName() *string {
	return s.FigureName
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigure) GetLocation() []*ImageQueueModerationResponseBodyDataExtPublicFigureLocation {
	return s.Location
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigure) SetFigureId(v string) *ImageQueueModerationResponseBodyDataExtPublicFigure {
	s.FigureId = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigure) SetFigureName(v string) *ImageQueueModerationResponseBodyDataExtPublicFigure {
	s.FigureName = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigure) SetLocation(v []*ImageQueueModerationResponseBodyDataExtPublicFigureLocation) *ImageQueueModerationResponseBodyDataExtPublicFigure {
	s.Location = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigure) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtPublicFigureLocation struct {
	// example:
	//
	// 440
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 330
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtPublicFigureLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtPublicFigureLocation) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigureLocation) GetH() *int32 {
	return s.H
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigureLocation) GetW() *int32 {
	return s.W
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigureLocation) GetX() *int32 {
	return s.X
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigureLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigureLocation) SetH(v int32) *ImageQueueModerationResponseBodyDataExtPublicFigureLocation {
	s.H = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigureLocation) SetW(v int32) *ImageQueueModerationResponseBodyDataExtPublicFigureLocation {
	s.W = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigureLocation) SetX(v int32) *ImageQueueModerationResponseBodyDataExtPublicFigureLocation {
	s.X = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigureLocation) SetY(v int32) *ImageQueueModerationResponseBodyDataExtPublicFigureLocation {
	s.Y = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtPublicFigureLocation) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtRecognition struct {
	// example:
	//
	// 办公大楼
	Classification *string `json:"Classification,omitempty" xml:"Classification,omitempty"`
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtRecognition) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtRecognition) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtRecognition) GetClassification() *string {
	return s.Classification
}

func (s *ImageQueueModerationResponseBodyDataExtRecognition) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataExtRecognition) SetClassification(v string) *ImageQueueModerationResponseBodyDataExtRecognition {
	s.Classification = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtRecognition) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataExtRecognition {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtRecognition) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtTextInImage struct {
	CustomText []*ImageQueueModerationResponseBodyDataExtTextInImageCustomText `json:"CustomText,omitempty" xml:"CustomText,omitempty" type:"Repeated"`
	OcrResult  []*ImageQueueModerationResponseBodyDataExtTextInImageOcrResult  `json:"OcrResult,omitempty" xml:"OcrResult,omitempty" type:"Repeated"`
	RiskWord   []*string                                                       `json:"RiskWord,omitempty" xml:"RiskWord,omitempty" type:"Repeated"`
}

func (s ImageQueueModerationResponseBodyDataExtTextInImage) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtTextInImage) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImage) GetCustomText() []*ImageQueueModerationResponseBodyDataExtTextInImageCustomText {
	return s.CustomText
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImage) GetOcrResult() []*ImageQueueModerationResponseBodyDataExtTextInImageOcrResult {
	return s.OcrResult
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImage) GetRiskWord() []*string {
	return s.RiskWord
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImage) SetCustomText(v []*ImageQueueModerationResponseBodyDataExtTextInImageCustomText) *ImageQueueModerationResponseBodyDataExtTextInImage {
	s.CustomText = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImage) SetOcrResult(v []*ImageQueueModerationResponseBodyDataExtTextInImageOcrResult) *ImageQueueModerationResponseBodyDataExtTextInImage {
	s.OcrResult = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImage) SetRiskWord(v []*string) *ImageQueueModerationResponseBodyDataExtTextInImage {
	s.RiskWord = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImage) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtTextInImageCustomText struct {
	// example:
	//
	// 自定义词1,自定义词2
	KeyWords *string `json:"KeyWords,omitempty" xml:"KeyWords,omitempty"`
	// example:
	//
	// 123456
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// 自定义库1
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtTextInImageCustomText) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtTextInImageCustomText) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageCustomText) GetKeyWords() *string {
	return s.KeyWords
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageCustomText) GetLibId() *string {
	return s.LibId
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageCustomText) GetLibName() *string {
	return s.LibName
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageCustomText) SetKeyWords(v string) *ImageQueueModerationResponseBodyDataExtTextInImageCustomText {
	s.KeyWords = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageCustomText) SetLibId(v string) *ImageQueueModerationResponseBodyDataExtTextInImageCustomText {
	s.LibId = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageCustomText) SetLibName(v string) *ImageQueueModerationResponseBodyDataExtTextInImageCustomText {
	s.LibName = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageCustomText) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtTextInImageOcrResult struct {
	Location *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation `json:"Location,omitempty" xml:"Location,omitempty" type:"Struct"`
	// example:
	//
	// abcd
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtTextInImageOcrResult) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtTextInImageOcrResult) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResult) GetLocation() *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation {
	return s.Location
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResult) GetText() *string {
	return s.Text
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResult) SetLocation(v *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) *ImageQueueModerationResponseBodyDataExtTextInImageOcrResult {
	s.Location = v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResult) SetText(v string) *ImageQueueModerationResponseBodyDataExtTextInImageOcrResult {
	s.Text = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResult) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation struct {
	// example:
	//
	// 33
	H *int32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 44
	W *int32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 11
	X *int32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 22
	Y *int32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) GetH() *int32 {
	return s.H
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) GetW() *int32 {
	return s.W
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) GetX() *int32 {
	return s.X
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) GetY() *int32 {
	return s.Y
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) SetH(v int32) *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.H = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) SetW(v int32) *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.W = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) SetX(v int32) *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.X = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) SetY(v int32) *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation {
	s.Y = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtTextInImageOcrResultLocation) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataExtVlContent struct {
	// example:
	//
	// 这是一段描述
	OutputText *string `json:"OutputText,omitempty" xml:"OutputText,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataExtVlContent) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataExtVlContent) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataExtVlContent) GetOutputText() *string {
	return s.OutputText
}

func (s *ImageQueueModerationResponseBodyDataExtVlContent) SetOutputText(v string) *ImageQueueModerationResponseBodyDataExtVlContent {
	s.OutputText = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataExtVlContent) Validate() error {
	return dara.Validate(s)
}

type ImageQueueModerationResponseBodyDataResult struct {
	// example:
	//
	// 81.22
	Confidence *float32 `json:"Confidence,omitempty" xml:"Confidence,omitempty"`
	// example:
	//
	// 未检测出风险
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// violent_explosion
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// high
	RiskLevel *string `json:"RiskLevel,omitempty" xml:"RiskLevel,omitempty"`
}

func (s ImageQueueModerationResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationResponseBodyDataResult) GetConfidence() *float32 {
	return s.Confidence
}

func (s *ImageQueueModerationResponseBodyDataResult) GetDescription() *string {
	return s.Description
}

func (s *ImageQueueModerationResponseBodyDataResult) GetLabel() *string {
	return s.Label
}

func (s *ImageQueueModerationResponseBodyDataResult) GetRiskLevel() *string {
	return s.RiskLevel
}

func (s *ImageQueueModerationResponseBodyDataResult) SetConfidence(v float32) *ImageQueueModerationResponseBodyDataResult {
	s.Confidence = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataResult) SetDescription(v string) *ImageQueueModerationResponseBodyDataResult {
	s.Description = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataResult) SetLabel(v string) *ImageQueueModerationResponseBodyDataResult {
	s.Label = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataResult) SetRiskLevel(v string) *ImageQueueModerationResponseBodyDataResult {
	s.RiskLevel = &v
	return s
}

func (s *ImageQueueModerationResponseBodyDataResult) Validate() error {
	return dara.Validate(s)
}
