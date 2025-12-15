// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextSparseEmbeddingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v []*string) *GetTextSparseEmbeddingRequest
	GetInput() []*string
	SetInputType(v string) *GetTextSparseEmbeddingRequest
	GetInputType() *string
	SetReturnToken(v bool) *GetTextSparseEmbeddingRequest
	GetReturnToken() *bool
}

type GetTextSparseEmbeddingRequest struct {
	// This parameter is required.
	Input []*string `json:"input,omitempty" xml:"input,omitempty" type:"Repeated"`
	// example:
	//
	// document
	InputType   *string `json:"input_type,omitempty" xml:"input_type,omitempty"`
	ReturnToken *bool   `json:"return_token,omitempty" xml:"return_token,omitempty"`
}

func (s GetTextSparseEmbeddingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTextSparseEmbeddingRequest) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingRequest) GetInput() []*string {
	return s.Input
}

func (s *GetTextSparseEmbeddingRequest) GetInputType() *string {
	return s.InputType
}

func (s *GetTextSparseEmbeddingRequest) GetReturnToken() *bool {
	return s.ReturnToken
}

func (s *GetTextSparseEmbeddingRequest) SetInput(v []*string) *GetTextSparseEmbeddingRequest {
	s.Input = v
	return s
}

func (s *GetTextSparseEmbeddingRequest) SetInputType(v string) *GetTextSparseEmbeddingRequest {
	s.InputType = &v
	return s
}

func (s *GetTextSparseEmbeddingRequest) SetReturnToken(v bool) *GetTextSparseEmbeddingRequest {
	s.ReturnToken = &v
	return s
}

func (s *GetTextSparseEmbeddingRequest) Validate() error {
	return dara.Validate(s)
}
