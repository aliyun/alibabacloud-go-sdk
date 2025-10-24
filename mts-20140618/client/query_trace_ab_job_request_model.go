// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceAbJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *QueryTraceAbJobRequest
	GetJobId() *string
	SetMediaId(v string) *QueryTraceAbJobRequest
	GetMediaId() *string
}

type QueryTraceAbJobRequest struct {
	// example:
	//
	// 31fa3c9ca8134fb4b0b0f7878301****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// 3e6149d5a8c944c09b1a8d2dc3e4****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
}

func (s QueryTraceAbJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceAbJobRequest) GoString() string {
	return s.String()
}

func (s *QueryTraceAbJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryTraceAbJobRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryTraceAbJobRequest) SetJobId(v string) *QueryTraceAbJobRequest {
	s.JobId = &v
	return s
}

func (s *QueryTraceAbJobRequest) SetMediaId(v string) *QueryTraceAbJobRequest {
	s.MediaId = &v
	return s
}

func (s *QueryTraceAbJobRequest) Validate() error {
	return dara.Validate(s)
}
