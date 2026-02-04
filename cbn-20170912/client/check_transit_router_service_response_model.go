// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTransitRouterServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckTransitRouterServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckTransitRouterServiceResponse
	GetStatusCode() *int32
	SetBody(v *CheckTransitRouterServiceResponseBody) *CheckTransitRouterServiceResponse
	GetBody() *CheckTransitRouterServiceResponseBody
}

type CheckTransitRouterServiceResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckTransitRouterServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckTransitRouterServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckTransitRouterServiceResponse) GoString() string {
	return s.String()
}

func (s *CheckTransitRouterServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckTransitRouterServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckTransitRouterServiceResponse) GetBody() *CheckTransitRouterServiceResponseBody {
	return s.Body
}

func (s *CheckTransitRouterServiceResponse) SetHeaders(v map[string]*string) *CheckTransitRouterServiceResponse {
	s.Headers = v
	return s
}

func (s *CheckTransitRouterServiceResponse) SetStatusCode(v int32) *CheckTransitRouterServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckTransitRouterServiceResponse) SetBody(v *CheckTransitRouterServiceResponseBody) *CheckTransitRouterServiceResponse {
	s.Body = v
	return s
}

func (s *CheckTransitRouterServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
