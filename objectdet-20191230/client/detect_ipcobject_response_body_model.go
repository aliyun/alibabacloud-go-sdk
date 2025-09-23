// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectIPCObjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectIPCObjectResponseBodyData) *DetectIPCObjectResponseBody
	GetData() *DetectIPCObjectResponseBodyData
	SetRequestId(v string) *DetectIPCObjectResponseBody
	GetRequestId() *string
}

type DetectIPCObjectResponseBody struct {
	Data *DetectIPCObjectResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 7AE23740-A3E5-5607-8E10-895DCBD4C260
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectIPCObjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectIPCObjectResponseBody) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponseBody) GetData() *DetectIPCObjectResponseBodyData {
	return s.Data
}

func (s *DetectIPCObjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectIPCObjectResponseBody) SetData(v *DetectIPCObjectResponseBodyData) *DetectIPCObjectResponseBody {
	s.Data = v
	return s
}

func (s *DetectIPCObjectResponseBody) SetRequestId(v string) *DetectIPCObjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectIPCObjectResponseBody) Validate() error {
	return dara.Validate(s)
}

type DetectIPCObjectResponseBodyData struct {
	Elements []*DetectIPCObjectResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// example:
	//
	// 1200
	Height *int64 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 1600
	Width *int64 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectIPCObjectResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectIPCObjectResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponseBodyData) GetElements() []*DetectIPCObjectResponseBodyDataElements {
	return s.Elements
}

func (s *DetectIPCObjectResponseBodyData) GetHeight() *int64 {
	return s.Height
}

func (s *DetectIPCObjectResponseBodyData) GetWidth() *int64 {
	return s.Width
}

func (s *DetectIPCObjectResponseBodyData) SetElements(v []*DetectIPCObjectResponseBodyDataElements) *DetectIPCObjectResponseBodyData {
	s.Elements = v
	return s
}

func (s *DetectIPCObjectResponseBodyData) SetHeight(v int64) *DetectIPCObjectResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectIPCObjectResponseBodyData) SetWidth(v int64) *DetectIPCObjectResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectIPCObjectResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DetectIPCObjectResponseBodyDataElements struct {
	Box []*int64 `json:"Box,omitempty" xml:"Box,omitempty" type:"Repeated"`
	// example:
	//
	// 0.7138671875
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
	// example:
	//
	// 0.8566723958333333
	TargetRate *float32 `json:"TargetRate,omitempty" xml:"TargetRate,omitempty"`
	// example:
	//
	// DOG
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DetectIPCObjectResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s DetectIPCObjectResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponseBodyDataElements) GetBox() []*int64 {
	return s.Box
}

func (s *DetectIPCObjectResponseBodyDataElements) GetScore() *float32 {
	return s.Score
}

func (s *DetectIPCObjectResponseBodyDataElements) GetTargetRate() *float32 {
	return s.TargetRate
}

func (s *DetectIPCObjectResponseBodyDataElements) GetType() *string {
	return s.Type
}

func (s *DetectIPCObjectResponseBodyDataElements) SetBox(v []*int64) *DetectIPCObjectResponseBodyDataElements {
	s.Box = v
	return s
}

func (s *DetectIPCObjectResponseBodyDataElements) SetScore(v float32) *DetectIPCObjectResponseBodyDataElements {
	s.Score = &v
	return s
}

func (s *DetectIPCObjectResponseBodyDataElements) SetTargetRate(v float32) *DetectIPCObjectResponseBodyDataElements {
	s.TargetRate = &v
	return s
}

func (s *DetectIPCObjectResponseBodyDataElements) SetType(v string) *DetectIPCObjectResponseBodyDataElements {
	s.Type = &v
	return s
}

func (s *DetectIPCObjectResponseBodyDataElements) Validate() error {
	return dara.Validate(s)
}
