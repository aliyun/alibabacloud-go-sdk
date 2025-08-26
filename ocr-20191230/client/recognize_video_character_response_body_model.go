// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizeVideoCharacterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizeVideoCharacterResponseBodyData) *RecognizeVideoCharacterResponseBody
	GetData() *RecognizeVideoCharacterResponseBodyData
	SetMessage(v string) *RecognizeVideoCharacterResponseBody
	GetMessage() *string
	SetRequestId(v string) *RecognizeVideoCharacterResponseBody
	GetRequestId() *string
}

type RecognizeVideoCharacterResponseBody struct {
	Data    *RecognizeVideoCharacterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D3F5BA69-79C4-46A4-B02B-58C4EEBC4C33
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizeVideoCharacterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBody) GetData() *RecognizeVideoCharacterResponseBodyData {
	return s.Data
}

func (s *RecognizeVideoCharacterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RecognizeVideoCharacterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizeVideoCharacterResponseBody) SetData(v *RecognizeVideoCharacterResponseBodyData) *RecognizeVideoCharacterResponseBody {
	s.Data = v
	return s
}

func (s *RecognizeVideoCharacterResponseBody) SetMessage(v string) *RecognizeVideoCharacterResponseBody {
	s.Message = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBody) SetRequestId(v string) *RecognizeVideoCharacterResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBody) Validate() error {
	return dara.Validate(s)
}

type RecognizeVideoCharacterResponseBodyData struct {
	Frames []*RecognizeVideoCharacterResponseBodyDataFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	// example:
	//
	// 1080
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// oss://my-bucket/a/b/c.mp4
	InputFile *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	// example:
	//
	// 1920
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeVideoCharacterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBodyData) GetFrames() []*RecognizeVideoCharacterResponseBodyDataFrames {
	return s.Frames
}

func (s *RecognizeVideoCharacterResponseBodyData) GetHeight() *int64 {
	return s.Height
}

func (s *RecognizeVideoCharacterResponseBodyData) GetInputFile() *string {
	return s.InputFile
}

func (s *RecognizeVideoCharacterResponseBodyData) GetWidth() *int64 {
	return s.Width
}

func (s *RecognizeVideoCharacterResponseBodyData) SetFrames(v []*RecognizeVideoCharacterResponseBodyDataFrames) *RecognizeVideoCharacterResponseBodyData {
	s.Frames = v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyData) SetHeight(v int64) *RecognizeVideoCharacterResponseBodyData {
	s.Height = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyData) SetInputFile(v string) *RecognizeVideoCharacterResponseBodyData {
	s.InputFile = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyData) SetWidth(v int64) *RecognizeVideoCharacterResponseBodyData {
	s.Width = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type RecognizeVideoCharacterResponseBodyDataFrames struct {
	Elements []*RecognizeVideoCharacterResponseBodyDataFramesElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// example:
	//
	// 6124533574
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s RecognizeVideoCharacterResponseBodyDataFrames) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBodyDataFrames) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBodyDataFrames) GetElements() []*RecognizeVideoCharacterResponseBodyDataFramesElements {
	return s.Elements
}

func (s *RecognizeVideoCharacterResponseBodyDataFrames) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *RecognizeVideoCharacterResponseBodyDataFrames) SetElements(v []*RecognizeVideoCharacterResponseBodyDataFramesElements) *RecognizeVideoCharacterResponseBodyDataFrames {
	s.Elements = v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFrames) SetTimestamp(v int64) *RecognizeVideoCharacterResponseBodyDataFrames {
	s.Timestamp = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFrames) Validate() error {
	return dara.Validate(s)
}

type RecognizeVideoCharacterResponseBodyDataFramesElements struct {
	// example:
	//
	// 0.99
	Score          *float32                                                               `json:"Score,omitempty" xml:"Score,omitempty"`
	Text           *string                                                                `json:"Text,omitempty" xml:"Text,omitempty"`
	TextRectangles []*RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles `json:"TextRectangles,omitempty" xml:"TextRectangles,omitempty" type:"Repeated"`
}

func (s RecognizeVideoCharacterResponseBodyDataFramesElements) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBodyDataFramesElements) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) GetScore() *float32 {
	return s.Score
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) GetText() *string {
	return s.Text
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) GetTextRectangles() []*RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	return s.TextRectangles
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) SetScore(v float32) *RecognizeVideoCharacterResponseBodyDataFramesElements {
	s.Score = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) SetText(v string) *RecognizeVideoCharacterResponseBodyDataFramesElements {
	s.Text = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) SetTextRectangles(v []*RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) *RecognizeVideoCharacterResponseBodyDataFramesElements {
	s.TextRectangles = v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElements) Validate() error {
	return dara.Validate(s)
}

type RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles struct {
	// example:
	//
	// -90
	Angle *int64 `json:"Angle,omitempty" xml:"Angle,omitempty"`
	// example:
	//
	// 213
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 213
	Left *int64 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 98
	Top *int64 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 46
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) String() string {
	return dara.Prettify(s)
}

func (s RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) GoString() string {
	return s.String()
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) GetAngle() *int64 {
	return s.Angle
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) GetHeight() *int64 {
	return s.Height
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) GetLeft() *int64 {
	return s.Left
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) GetTop() *int64 {
	return s.Top
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) GetWidth() *int64 {
	return s.Width
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetAngle(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Angle = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetHeight(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Height = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetLeft(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Left = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetTop(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Top = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) SetWidth(v int64) *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles {
	s.Width = &v
	return s
}

func (s *RecognizeVideoCharacterResponseBodyDataFramesElementsTextRectangles) Validate() error {
	return dara.Validate(s)
}
