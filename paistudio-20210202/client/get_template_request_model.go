// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVerbose(v bool) *GetTemplateRequest
	GetVerbose() *bool
}

type GetTemplateRequest struct {
	Verbose *bool `json:"Verbose,omitempty" xml:"Verbose,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) GetVerbose() *bool {
	return s.Verbose
}

func (s *GetTemplateRequest) SetVerbose(v bool) *GetTemplateRequest {
	s.Verbose = &v
	return s
}

func (s *GetTemplateRequest) Validate() error {
	return dara.Validate(s)
}
