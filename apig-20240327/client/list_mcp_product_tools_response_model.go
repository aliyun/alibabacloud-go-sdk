// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcpProductToolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcpProductToolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcpProductToolsResponse
	GetStatusCode() *int32
	SetBody(v *ListMcpProductToolsResponseBody) *ListMcpProductToolsResponse
	GetBody() *ListMcpProductToolsResponseBody
}

type ListMcpProductToolsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcpProductToolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcpProductToolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcpProductToolsResponse) GoString() string {
	return s.String()
}

func (s *ListMcpProductToolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcpProductToolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcpProductToolsResponse) GetBody() *ListMcpProductToolsResponseBody {
	return s.Body
}

func (s *ListMcpProductToolsResponse) SetHeaders(v map[string]*string) *ListMcpProductToolsResponse {
	s.Headers = v
	return s
}

func (s *ListMcpProductToolsResponse) SetStatusCode(v int32) *ListMcpProductToolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcpProductToolsResponse) SetBody(v *ListMcpProductToolsResponseBody) *ListMcpProductToolsResponse {
	s.Body = v
	return s
}

func (s *ListMcpProductToolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
