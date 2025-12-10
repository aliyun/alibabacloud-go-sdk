// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetJobRequest
	GetVerbose() *bool
}

type GetJobRequest struct {
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetJobRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobRequest) GoString() string {
	return s.String()
}

func (s *GetJobRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetJobRequest) SetVerbose(v bool) *GetJobRequest {
	s.Verbose = &v
	return s
}

func (s *GetJobRequest) Validate() error {
	return dara.Validate(s)
}
