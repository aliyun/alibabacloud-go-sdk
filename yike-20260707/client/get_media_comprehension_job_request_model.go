// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaComprehensionJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetMediaComprehensionJobRequest
	GetJobId() *string
}

type GetMediaComprehensionJobRequest struct {
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetMediaComprehensionJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaComprehensionJobRequest) GoString() string {
	return s.String()
}

func (s *GetMediaComprehensionJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetMediaComprehensionJobRequest) SetJobId(v string) *GetMediaComprehensionJobRequest {
	s.JobId = &v
	return s
}

func (s *GetMediaComprehensionJobRequest) Validate() error {
	return dara.Validate(s)
}
