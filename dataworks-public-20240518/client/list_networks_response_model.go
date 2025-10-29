// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNetworksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetworksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetworksResponse
	GetStatusCode() *int32
	SetBody(v *ListNetworksResponseBody) *ListNetworksResponse
	GetBody() *ListNetworksResponseBody
}

type ListNetworksResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetworksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetworksResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetworksResponse) GoString() string {
	return s.String()
}

func (s *ListNetworksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetworksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetworksResponse) GetBody() *ListNetworksResponseBody {
	return s.Body
}

func (s *ListNetworksResponse) SetHeaders(v map[string]*string) *ListNetworksResponse {
	s.Headers = v
	return s
}

func (s *ListNetworksResponse) SetStatusCode(v int32) *ListNetworksResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetworksResponse) SetBody(v *ListNetworksResponseBody) *ListNetworksResponse {
	s.Body = v
	return s
}

func (s *ListNetworksResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
