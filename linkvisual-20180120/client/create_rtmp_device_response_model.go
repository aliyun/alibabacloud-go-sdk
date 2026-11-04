// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRtmpDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRtmpDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRtmpDeviceResponse
	GetStatusCode() *int32
	SetBody(v *CreateRtmpDeviceResponseBody) *CreateRtmpDeviceResponse
	GetBody() *CreateRtmpDeviceResponseBody
}

type CreateRtmpDeviceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRtmpDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRtmpDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRtmpDeviceResponse) GoString() string {
	return s.String()
}

func (s *CreateRtmpDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRtmpDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRtmpDeviceResponse) GetBody() *CreateRtmpDeviceResponseBody {
	return s.Body
}

func (s *CreateRtmpDeviceResponse) SetHeaders(v map[string]*string) *CreateRtmpDeviceResponse {
	s.Headers = v
	return s
}

func (s *CreateRtmpDeviceResponse) SetStatusCode(v int32) *CreateRtmpDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRtmpDeviceResponse) SetBody(v *CreateRtmpDeviceResponseBody) *CreateRtmpDeviceResponse {
	s.Body = v
	return s
}

func (s *CreateRtmpDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
