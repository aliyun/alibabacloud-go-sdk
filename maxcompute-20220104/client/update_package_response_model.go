// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePackageResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePackageResponseBody) *UpdatePackageResponse
	GetBody() *UpdatePackageResponseBody
}

type UpdatePackageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePackageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePackageResponse) GoString() string {
	return s.String()
}

func (s *UpdatePackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePackageResponse) GetBody() *UpdatePackageResponseBody {
	return s.Body
}

func (s *UpdatePackageResponse) SetHeaders(v map[string]*string) *UpdatePackageResponse {
	s.Headers = v
	return s
}

func (s *UpdatePackageResponse) SetStatusCode(v int32) *UpdatePackageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePackageResponse) SetBody(v *UpdatePackageResponseBody) *UpdatePackageResponse {
	s.Body = v
	return s
}

func (s *UpdatePackageResponse) Validate() error {
	return dara.Validate(s)
}
