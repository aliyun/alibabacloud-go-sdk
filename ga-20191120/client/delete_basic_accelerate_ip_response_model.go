// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBasicAccelerateIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBasicAccelerateIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBasicAccelerateIpResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBasicAccelerateIpResponseBody) *DeleteBasicAccelerateIpResponse
	GetBody() *DeleteBasicAccelerateIpResponseBody
}

type DeleteBasicAccelerateIpResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBasicAccelerateIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBasicAccelerateIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBasicAccelerateIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteBasicAccelerateIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBasicAccelerateIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBasicAccelerateIpResponse) GetBody() *DeleteBasicAccelerateIpResponseBody {
	return s.Body
}

func (s *DeleteBasicAccelerateIpResponse) SetHeaders(v map[string]*string) *DeleteBasicAccelerateIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteBasicAccelerateIpResponse) SetStatusCode(v int32) *DeleteBasicAccelerateIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBasicAccelerateIpResponse) SetBody(v *DeleteBasicAccelerateIpResponseBody) *DeleteBasicAccelerateIpResponse {
	s.Body = v
	return s
}

func (s *DeleteBasicAccelerateIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
