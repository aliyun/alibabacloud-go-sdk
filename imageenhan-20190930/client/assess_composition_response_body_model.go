// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssessCompositionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AssessCompositionResponseBodyData) *AssessCompositionResponseBody
	GetData() *AssessCompositionResponseBodyData
	SetRequestId(v string) *AssessCompositionResponseBody
	GetRequestId() *string
}

type AssessCompositionResponseBody struct {
	Data *AssessCompositionResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// CCAD9435-AEDB-49E4-BCCC-99B65ECC6693
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssessCompositionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssessCompositionResponseBody) GoString() string {
	return s.String()
}

func (s *AssessCompositionResponseBody) GetData() *AssessCompositionResponseBodyData {
	return s.Data
}

func (s *AssessCompositionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssessCompositionResponseBody) SetData(v *AssessCompositionResponseBodyData) *AssessCompositionResponseBody {
	s.Data = v
	return s
}

func (s *AssessCompositionResponseBody) SetRequestId(v string) *AssessCompositionResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssessCompositionResponseBody) Validate() error {
	return dara.Validate(s)
}

type AssessCompositionResponseBodyData struct {
	// example:
	//
	// 4.2551436
	Score *float32 `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s AssessCompositionResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AssessCompositionResponseBodyData) GoString() string {
	return s.String()
}

func (s *AssessCompositionResponseBodyData) GetScore() *float32 {
	return s.Score
}

func (s *AssessCompositionResponseBodyData) SetScore(v float32) *AssessCompositionResponseBodyData {
	s.Score = &v
	return s
}

func (s *AssessCompositionResponseBodyData) Validate() error {
	return dara.Validate(s)
}
