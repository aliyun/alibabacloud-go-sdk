// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualTryOnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VirtualTryOnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VirtualTryOnResponse
	GetStatusCode() *int32
	SetBody(v *VirtualTryOnResponseBody) *VirtualTryOnResponse
	GetBody() *VirtualTryOnResponseBody
}

type VirtualTryOnResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VirtualTryOnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VirtualTryOnResponse) String() string {
	return dara.Prettify(s)
}

func (s VirtualTryOnResponse) GoString() string {
	return s.String()
}

func (s *VirtualTryOnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VirtualTryOnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VirtualTryOnResponse) GetBody() *VirtualTryOnResponseBody {
	return s.Body
}

func (s *VirtualTryOnResponse) SetHeaders(v map[string]*string) *VirtualTryOnResponse {
	s.Headers = v
	return s
}

func (s *VirtualTryOnResponse) SetStatusCode(v int32) *VirtualTryOnResponse {
	s.StatusCode = &v
	return s
}

func (s *VirtualTryOnResponse) SetBody(v *VirtualTryOnResponseBody) *VirtualTryOnResponse {
	s.Body = v
	return s
}

func (s *VirtualTryOnResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
