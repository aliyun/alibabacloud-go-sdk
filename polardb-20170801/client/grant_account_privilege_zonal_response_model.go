// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAccountPrivilegeZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantAccountPrivilegeZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantAccountPrivilegeZonalResponse
	GetStatusCode() *int32
	SetBody(v *GrantAccountPrivilegeZonalResponseBody) *GrantAccountPrivilegeZonalResponse
	GetBody() *GrantAccountPrivilegeZonalResponseBody
}

type GrantAccountPrivilegeZonalResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantAccountPrivilegeZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantAccountPrivilegeZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantAccountPrivilegeZonalResponse) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantAccountPrivilegeZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantAccountPrivilegeZonalResponse) GetBody() *GrantAccountPrivilegeZonalResponseBody {
	return s.Body
}

func (s *GrantAccountPrivilegeZonalResponse) SetHeaders(v map[string]*string) *GrantAccountPrivilegeZonalResponse {
	s.Headers = v
	return s
}

func (s *GrantAccountPrivilegeZonalResponse) SetStatusCode(v int32) *GrantAccountPrivilegeZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantAccountPrivilegeZonalResponse) SetBody(v *GrantAccountPrivilegeZonalResponseBody) *GrantAccountPrivilegeZonalResponse {
	s.Body = v
	return s
}

func (s *GrantAccountPrivilegeZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
