// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterResponseBody) *CreateTransitRouterResponse
	GetBody() *CreateTransitRouterResponseBody
}

type CreateTransitRouterResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterResponse) GetBody() *CreateTransitRouterResponseBody {
	return s.Body
}

func (s *CreateTransitRouterResponse) SetHeaders(v map[string]*string) *CreateTransitRouterResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterResponse) SetStatusCode(v int32) *CreateTransitRouterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterResponse) SetBody(v *CreateTransitRouterResponseBody) *CreateTransitRouterResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
