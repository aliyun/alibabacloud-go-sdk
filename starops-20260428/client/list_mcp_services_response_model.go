// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcpServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcpServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListMcpServicesResponseBody) *ListMcpServicesResponse
	GetBody() *ListMcpServicesResponseBody
}

type ListMcpServicesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcpServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcpServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServicesResponse) GoString() string {
	return s.String()
}

func (s *ListMcpServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcpServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcpServicesResponse) GetBody() *ListMcpServicesResponseBody {
	return s.Body
}

func (s *ListMcpServicesResponse) SetHeaders(v map[string]*string) *ListMcpServicesResponse {
	s.Headers = v
	return s
}

func (s *ListMcpServicesResponse) SetStatusCode(v int32) *ListMcpServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcpServicesResponse) SetBody(v *ListMcpServicesResponseBody) *ListMcpServicesResponse {
	s.Body = v
	return s
}

func (s *ListMcpServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
