// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceListResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceListResponseBody) *QueryDeviceListResponse
	GetBody() *QueryDeviceListResponseBody
}

type QueryDeviceListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceListResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceListResponse) GetBody() *QueryDeviceListResponseBody {
	return s.Body
}

func (s *QueryDeviceListResponse) SetHeaders(v map[string]*string) *QueryDeviceListResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceListResponse) SetStatusCode(v int32) *QueryDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceListResponse) SetBody(v *QueryDeviceListResponseBody) *QueryDeviceListResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceListResponse) Validate() error {
	return dara.Validate(s)
}
