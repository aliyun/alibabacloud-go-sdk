// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPipelineRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPipelineRelationsResponse
	GetStatusCode() *int32
	SetBody(v *ListPipelineRelationsResponseBody) *ListPipelineRelationsResponse
	GetBody() *ListPipelineRelationsResponseBody
}

type ListPipelineRelationsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPipelineRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPipelineRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineRelationsResponse) GoString() string {
	return s.String()
}

func (s *ListPipelineRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPipelineRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPipelineRelationsResponse) GetBody() *ListPipelineRelationsResponseBody {
	return s.Body
}

func (s *ListPipelineRelationsResponse) SetHeaders(v map[string]*string) *ListPipelineRelationsResponse {
	s.Headers = v
	return s
}

func (s *ListPipelineRelationsResponse) SetStatusCode(v int32) *ListPipelineRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPipelineRelationsResponse) SetBody(v *ListPipelineRelationsResponseBody) *ListPipelineRelationsResponse {
	s.Body = v
	return s
}

func (s *ListPipelineRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
