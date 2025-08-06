// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAICaptionExtractionJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *SubmitAICaptionExtractionJobResponseBody
	GetJobId() *string
	SetRequestId(v string) *SubmitAICaptionExtractionJobResponseBody
	GetRequestId() *string
}

type SubmitAICaptionExtractionJobResponseBody struct {
	JobId     *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitAICaptionExtractionJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICaptionExtractionJobResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitAICaptionExtractionJobResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *SubmitAICaptionExtractionJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitAICaptionExtractionJobResponseBody) SetJobId(v string) *SubmitAICaptionExtractionJobResponseBody {
	s.JobId = &v
	return s
}

func (s *SubmitAICaptionExtractionJobResponseBody) SetRequestId(v string) *SubmitAICaptionExtractionJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitAICaptionExtractionJobResponseBody) Validate() error {
	return dara.Validate(s)
}
