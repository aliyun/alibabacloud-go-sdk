// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllNodeResponse
	GetStatusCode() *int32
	SetBody(v *ListAllNodeResponseBody) *ListAllNodeResponse
	GetBody() *ListAllNodeResponseBody
}

type ListAllNodeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllNodeResponse) GoString() string {
	return s.String()
}

func (s *ListAllNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllNodeResponse) GetBody() *ListAllNodeResponseBody {
	return s.Body
}

func (s *ListAllNodeResponse) SetHeaders(v map[string]*string) *ListAllNodeResponse {
	s.Headers = v
	return s
}

func (s *ListAllNodeResponse) SetStatusCode(v int32) *ListAllNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllNodeResponse) SetBody(v *ListAllNodeResponseBody) *ListAllNodeResponse {
	s.Body = v
	return s
}

func (s *ListAllNodeResponse) Validate() error {
	return dara.Validate(s)
}
