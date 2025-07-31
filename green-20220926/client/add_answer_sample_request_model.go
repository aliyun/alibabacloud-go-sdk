// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAnswerSampleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibId(v string) *AddAnswerSampleRequest
	GetLibId() *string
	SetRegionId(v string) *AddAnswerSampleRequest
	GetRegionId() *string
	SetSampleObject(v string) *AddAnswerSampleRequest
	GetSampleObject() *string
	SetSamples(v string) *AddAnswerSampleRequest
	GetSamples() *string
}

type AddAnswerSampleRequest struct {
	// example:
	//
	// alxxxx
	LibId *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// data/xxx.xlsx
	SampleObject *string `json:"SampleObject,omitempty" xml:"SampleObject,omitempty"`
	Samples      *string `json:"Samples,omitempty" xml:"Samples,omitempty"`
}

func (s AddAnswerSampleRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAnswerSampleRequest) GoString() string {
	return s.String()
}

func (s *AddAnswerSampleRequest) GetLibId() *string {
	return s.LibId
}

func (s *AddAnswerSampleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddAnswerSampleRequest) GetSampleObject() *string {
	return s.SampleObject
}

func (s *AddAnswerSampleRequest) GetSamples() *string {
	return s.Samples
}

func (s *AddAnswerSampleRequest) SetLibId(v string) *AddAnswerSampleRequest {
	s.LibId = &v
	return s
}

func (s *AddAnswerSampleRequest) SetRegionId(v string) *AddAnswerSampleRequest {
	s.RegionId = &v
	return s
}

func (s *AddAnswerSampleRequest) SetSampleObject(v string) *AddAnswerSampleRequest {
	s.SampleObject = &v
	return s
}

func (s *AddAnswerSampleRequest) SetSamples(v string) *AddAnswerSampleRequest {
	s.Samples = &v
	return s
}

func (s *AddAnswerSampleRequest) Validate() error {
	return dara.Validate(s)
}
