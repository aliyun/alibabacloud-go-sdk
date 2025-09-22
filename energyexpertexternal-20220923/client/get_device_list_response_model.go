// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceListResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceListResponseBody) *GetDeviceListResponse
	GetBody() *GetDeviceListResponseBody
}

type GetDeviceListResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceListResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceListResponse) GetBody() *GetDeviceListResponseBody {
	return s.Body
}

func (s *GetDeviceListResponse) SetHeaders(v map[string]*string) *GetDeviceListResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceListResponse) SetStatusCode(v int32) *GetDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceListResponse) SetBody(v *GetDeviceListResponseBody) *GetDeviceListResponse {
	s.Body = v
	return s
}

func (s *GetDeviceListResponse) Validate() error {
	return dara.Validate(s)
}
