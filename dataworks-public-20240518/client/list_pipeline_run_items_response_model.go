// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRunItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelineRunItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelineRunItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelineRunItemsResponseBody) *ListPipelineRunItemsResponse
	GetBody() *ListPipelineRunItemsResponseBody
}

type ListPipelineRunItemsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineRunItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineRunItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRunItemsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineRunItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelineRunItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelineRunItemsResponse) GetBody() *ListPipelineRunItemsResponseBody {
	return s.Body
}

func (s *ListPipelineRunItemsResponse) SetHeaders(v map[string]*string) *ListPipelineRunItemsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineRunItemsResponse) SetStatusCode(v int32) *ListPipelineRunItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineRunItemsResponse) SetBody(v *ListPipelineRunItemsResponseBody) *ListPipelineRunItemsResponse {
	s.Body = v
	return s
}

func (s *ListPipelineRunItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
