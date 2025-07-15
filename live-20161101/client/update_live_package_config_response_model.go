// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePackageConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLivePackageConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLivePackageConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLivePackageConfigResponseBody) *UpdateLivePackageConfigResponse
	GetBody() *UpdateLivePackageConfigResponseBody
}

type UpdateLivePackageConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLivePackageConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLivePackageConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePackageConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLivePackageConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLivePackageConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLivePackageConfigResponse) GetBody() *UpdateLivePackageConfigResponseBody {
	return s.Body
}

func (s *UpdateLivePackageConfigResponse) SetHeaders(v map[string]*string) *UpdateLivePackageConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLivePackageConfigResponse) SetStatusCode(v int32) *UpdateLivePackageConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLivePackageConfigResponse) SetBody(v *UpdateLivePackageConfigResponseBody) *UpdateLivePackageConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLivePackageConfigResponse) Validate() error {
	return dara.Validate(s)
}
