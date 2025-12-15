// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMultiModalEmbeddingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMultiModalEmbeddingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMultiModalEmbeddingResponse
	GetStatusCode() *int32
	SetBody(v *GetMultiModalEmbeddingResponseBody) *GetMultiModalEmbeddingResponse
	GetBody() *GetMultiModalEmbeddingResponseBody
}

type GetMultiModalEmbeddingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMultiModalEmbeddingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMultiModalEmbeddingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMultiModalEmbeddingResponse) GoString() string {
	return s.String()
}

func (s *GetMultiModalEmbeddingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMultiModalEmbeddingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMultiModalEmbeddingResponse) GetBody() *GetMultiModalEmbeddingResponseBody {
	return s.Body
}

func (s *GetMultiModalEmbeddingResponse) SetHeaders(v map[string]*string) *GetMultiModalEmbeddingResponse {
	s.Headers = v
	return s
}

func (s *GetMultiModalEmbeddingResponse) SetStatusCode(v int32) *GetMultiModalEmbeddingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMultiModalEmbeddingResponse) SetBody(v *GetMultiModalEmbeddingResponseBody) *GetMultiModalEmbeddingResponse {
	s.Body = v
	return s
}

func (s *GetMultiModalEmbeddingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
