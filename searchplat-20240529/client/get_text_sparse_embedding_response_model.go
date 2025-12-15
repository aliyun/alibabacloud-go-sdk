// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextSparseEmbeddingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTextSparseEmbeddingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTextSparseEmbeddingResponse
	GetStatusCode() *int32
	SetBody(v *GetTextSparseEmbeddingResponseBody) *GetTextSparseEmbeddingResponse
	GetBody() *GetTextSparseEmbeddingResponseBody
}

type GetTextSparseEmbeddingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextSparseEmbeddingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextSparseEmbeddingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTextSparseEmbeddingResponse) GoString() string {
	return s.String()
}

func (s *GetTextSparseEmbeddingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTextSparseEmbeddingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTextSparseEmbeddingResponse) GetBody() *GetTextSparseEmbeddingResponseBody {
	return s.Body
}

func (s *GetTextSparseEmbeddingResponse) SetHeaders(v map[string]*string) *GetTextSparseEmbeddingResponse {
	s.Headers = v
	return s
}

func (s *GetTextSparseEmbeddingResponse) SetStatusCode(v int32) *GetTextSparseEmbeddingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextSparseEmbeddingResponse) SetBody(v *GetTextSparseEmbeddingResponseBody) *GetTextSparseEmbeddingResponse {
	s.Body = v
	return s
}

func (s *GetTextSparseEmbeddingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
