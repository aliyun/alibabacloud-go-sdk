// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableCredentialResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableCredentialResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableCredentialResponse
	GetStatusCode() *int32
	SetBody(v *DisableCredentialResponseBody) *DisableCredentialResponse
	GetBody() *DisableCredentialResponseBody
}

type DisableCredentialResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableCredentialResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableCredentialResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableCredentialResponse) GoString() string {
	return s.String()
}

func (s *DisableCredentialResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableCredentialResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableCredentialResponse) GetBody() *DisableCredentialResponseBody {
	return s.Body
}

func (s *DisableCredentialResponse) SetHeaders(v map[string]*string) *DisableCredentialResponse {
	s.Headers = v
	return s
}

func (s *DisableCredentialResponse) SetStatusCode(v int32) *DisableCredentialResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableCredentialResponse) SetBody(v *DisableCredentialResponseBody) *DisableCredentialResponse {
	s.Body = v
	return s
}

func (s *DisableCredentialResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
