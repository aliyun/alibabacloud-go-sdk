// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaConvertJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetMediaConvertJobRequest
	GetJobId() *string
}

type GetMediaConvertJobRequest struct {
	// The ID of the transcoding task.
	//
	// example:
	//
	// ****d80e4e4044975745c14b****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetMediaConvertJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMediaConvertJobRequest) GoString() string {
	return s.String()
}

func (s *GetMediaConvertJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetMediaConvertJobRequest) SetJobId(v string) *GetMediaConvertJobRequest {
	s.JobId = &v
	return s
}

func (s *GetMediaConvertJobRequest) Validate() error {
	return dara.Validate(s)
}
