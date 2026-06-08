// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProcessDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteProcessDefinitionRequest
	GetId() *string
}

type DeleteProcessDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f0d6d578-a305-40ac-ba1e-0a09f64cbc69
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteProcessDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProcessDefinitionRequest) GoString() string {
	return s.String()
}

func (s *DeleteProcessDefinitionRequest) GetId() *string {
	return s.Id
}

func (s *DeleteProcessDefinitionRequest) SetId(v string) *DeleteProcessDefinitionRequest {
	s.Id = &v
	return s
}

func (s *DeleteProcessDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
