// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogDeviceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTlogDeviceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTlogDeviceInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetTlogDeviceInfoResponseBody) *GetTlogDeviceInfoResponse
	GetBody() *GetTlogDeviceInfoResponseBody
}

type GetTlogDeviceInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTlogDeviceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTlogDeviceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTlogDeviceInfoResponse) GoString() string {
	return s.String()
}

func (s *GetTlogDeviceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTlogDeviceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTlogDeviceInfoResponse) GetBody() *GetTlogDeviceInfoResponseBody {
	return s.Body
}

func (s *GetTlogDeviceInfoResponse) SetHeaders(v map[string]*string) *GetTlogDeviceInfoResponse {
	s.Headers = v
	return s
}

func (s *GetTlogDeviceInfoResponse) SetStatusCode(v int32) *GetTlogDeviceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTlogDeviceInfoResponse) SetBody(v *GetTlogDeviceInfoResponseBody) *GetTlogDeviceInfoResponse {
	s.Body = v
	return s
}

func (s *GetTlogDeviceInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
