// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextEmbeddingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v []*string) *GetTextEmbeddingRequest
	GetInput() []*string
	SetInputType(v string) *GetTextEmbeddingRequest
	GetInputType() *string
}

type GetTextEmbeddingRequest struct {
	// This parameter is required.
	Input []*string `json:"input,omitempty" xml:"input,omitempty" type:"Repeated"`
	// example:
	//
	// document
	InputType *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
}

func (s GetTextEmbeddingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTextEmbeddingRequest) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingRequest) GetInput() []*string {
	return s.Input
}

func (s *GetTextEmbeddingRequest) GetInputType() *string {
	return s.InputType
}

func (s *GetTextEmbeddingRequest) SetInput(v []*string) *GetTextEmbeddingRequest {
	s.Input = v
	return s
}

func (s *GetTextEmbeddingRequest) SetInputType(v string) *GetTextEmbeddingRequest {
	s.InputType = &v
	return s
}

func (s *GetTextEmbeddingRequest) Validate() error {
	return dara.Validate(s)
}
