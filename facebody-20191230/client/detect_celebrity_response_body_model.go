// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectCelebrityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DetectCelebrityResponseBodyData) *DetectCelebrityResponseBody
	GetData() *DetectCelebrityResponseBodyData
	SetRequestId(v string) *DetectCelebrityResponseBody
	GetRequestId() *string
}

type DetectCelebrityResponseBody struct {
	Data *DetectCelebrityResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 4E92C0D4-BB0F-4F25-AD13-E87D02FAA205
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetectCelebrityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetectCelebrityResponseBody) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponseBody) GetData() *DetectCelebrityResponseBodyData {
	return s.Data
}

func (s *DetectCelebrityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetectCelebrityResponseBody) SetData(v *DetectCelebrityResponseBodyData) *DetectCelebrityResponseBody {
	s.Data = v
	return s
}

func (s *DetectCelebrityResponseBody) SetRequestId(v string) *DetectCelebrityResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetectCelebrityResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DetectCelebrityResponseBodyData struct {
	FaceRecognizeResults []*DetectCelebrityResponseBodyDataFaceRecognizeResults `json:"FaceRecognizeResults,omitempty" xml:"FaceRecognizeResults,omitempty" type:"Repeated"`
	// example:
	//
	// 1000
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 1000
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s DetectCelebrityResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DetectCelebrityResponseBodyData) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponseBodyData) GetFaceRecognizeResults() []*DetectCelebrityResponseBodyDataFaceRecognizeResults {
	return s.FaceRecognizeResults
}

func (s *DetectCelebrityResponseBodyData) GetHeight() *int32 {
	return s.Height
}

func (s *DetectCelebrityResponseBodyData) GetWidth() *int32 {
	return s.Width
}

func (s *DetectCelebrityResponseBodyData) SetFaceRecognizeResults(v []*DetectCelebrityResponseBodyDataFaceRecognizeResults) *DetectCelebrityResponseBodyData {
	s.FaceRecognizeResults = v
	return s
}

func (s *DetectCelebrityResponseBodyData) SetHeight(v int32) *DetectCelebrityResponseBodyData {
	s.Height = &v
	return s
}

func (s *DetectCelebrityResponseBodyData) SetWidth(v int32) *DetectCelebrityResponseBodyData {
	s.Width = &v
	return s
}

func (s *DetectCelebrityResponseBodyData) Validate() error {
	if s.FaceRecognizeResults != nil {
		for _, item := range s.FaceRecognizeResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DetectCelebrityResponseBodyDataFaceRecognizeResults struct {
	// 1
	FaceBoxes []*float32 `json:"FaceBoxes,omitempty" xml:"FaceBoxes,omitempty" type:"Repeated"`
	Name      *string    `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DetectCelebrityResponseBodyDataFaceRecognizeResults) String() string {
	return dara.Prettify(s)
}

func (s DetectCelebrityResponseBodyDataFaceRecognizeResults) GoString() string {
	return s.String()
}

func (s *DetectCelebrityResponseBodyDataFaceRecognizeResults) GetFaceBoxes() []*float32 {
	return s.FaceBoxes
}

func (s *DetectCelebrityResponseBodyDataFaceRecognizeResults) GetName() *string {
	return s.Name
}

func (s *DetectCelebrityResponseBodyDataFaceRecognizeResults) SetFaceBoxes(v []*float32) *DetectCelebrityResponseBodyDataFaceRecognizeResults {
	s.FaceBoxes = v
	return s
}

func (s *DetectCelebrityResponseBodyDataFaceRecognizeResults) SetName(v string) *DetectCelebrityResponseBodyDataFaceRecognizeResults {
	s.Name = &v
	return s
}

func (s *DetectCelebrityResponseBodyDataFaceRecognizeResults) Validate() error {
	return dara.Validate(s)
}
