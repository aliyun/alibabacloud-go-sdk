// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDhcpOptionsSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDhcpOptionsSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDhcpOptionsSetsResponse
	GetStatusCode() *int32
	SetBody(v *ListDhcpOptionsSetsResponseBody) *ListDhcpOptionsSetsResponse
	GetBody() *ListDhcpOptionsSetsResponseBody
}

type ListDhcpOptionsSetsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDhcpOptionsSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDhcpOptionsSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDhcpOptionsSetsResponse) GoString() string {
	return s.String()
}

func (s *ListDhcpOptionsSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDhcpOptionsSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDhcpOptionsSetsResponse) GetBody() *ListDhcpOptionsSetsResponseBody {
	return s.Body
}

func (s *ListDhcpOptionsSetsResponse) SetHeaders(v map[string]*string) *ListDhcpOptionsSetsResponse {
	s.Headers = v
	return s
}

func (s *ListDhcpOptionsSetsResponse) SetStatusCode(v int32) *ListDhcpOptionsSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDhcpOptionsSetsResponse) SetBody(v *ListDhcpOptionsSetsResponseBody) *ListDhcpOptionsSetsResponse {
	s.Body = v
	return s
}

func (s *ListDhcpOptionsSetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
