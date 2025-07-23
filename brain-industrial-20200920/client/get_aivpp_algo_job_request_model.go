// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAivppAlgoJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *GetAivppAlgoJobRequest
	GetJobId() *string
}

type GetAivppAlgoJobRequest struct {
	// example:
	//
	// 5854bfa6-f002-43c2-8e1d-e9b2c28f9384
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetAivppAlgoJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAivppAlgoJobRequest) GoString() string {
	return s.String()
}

func (s *GetAivppAlgoJobRequest) GetJobId() *string {
	return s.JobId
}

func (s *GetAivppAlgoJobRequest) SetJobId(v string) *GetAivppAlgoJobRequest {
	s.JobId = &v
	return s
}

func (s *GetAivppAlgoJobRequest) Validate() error {
	return dara.Validate(s)
}
