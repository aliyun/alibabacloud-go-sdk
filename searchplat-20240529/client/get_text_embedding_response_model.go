// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTextEmbeddingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTextEmbeddingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTextEmbeddingResponse
	GetStatusCode() *int32
	SetBody(v *GetTextEmbeddingResponseBody) *GetTextEmbeddingResponse
	GetBody() *GetTextEmbeddingResponseBody
}

type GetTextEmbeddingResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTextEmbeddingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTextEmbeddingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTextEmbeddingResponse) GoString() string {
	return s.String()
}

func (s *GetTextEmbeddingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTextEmbeddingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTextEmbeddingResponse) GetBody() *GetTextEmbeddingResponseBody {
	return s.Body
}

func (s *GetTextEmbeddingResponse) SetHeaders(v map[string]*string) *GetTextEmbeddingResponse {
	s.Headers = v
	return s
}

func (s *GetTextEmbeddingResponse) SetStatusCode(v int32) *GetTextEmbeddingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTextEmbeddingResponse) SetBody(v *GetTextEmbeddingResponseBody) *GetTextEmbeddingResponse {
	s.Body = v
	return s
}

func (s *GetTextEmbeddingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
