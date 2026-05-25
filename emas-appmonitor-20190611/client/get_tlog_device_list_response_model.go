// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTlogDeviceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTlogDeviceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTlogDeviceListResponse
	GetStatusCode() *int32
	SetBody(v *GetTlogDeviceListResponseBody) *GetTlogDeviceListResponse
	GetBody() *GetTlogDeviceListResponseBody
}

type GetTlogDeviceListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTlogDeviceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTlogDeviceListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTlogDeviceListResponse) GoString() string {
	return s.String()
}

func (s *GetTlogDeviceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTlogDeviceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTlogDeviceListResponse) GetBody() *GetTlogDeviceListResponseBody {
	return s.Body
}

func (s *GetTlogDeviceListResponse) SetHeaders(v map[string]*string) *GetTlogDeviceListResponse {
	s.Headers = v
	return s
}

func (s *GetTlogDeviceListResponse) SetStatusCode(v int32) *GetTlogDeviceListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTlogDeviceListResponse) SetBody(v *GetTlogDeviceListResponseBody) *GetTlogDeviceListResponse {
	s.Body = v
	return s
}

func (s *GetTlogDeviceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
