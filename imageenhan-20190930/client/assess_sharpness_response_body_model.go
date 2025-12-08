// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssessSharpnessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AssessSharpnessResponseBodyData) *AssessSharpnessResponseBody
	GetData() *AssessSharpnessResponseBodyData
	SetRequestId(v string) *AssessSharpnessResponseBody
	GetRequestId() *string
}

type AssessSharpnessResponseBody struct {
	Data *AssessSharpnessResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// C0B594A1-383E-4F97-A740-0D51CF8E37D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssessSharpnessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssessSharpnessResponseBody) GoString() string {
	return s.String()
}

func (s *AssessSharpnessResponseBody) GetData() *AssessSharpnessResponseBodyData {
	return s.Data
}

func (s *AssessSharpnessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssessSharpnessResponseBody) SetData(v *AssessSharpnessResponseBodyData) *AssessSharpnessResponseBody {
	s.Data = v
	return s
}

func (s *AssessSharpnessResponseBody) SetRequestId(v string) *AssessSharpnessResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssessSharpnessResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AssessSharpnessResponseBodyData struct {
	// example:
	//
	// 0.1
	Sharpness *float32 `json:"Sharpness,omitempty" xml:"Sharpness,omitempty"`
}

func (s AssessSharpnessResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AssessSharpnessResponseBodyData) GoString() string {
	return s.String()
}

func (s *AssessSharpnessResponseBodyData) GetSharpness() *float32 {
	return s.Sharpness
}

func (s *AssessSharpnessResponseBodyData) SetSharpness(v float32) *AssessSharpnessResponseBodyData {
	s.Sharpness = &v
	return s
}

func (s *AssessSharpnessResponseBodyData) Validate() error {
	return dara.Validate(s)
}
