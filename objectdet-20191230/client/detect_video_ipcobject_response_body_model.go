// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVideoIPCObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectVideoIPCObjectResponseBodyData) *DetectVideoIPCObjectResponseBody
	GetData() *DetectVideoIPCObjectResponseBodyData
	SetMessage(v string) *DetectVideoIPCObjectResponseBody
	GetMessage() *string
	SetRequestId(v string) *DetectVideoIPCObjectResponseBody
	GetRequestId() *string
}

type DetectVideoIPCObjectResponseBody struct {
	Data    *DetectVideoIPCObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 35B11E1B-800C-4598-B5AA-577E3BDBD917
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectVideoIPCObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoIPCObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponseBody) GetData() *DetectVideoIPCObjectResponseBodyData {
	return s.Data
}

func (s *DetectVideoIPCObjectResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DetectVideoIPCObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectVideoIPCObjectResponseBody) SetData(v *DetectVideoIPCObjectResponseBodyData) *DetectVideoIPCObjectResponseBody {
	s.Data = v
	return s
}

func (s *DetectVideoIPCObjectResponseBody) SetMessage(v string) *DetectVideoIPCObjectResponseBody {
	s.Message = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBody) SetRequestId(v string) *DetectVideoIPCObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectVideoIPCObjectResponseBodyData struct {
	Frames []*DetectVideoIPCObjectResponseBodyDataFrames `json:"Frames,omitempty" xml:"Frames,omitempty" type:"Repeated"`
	// example:
	//
	// 720
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// oss://viapi-test/viapi-3.0domepic/objectdet/DetectVideoIPCObject/DetectVideoIPCObject4.mp4
	InputFile *string `json:"InputFile,omitempty" xml:"InputFile,omitempty"`
	// example:
	//
	// 1280
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectVideoIPCObjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoIPCObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponseBodyData) GetFrames() []*DetectVideoIPCObjectResponseBodyDataFrames {
	return s.Frames
}

func (s *DetectVideoIPCObjectResponseBodyData) GetHeight() *int64 {
	return s.Height
}

func (s *DetectVideoIPCObjectResponseBodyData) GetInputFile() *string {
	return s.InputFile
}

func (s *DetectVideoIPCObjectResponseBodyData) GetWidth() *int64 {
	return s.Width
}

func (s *DetectVideoIPCObjectResponseBodyData) SetFrames(v []*DetectVideoIPCObjectResponseBodyDataFrames) *DetectVideoIPCObjectResponseBodyData {
	s.Frames = v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyData) SetHeight(v int64) *DetectVideoIPCObjectResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyData) SetInputFile(v string) *DetectVideoIPCObjectResponseBodyData {
	s.InputFile = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyData) SetWidth(v int64) *DetectVideoIPCObjectResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DetectVideoIPCObjectResponseBodyDataFrames struct {
	Elements []*DetectVideoIPCObjectResponseBodyDataFramesElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// example:
	//
	// 6124533574
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
}

func (s DetectVideoIPCObjectResponseBodyDataFrames) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoIPCObjectResponseBodyDataFrames) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponseBodyDataFrames) GetElements() []*DetectVideoIPCObjectResponseBodyDataFramesElements {
	return s.Elements
}

func (s *DetectVideoIPCObjectResponseBodyDataFrames) GetTime() *int64 {
	return s.Time
}

func (s *DetectVideoIPCObjectResponseBodyDataFrames) SetElements(v []*DetectVideoIPCObjectResponseBodyDataFramesElements) *DetectVideoIPCObjectResponseBodyDataFrames {
	s.Elements = v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFrames) SetTime(v int64) *DetectVideoIPCObjectResponseBodyDataFrames {
	s.Time = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFrames) Validate() error {
	return dara.Validate(s)
}

type DetectVideoIPCObjectResponseBodyDataFramesElements struct {
	// example:
	//
	// 156
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 0.7812
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// PERSON
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 100
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
	// example:
	//
	// 289
	X *int64 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 271
	Y *int64 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s DetectVideoIPCObjectResponseBodyDataFramesElements) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoIPCObjectResponseBodyDataFramesElements) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) GetHeight() *int64 {
	return s.Height
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) GetScore() *float32 {
	return s.Score
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) GetType() *string {
	return s.Type
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) GetWidth() *int64 {
	return s.Width
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) GetX() *int64 {
	return s.X
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) GetY() *int64 {
	return s.Y
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetHeight(v int64) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Height = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetScore(v float32) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Score = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetType(v string) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Type = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetWidth(v int64) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Width = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetX(v int64) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.X = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) SetY(v int64) *DetectVideoIPCObjectResponseBodyDataFramesElements {
	s.Y = &v
	return s
}

func (s *DetectVideoIPCObjectResponseBodyDataFramesElements) Validate() error {
	return dara.Validate(s)
}
