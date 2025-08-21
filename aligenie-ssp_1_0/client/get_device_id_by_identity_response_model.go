// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceIdByIdentityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceIdByIdentityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceIdByIdentityResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceIdByIdentityResponseBody) *GetDeviceIdByIdentityResponse
	GetBody() *GetDeviceIdByIdentityResponseBody
}

type GetDeviceIdByIdentityResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceIdByIdentityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceIdByIdentityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceIdByIdentityResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceIdByIdentityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceIdByIdentityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceIdByIdentityResponse) GetBody() *GetDeviceIdByIdentityResponseBody {
	return s.Body
}

func (s *GetDeviceIdByIdentityResponse) SetHeaders(v map[string]*string) *GetDeviceIdByIdentityResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceIdByIdentityResponse) SetStatusCode(v int32) *GetDeviceIdByIdentityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceIdByIdentityResponse) SetBody(v *GetDeviceIdByIdentityResponseBody) *GetDeviceIdByIdentityResponse {
	s.Body = v
	return s
}

func (s *GetDeviceIdByIdentityResponse) Validate() error {
	return dara.Validate(s)
}
