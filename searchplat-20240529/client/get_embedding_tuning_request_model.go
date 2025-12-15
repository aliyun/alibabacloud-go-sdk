// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmbeddingTuningRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInput(v [][]*float32) *GetEmbeddingTuningRequest
	GetInput() [][]*float32
	SetParameters(v map[string]interface{}) *GetEmbeddingTuningRequest
	GetParameters() map[string]interface{}
}

type GetEmbeddingTuningRequest struct {
	Input      [][]*float32           `json:"input,omitempty" xml:"input,omitempty" type:"Repeated"`
	Parameters map[string]interface{} `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s GetEmbeddingTuningRequest) String() string {
	return dara.Prettify(s)
}

func (s GetEmbeddingTuningRequest) GoString() string {
	return s.String()
}

func (s *GetEmbeddingTuningRequest) GetInput() [][]*float32 {
	return s.Input
}

func (s *GetEmbeddingTuningRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *GetEmbeddingTuningRequest) SetInput(v [][]*float32) *GetEmbeddingTuningRequest {
	s.Input = v
	return s
}

func (s *GetEmbeddingTuningRequest) SetParameters(v map[string]interface{}) *GetEmbeddingTuningRequest {
	s.Parameters = v
	return s
}

func (s *GetEmbeddingTuningRequest) Validate() error {
	return dara.Validate(s)
}
