// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiTryOnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AiTryOnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AiTryOnResponse
	GetStatusCode() *int32
	SetBody(v *AiTryOnResponseBody) *AiTryOnResponse
	GetBody() *AiTryOnResponseBody
}

type AiTryOnResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AiTryOnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AiTryOnResponse) String() string {
	return dara.Prettify(s)
}

func (s AiTryOnResponse) GoString() string {
	return s.String()
}

func (s *AiTryOnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AiTryOnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AiTryOnResponse) GetBody() *AiTryOnResponseBody {
	return s.Body
}

func (s *AiTryOnResponse) SetHeaders(v map[string]*string) *AiTryOnResponse {
	s.Headers = v
	return s
}

func (s *AiTryOnResponse) SetStatusCode(v int32) *AiTryOnResponse {
	s.StatusCode = &v
	return s
}

func (s *AiTryOnResponse) SetBody(v *AiTryOnResponseBody) *AiTryOnResponse {
	s.Body = v
	return s
}

func (s *AiTryOnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
