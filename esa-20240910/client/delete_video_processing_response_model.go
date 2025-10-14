// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVideoProcessingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVideoProcessingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVideoProcessingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVideoProcessingResponseBody) *DeleteVideoProcessingResponse
	GetBody() *DeleteVideoProcessingResponseBody
}

type DeleteVideoProcessingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVideoProcessingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVideoProcessingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVideoProcessingResponse) GoString() string {
	return s.String()
}

func (s *DeleteVideoProcessingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVideoProcessingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVideoProcessingResponse) GetBody() *DeleteVideoProcessingResponseBody {
	return s.Body
}

func (s *DeleteVideoProcessingResponse) SetHeaders(v map[string]*string) *DeleteVideoProcessingResponse {
	s.Headers = v
	return s
}

func (s *DeleteVideoProcessingResponse) SetStatusCode(v int32) *DeleteVideoProcessingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVideoProcessingResponse) SetBody(v *DeleteVideoProcessingResponseBody) *DeleteVideoProcessingResponse {
	s.Body = v
	return s
}

func (s *DeleteVideoProcessingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
