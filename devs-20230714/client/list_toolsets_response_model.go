// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListToolsetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListToolsetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListToolsetsResponse
	GetStatusCode() *int32
	SetBody(v *ListToolsetsResponseBody) *ListToolsetsResponse
	GetBody() *ListToolsetsResponseBody
}

type ListToolsetsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListToolsetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListToolsetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListToolsetsResponse) GoString() string {
	return s.String()
}

func (s *ListToolsetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListToolsetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListToolsetsResponse) GetBody() *ListToolsetsResponseBody {
	return s.Body
}

func (s *ListToolsetsResponse) SetHeaders(v map[string]*string) *ListToolsetsResponse {
	s.Headers = v
	return s
}

func (s *ListToolsetsResponse) SetStatusCode(v int32) *ListToolsetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListToolsetsResponse) SetBody(v *ListToolsetsResponseBody) *ListToolsetsResponse {
	s.Body = v
	return s
}

func (s *ListToolsetsResponse) Validate() error {
	return dara.Validate(s)
}
