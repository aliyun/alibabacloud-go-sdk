// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceStatusResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceStatusResponseBody) *QueryDeviceStatusResponse
	GetBody() *QueryDeviceStatusResponseBody
}

type QueryDeviceStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceStatusResponse) GetBody() *QueryDeviceStatusResponseBody {
	return s.Body
}

func (s *QueryDeviceStatusResponse) SetHeaders(v map[string]*string) *QueryDeviceStatusResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceStatusResponse) SetStatusCode(v int32) *QueryDeviceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceStatusResponse) SetBody(v *QueryDeviceStatusResponseBody) *QueryDeviceStatusResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
