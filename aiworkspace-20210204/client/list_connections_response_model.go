// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConnectionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConnectionsResponse
	GetStatusCode() *int32
	SetBody(v *ListConnectionsResponseBody) *ListConnectionsResponse
	GetBody() *ListConnectionsResponseBody
}

type ListConnectionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConnectionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConnectionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConnectionsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConnectionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConnectionsResponse) GetBody() *ListConnectionsResponseBody {
	return s.Body
}

func (s *ListConnectionsResponse) SetHeaders(v map[string]*string) *ListConnectionsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectionsResponse) SetStatusCode(v int32) *ListConnectionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectionsResponse) SetBody(v *ListConnectionsResponseBody) *ListConnectionsResponse {
	s.Body = v
	return s
}

func (s *ListConnectionsResponse) Validate() error {
	return dara.Validate(s)
}
