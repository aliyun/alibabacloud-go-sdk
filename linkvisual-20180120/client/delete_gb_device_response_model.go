// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGbDeviceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGbDeviceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGbDeviceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGbDeviceResponseBody) *DeleteGbDeviceResponse
	GetBody() *DeleteGbDeviceResponseBody
}

type DeleteGbDeviceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGbDeviceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGbDeviceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGbDeviceResponse) GoString() string {
	return s.String()
}

func (s *DeleteGbDeviceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGbDeviceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGbDeviceResponse) GetBody() *DeleteGbDeviceResponseBody {
	return s.Body
}

func (s *DeleteGbDeviceResponse) SetHeaders(v map[string]*string) *DeleteGbDeviceResponse {
	s.Headers = v
	return s
}

func (s *DeleteGbDeviceResponse) SetStatusCode(v int32) *DeleteGbDeviceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGbDeviceResponse) SetBody(v *DeleteGbDeviceResponseBody) *DeleteGbDeviceResponse {
	s.Body = v
	return s
}

func (s *DeleteGbDeviceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
