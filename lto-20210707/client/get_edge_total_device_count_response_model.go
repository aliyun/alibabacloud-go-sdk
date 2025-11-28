// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeTotalDeviceCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEdgeTotalDeviceCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEdgeTotalDeviceCountResponse
	GetStatusCode() *int32
	SetBody(v *GetEdgeTotalDeviceCountResponseBody) *GetEdgeTotalDeviceCountResponse
	GetBody() *GetEdgeTotalDeviceCountResponseBody
}

type GetEdgeTotalDeviceCountResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEdgeTotalDeviceCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEdgeTotalDeviceCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeTotalDeviceCountResponse) GoString() string {
	return s.String()
}

func (s *GetEdgeTotalDeviceCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEdgeTotalDeviceCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEdgeTotalDeviceCountResponse) GetBody() *GetEdgeTotalDeviceCountResponseBody {
	return s.Body
}

func (s *GetEdgeTotalDeviceCountResponse) SetHeaders(v map[string]*string) *GetEdgeTotalDeviceCountResponse {
	s.Headers = v
	return s
}

func (s *GetEdgeTotalDeviceCountResponse) SetStatusCode(v int32) *GetEdgeTotalDeviceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEdgeTotalDeviceCountResponse) SetBody(v *GetEdgeTotalDeviceCountResponseBody) *GetEdgeTotalDeviceCountResponse {
	s.Body = v
	return s
}

func (s *GetEdgeTotalDeviceCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
