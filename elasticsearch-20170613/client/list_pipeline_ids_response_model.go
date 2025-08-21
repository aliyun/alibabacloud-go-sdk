// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineIdsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelineIdsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelineIdsResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelineIdsResponseBody) *ListPipelineIdsResponse
	GetBody() *ListPipelineIdsResponseBody
}

type ListPipelineIdsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineIdsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineIdsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineIdsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineIdsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelineIdsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelineIdsResponse) GetBody() *ListPipelineIdsResponseBody {
	return s.Body
}

func (s *ListPipelineIdsResponse) SetHeaders(v map[string]*string) *ListPipelineIdsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineIdsResponse) SetStatusCode(v int32) *ListPipelineIdsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineIdsResponse) SetBody(v *ListPipelineIdsResponseBody) *ListPipelineIdsResponse {
	s.Body = v
	return s
}

func (s *ListPipelineIdsResponse) Validate() error {
	return dara.Validate(s)
}
