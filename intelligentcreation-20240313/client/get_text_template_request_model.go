// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIndustry(v string) *GetTextTemplateRequest
	GetIndustry() *string
}

type GetTextTemplateRequest struct {
	// example:
	//
	// Car
	Industry *string `json:"industry,omitempty" xml:"industry,omitempty"`
}

func (s GetTextTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTextTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTextTemplateRequest) GetIndustry() *string {
	return s.Industry
}

func (s *GetTextTemplateRequest) SetIndustry(v string) *GetTextTemplateRequest {
	s.Industry = &v
	return s
}

func (s *GetTextTemplateRequest) Validate() error {
	return dara.Validate(s)
}
