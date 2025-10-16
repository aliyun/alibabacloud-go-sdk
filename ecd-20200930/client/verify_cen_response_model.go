// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyCenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VerifyCenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VerifyCenResponse
	GetStatusCode() *int32
	SetBody(v *VerifyCenResponseBody) *VerifyCenResponse
	GetBody() *VerifyCenResponseBody
}

type VerifyCenResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VerifyCenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VerifyCenResponse) String() string {
	return dara.Prettify(s)
}

func (s VerifyCenResponse) GoString() string {
	return s.String()
}

func (s *VerifyCenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VerifyCenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VerifyCenResponse) GetBody() *VerifyCenResponseBody {
	return s.Body
}

func (s *VerifyCenResponse) SetHeaders(v map[string]*string) *VerifyCenResponse {
	s.Headers = v
	return s
}

func (s *VerifyCenResponse) SetStatusCode(v int32) *VerifyCenResponse {
	s.StatusCode = &v
	return s
}

func (s *VerifyCenResponse) SetBody(v *VerifyCenResponseBody) *VerifyCenResponse {
	s.Body = v
	return s
}

func (s *VerifyCenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
