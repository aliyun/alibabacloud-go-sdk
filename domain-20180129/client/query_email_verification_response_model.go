// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEmailVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEmailVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEmailVerificationResponse
	GetStatusCode() *int32
	SetBody(v *QueryEmailVerificationResponseBody) *QueryEmailVerificationResponse
	GetBody() *QueryEmailVerificationResponseBody
}

type QueryEmailVerificationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEmailVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *QueryEmailVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEmailVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEmailVerificationResponse) GetBody() *QueryEmailVerificationResponseBody {
	return s.Body
}

func (s *QueryEmailVerificationResponse) SetHeaders(v map[string]*string) *QueryEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *QueryEmailVerificationResponse) SetStatusCode(v int32) *QueryEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEmailVerificationResponse) SetBody(v *QueryEmailVerificationResponseBody) *QueryEmailVerificationResponse {
	s.Body = v
	return s
}

func (s *QueryEmailVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
