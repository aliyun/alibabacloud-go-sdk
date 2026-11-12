// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetDatabaseAccountPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetDatabaseAccountPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetDatabaseAccountPasswordResponse
	GetStatusCode() *int32
	SetBody(v *ResetDatabaseAccountPasswordResponseBody) *ResetDatabaseAccountPasswordResponse
	GetBody() *ResetDatabaseAccountPasswordResponseBody
}

type ResetDatabaseAccountPasswordResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetDatabaseAccountPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetDatabaseAccountPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetDatabaseAccountPasswordResponse) GoString() string {
	return s.String()
}

func (s *ResetDatabaseAccountPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetDatabaseAccountPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetDatabaseAccountPasswordResponse) GetBody() *ResetDatabaseAccountPasswordResponseBody {
	return s.Body
}

func (s *ResetDatabaseAccountPasswordResponse) SetHeaders(v map[string]*string) *ResetDatabaseAccountPasswordResponse {
	s.Headers = v
	return s
}

func (s *ResetDatabaseAccountPasswordResponse) SetStatusCode(v int32) *ResetDatabaseAccountPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetDatabaseAccountPasswordResponse) SetBody(v *ResetDatabaseAccountPasswordResponseBody) *ResetDatabaseAccountPasswordResponse {
	s.Body = v
	return s
}

func (s *ResetDatabaseAccountPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
