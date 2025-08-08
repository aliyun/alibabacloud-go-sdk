// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMcubeWhitelistsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMcubeWhitelistsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMcubeWhitelistsResponse
	GetStatusCode() *int32
	SetBody(v *ListMcubeWhitelistsResponseBody) *ListMcubeWhitelistsResponse
	GetBody() *ListMcubeWhitelistsResponseBody
}

type ListMcubeWhitelistsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMcubeWhitelistsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMcubeWhitelistsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMcubeWhitelistsResponse) GoString() string {
	return s.String()
}

func (s *ListMcubeWhitelistsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMcubeWhitelistsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMcubeWhitelistsResponse) GetBody() *ListMcubeWhitelistsResponseBody {
	return s.Body
}

func (s *ListMcubeWhitelistsResponse) SetHeaders(v map[string]*string) *ListMcubeWhitelistsResponse {
	s.Headers = v
	return s
}

func (s *ListMcubeWhitelistsResponse) SetStatusCode(v int32) *ListMcubeWhitelistsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMcubeWhitelistsResponse) SetBody(v *ListMcubeWhitelistsResponseBody) *ListMcubeWhitelistsResponse {
	s.Body = v
	return s
}

func (s *ListMcubeWhitelistsResponse) Validate() error {
	return dara.Validate(s)
}
