// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpcDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIpcDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIpcDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetIpcDeviceInfoResponseBody) *GetIpcDeviceInfoResponse
	GetBody() *GetIpcDeviceInfoResponseBody
}

type GetIpcDeviceInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpcDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpcDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIpcDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetIpcDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIpcDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIpcDeviceInfoResponse) GetBody() *GetIpcDeviceInfoResponseBody {
	return s.Body
}

func (s *GetIpcDeviceInfoResponse) SetHeaders(v map[string]*string) *GetIpcDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetIpcDeviceInfoResponse) SetStatusCode(v int32) *GetIpcDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpcDeviceInfoResponse) SetBody(v *GetIpcDeviceInfoResponseBody) *GetIpcDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *GetIpcDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
