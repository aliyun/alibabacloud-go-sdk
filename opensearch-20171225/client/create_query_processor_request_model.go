// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQueryProcessorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v interface{}) *CreateQueryProcessorRequest
	GetBody() interface{}
	SetDryRun(v bool) *CreateQueryProcessorRequest
	GetDryRun() *bool
}

type CreateQueryProcessorRequest struct {
	// The request body.
	//
	// example:
	//
	// {}
	Body interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s CreateQueryProcessorRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateQueryProcessorRequest) GoString() string {
	return s.String()
}

func (s *CreateQueryProcessorRequest) GetBody() interface{} {
	return s.Body
}

func (s *CreateQueryProcessorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateQueryProcessorRequest) SetBody(v interface{}) *CreateQueryProcessorRequest {
	s.Body = v
	return s
}

func (s *CreateQueryProcessorRequest) SetDryRun(v bool) *CreateQueryProcessorRequest {
	s.DryRun = &v
	return s
}

func (s *CreateQueryProcessorRequest) Validate() error {
	return dara.Validate(s)
}
