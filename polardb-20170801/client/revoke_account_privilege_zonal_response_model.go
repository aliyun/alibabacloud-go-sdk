// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeAccountPrivilegeZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeAccountPrivilegeZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeAccountPrivilegeZonalResponse
	GetStatusCode() *int32
	SetBody(v *RevokeAccountPrivilegeZonalResponseBody) *RevokeAccountPrivilegeZonalResponse
	GetBody() *RevokeAccountPrivilegeZonalResponseBody
}

type RevokeAccountPrivilegeZonalResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeAccountPrivilegeZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeAccountPrivilegeZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeAccountPrivilegeZonalResponse) GoString() string {
	return s.String()
}

func (s *RevokeAccountPrivilegeZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeAccountPrivilegeZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeAccountPrivilegeZonalResponse) GetBody() *RevokeAccountPrivilegeZonalResponseBody {
	return s.Body
}

func (s *RevokeAccountPrivilegeZonalResponse) SetHeaders(v map[string]*string) *RevokeAccountPrivilegeZonalResponse {
	s.Headers = v
	return s
}

func (s *RevokeAccountPrivilegeZonalResponse) SetStatusCode(v int32) *RevokeAccountPrivilegeZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeAccountPrivilegeZonalResponse) SetBody(v *RevokeAccountPrivilegeZonalResponseBody) *RevokeAccountPrivilegeZonalResponse {
	s.Body = v
	return s
}

func (s *RevokeAccountPrivilegeZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
