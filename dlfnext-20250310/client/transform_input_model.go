// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformInput interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v interface{}) *TransformInput
	GetInput() interface{}
	SetType(v string) *TransformInput
	GetType() *string
}

type TransformInput struct {
	Input interface{} `json:"input,omitempty" xml:"input,omitempty"`
	Type  *string     `json:"type,omitempty" xml:"type,omitempty"`
}

func (s TransformInput) String() string {
	return dara.Prettify(s)
}

func (s TransformInput) GoString() string {
	return s.String()
}

func (s *TransformInput) GetInput() interface{} {
	return s.Input
}

func (s *TransformInput) GetType() *string {
	return s.Type
}

func (s *TransformInput) SetInput(v interface{}) *TransformInput {
	s.Input = v
	return s
}

func (s *TransformInput) SetType(v string) *TransformInput {
	s.Type = &v
	return s
}

func (s *TransformInput) Validate() error {
	return dara.Validate(s)
}
