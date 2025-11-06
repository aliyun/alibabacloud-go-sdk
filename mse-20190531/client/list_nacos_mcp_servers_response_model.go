// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNacosMcpServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNacosMcpServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNacosMcpServersResponse
	GetStatusCode() *int32
	SetBody(v *ListNacosMcpServersResponseBody) *ListNacosMcpServersResponse
	GetBody() *ListNacosMcpServersResponseBody
}

type ListNacosMcpServersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNacosMcpServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNacosMcpServersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNacosMcpServersResponse) GoString() string {
	return s.String()
}

func (s *ListNacosMcpServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNacosMcpServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNacosMcpServersResponse) GetBody() *ListNacosMcpServersResponseBody {
	return s.Body
}

func (s *ListNacosMcpServersResponse) SetHeaders(v map[string]*string) *ListNacosMcpServersResponse {
	s.Headers = v
	return s
}

func (s *ListNacosMcpServersResponse) SetStatusCode(v int32) *ListNacosMcpServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNacosMcpServersResponse) SetBody(v *ListNacosMcpServersResponseBody) *ListNacosMcpServersResponse {
	s.Body = v
	return s
}

func (s *ListNacosMcpServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
