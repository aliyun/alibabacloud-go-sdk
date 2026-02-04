// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeInstanceFromTransitRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokeInstanceFromTransitRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokeInstanceFromTransitRouterResponse
	GetStatusCode() *int32
	SetBody(v *RevokeInstanceFromTransitRouterResponseBody) *RevokeInstanceFromTransitRouterResponse
	GetBody() *RevokeInstanceFromTransitRouterResponseBody
}

type RevokeInstanceFromTransitRouterResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokeInstanceFromTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokeInstanceFromTransitRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokeInstanceFromTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *RevokeInstanceFromTransitRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokeInstanceFromTransitRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokeInstanceFromTransitRouterResponse) GetBody() *RevokeInstanceFromTransitRouterResponseBody {
	return s.Body
}

func (s *RevokeInstanceFromTransitRouterResponse) SetHeaders(v map[string]*string) *RevokeInstanceFromTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *RevokeInstanceFromTransitRouterResponse) SetStatusCode(v int32) *RevokeInstanceFromTransitRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokeInstanceFromTransitRouterResponse) SetBody(v *RevokeInstanceFromTransitRouterResponseBody) *RevokeInstanceFromTransitRouterResponse {
	s.Body = v
	return s
}

func (s *RevokeInstanceFromTransitRouterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
