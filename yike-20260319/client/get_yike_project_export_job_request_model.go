// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetYikeProjectExportJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetYikeProjectExportJobRequest
	GetJobId() *string
}

type GetYikeProjectExportJobRequest struct {
	// The task ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetYikeProjectExportJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetYikeProjectExportJobRequest) GoString() string {
	return s.String()
}

func (s *GetYikeProjectExportJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetYikeProjectExportJobRequest) SetJobId(v string) *GetYikeProjectExportJobRequest {
	s.JobId = &v
	return s
}

func (s *GetYikeProjectExportJobRequest) Validate() error {
	return dara.Validate(s)
}
