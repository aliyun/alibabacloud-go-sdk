// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageGenerationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetImageGenerationJobRequest
	GetJobId() *string
}

type GetImageGenerationJobRequest struct {
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetImageGenerationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageGenerationJobRequest) GoString() string {
	return s.String()
}

func (s *GetImageGenerationJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetImageGenerationJobRequest) SetJobId(v string) *GetImageGenerationJobRequest {
	s.JobId = &v
	return s
}

func (s *GetImageGenerationJobRequest) Validate() error {
	return dara.Validate(s)
}
