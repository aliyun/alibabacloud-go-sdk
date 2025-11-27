// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeRCSecurityGroupPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeRCSecurityGroupPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeRCSecurityGroupPermissionResponse
	GetStatusCode() *int32
	SetBody(v *RevokeRCSecurityGroupPermissionResponseBody) *RevokeRCSecurityGroupPermissionResponse
	GetBody() *RevokeRCSecurityGroupPermissionResponseBody
}

type RevokeRCSecurityGroupPermissionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeRCSecurityGroupPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeRCSecurityGroupPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeRCSecurityGroupPermissionResponse) GoString() string {
	return s.String()
}

func (s *RevokeRCSecurityGroupPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeRCSecurityGroupPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeRCSecurityGroupPermissionResponse) GetBody() *RevokeRCSecurityGroupPermissionResponseBody {
	return s.Body
}

func (s *RevokeRCSecurityGroupPermissionResponse) SetHeaders(v map[string]*string) *RevokeRCSecurityGroupPermissionResponse {
	s.Headers = v
	return s
}

func (s *RevokeRCSecurityGroupPermissionResponse) SetStatusCode(v int32) *RevokeRCSecurityGroupPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeRCSecurityGroupPermissionResponse) SetBody(v *RevokeRCSecurityGroupPermissionResponseBody) *RevokeRCSecurityGroupPermissionResponse {
	s.Body = v
	return s
}

func (s *RevokeRCSecurityGroupPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
