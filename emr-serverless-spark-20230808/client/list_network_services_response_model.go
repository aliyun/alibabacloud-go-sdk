// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkServicesResponseBody) *ListNetworkServicesResponse
	GetBody() *ListNetworkServicesResponseBody
}

type ListNetworkServicesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkServicesResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkServicesResponse) GetBody() *ListNetworkServicesResponseBody {
	return s.Body
}

func (s *ListNetworkServicesResponse) SetHeaders(v map[string]*string) *ListNetworkServicesResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkServicesResponse) SetStatusCode(v int32) *ListNetworkServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkServicesResponse) SetBody(v *ListNetworkServicesResponseBody) *ListNetworkServicesResponse {
	s.Body = v
	return s
}

func (s *ListNetworkServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
