// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeCoordinatePrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeCoordinatePrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeCoordinatePrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *RevokeCoordinatePrivilegeResponseBody) *RevokeCoordinatePrivilegeResponse
	GetBody() *RevokeCoordinatePrivilegeResponseBody
}

type RevokeCoordinatePrivilegeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeCoordinatePrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeCoordinatePrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeCoordinatePrivilegeResponse) GoString() string {
	return s.String()
}

func (s *RevokeCoordinatePrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeCoordinatePrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeCoordinatePrivilegeResponse) GetBody() *RevokeCoordinatePrivilegeResponseBody {
	return s.Body
}

func (s *RevokeCoordinatePrivilegeResponse) SetHeaders(v map[string]*string) *RevokeCoordinatePrivilegeResponse {
	s.Headers = v
	return s
}

func (s *RevokeCoordinatePrivilegeResponse) SetStatusCode(v int32) *RevokeCoordinatePrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeCoordinatePrivilegeResponse) SetBody(v *RevokeCoordinatePrivilegeResponseBody) *RevokeCoordinatePrivilegeResponse {
	s.Body = v
	return s
}

func (s *RevokeCoordinatePrivilegeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
