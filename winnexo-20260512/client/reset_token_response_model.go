// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetTokenResponse
	GetStatusCode() *int32
	SetBody(v *ResetTokenResponseBody) *ResetTokenResponse
	GetBody() *ResetTokenResponseBody
}

type ResetTokenResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetTokenResponse) GoString() string {
	return s.String()
}

func (s *ResetTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetTokenResponse) GetBody() *ResetTokenResponseBody {
	return s.Body
}

func (s *ResetTokenResponse) SetHeaders(v map[string]*string) *ResetTokenResponse {
	s.Headers = v
	return s
}

func (s *ResetTokenResponse) SetStatusCode(v int32) *ResetTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetTokenResponse) SetBody(v *ResetTokenResponseBody) *ResetTokenResponse {
	s.Body = v
	return s
}

func (s *ResetTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
