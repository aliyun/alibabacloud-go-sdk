// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryDeviceInfoResponseBody) *QueryDeviceInfoResponse
	GetBody() *QueryDeviceInfoResponseBody
}

type QueryDeviceInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDeviceInfoResponse) GetBody() *QueryDeviceInfoResponseBody {
	return s.Body
}

func (s *QueryDeviceInfoResponse) SetHeaders(v map[string]*string) *QueryDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryDeviceInfoResponse) SetStatusCode(v int32) *QueryDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDeviceInfoResponse) SetBody(v *QueryDeviceInfoResponseBody) *QueryDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *QueryDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
