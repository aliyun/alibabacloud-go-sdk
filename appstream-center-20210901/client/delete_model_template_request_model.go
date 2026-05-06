// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetModelTemplateId(v string) *DeleteModelTemplateRequest
	GetModelTemplateId() *string
}

type DeleteModelTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mt-xxxx
	ModelTemplateId *string `json:"ModelTemplateId,omitempty" xml:"ModelTemplateId,omitempty"`
}

func (s DeleteModelTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelTemplateRequest) GetModelTemplateId() *string {
	return s.ModelTemplateId
}

func (s *DeleteModelTemplateRequest) SetModelTemplateId(v string) *DeleteModelTemplateRequest {
	s.ModelTemplateId = &v
	return s
}

func (s *DeleteModelTemplateRequest) Validate() error {
	return dara.Validate(s)
}
