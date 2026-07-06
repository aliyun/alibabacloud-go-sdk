// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignOutOrgAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SignOutOrgAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SignOutOrgAccountResponse
	GetStatusCode() *int32
	SetBody(v *SignOutOrgAccountResponseBody) *SignOutOrgAccountResponse
	GetBody() *SignOutOrgAccountResponseBody
}

type SignOutOrgAccountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SignOutOrgAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SignOutOrgAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s SignOutOrgAccountResponse) GoString() string {
	return s.String()
}

func (s *SignOutOrgAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SignOutOrgAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SignOutOrgAccountResponse) GetBody() *SignOutOrgAccountResponseBody {
	return s.Body
}

func (s *SignOutOrgAccountResponse) SetHeaders(v map[string]*string) *SignOutOrgAccountResponse {
	s.Headers = v
	return s
}

func (s *SignOutOrgAccountResponse) SetStatusCode(v int32) *SignOutOrgAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *SignOutOrgAccountResponse) SetBody(v *SignOutOrgAccountResponseBody) *SignOutOrgAccountResponse {
	s.Body = v
	return s
}

func (s *SignOutOrgAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
