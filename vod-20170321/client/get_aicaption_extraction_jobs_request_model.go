// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICaptionExtractionJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobIds(v string) *GetAICaptionExtractionJobsRequest
	GetJobIds() *string
}

type GetAICaptionExtractionJobsRequest struct {
	// This parameter is required.
	JobIds *string `json:"JobIds,omitempty" xml:"JobIds,omitempty"`
}

func (s GetAICaptionExtractionJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAICaptionExtractionJobsRequest) GoString() string {
	return s.String()
}

func (s *GetAICaptionExtractionJobsRequest) GetJobIds() *string {
	return s.JobIds
}

func (s *GetAICaptionExtractionJobsRequest) SetJobIds(v string) *GetAICaptionExtractionJobsRequest {
	s.JobIds = &v
	return s
}

func (s *GetAICaptionExtractionJobsRequest) Validate() error {
	return dara.Validate(s)
}
