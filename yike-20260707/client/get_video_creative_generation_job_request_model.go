// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoCreativeGenerationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetVideoCreativeGenerationJobRequest
	GetJobId() *string
}

type GetVideoCreativeGenerationJobRequest struct {
	// The job ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2ff6142097984dde9591925df43291a9
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetVideoCreativeGenerationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoCreativeGenerationJobRequest) GoString() string {
	return s.String()
}

func (s *GetVideoCreativeGenerationJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoCreativeGenerationJobRequest) SetJobId(v string) *GetVideoCreativeGenerationJobRequest {
	s.JobId = &v
	return s
}

func (s *GetVideoCreativeGenerationJobRequest) Validate() error {
	return dara.Validate(s)
}
