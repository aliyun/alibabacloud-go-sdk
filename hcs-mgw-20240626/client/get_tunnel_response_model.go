// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTunnelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTunnelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTunnelResponse
	GetStatusCode() *int32
	SetBody(v *GetTunnelResponseBody) *GetTunnelResponse
	GetBody() *GetTunnelResponseBody
}

type GetTunnelResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTunnelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTunnelResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTunnelResponse) GoString() string {
	return s.String()
}

func (s *GetTunnelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTunnelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTunnelResponse) GetBody() *GetTunnelResponseBody {
	return s.Body
}

func (s *GetTunnelResponse) SetHeaders(v map[string]*string) *GetTunnelResponse {
	s.Headers = v
	return s
}

func (s *GetTunnelResponse) SetStatusCode(v int32) *GetTunnelResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTunnelResponse) SetBody(v *GetTunnelResponseBody) *GetTunnelResponse {
	s.Body = v
	return s
}

func (s *GetTunnelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
