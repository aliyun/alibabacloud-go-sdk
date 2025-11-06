// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEmailVerificationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEmailVerificationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEmailVerificationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEmailVerificationResponseBody) *DeleteEmailVerificationResponse
	GetBody() *DeleteEmailVerificationResponseBody
}

type DeleteEmailVerificationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEmailVerificationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEmailVerificationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEmailVerificationResponse) GoString() string {
	return s.String()
}

func (s *DeleteEmailVerificationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEmailVerificationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEmailVerificationResponse) GetBody() *DeleteEmailVerificationResponseBody {
	return s.Body
}

func (s *DeleteEmailVerificationResponse) SetHeaders(v map[string]*string) *DeleteEmailVerificationResponse {
	s.Headers = v
	return s
}

func (s *DeleteEmailVerificationResponse) SetStatusCode(v int32) *DeleteEmailVerificationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEmailVerificationResponse) SetBody(v *DeleteEmailVerificationResponseBody) *DeleteEmailVerificationResponse {
	s.Body = v
	return s
}

func (s *DeleteEmailVerificationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
