// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantAccountPrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GrantAccountPrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GrantAccountPrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *GrantAccountPrivilegeResponseBody) *GrantAccountPrivilegeResponse
	GetBody() *GrantAccountPrivilegeResponseBody
}

type GrantAccountPrivilegeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GrantAccountPrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GrantAccountPrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s GrantAccountPrivilegeResponse) GoString() string {
	return s.String()
}

func (s *GrantAccountPrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GrantAccountPrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GrantAccountPrivilegeResponse) GetBody() *GrantAccountPrivilegeResponseBody {
	return s.Body
}

func (s *GrantAccountPrivilegeResponse) SetHeaders(v map[string]*string) *GrantAccountPrivilegeResponse {
	s.Headers = v
	return s
}

func (s *GrantAccountPrivilegeResponse) SetStatusCode(v int32) *GrantAccountPrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *GrantAccountPrivilegeResponse) SetBody(v *GrantAccountPrivilegeResponseBody) *GrantAccountPrivilegeResponse {
	s.Body = v
	return s
}

func (s *GrantAccountPrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
