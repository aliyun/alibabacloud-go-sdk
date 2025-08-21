// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindMFADeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindMFADeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindMFADeviceResponse
	GetStatusCode() *int32
	SetBody(v *BindMFADeviceResponseBody) *BindMFADeviceResponse
	GetBody() *BindMFADeviceResponseBody
}

type BindMFADeviceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindMFADeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindMFADeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s BindMFADeviceResponse) GoString() string {
	return s.String()
}

func (s *BindMFADeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindMFADeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindMFADeviceResponse) GetBody() *BindMFADeviceResponseBody {
	return s.Body
}

func (s *BindMFADeviceResponse) SetHeaders(v map[string]*string) *BindMFADeviceResponse {
	s.Headers = v
	return s
}

func (s *BindMFADeviceResponse) SetStatusCode(v int32) *BindMFADeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *BindMFADeviceResponse) SetBody(v *BindMFADeviceResponseBody) *BindMFADeviceResponse {
	s.Body = v
	return s
}

func (s *BindMFADeviceResponse) Validate() error {
	return dara.Validate(s)
}
