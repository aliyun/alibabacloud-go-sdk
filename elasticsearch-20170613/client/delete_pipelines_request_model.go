// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeletePipelinesRequest
	GetClientToken() *string
	SetPipelineIds(v string) *DeletePipelinesRequest
	GetPipelineIds() *string
}

type DeletePipelinesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the pipeline.
	//
	// example:
	//
	// pipeline-test
	PipelineIds *string `json:"pipelineIds,omitempty" xml:"pipelineIds,omitempty"`
}

func (s DeletePipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelinesRequest) GoString() string {
	return s.String()
}

func (s *DeletePipelinesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeletePipelinesRequest) GetPipelineIds() *string {
	return s.PipelineIds
}

func (s *DeletePipelinesRequest) SetClientToken(v string) *DeletePipelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeletePipelinesRequest) SetPipelineIds(v string) *DeletePipelinesRequest {
	s.PipelineIds = &v
	return s
}

func (s *DeletePipelinesRequest) Validate() error {
	return dara.Validate(s)
}
