// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTransitRouterCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTransitRouterCidrResponse
	GetStatusCode() *int32
	SetBody(v *CreateTransitRouterCidrResponseBody) *CreateTransitRouterCidrResponse
	GetBody() *CreateTransitRouterCidrResponseBody
}

type CreateTransitRouterCidrResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTransitRouterCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTransitRouterCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterCidrResponse) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTransitRouterCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTransitRouterCidrResponse) GetBody() *CreateTransitRouterCidrResponseBody {
	return s.Body
}

func (s *CreateTransitRouterCidrResponse) SetHeaders(v map[string]*string) *CreateTransitRouterCidrResponse {
	s.Headers = v
	return s
}

func (s *CreateTransitRouterCidrResponse) SetStatusCode(v int32) *CreateTransitRouterCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTransitRouterCidrResponse) SetBody(v *CreateTransitRouterCidrResponseBody) *CreateTransitRouterCidrResponse {
	s.Body = v
	return s
}

func (s *CreateTransitRouterCidrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
