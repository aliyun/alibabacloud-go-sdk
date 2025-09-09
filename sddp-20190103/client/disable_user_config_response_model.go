// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableUserConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableUserConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableUserConfigResponse
	GetStatusCode() *int32
	SetBody(v *DisableUserConfigResponseBody) *DisableUserConfigResponse
	GetBody() *DisableUserConfigResponseBody
}

type DisableUserConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableUserConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableUserConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableUserConfigResponse) GoString() string {
	return s.String()
}

func (s *DisableUserConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableUserConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableUserConfigResponse) GetBody() *DisableUserConfigResponseBody {
	return s.Body
}

func (s *DisableUserConfigResponse) SetHeaders(v map[string]*string) *DisableUserConfigResponse {
	s.Headers = v
	return s
}

func (s *DisableUserConfigResponse) SetStatusCode(v int32) *DisableUserConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableUserConfigResponse) SetBody(v *DisableUserConfigResponseBody) *DisableUserConfigResponse {
	s.Body = v
	return s
}

func (s *DisableUserConfigResponse) Validate() error {
	return dara.Validate(s)
}
