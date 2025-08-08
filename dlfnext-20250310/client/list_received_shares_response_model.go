// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListReceivedSharesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListReceivedSharesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListReceivedSharesResponse
	GetStatusCode() *int32
	SetBody(v *ListReceivedSharesResponseBody) *ListReceivedSharesResponse
	GetBody() *ListReceivedSharesResponseBody
}

type ListReceivedSharesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReceivedSharesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReceivedSharesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListReceivedSharesResponse) GoString() string {
	return s.String()
}

func (s *ListReceivedSharesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListReceivedSharesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListReceivedSharesResponse) GetBody() *ListReceivedSharesResponseBody {
	return s.Body
}

func (s *ListReceivedSharesResponse) SetHeaders(v map[string]*string) *ListReceivedSharesResponse {
	s.Headers = v
	return s
}

func (s *ListReceivedSharesResponse) SetStatusCode(v int32) *ListReceivedSharesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReceivedSharesResponse) SetBody(v *ListReceivedSharesResponseBody) *ListReceivedSharesResponse {
	s.Body = v
	return s
}

func (s *ListReceivedSharesResponse) Validate() error {
	return dara.Validate(s)
}
