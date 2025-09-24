// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetExperimentRequest
	GetVerbose() *bool
}

type GetExperimentRequest struct {
	// Specifies whether to obtain the latest run information associated with the experiment
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetExperimentRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentRequest) GoString() string {
	return s.String()
}

func (s *GetExperimentRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetExperimentRequest) SetVerbose(v bool) *GetExperimentRequest {
	s.Verbose = &v
	return s
}

func (s *GetExperimentRequest) Validate() error {
	return dara.Validate(s)
}
