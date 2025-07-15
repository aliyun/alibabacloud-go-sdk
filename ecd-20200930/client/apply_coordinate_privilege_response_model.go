// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyCoordinatePrivilegeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApplyCoordinatePrivilegeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApplyCoordinatePrivilegeResponse
	GetStatusCode() *int32
	SetBody(v *ApplyCoordinatePrivilegeResponseBody) *ApplyCoordinatePrivilegeResponse
	GetBody() *ApplyCoordinatePrivilegeResponseBody
}

type ApplyCoordinatePrivilegeResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApplyCoordinatePrivilegeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApplyCoordinatePrivilegeResponse) String() string {
	return dara.Prettify(s)
}

func (s ApplyCoordinatePrivilegeResponse) GoString() string {
	return s.String()
}

func (s *ApplyCoordinatePrivilegeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApplyCoordinatePrivilegeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApplyCoordinatePrivilegeResponse) GetBody() *ApplyCoordinatePrivilegeResponseBody {
	return s.Body
}

func (s *ApplyCoordinatePrivilegeResponse) SetHeaders(v map[string]*string) *ApplyCoordinatePrivilegeResponse {
	s.Headers = v
	return s
}

func (s *ApplyCoordinatePrivilegeResponse) SetStatusCode(v int32) *ApplyCoordinatePrivilegeResponse {
	s.StatusCode = &v
	return s
}

func (s *ApplyCoordinatePrivilegeResponse) SetBody(v *ApplyCoordinatePrivilegeResponseBody) *ApplyCoordinatePrivilegeResponse {
	s.Body = v
	return s
}

func (s *ApplyCoordinatePrivilegeResponse) Validate() error {
	return dara.Validate(s)
}
