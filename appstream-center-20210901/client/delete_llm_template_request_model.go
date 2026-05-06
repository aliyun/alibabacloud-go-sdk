// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLlmTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLlmTemplateId(v string) *DeleteLlmTemplateRequest
	GetLlmTemplateId() *string
}

type DeleteLlmTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// llmt-xxxx
	LlmTemplateId *string `json:"LlmTemplateId,omitempty" xml:"LlmTemplateId,omitempty"`
}

func (s DeleteLlmTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLlmTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteLlmTemplateRequest) GetLlmTemplateId() *string {
	return s.LlmTemplateId
}

func (s *DeleteLlmTemplateRequest) SetLlmTemplateId(v string) *DeleteLlmTemplateRequest {
	s.LlmTemplateId = &v
	return s
}

func (s *DeleteLlmTemplateRequest) Validate() error {
	return dara.Validate(s)
}
