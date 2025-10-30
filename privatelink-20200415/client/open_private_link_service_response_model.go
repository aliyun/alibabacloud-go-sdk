// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenPrivateLinkServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenPrivateLinkServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenPrivateLinkServiceResponse
	GetStatusCode() *int32
	SetBody(v *OpenPrivateLinkServiceResponseBody) *OpenPrivateLinkServiceResponse
	GetBody() *OpenPrivateLinkServiceResponseBody
}

type OpenPrivateLinkServiceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenPrivateLinkServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenPrivateLinkServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenPrivateLinkServiceResponse) GoString() string {
	return s.String()
}

func (s *OpenPrivateLinkServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenPrivateLinkServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenPrivateLinkServiceResponse) GetBody() *OpenPrivateLinkServiceResponseBody {
	return s.Body
}

func (s *OpenPrivateLinkServiceResponse) SetHeaders(v map[string]*string) *OpenPrivateLinkServiceResponse {
	s.Headers = v
	return s
}

func (s *OpenPrivateLinkServiceResponse) SetStatusCode(v int32) *OpenPrivateLinkServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenPrivateLinkServiceResponse) SetBody(v *OpenPrivateLinkServiceResponseBody) *OpenPrivateLinkServiceResponse {
	s.Body = v
	return s
}

func (s *OpenPrivateLinkServiceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
