// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAccountPrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeAccountPrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeAccountPrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *RevokeAccountPrivilegeResponseBody) *RevokeAccountPrivilegeResponse
	GetBody() *RevokeAccountPrivilegeResponseBody
}

type RevokeAccountPrivilegeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeAccountPrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeAccountPrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeAccountPrivilegeResponse) GetBody() *RevokeAccountPrivilegeResponseBody {
	return s.Body
}

func (s *RevokeAccountPrivilegeResponse) SetHeaders(v map[string]*string) *RevokeAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *RevokeAccountPrivilegeResponse) SetStatusCode(v int32) *RevokeAccountPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeAccountPrivilegeResponse) SetBody(v *RevokeAccountPrivilegeResponseBody) *RevokeAccountPrivilegeResponse {
	s.Body = v
	return s
}

func (s *RevokeAccountPrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
