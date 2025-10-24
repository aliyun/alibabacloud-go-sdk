// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCopyrightExtractJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *QueryCopyrightExtractJobRequest
	GetJobId() *string
}

type QueryCopyrightExtractJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 2288c6ca184c0e47098a5b665e2a12****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s QueryCopyrightExtractJobRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCopyrightExtractJobRequest) GoString() string {
	return s.String()
}

func (s *QueryCopyrightExtractJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *QueryCopyrightExtractJobRequest) SetJobId(v string) *QueryCopyrightExtractJobRequest {
	s.JobId = &v
	return s
}

func (s *QueryCopyrightExtractJobRequest) Validate() error {
	return dara.Validate(s)
}
