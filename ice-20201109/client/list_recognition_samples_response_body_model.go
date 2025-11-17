// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecognitionSamplesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *ListRecognitionSamplesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListRecognitionSamplesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListRecognitionSamplesResponseBody
	GetRequestId() *string
	SetSamples(v *ListRecognitionSamplesResponseBodySamples) *ListRecognitionSamplesResponseBody
	GetSamples() *ListRecognitionSamplesResponseBodySamples
	SetTotalCount(v int64) *ListRecognitionSamplesResponseBody
	GetTotalCount() *int64
}

type ListRecognitionSamplesResponseBody struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// xxx-xxxx-xxxxx-xxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The samples.
	Samples *ListRecognitionSamplesResponseBodySamples `json:"Samples,omitempty" xml:"Samples,omitempty" type:"Struct"`
	// The total number of samples.
	//
	// example:
	//
	// 5
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListRecognitionSamplesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionSamplesResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecognitionSamplesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListRecognitionSamplesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListRecognitionSamplesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRecognitionSamplesResponseBody) GetSamples() *ListRecognitionSamplesResponseBodySamples {
	return s.Samples
}

func (s *ListRecognitionSamplesResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListRecognitionSamplesResponseBody) SetPageNumber(v int32) *ListRecognitionSamplesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListRecognitionSamplesResponseBody) SetPageSize(v int32) *ListRecognitionSamplesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListRecognitionSamplesResponseBody) SetRequestId(v string) *ListRecognitionSamplesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRecognitionSamplesResponseBody) SetSamples(v *ListRecognitionSamplesResponseBodySamples) *ListRecognitionSamplesResponseBody {
	s.Samples = v
	return s
}

func (s *ListRecognitionSamplesResponseBody) SetTotalCount(v int64) *ListRecognitionSamplesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListRecognitionSamplesResponseBody) Validate() error {
	if s.Samples != nil {
		if err := s.Samples.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListRecognitionSamplesResponseBodySamples struct {
	Sample []*ListRecognitionSamplesResponseBodySamplesSample `json:"Sample,omitempty" xml:"Sample,omitempty" type:"Repeated"`
}

func (s ListRecognitionSamplesResponseBodySamples) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionSamplesResponseBodySamples) GoString() string {
	return s.String()
}

func (s *ListRecognitionSamplesResponseBodySamples) GetSample() []*ListRecognitionSamplesResponseBodySamplesSample {
	return s.Sample
}

func (s *ListRecognitionSamplesResponseBodySamples) SetSample(v []*ListRecognitionSamplesResponseBodySamplesSample) *ListRecognitionSamplesResponseBodySamples {
	s.Sample = v
	return s
}

func (s *ListRecognitionSamplesResponseBodySamples) Validate() error {
	if s.Sample != nil {
		for _, item := range s.Sample {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRecognitionSamplesResponseBodySamplesSample struct {
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The URL of the image sample.
	//
	// example:
	//
	// https://example.com/sample.png
	ImageUrl *string `json:"ImageUrl,omitempty" xml:"ImageUrl,omitempty"`
	LibId    *string `json:"LibId,omitempty" xml:"LibId,omitempty"`
	// The sample ID.
	//
	// example:
	//
	// xxxxxxxxxxxxx
	SampleId *string `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
}

func (s ListRecognitionSamplesResponseBodySamplesSample) String() string {
	return dara.Prettify(s)
}

func (s ListRecognitionSamplesResponseBodySamplesSample) GoString() string {
	return s.String()
}

func (s *ListRecognitionSamplesResponseBodySamplesSample) GetEntityId() *string {
	return s.EntityId
}

func (s *ListRecognitionSamplesResponseBodySamplesSample) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *ListRecognitionSamplesResponseBodySamplesSample) GetLibId() *string {
	return s.LibId
}

func (s *ListRecognitionSamplesResponseBodySamplesSample) GetSampleId() *string {
	return s.SampleId
}

func (s *ListRecognitionSamplesResponseBodySamplesSample) SetEntityId(v string) *ListRecognitionSamplesResponseBodySamplesSample {
	s.EntityId = &v
	return s
}

func (s *ListRecognitionSamplesResponseBodySamplesSample) SetImageUrl(v string) *ListRecognitionSamplesResponseBodySamplesSample {
	s.ImageUrl = &v
	return s
}

func (s *ListRecognitionSamplesResponseBodySamplesSample) SetLibId(v string) *ListRecognitionSamplesResponseBodySamplesSample {
	s.LibId = &v
	return s
}

func (s *ListRecognitionSamplesResponseBodySamplesSample) SetSampleId(v string) *ListRecognitionSamplesResponseBodySamplesSample {
	s.SampleId = &v
	return s
}

func (s *ListRecognitionSamplesResponseBodySamplesSample) Validate() error {
	return dara.Validate(s)
}
