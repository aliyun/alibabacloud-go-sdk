// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiMcpServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiMcpServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiMcpServersResponse
	GetStatusCode() *int32
	SetBody(v *ListApiMcpServersResponseBody) *ListApiMcpServersResponse
	GetBody() *ListApiMcpServersResponseBody
}

type ListApiMcpServersResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiMcpServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiMcpServersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiMcpServersResponse) GoString() string {
	return s.String()
}

func (s *ListApiMcpServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiMcpServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiMcpServersResponse) GetBody() *ListApiMcpServersResponseBody {
	return s.Body
}

func (s *ListApiMcpServersResponse) SetHeaders(v map[string]*string) *ListApiMcpServersResponse {
	s.Headers = v
	return s
}

func (s *ListApiMcpServersResponse) SetStatusCode(v int32) *ListApiMcpServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiMcpServersResponse) SetBody(v *ListApiMcpServersResponseBody) *ListApiMcpServersResponse {
	s.Body = v
	return s
}

func (s *ListApiMcpServersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
