// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRunRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetRunRequest
	GetVerbose() *bool
}

type GetRunRequest struct {
	// Specifies whether to obtain the Metrics, Params, and Labels information. Default value: false.
	//
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetRunRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRunRequest) GoString() string {
	return s.String()
}

func (s *GetRunRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetRunRequest) SetVerbose(v bool) *GetRunRequest {
	s.Verbose = &v
	return s
}

func (s *GetRunRequest) Validate() error {
	return dara.Validate(s)
}
