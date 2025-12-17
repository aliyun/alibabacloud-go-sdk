// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSanityCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetSanityCheckTaskRequest
	GetVerbose() *bool
}

type GetSanityCheckTaskRequest struct {
	// example:
	//
	// true
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetSanityCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSanityCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *GetSanityCheckTaskRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetSanityCheckTaskRequest) SetVerbose(v bool) *GetSanityCheckTaskRequest {
	s.Verbose = &v
	return s
}

func (s *GetSanityCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
