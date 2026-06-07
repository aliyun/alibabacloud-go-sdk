// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcpServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcpServersResponse
	GetStatusCode() *int32
	SetBody(v *ListMcpServersResponseBody) *ListMcpServersResponse
	GetBody() *ListMcpServersResponseBody
}

type ListMcpServersResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcpServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcpServersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcpServersResponse) GoString() string {
	return s.String()
}

func (s *ListMcpServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcpServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcpServersResponse) GetBody() *ListMcpServersResponseBody {
	return s.Body
}

func (s *ListMcpServersResponse) SetHeaders(v map[string]*string) *ListMcpServersResponse {
	s.Headers = v
	return s
}

func (s *ListMcpServersResponse) SetStatusCode(v int32) *ListMcpServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcpServersResponse) SetBody(v *ListMcpServersResponseBody) *ListMcpServersResponse {
	s.Body = v
	return s
}

func (s *ListMcpServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
