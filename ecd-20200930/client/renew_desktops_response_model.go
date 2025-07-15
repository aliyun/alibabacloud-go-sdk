// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewDesktopsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewDesktopsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewDesktopsResponse
	GetStatusCode() *int32
	SetBody(v *RenewDesktopsResponseBody) *RenewDesktopsResponse
	GetBody() *RenewDesktopsResponseBody
}

type RenewDesktopsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewDesktopsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewDesktopsResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewDesktopsResponse) GoString() string {
	return s.String()
}

func (s *RenewDesktopsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewDesktopsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewDesktopsResponse) GetBody() *RenewDesktopsResponseBody {
	return s.Body
}

func (s *RenewDesktopsResponse) SetHeaders(v map[string]*string) *RenewDesktopsResponse {
	s.Headers = v
	return s
}

func (s *RenewDesktopsResponse) SetStatusCode(v int32) *RenewDesktopsResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewDesktopsResponse) SetBody(v *RenewDesktopsResponseBody) *RenewDesktopsResponse {
	s.Body = v
	return s
}

func (s *RenewDesktopsResponse) Validate() error {
	return dara.Validate(s)
}
