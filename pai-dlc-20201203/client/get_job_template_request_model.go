// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVersion(v string) *GetJobTemplateRequest
	GetVersion() *string
}

type GetJobTemplateRequest struct {
	// The version to retrieve. If omitted, the default version is returned. Specify `all` to retrieve all versions.
	//
	// example:
	//
	// all
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s GetJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetJobTemplateRequest) GetVersion() *string {
	return s.Version
}

func (s *GetJobTemplateRequest) SetVersion(v string) *GetJobTemplateRequest {
	s.Version = &v
	return s
}

func (s *GetJobTemplateRequest) Validate() error {
	return dara.Validate(s)
}
