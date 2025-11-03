// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNatIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNatIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNatIpResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNatIpResponseBody) *DeleteNatIpResponse
	GetBody() *DeleteNatIpResponseBody
}

type DeleteNatIpResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNatIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNatIpResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNatIpResponse) GoString() string {
	return s.String()
}

func (s *DeleteNatIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNatIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNatIpResponse) GetBody() *DeleteNatIpResponseBody {
	return s.Body
}

func (s *DeleteNatIpResponse) SetHeaders(v map[string]*string) *DeleteNatIpResponse {
	s.Headers = v
	return s
}

func (s *DeleteNatIpResponse) SetStatusCode(v int32) *DeleteNatIpResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNatIpResponse) SetBody(v *DeleteNatIpResponseBody) *DeleteNatIpResponse {
	s.Body = v
	return s
}

func (s *DeleteNatIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
