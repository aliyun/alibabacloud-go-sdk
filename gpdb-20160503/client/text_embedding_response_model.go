// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTextEmbeddingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TextEmbeddingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TextEmbeddingResponse
	GetStatusCode() *int32
	SetBody(v *TextEmbeddingResponseBody) *TextEmbeddingResponse
	GetBody() *TextEmbeddingResponseBody
}

type TextEmbeddingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TextEmbeddingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TextEmbeddingResponse) String() string {
	return dara.Prettify(s)
}

func (s TextEmbeddingResponse) GoString() string {
	return s.String()
}

func (s *TextEmbeddingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TextEmbeddingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TextEmbeddingResponse) GetBody() *TextEmbeddingResponseBody {
	return s.Body
}

func (s *TextEmbeddingResponse) SetHeaders(v map[string]*string) *TextEmbeddingResponse {
	s.Headers = v
	return s
}

func (s *TextEmbeddingResponse) SetStatusCode(v int32) *TextEmbeddingResponse {
	s.StatusCode = &v
	return s
}

func (s *TextEmbeddingResponse) SetBody(v *TextEmbeddingResponseBody) *TextEmbeddingResponse {
	s.Body = v
	return s
}

func (s *TextEmbeddingResponse) Validate() error {
	return dara.Validate(s)
}
