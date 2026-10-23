// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunPipelinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunPipelinesRequest
	GetClientToken() *string
	SetBody(v string) *RunPipelinesRequest
	GetBody() *string
}

type RunPipelinesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 5A2CFF0E-5718-45B5-9D4D-70B3FF****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of the Logstash pipelines to start, represented as a JSON array string.
	//
	// example:
	//
	// ["PipelineId1","PipelineId2"]
	Body *string `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunPipelinesRequest) String() string {
	return dara.Prettify(s)
}

func (s RunPipelinesRequest) GoString() string {
	return s.String()
}

func (s *RunPipelinesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunPipelinesRequest) GetBody() *string {
	return s.Body
}

func (s *RunPipelinesRequest) SetClientToken(v string) *RunPipelinesRequest {
	s.ClientToken = &v
	return s
}

func (s *RunPipelinesRequest) SetBody(v string) *RunPipelinesRequest {
	s.Body = &v
	return s
}

func (s *RunPipelinesRequest) Validate() error {
	return dara.Validate(s)
}
