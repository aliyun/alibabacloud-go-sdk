// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoGenerationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetVideoGenerationJobRequest
	GetClientToken() *string
	SetJobId(v string) *GetVideoGenerationJobRequest
	GetJobId() *string
}

type GetVideoGenerationJobRequest struct {
	// example:
	//
	// xxxx
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 68ca759e798b40b4903b255*******
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetVideoGenerationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoGenerationJobRequest) GoString() string {
	return s.String()
}

func (s *GetVideoGenerationJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetVideoGenerationJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoGenerationJobRequest) SetClientToken(v string) *GetVideoGenerationJobRequest {
	s.ClientToken = &v
	return s
}

func (s *GetVideoGenerationJobRequest) SetJobId(v string) *GetVideoGenerationJobRequest {
	s.JobId = &v
	return s
}

func (s *GetVideoGenerationJobRequest) Validate() error {
	return dara.Validate(s)
}
