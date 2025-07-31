// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAnswerLibRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLibName(v string) *CreateAnswerLibRequest
	GetLibName() *string
	SetRegionId(v string) *CreateAnswerLibRequest
	GetRegionId() *string
	SetSampleBucket(v string) *CreateAnswerLibRequest
	GetSampleBucket() *string
	SetSampleObject(v string) *CreateAnswerLibRequest
	GetSampleObject() *string
	SetSamples(v string) *CreateAnswerLibRequest
	GetSamples() *string
}

type CreateAnswerLibRequest struct {
	LibName *string `json:"LibName,omitempty" xml:"LibName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// oss-cip-shanghai
	SampleBucket *string `json:"SampleBucket,omitempty" xml:"SampleBucket,omitempty"`
	// example:
	//
	// data/xxx.xlsx
	SampleObject *string `json:"SampleObject,omitempty" xml:"SampleObject,omitempty"`
	Samples      *string `json:"Samples,omitempty" xml:"Samples,omitempty"`
}

func (s CreateAnswerLibRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAnswerLibRequest) GoString() string {
	return s.String()
}

func (s *CreateAnswerLibRequest) GetLibName() *string {
	return s.LibName
}

func (s *CreateAnswerLibRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateAnswerLibRequest) GetSampleBucket() *string {
	return s.SampleBucket
}

func (s *CreateAnswerLibRequest) GetSampleObject() *string {
	return s.SampleObject
}

func (s *CreateAnswerLibRequest) GetSamples() *string {
	return s.Samples
}

func (s *CreateAnswerLibRequest) SetLibName(v string) *CreateAnswerLibRequest {
	s.LibName = &v
	return s
}

func (s *CreateAnswerLibRequest) SetRegionId(v string) *CreateAnswerLibRequest {
	s.RegionId = &v
	return s
}

func (s *CreateAnswerLibRequest) SetSampleBucket(v string) *CreateAnswerLibRequest {
	s.SampleBucket = &v
	return s
}

func (s *CreateAnswerLibRequest) SetSampleObject(v string) *CreateAnswerLibRequest {
	s.SampleObject = &v
	return s
}

func (s *CreateAnswerLibRequest) SetSamples(v string) *CreateAnswerLibRequest {
	s.Samples = &v
	return s
}

func (s *CreateAnswerLibRequest) Validate() error {
	return dara.Validate(s)
}
