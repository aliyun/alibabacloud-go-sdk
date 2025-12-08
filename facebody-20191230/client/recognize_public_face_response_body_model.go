// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecognizePublicFaceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RecognizePublicFaceResponseBodyData) *RecognizePublicFaceResponseBody
	GetData() *RecognizePublicFaceResponseBodyData
	SetRequestId(v string) *RecognizePublicFaceResponseBody
	GetRequestId() *string
}

type RecognizePublicFaceResponseBody struct {
	Data *RecognizePublicFaceResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// AC4D107C-29E3-4873-A719-0D2217EA28A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecognizePublicFaceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBody) GetData() *RecognizePublicFaceResponseBodyData {
	return s.Data
}

func (s *RecognizePublicFaceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecognizePublicFaceResponseBody) SetData(v *RecognizePublicFaceResponseBodyData) *RecognizePublicFaceResponseBody {
	s.Data = v
	return s
}

func (s *RecognizePublicFaceResponseBody) SetRequestId(v string) *RecognizePublicFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecognizePublicFaceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RecognizePublicFaceResponseBodyData struct {
	Elements []*RecognizePublicFaceResponseBodyDataElements `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
}

func (s RecognizePublicFaceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyData) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyData) GetElements() []*RecognizePublicFaceResponseBodyDataElements {
	return s.Elements
}

func (s *RecognizePublicFaceResponseBodyData) SetElements(v []*RecognizePublicFaceResponseBodyDataElements) *RecognizePublicFaceResponseBodyData {
	s.Elements = v
	return s
}

func (s *RecognizePublicFaceResponseBodyData) Validate() error {
	if s.Elements != nil {
		for _, item := range s.Elements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizePublicFaceResponseBodyDataElements struct {
	// example:
	//
	// https://viapi-oss.oss-cn-shanghai.aliyuncs.com/doc/facebody/xxx.jpg
	ImageURL *string                                               `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
	Results  []*RecognizePublicFaceResponseBodyDataElementsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// example:
	//
	// img3dhYqt1e4wO77Wnf2y@t@E-1tYAEt
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElements) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElements) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElements) GetImageURL() *string {
	return s.ImageURL
}

func (s *RecognizePublicFaceResponseBodyDataElements) GetResults() []*RecognizePublicFaceResponseBodyDataElementsResults {
	return s.Results
}

func (s *RecognizePublicFaceResponseBodyDataElements) GetTaskId() *string {
	return s.TaskId
}

func (s *RecognizePublicFaceResponseBodyDataElements) SetImageURL(v string) *RecognizePublicFaceResponseBodyDataElements {
	s.ImageURL = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElements) SetResults(v []*RecognizePublicFaceResponseBodyDataElementsResults) *RecognizePublicFaceResponseBodyDataElements {
	s.Results = v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElements) SetTaskId(v string) *RecognizePublicFaceResponseBodyDataElements {
	s.TaskId = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElements) Validate() error {
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizePublicFaceResponseBodyDataElementsResults struct {
	// example:
	//
	// sface
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// example:
	//
	// 98.35
	Rate       *float32                                                        `json:"Rate,omitempty" xml:"Rate,omitempty"`
	SubResults []*RecognizePublicFaceResponseBodyDataElementsResultsSubResults `json:"SubResults,omitempty" xml:"SubResults,omitempty" type:"Repeated"`
	// example:
	//
	// review
	Suggestion *string `json:"Suggestion,omitempty" xml:"Suggestion,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElementsResults) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElementsResults) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) GetLabel() *string {
	return s.Label
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) GetRate() *float32 {
	return s.Rate
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) GetSubResults() []*RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	return s.SubResults
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) GetSuggestion() *string {
	return s.Suggestion
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetLabel(v string) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.Label = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetRate(v float32) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.Rate = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetSubResults(v []*RecognizePublicFaceResponseBodyDataElementsResultsSubResults) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.SubResults = v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) SetSuggestion(v string) *RecognizePublicFaceResponseBodyDataElementsResults {
	s.Suggestion = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResults) Validate() error {
	if s.SubResults != nil {
		for _, item := range s.SubResults {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizePublicFaceResponseBodyDataElementsResultsSubResults struct {
	Faces []*RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
	// example:
	//
	// 153
	H *float32 `json:"H,omitempty" xml:"H,omitempty"`
	// example:
	//
	// 132
	W *float32 `json:"W,omitempty" xml:"W,omitempty"`
	// example:
	//
	// 182
	X *float32 `json:"X,omitempty" xml:"X,omitempty"`
	// example:
	//
	// 153
	Y *float32 `json:"Y,omitempty" xml:"Y,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResults) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResults) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) GetFaces() []*RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	return s.Faces
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) GetH() *float32 {
	return s.H
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) GetW() *float32 {
	return s.W
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) GetX() *float32 {
	return s.X
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) GetY() *float32 {
	return s.Y
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetFaces(v []*RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.Faces = v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetH(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.H = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetW(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.W = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetX(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.X = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) SetY(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResults {
	s.Y = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResults) Validate() error {
	if s.Faces != nil {
		for _, item := range s.Faces {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces struct {
	// example:
	//
	// AliFace_0006272
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 98.35
	Rate *float32 `json:"Rate,omitempty" xml:"Rate,omitempty"`
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) String() string {
	return dara.Prettify(s)
}

func (s RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) GoString() string {
	return s.String()
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) GetId() *string {
	return s.Id
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) GetName() *string {
	return s.Name
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) GetRate() *float32 {
	return s.Rate
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) SetId(v string) *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	s.Id = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) SetName(v string) *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	s.Name = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) SetRate(v float32) *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces {
	s.Rate = &v
	return s
}

func (s *RecognizePublicFaceResponseBodyDataElementsResultsSubResultsFaces) Validate() error {
	return dara.Validate(s)
}
