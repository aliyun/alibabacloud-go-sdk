// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDesktopMaintenanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDesktopMaintenanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDesktopMaintenanceResponse
	GetStatusCode() *int32
	SetBody(v *SetDesktopMaintenanceResponseBody) *SetDesktopMaintenanceResponse
	GetBody() *SetDesktopMaintenanceResponseBody
}

type SetDesktopMaintenanceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDesktopMaintenanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDesktopMaintenanceResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDesktopMaintenanceResponse) GoString() string {
	return s.String()
}

func (s *SetDesktopMaintenanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDesktopMaintenanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDesktopMaintenanceResponse) GetBody() *SetDesktopMaintenanceResponseBody {
	return s.Body
}

func (s *SetDesktopMaintenanceResponse) SetHeaders(v map[string]*string) *SetDesktopMaintenanceResponse {
	s.Headers = v
	return s
}

func (s *SetDesktopMaintenanceResponse) SetStatusCode(v int32) *SetDesktopMaintenanceResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDesktopMaintenanceResponse) SetBody(v *SetDesktopMaintenanceResponseBody) *SetDesktopMaintenanceResponse {
	s.Body = v
	return s
}

func (s *SetDesktopMaintenanceResponse) Validate() error {
	return dara.Validate(s)
}
