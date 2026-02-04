// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenTransitRouterServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenTransitRouterServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenTransitRouterServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenTransitRouterServiceResponseBody) *OpenTransitRouterServiceResponse
	GetBody() *OpenTransitRouterServiceResponseBody
}

type OpenTransitRouterServiceResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenTransitRouterServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenTransitRouterServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenTransitRouterServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenTransitRouterServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenTransitRouterServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenTransitRouterServiceResponse) GetBody() *OpenTransitRouterServiceResponseBody {
	return s.Body
}

func (s *OpenTransitRouterServiceResponse) SetHeaders(v map[string]*string) *OpenTransitRouterServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenTransitRouterServiceResponse) SetStatusCode(v int32) *OpenTransitRouterServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenTransitRouterServiceResponse) SetBody(v *OpenTransitRouterServiceResponseBody) *OpenTransitRouterServiceResponse {
	s.Body = v
	return s
}

func (s *OpenTransitRouterServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
