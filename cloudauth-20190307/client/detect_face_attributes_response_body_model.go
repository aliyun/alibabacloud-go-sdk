// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectFaceAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DetectFaceAttributesResponseBody
	GetCode() *string
	SetData(v *DetectFaceAttributesResponseBodyData) *DetectFaceAttributesResponseBody
	GetData() *DetectFaceAttributesResponseBodyData
	SetMessage(v string) *DetectFaceAttributesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetectFaceAttributesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DetectFaceAttributesResponseBody
	GetSuccess() *bool
}

type DetectFaceAttributesResponseBody struct {
	// HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Returned data.
	Data *DetectFaceAttributesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error code.
	//
	// example:
	//
	// Error.InternalError
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// ID of this request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the response was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DetectFaceAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DetectFaceAttributesResponseBody) GetData() *DetectFaceAttributesResponseBodyData {
	return s.Data
}

func (s *DetectFaceAttributesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetectFaceAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectFaceAttributesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DetectFaceAttributesResponseBody) SetCode(v string) *DetectFaceAttributesResponseBody {
	s.Code = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetData(v *DetectFaceAttributesResponseBodyData) *DetectFaceAttributesResponseBody {
	s.Data = v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetMessage(v string) *DetectFaceAttributesResponseBody {
	s.Message = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetRequestId(v string) *DetectFaceAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) SetSuccess(v bool) *DetectFaceAttributesResponseBody {
	s.Success = &v
	return s
}

func (s *DetectFaceAttributesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectFaceAttributesResponseBodyData struct {
	// Face detection results.
	FaceInfos *DetectFaceAttributesResponseBodyDataFaceInfos `json:"FaceInfos,omitempty" xml:"FaceInfos,omitempty" type:"Struct"`
	// Original image height, in pixels.
	//
	// example:
	//
	// 1920
	ImgHeight *int32 `json:"ImgHeight,omitempty" xml:"ImgHeight,omitempty"`
	// Original image width, in pixels.
	//
	// example:
	//
	// 1080
	ImgWidth *int32 `json:"ImgWidth,omitempty" xml:"ImgWidth,omitempty"`
}

func (s DetectFaceAttributesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyData) GetFaceInfos() *DetectFaceAttributesResponseBodyDataFaceInfos {
	return s.FaceInfos
}

func (s *DetectFaceAttributesResponseBodyData) GetImgHeight() *int32 {
	return s.ImgHeight
}

func (s *DetectFaceAttributesResponseBodyData) GetImgWidth() *int32 {
	return s.ImgWidth
}

func (s *DetectFaceAttributesResponseBodyData) SetFaceInfos(v *DetectFaceAttributesResponseBodyDataFaceInfos) *DetectFaceAttributesResponseBodyData {
	s.FaceInfos = v
	return s
}

func (s *DetectFaceAttributesResponseBodyData) SetImgHeight(v int32) *DetectFaceAttributesResponseBodyData {
	s.ImgHeight = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyData) SetImgWidth(v int32) *DetectFaceAttributesResponseBodyData {
	s.ImgWidth = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyData) Validate() error {
	if s.FaceInfos != nil {
		if err := s.FaceInfos.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectFaceAttributesResponseBodyDataFaceInfos struct {
	FaceAttributesDetectInfo []*DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo `json:"FaceAttributesDetectInfo,omitempty" xml:"FaceAttributesDetectInfo,omitempty" type:"Repeated"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfos) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfos) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfos) GetFaceAttributesDetectInfo() []*DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo {
	return s.FaceAttributesDetectInfo
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfos) SetFaceAttributesDetectInfo(v []*DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) *DetectFaceAttributesResponseBodyDataFaceInfos {
	s.FaceAttributesDetectInfo = v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfos) Validate() error {
	if s.FaceAttributesDetectInfo != nil {
		for _, item := range s.FaceAttributesDetectInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo struct {
	// Face attributes.
	FaceAttributes *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes `json:"FaceAttributes,omitempty" xml:"FaceAttributes,omitempty" type:"Struct"`
	// Position of the face in the original image.
	FaceRect *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect `json:"FaceRect,omitempty" xml:"FaceRect,omitempty" type:"Struct"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) GetFaceAttributes() *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	return s.FaceAttributes
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) GetFaceRect() *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	return s.FaceRect
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) SetFaceAttributes(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo {
	s.FaceAttributes = v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) SetFaceRect(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo {
	s.FaceRect = v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfo) Validate() error {
	if s.FaceAttributes != nil {
		if err := s.FaceAttributes.Validate(); err != nil {
			return err
		}
	}
	if s.FaceRect != nil {
		if err := s.FaceRect.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes struct {
	// Face blur level, with higher values indicating more blurriness. Typically, a value ≥2.0 is considered quite blurry. It is recommended to adjust based on actual business data testing.
	//
	// example:
	//
	// 0.1419367790222168
	Blur *float32 `json:"Blur,omitempty" xml:"Blur,omitempty"`
	// A score ranging from 0 to 100. The higher the score, the better the quality. It is recommended to consider a score of 50 or above as acceptable quality.
	//
	// example:
	//
	// 60
	Facequal *float32 `json:"Facequal,omitempty" xml:"Facequal,omitempty"`
	// Whether it is a human face. Values:
	//
	// - **None**: Not a human face
	//
	// - **Face**: Human face
	//
	// - **Profile**: Profile (head turned left or right by more than 30°)
	//
	// > If no face is detected, the `faceInfos` in the response will be empty; here, `None` means that a face was detected but is considered to be a cartoon, pet, etc.
	//
	// example:
	//
	// Face
	Facetype *string `json:"Facetype,omitempty" xml:"Facetype,omitempty"`
	// Whether wearing glasses. Values:
	//
	// - **None**: Not wearing glasses
	//
	// - **Wear**: Wearing regular glasses
	//
	// - **Sunglass**: Wearing sunglasses
	//
	// example:
	//
	// None
	Glasses *string `json:"Glasses,omitempty" xml:"Glasses,omitempty"`
	// Face pose.
	Headpose *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose `json:"Headpose,omitempty" xml:"Headpose,omitempty" type:"Struct"`
	// A score ranging from 0 to 100. The higher the score, the more complete the face. It is recommended to consider a score of 70 or above as acceptable completeness.
	//
	// example:
	//
	// 70
	Integrity *int32 `json:"Integrity,omitempty" xml:"Integrity,omitempty"`
	// Whether wearing a mask. Values:
	//
	// - Wear: Wearing a mask.
	//
	// - None: Not wearing a mask.
	//
	// example:
	//
	// Wear
	Respirator *string `json:"Respirator,omitempty" xml:"Respirator,omitempty"`
	// Whether smiling.
	Smiling *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling `json:"Smiling,omitempty" xml:"Smiling,omitempty" type:"Struct"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GetBlur() *float32 {
	return s.Blur
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GetFacequal() *float32 {
	return s.Facequal
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GetFacetype() *string {
	return s.Facetype
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GetGlasses() *string {
	return s.Glasses
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GetHeadpose() *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	return s.Headpose
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GetIntegrity() *int32 {
	return s.Integrity
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GetRespirator() *string {
	return s.Respirator
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) GetSmiling() *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	return s.Smiling
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetBlur(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Blur = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacequal(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facequal = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetFacetype(v string) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Facetype = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetGlasses(v string) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Glasses = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetHeadpose(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Headpose = v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetIntegrity(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Integrity = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetRespirator(v string) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Respirator = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) SetSmiling(v *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes {
	s.Smiling = v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributes) Validate() error {
	if s.Headpose != nil {
		if err := s.Headpose.Validate(); err != nil {
			return err
		}
	}
	if s.Smiling != nil {
		if err := s.Smiling.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose struct {
	// Pitch angle, in degrees.
	//
	// example:
	//
	// -1.5683923959732056
	PitchAngle *float32 `json:"PitchAngle,omitempty" xml:"PitchAngle,omitempty"`
	// Roll angle, in degrees.
	//
	// example:
	//
	// 7.163370132446289
	RollAngle *float32 `json:"RollAngle,omitempty" xml:"RollAngle,omitempty"`
	// Yaw angle, in degrees.
	//
	// example:
	//
	// -6.925303936004639
	YawAngle *float32 `json:"YawAngle,omitempty" xml:"YawAngle,omitempty"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) GetPitchAngle() *float32 {
	return s.PitchAngle
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) GetRollAngle() *float32 {
	return s.RollAngle
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) GetYawAngle() *float32 {
	return s.YawAngle
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetPitchAngle(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.PitchAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetRollAngle(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.RollAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) SetYawAngle(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose {
	s.YawAngle = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesHeadpose) Validate() error {
	return dara.Validate(s)
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling struct {
	// Smile threshold.
	//
	// example:
	//
	// 95
	Threshold *float32 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// Smile score.
	//
	// example:
	//
	// 97
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) GetThreshold() *float32 {
	return s.Threshold
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) GetValue() *float32 {
	return s.Value
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetThreshold(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Threshold = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) SetValue(v float32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling {
	s.Value = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceAttributesSmiling) Validate() error {
	return dara.Validate(s)
}

type DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect struct {
	// Height of the face rectangle.
	//
	// example:
	//
	// 473
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// Distance from the top-left corner of the face rectangle to the left edge of the original image, in pixels.
	//
	// example:
	//
	// 354
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// Distance from the top-left corner of the face rectangle to the top edge of the original image, in pixels.
	//
	// example:
	//
	// 453
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// Width of the face rectangle.
	//
	// example:
	//
	// 473
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) String() string {
	return dara.Prettify(s)
}

func (s DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) GoString() string {
	return s.String()
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) GetHeight() *int32 {
	return s.Height
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) GetLeft() *int32 {
	return s.Left
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) GetTop() *int32 {
	return s.Top
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) GetWidth() *int32 {
	return s.Width
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetHeight(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Height = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetLeft(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Left = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetTop(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Top = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) SetWidth(v int32) *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect {
	s.Width = &v
	return s
}

func (s *DetectFaceAttributesResponseBodyDataFaceInfosFaceAttributesDetectInfoFaceRect) Validate() error {
	return dara.Validate(s)
}
