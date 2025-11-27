// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSmsAuthorizationLetterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSmsAuthorizationLetterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSmsAuthorizationLetterResponse
	GetStatusCode() *int32
	SetBody(v *CreateSmsAuthorizationLetterResponseBody) *CreateSmsAuthorizationLetterResponse
	GetBody() *CreateSmsAuthorizationLetterResponseBody
}

type CreateSmsAuthorizationLetterResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSmsAuthorizationLetterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSmsAuthorizationLetterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSmsAuthorizationLetterResponse) GoString() string {
	return s.String()
}

func (s *CreateSmsAuthorizationLetterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSmsAuthorizationLetterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSmsAuthorizationLetterResponse) GetBody() *CreateSmsAuthorizationLetterResponseBody {
	return s.Body
}

func (s *CreateSmsAuthorizationLetterResponse) SetHeaders(v map[string]*string) *CreateSmsAuthorizationLetterResponse {
	s.Headers = v
	return s
}

func (s *CreateSmsAuthorizationLetterResponse) SetStatusCode(v int32) *CreateSmsAuthorizationLetterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSmsAuthorizationLetterResponse) SetBody(v *CreateSmsAuthorizationLetterResponseBody) *CreateSmsAuthorizationLetterResponse {
	s.Body = v
	return s
}

func (s *CreateSmsAuthorizationLetterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
