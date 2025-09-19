// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMonitorAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMonitorAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMonitorAccountResponse
	GetStatusCode() *int32
	SetBody(v *CreateMonitorAccountResponseBody) *CreateMonitorAccountResponse
	GetBody() *CreateMonitorAccountResponseBody
}

type CreateMonitorAccountResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMonitorAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMonitorAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMonitorAccountResponse) GoString() string {
	return s.String()
}

func (s *CreateMonitorAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMonitorAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMonitorAccountResponse) GetBody() *CreateMonitorAccountResponseBody {
	return s.Body
}

func (s *CreateMonitorAccountResponse) SetHeaders(v map[string]*string) *CreateMonitorAccountResponse {
	s.Headers = v
	return s
}

func (s *CreateMonitorAccountResponse) SetStatusCode(v int32) *CreateMonitorAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMonitorAccountResponse) SetBody(v *CreateMonitorAccountResponseBody) *CreateMonitorAccountResponse {
	s.Body = v
	return s
}

func (s *CreateMonitorAccountResponse) Validate() error {
	return dara.Validate(s)
}
