// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDeviceWorkloadTrendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserDeviceWorkloadTrendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserDeviceWorkloadTrendResponse
	GetStatusCode() *int32
	SetBody(v *GetUserDeviceWorkloadTrendResponseBody) *GetUserDeviceWorkloadTrendResponse
	GetBody() *GetUserDeviceWorkloadTrendResponseBody
}

type GetUserDeviceWorkloadTrendResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserDeviceWorkloadTrendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserDeviceWorkloadTrendResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceWorkloadTrendResponse) GoString() string {
	return s.String()
}

func (s *GetUserDeviceWorkloadTrendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserDeviceWorkloadTrendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserDeviceWorkloadTrendResponse) GetBody() *GetUserDeviceWorkloadTrendResponseBody {
	return s.Body
}

func (s *GetUserDeviceWorkloadTrendResponse) SetHeaders(v map[string]*string) *GetUserDeviceWorkloadTrendResponse {
	s.Headers = v
	return s
}

func (s *GetUserDeviceWorkloadTrendResponse) SetStatusCode(v int32) *GetUserDeviceWorkloadTrendResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendResponse) SetBody(v *GetUserDeviceWorkloadTrendResponseBody) *GetUserDeviceWorkloadTrendResponse {
	s.Body = v
	return s
}

func (s *GetUserDeviceWorkloadTrendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
