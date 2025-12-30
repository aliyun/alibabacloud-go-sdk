// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultimodalEmbeddingModelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultimodalEmbeddingModelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultimodalEmbeddingModelResponse
	GetStatusCode() *int32
	SetBody(v *ListMultimodalEmbeddingModelResponseBody) *ListMultimodalEmbeddingModelResponse
	GetBody() *ListMultimodalEmbeddingModelResponseBody
}

type ListMultimodalEmbeddingModelResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultimodalEmbeddingModelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultimodalEmbeddingModelResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultimodalEmbeddingModelResponse) GoString() string {
	return s.String()
}

func (s *ListMultimodalEmbeddingModelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultimodalEmbeddingModelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultimodalEmbeddingModelResponse) GetBody() *ListMultimodalEmbeddingModelResponseBody {
	return s.Body
}

func (s *ListMultimodalEmbeddingModelResponse) SetHeaders(v map[string]*string) *ListMultimodalEmbeddingModelResponse {
	s.Headers = v
	return s
}

func (s *ListMultimodalEmbeddingModelResponse) SetStatusCode(v int32) *ListMultimodalEmbeddingModelResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultimodalEmbeddingModelResponse) SetBody(v *ListMultimodalEmbeddingModelResponseBody) *ListMultimodalEmbeddingModelResponse {
	s.Body = v
	return s
}

func (s *ListMultimodalEmbeddingModelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
