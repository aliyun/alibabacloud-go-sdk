// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateDatabaseUserTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ValidateDatabaseUserTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ValidateDatabaseUserTokenResponse
	GetStatusCode() *int32
	SetBody(v *ValidateDatabaseUserTokenResponseBody) *ValidateDatabaseUserTokenResponse
	GetBody() *ValidateDatabaseUserTokenResponseBody
}

type ValidateDatabaseUserTokenResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ValidateDatabaseUserTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ValidateDatabaseUserTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s ValidateDatabaseUserTokenResponse) GoString() string {
	return s.String()
}

func (s *ValidateDatabaseUserTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ValidateDatabaseUserTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ValidateDatabaseUserTokenResponse) GetBody() *ValidateDatabaseUserTokenResponseBody {
	return s.Body
}

func (s *ValidateDatabaseUserTokenResponse) SetHeaders(v map[string]*string) *ValidateDatabaseUserTokenResponse {
	s.Headers = v
	return s
}

func (s *ValidateDatabaseUserTokenResponse) SetStatusCode(v int32) *ValidateDatabaseUserTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *ValidateDatabaseUserTokenResponse) SetBody(v *ValidateDatabaseUserTokenResponseBody) *ValidateDatabaseUserTokenResponse {
	s.Body = v
	return s
}

func (s *ValidateDatabaseUserTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
