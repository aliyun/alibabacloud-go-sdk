// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalEmbeddingModelModeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultimodalEmbeddingModelModeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultimodalEmbeddingModelModeResponse
	GetStatusCode() *int32
	SetBody(v *ListMultimodalEmbeddingModelModeResponseBody) *ListMultimodalEmbeddingModelModeResponse
	GetBody() *ListMultimodalEmbeddingModelModeResponseBody
}

type ListMultimodalEmbeddingModelModeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultimodalEmbeddingModelModeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultimodalEmbeddingModelModeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalEmbeddingModelModeResponse) GoString() string {
	return s.String()
}

func (s *ListMultimodalEmbeddingModelModeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultimodalEmbeddingModelModeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultimodalEmbeddingModelModeResponse) GetBody() *ListMultimodalEmbeddingModelModeResponseBody {
	return s.Body
}

func (s *ListMultimodalEmbeddingModelModeResponse) SetHeaders(v map[string]*string) *ListMultimodalEmbeddingModelModeResponse {
	s.Headers = v
	return s
}

func (s *ListMultimodalEmbeddingModelModeResponse) SetStatusCode(v int32) *ListMultimodalEmbeddingModelModeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultimodalEmbeddingModelModeResponse) SetBody(v *ListMultimodalEmbeddingModelModeResponseBody) *ListMultimodalEmbeddingModelModeResponse {
	s.Body = v
	return s
}

func (s *ListMultimodalEmbeddingModelModeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
