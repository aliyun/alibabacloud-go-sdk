// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMultimodalEmbeddingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMultimodalEmbeddingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMultimodalEmbeddingResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMultimodalEmbeddingResponseBody) *DeleteMultimodalEmbeddingResponse
	GetBody() *DeleteMultimodalEmbeddingResponseBody
}

type DeleteMultimodalEmbeddingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMultimodalEmbeddingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMultimodalEmbeddingResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMultimodalEmbeddingResponse) GoString() string {
	return s.String()
}

func (s *DeleteMultimodalEmbeddingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMultimodalEmbeddingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMultimodalEmbeddingResponse) GetBody() *DeleteMultimodalEmbeddingResponseBody {
	return s.Body
}

func (s *DeleteMultimodalEmbeddingResponse) SetHeaders(v map[string]*string) *DeleteMultimodalEmbeddingResponse {
	s.Headers = v
	return s
}

func (s *DeleteMultimodalEmbeddingResponse) SetStatusCode(v int32) *DeleteMultimodalEmbeddingResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMultimodalEmbeddingResponse) SetBody(v *DeleteMultimodalEmbeddingResponseBody) *DeleteMultimodalEmbeddingResponse {
	s.Body = v
	return s
}

func (s *DeleteMultimodalEmbeddingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
