// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConnectorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConnectorsResponse
	GetStatusCode() *int32
	SetBody(v *ListConnectorsResponseBody) *ListConnectorsResponse
	GetBody() *ListConnectorsResponseBody
}

type ListConnectorsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConnectorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConnectorsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsResponse) GoString() string {
	return s.String()
}

func (s *ListConnectorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConnectorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConnectorsResponse) GetBody() *ListConnectorsResponseBody {
	return s.Body
}

func (s *ListConnectorsResponse) SetHeaders(v map[string]*string) *ListConnectorsResponse {
	s.Headers = v
	return s
}

func (s *ListConnectorsResponse) SetStatusCode(v int32) *ListConnectorsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConnectorsResponse) SetBody(v *ListConnectorsResponseBody) *ListConnectorsResponse {
	s.Body = v
	return s
}

func (s *ListConnectorsResponse) Validate() error {
	return dara.Validate(s)
}
