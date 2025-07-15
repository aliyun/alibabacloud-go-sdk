// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePackageStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstancePackageStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstancePackageStateResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstancePackageStateResponseBody) *UpdateInstancePackageStateResponse
	GetBody() *UpdateInstancePackageStateResponseBody
}

type UpdateInstancePackageStateResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstancePackageStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstancePackageStateResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePackageStateResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstancePackageStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstancePackageStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstancePackageStateResponse) GetBody() *UpdateInstancePackageStateResponseBody {
	return s.Body
}

func (s *UpdateInstancePackageStateResponse) SetHeaders(v map[string]*string) *UpdateInstancePackageStateResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstancePackageStateResponse) SetStatusCode(v int32) *UpdateInstancePackageStateResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstancePackageStateResponse) SetBody(v *UpdateInstancePackageStateResponseBody) *UpdateInstancePackageStateResponse {
	s.Body = v
	return s
}

func (s *UpdateInstancePackageStateResponse) Validate() error {
	return dara.Validate(s)
}
