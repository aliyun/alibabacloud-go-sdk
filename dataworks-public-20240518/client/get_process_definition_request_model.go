// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetProcessDefinitionRequest
	GetId() *string
}

type GetProcessDefinitionRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f0d6d578-a305-40ac-ba1e-0a09f64cbc69
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetProcessDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetProcessDefinitionRequest) GoString() string {
	return s.String()
}

func (s *GetProcessDefinitionRequest) GetId() *string {
	return s.Id
}

func (s *GetProcessDefinitionRequest) SetId(v string) *GetProcessDefinitionRequest {
	s.Id = &v
	return s
}

func (s *GetProcessDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
