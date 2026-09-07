// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoTranslationJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetVideoTranslationJobRequest
	GetJobId() *string
}

type GetVideoTranslationJobRequest struct {
	// The video translation job ID returned by SubmitVideoTranslationJob.
	//
	// This parameter is required.
	//
	// example:
	//
	// vtj_0123456789abcdef0123456789abcdef
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetVideoTranslationJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoTranslationJobRequest) GoString() string {
	return s.String()
}

func (s *GetVideoTranslationJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetVideoTranslationJobRequest) SetJobId(v string) *GetVideoTranslationJobRequest {
	s.JobId = &v
	return s
}

func (s *GetVideoTranslationJobRequest) Validate() error {
	return dara.Validate(s)
}
