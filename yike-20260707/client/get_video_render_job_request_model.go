// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoRenderJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetVideoRenderJobRequest
	GetJobId() *string
}

type GetVideoRenderJobRequest struct {
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetVideoRenderJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoRenderJobRequest) GoString() string {
	return s.String()
}

func (s *GetVideoRenderJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoRenderJobRequest) SetJobId(v string) *GetVideoRenderJobRequest {
	s.JobId = &v
	return s
}

func (s *GetVideoRenderJobRequest) Validate() error {
	return dara.Validate(s)
}
