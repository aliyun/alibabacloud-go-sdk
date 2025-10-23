// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetDeviceInfoResponseBody) *GetDeviceInfoResponse
	GetBody() *GetDeviceInfoResponseBody
}

type GetDeviceInfoResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeviceInfoResponse) GetBody() *GetDeviceInfoResponseBody {
	return s.Body
}

func (s *GetDeviceInfoResponse) SetHeaders(v map[string]*string) *GetDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetDeviceInfoResponse) SetStatusCode(v int32) *GetDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeviceInfoResponse) SetBody(v *GetDeviceInfoResponseBody) *GetDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *GetDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
