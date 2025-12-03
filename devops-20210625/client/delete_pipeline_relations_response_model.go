// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePipelineRelationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePipelineRelationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePipelineRelationsResponse
	GetStatusCode() *int32
	SetBody(v *DeletePipelineRelationsResponseBody) *DeletePipelineRelationsResponse
	GetBody() *DeletePipelineRelationsResponseBody
}

type DeletePipelineRelationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePipelineRelationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePipelineRelationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePipelineRelationsResponse) GoString() string {
	return s.String()
}

func (s *DeletePipelineRelationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePipelineRelationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePipelineRelationsResponse) GetBody() *DeletePipelineRelationsResponseBody {
	return s.Body
}

func (s *DeletePipelineRelationsResponse) SetHeaders(v map[string]*string) *DeletePipelineRelationsResponse {
	s.Headers = v
	return s
}

func (s *DeletePipelineRelationsResponse) SetStatusCode(v int32) *DeletePipelineRelationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePipelineRelationsResponse) SetBody(v *DeletePipelineRelationsResponseBody) *DeletePipelineRelationsResponse {
	s.Body = v
	return s
}

func (s *DeletePipelineRelationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
