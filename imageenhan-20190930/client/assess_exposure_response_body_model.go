// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssessExposureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AssessExposureResponseBodyData) *AssessExposureResponseBody
	GetData() *AssessExposureResponseBodyData
	SetRequestId(v string) *AssessExposureResponseBody
	GetRequestId() *string
}

type AssessExposureResponseBody struct {
	Data *AssessExposureResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 4EF3C65B-C3CC-425B-AFB3-2FE6B98C578B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssessExposureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssessExposureResponseBody) GoString() string {
	return s.String()
}

func (s *AssessExposureResponseBody) GetData() *AssessExposureResponseBodyData {
	return s.Data
}

func (s *AssessExposureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssessExposureResponseBody) SetData(v *AssessExposureResponseBodyData) *AssessExposureResponseBody {
	s.Data = v
	return s
}

func (s *AssessExposureResponseBody) SetRequestId(v string) *AssessExposureResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssessExposureResponseBody) Validate() error {
	return dara.Validate(s)
}

type AssessExposureResponseBodyData struct {
	// example:
	//
	// 0.1
	Exposure *float32 `json:"Exposure,omitempty" xml:"Exposure,omitempty"`
}

func (s AssessExposureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AssessExposureResponseBodyData) GoString() string {
	return s.String()
}

func (s *AssessExposureResponseBodyData) GetExposure() *float32 {
	return s.Exposure
}

func (s *AssessExposureResponseBodyData) SetExposure(v float32) *AssessExposureResponseBodyData {
	s.Exposure = &v
	return s
}

func (s *AssessExposureResponseBodyData) Validate() error {
	return dara.Validate(s)
}
