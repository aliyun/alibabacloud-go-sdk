// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListPipelinesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListPipelinesRequest
	GetNextToken() *string
	SetPipelineName(v string) *ListPipelinesRequest
	GetPipelineName() *string
}

type ListPipelinesRequest struct {
	// example:
	//
	// 100
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// example:
	//
	// fff3442dac1de7950f44d5afc0c735ebd12e27f603b21d17ec30cb1d5c735b1ba7c4fb3a1c124bce
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// pipeline-name-1
	PipelineName *string `json:"pipelineName,omitempty" xml:"pipelineName,omitempty"`
}

func (s ListPipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesRequest) GoString() string {
	return s.String()
}

func (s *ListPipelinesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPipelinesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPipelinesRequest) GetPipelineName() *string {
	return s.PipelineName
}

func (s *ListPipelinesRequest) SetMaxResults(v int32) *ListPipelinesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPipelinesRequest) SetNextToken(v string) *ListPipelinesRequest {
	s.NextToken = &v
	return s
}

func (s *ListPipelinesRequest) SetPipelineName(v string) *ListPipelinesRequest {
	s.PipelineName = &v
	return s
}

func (s *ListPipelinesRequest) Validate() error {
	return dara.Validate(s)
}
