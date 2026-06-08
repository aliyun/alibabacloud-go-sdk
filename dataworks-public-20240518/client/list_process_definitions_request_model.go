// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProcessDefinitionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetType(v string) *ListProcessDefinitionsRequest
	GetType() *string
}

type ListProcessDefinitionsRequest struct {
	// example:
	//
	// MaxCompute
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListProcessDefinitionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListProcessDefinitionsRequest) GoString() string {
	return s.String()
}

func (s *ListProcessDefinitionsRequest) GetType() *string {
	return s.Type
}

func (s *ListProcessDefinitionsRequest) SetType(v string) *ListProcessDefinitionsRequest {
	s.Type = &v
	return s
}

func (s *ListProcessDefinitionsRequest) Validate() error {
	return dara.Validate(s)
}
