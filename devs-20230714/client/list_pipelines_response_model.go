// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelinesResponse
	GetStatusCode() *int32
	SetBody(v []*Pipeline) *ListPipelinesResponse
	GetBody() []*Pipeline
}

type ListPipelinesResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*Pipeline        `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s ListPipelinesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelinesResponse) GoString() string {
	return s.String()
}

func (s *ListPipelinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelinesResponse) GetBody() []*Pipeline {
	return s.Body
}

func (s *ListPipelinesResponse) SetHeaders(v map[string]*string) *ListPipelinesResponse {
	s.Headers = v
	return s
}

func (s *ListPipelinesResponse) SetStatusCode(v int32) *ListPipelinesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelinesResponse) SetBody(v []*Pipeline) *ListPipelinesResponse {
	s.Body = v
	return s
}

func (s *ListPipelinesResponse) Validate() error {
	return dara.Validate(s)
}
