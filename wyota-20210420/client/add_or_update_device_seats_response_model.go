// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateDeviceSeatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddOrUpdateDeviceSeatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddOrUpdateDeviceSeatsResponse
	GetStatusCode() *int32
	SetBody(v *AddOrUpdateDeviceSeatsResponseBody) *AddOrUpdateDeviceSeatsResponse
	GetBody() *AddOrUpdateDeviceSeatsResponseBody
}

type AddOrUpdateDeviceSeatsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrUpdateDeviceSeatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrUpdateDeviceSeatsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateDeviceSeatsResponse) GoString() string {
	return s.String()
}

func (s *AddOrUpdateDeviceSeatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddOrUpdateDeviceSeatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrUpdateDeviceSeatsResponse) GetBody() *AddOrUpdateDeviceSeatsResponseBody {
	return s.Body
}

func (s *AddOrUpdateDeviceSeatsResponse) SetHeaders(v map[string]*string) *AddOrUpdateDeviceSeatsResponse {
	s.Headers = v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponse) SetStatusCode(v int32) *AddOrUpdateDeviceSeatsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponse) SetBody(v *AddOrUpdateDeviceSeatsResponseBody) *AddOrUpdateDeviceSeatsResponse {
	s.Body = v
	return s
}

func (s *AddOrUpdateDeviceSeatsResponse) Validate() error {
	return dara.Validate(s)
}
