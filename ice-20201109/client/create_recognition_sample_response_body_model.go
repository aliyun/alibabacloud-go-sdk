// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecognitionSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRecognitionSampleResponseBody
	GetRequestId() *string
	SetSampleId(v string) *CreateRecognitionSampleResponseBody
	GetSampleId() *string
}

type CreateRecognitionSampleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the sample.
	//
	// example:
	//
	// **************4d2ba728e2f**************
	SampleId *string `json:"SampleId,omitempty" xml:"SampleId,omitempty"`
}

func (s CreateRecognitionSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecognitionSampleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecognitionSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecognitionSampleResponseBody) GetSampleId() *string {
	return s.SampleId
}

func (s *CreateRecognitionSampleResponseBody) SetRequestId(v string) *CreateRecognitionSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecognitionSampleResponseBody) SetSampleId(v string) *CreateRecognitionSampleResponseBody {
	s.SampleId = &v
	return s
}

func (s *CreateRecognitionSampleResponseBody) Validate() error {
	return dara.Validate(s)
}
