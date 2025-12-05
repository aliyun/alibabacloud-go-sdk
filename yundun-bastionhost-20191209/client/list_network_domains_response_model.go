// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworkDomainsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworkDomainsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworkDomainsResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworkDomainsResponseBody) *ListNetworkDomainsResponse
	GetBody() *ListNetworkDomainsResponseBody
}

type ListNetworkDomainsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworkDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworkDomainsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworkDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListNetworkDomainsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworkDomainsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworkDomainsResponse) GetBody() *ListNetworkDomainsResponseBody {
	return s.Body
}

func (s *ListNetworkDomainsResponse) SetHeaders(v map[string]*string) *ListNetworkDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListNetworkDomainsResponse) SetStatusCode(v int32) *ListNetworkDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworkDomainsResponse) SetBody(v *ListNetworkDomainsResponseBody) *ListNetworkDomainsResponse {
	s.Body = v
	return s
}

func (s *ListNetworkDomainsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
