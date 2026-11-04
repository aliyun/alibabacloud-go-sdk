// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRtmpDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRtmpDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRtmpDeviceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRtmpDeviceResponseBody) *DeleteRtmpDeviceResponse
	GetBody() *DeleteRtmpDeviceResponseBody
}

type DeleteRtmpDeviceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRtmpDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRtmpDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRtmpDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteRtmpDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRtmpDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRtmpDeviceResponse) GetBody() *DeleteRtmpDeviceResponseBody {
	return s.Body
}

func (s *DeleteRtmpDeviceResponse) SetHeaders(v map[string]*string) *DeleteRtmpDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteRtmpDeviceResponse) SetStatusCode(v int32) *DeleteRtmpDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRtmpDeviceResponse) SetBody(v *DeleteRtmpDeviceResponseBody) *DeleteRtmpDeviceResponse {
	s.Body = v
	return s
}

func (s *DeleteRtmpDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
