// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendEmailVerificationForMessageContactResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendEmailVerificationForMessageContactResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendEmailVerificationForMessageContactResponse
	GetStatusCode() *int32
	SetBody(v *SendEmailVerificationForMessageContactResponseBody) *SendEmailVerificationForMessageContactResponse
	GetBody() *SendEmailVerificationForMessageContactResponseBody
}

type SendEmailVerificationForMessageContactResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendEmailVerificationForMessageContactResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendEmailVerificationForMessageContactResponse) String() string {
	return dara.Prettify(s)
}

func (s SendEmailVerificationForMessageContactResponse) GoString() string {
	return s.String()
}

func (s *SendEmailVerificationForMessageContactResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendEmailVerificationForMessageContactResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendEmailVerificationForMessageContactResponse) GetBody() *SendEmailVerificationForMessageContactResponseBody {
	return s.Body
}

func (s *SendEmailVerificationForMessageContactResponse) SetHeaders(v map[string]*string) *SendEmailVerificationForMessageContactResponse {
	s.Headers = v
	return s
}

func (s *SendEmailVerificationForMessageContactResponse) SetStatusCode(v int32) *SendEmailVerificationForMessageContactResponse {
	s.StatusCode = &v
	return s
}

func (s *SendEmailVerificationForMessageContactResponse) SetBody(v *SendEmailVerificationForMessageContactResponseBody) *SendEmailVerificationForMessageContactResponse {
	s.Body = v
	return s
}

func (s *SendEmailVerificationForMessageContactResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
