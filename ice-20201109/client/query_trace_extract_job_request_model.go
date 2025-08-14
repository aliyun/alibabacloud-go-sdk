// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryTraceExtractJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *QueryTraceExtractJobRequest
	GetJobId() *string
}

type QueryTraceExtractJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 31fa3c9ca8134fb4b0b0f7878301****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryTraceExtractJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryTraceExtractJobRequest) GoString() string {
	return s.String()
}

func (s *QueryTraceExtractJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryTraceExtractJobRequest) SetJobId(v string) *QueryTraceExtractJobRequest {
	s.JobId = &v
	return s
}

func (s *QueryTraceExtractJobRequest) Validate() error {
	return dara.Validate(s)
}
