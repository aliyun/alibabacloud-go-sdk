// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBodyPostureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BodyPostureResponseBodyData) *BodyPostureResponseBody
	GetData() *BodyPostureResponseBodyData
	SetRequestId(v string) *BodyPostureResponseBody
	GetRequestId() *string
}

type BodyPostureResponseBody struct {
	Data *BodyPostureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// B854094B-9231-4A54-89AB-C377CB0D237D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s BodyPostureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureResponseBody) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBody) GetData() *BodyPostureResponseBodyData {
	return s.Data
}

func (s *BodyPostureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BodyPostureResponseBody) SetData(v *BodyPostureResponseBodyData) *BodyPostureResponseBody {
	s.Data = v
	return s
}

func (s *BodyPostureResponseBody) SetRequestId(v string) *BodyPostureResponseBody {
	s.RequestId = &v
	return s
}

func (s *BodyPostureResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BodyPostureResponseBodyData struct {
	MetaObject *BodyPostureResponseBodyDataMetaObject `json:"MetaObject,omitempty" xml:"MetaObject,omitempty" type:"Struct"`
	Outputs    []*BodyPostureResponseBodyDataOutputs  `json:"Outputs,omitempty" xml:"Outputs,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureResponseBodyData) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyData) GetMetaObject() *BodyPostureResponseBodyDataMetaObject {
	return s.MetaObject
}

func (s *BodyPostureResponseBodyData) GetOutputs() []*BodyPostureResponseBodyDataOutputs {
	return s.Outputs
}

func (s *BodyPostureResponseBodyData) SetMetaObject(v *BodyPostureResponseBodyDataMetaObject) *BodyPostureResponseBodyData {
	s.MetaObject = v
	return s
}

func (s *BodyPostureResponseBodyData) SetOutputs(v []*BodyPostureResponseBodyDataOutputs) *BodyPostureResponseBodyData {
	s.Outputs = v
	return s
}

func (s *BodyPostureResponseBodyData) Validate() error {
	if s.MetaObject != nil {
		if err := s.MetaObject.Validate(); err != nil {
			return err
		}
	}
	if s.Outputs != nil {
		for _, item := range s.Outputs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BodyPostureResponseBodyDataMetaObject struct {
	// example:
	//
	// 500
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 500
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s BodyPostureResponseBodyDataMetaObject) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureResponseBodyDataMetaObject) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataMetaObject) GetHeight() *int32 {
	return s.Height
}

func (s *BodyPostureResponseBodyDataMetaObject) GetWidth() *int32 {
	return s.Width
}

func (s *BodyPostureResponseBodyDataMetaObject) SetHeight(v int32) *BodyPostureResponseBodyDataMetaObject {
	s.Height = &v
	return s
}

func (s *BodyPostureResponseBodyDataMetaObject) SetWidth(v int32) *BodyPostureResponseBodyDataMetaObject {
	s.Width = &v
	return s
}

func (s *BodyPostureResponseBodyDataMetaObject) Validate() error {
	return dara.Validate(s)
}

type BodyPostureResponseBodyDataOutputs struct {
	// example:
	//
	// 1
	HumanCount *int32                                       `json:"HumanCount,omitempty" xml:"HumanCount,omitempty"`
	Results    []*BodyPostureResponseBodyDataOutputsResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputs) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputs) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputs) GetHumanCount() *int32 {
	return s.HumanCount
}

func (s *BodyPostureResponseBodyDataOutputs) GetResults() []*BodyPostureResponseBodyDataOutputsResults {
	return s.Results
}

func (s *BodyPostureResponseBodyDataOutputs) SetHumanCount(v int32) *BodyPostureResponseBodyDataOutputs {
	s.HumanCount = &v
	return s
}

func (s *BodyPostureResponseBodyDataOutputs) SetResults(v []*BodyPostureResponseBodyDataOutputsResults) *BodyPostureResponseBodyDataOutputs {
	s.Results = v
	return s
}

func (s *BodyPostureResponseBodyDataOutputs) Validate() error {
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

type BodyPostureResponseBodyDataOutputsResults struct {
	Bodies []*BodyPostureResponseBodyDataOutputsResultsBodies `json:"Bodies,omitempty" xml:"Bodies,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputsResults) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputsResults) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputsResults) GetBodies() []*BodyPostureResponseBodyDataOutputsResultsBodies {
	return s.Bodies
}

func (s *BodyPostureResponseBodyDataOutputsResults) SetBodies(v []*BodyPostureResponseBodyDataOutputsResultsBodies) *BodyPostureResponseBodyDataOutputsResults {
	s.Bodies = v
	return s
}

func (s *BodyPostureResponseBodyDataOutputsResults) Validate() error {
	if s.Bodies != nil {
		for _, item := range s.Bodies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BodyPostureResponseBodyDataOutputsResultsBodies struct {
	// example:
	//
	// 0.91309475898742676
	Confident *float32 `json:"Confident,omitempty" xml:"Confident,omitempty"`
	// example:
	//
	// nose
	Label     *string                                                     `json:"Label,omitempty" xml:"Label,omitempty"`
	Positions []*BodyPostureResponseBodyDataOutputsResultsBodiesPositions `json:"Positions,omitempty" xml:"Positions,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputsResultsBodies) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputsResultsBodies) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) GetConfident() *float32 {
	return s.Confident
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) GetLabel() *string {
	return s.Label
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) GetPositions() []*BodyPostureResponseBodyDataOutputsResultsBodiesPositions {
	return s.Positions
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) SetConfident(v float32) *BodyPostureResponseBodyDataOutputsResultsBodies {
	s.Confident = &v
	return s
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) SetLabel(v string) *BodyPostureResponseBodyDataOutputsResultsBodies {
	s.Label = &v
	return s
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) SetPositions(v []*BodyPostureResponseBodyDataOutputsResultsBodiesPositions) *BodyPostureResponseBodyDataOutputsResultsBodies {
	s.Positions = v
	return s
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodies) Validate() error {
	if s.Positions != nil {
		for _, item := range s.Positions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BodyPostureResponseBodyDataOutputsResultsBodiesPositions struct {
	Points []*float32 `json:"Points,omitempty" xml:"Points,omitempty" type:"Repeated"`
}

func (s BodyPostureResponseBodyDataOutputsResultsBodiesPositions) String() string {
	return dara.Prettify(s)
}

func (s BodyPostureResponseBodyDataOutputsResultsBodiesPositions) GoString() string {
	return s.String()
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodiesPositions) GetPoints() []*float32 {
	return s.Points
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodiesPositions) SetPoints(v []*float32) *BodyPostureResponseBodyDataOutputsResultsBodiesPositions {
	s.Points = v
	return s
}

func (s *BodyPostureResponseBodyDataOutputsResultsBodiesPositions) Validate() error {
	return dara.Validate(s)
}
