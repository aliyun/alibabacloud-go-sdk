// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifySentenceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifySentenceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifySentenceResponse
	GetStatusCode() *int32
	SetBody(v *VerifySentenceResponseBody) *VerifySentenceResponse
	GetBody() *VerifySentenceResponseBody
}

type VerifySentenceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifySentenceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifySentenceResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifySentenceResponse) GoString() string {
	return s.String()
}

func (s *VerifySentenceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifySentenceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifySentenceResponse) GetBody() *VerifySentenceResponseBody {
	return s.Body
}

func (s *VerifySentenceResponse) SetHeaders(v map[string]*string) *VerifySentenceResponse {
	s.Headers = v
	return s
}

func (s *VerifySentenceResponse) SetStatusCode(v int32) *VerifySentenceResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifySentenceResponse) SetBody(v *VerifySentenceResponseBody) *VerifySentenceResponse {
	s.Body = v
	return s
}

func (s *VerifySentenceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
