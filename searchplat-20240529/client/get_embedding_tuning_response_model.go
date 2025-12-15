// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEmbeddingTuningResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEmbeddingTuningResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEmbeddingTuningResponse
	GetStatusCode() *int32
	SetBody(v *GetEmbeddingTuningResponseBody) *GetEmbeddingTuningResponse
	GetBody() *GetEmbeddingTuningResponseBody
}

type GetEmbeddingTuningResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEmbeddingTuningResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEmbeddingTuningResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEmbeddingTuningResponse) GoString() string {
	return s.String()
}

func (s *GetEmbeddingTuningResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEmbeddingTuningResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEmbeddingTuningResponse) GetBody() *GetEmbeddingTuningResponseBody {
	return s.Body
}

func (s *GetEmbeddingTuningResponse) SetHeaders(v map[string]*string) *GetEmbeddingTuningResponse {
	s.Headers = v
	return s
}

func (s *GetEmbeddingTuningResponse) SetStatusCode(v int32) *GetEmbeddingTuningResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEmbeddingTuningResponse) SetBody(v *GetEmbeddingTuningResponseBody) *GetEmbeddingTuningResponse {
	s.Body = v
	return s
}

func (s *GetEmbeddingTuningResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
