// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiPathFieldTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPath(v string) *GetApiPathFieldTemplateRequest
	GetPath() *string
}

type GetApiPathFieldTemplateRequest struct {
	// The gateway type.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s GetApiPathFieldTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiPathFieldTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetApiPathFieldTemplateRequest) GetPath() *string {
	return s.Path
}

func (s *GetApiPathFieldTemplateRequest) SetPath(v string) *GetApiPathFieldTemplateRequest {
	s.Path = &v
	return s
}

func (s *GetApiPathFieldTemplateRequest) Validate() error {
	return dara.Validate(s)
}
