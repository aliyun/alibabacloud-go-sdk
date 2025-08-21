// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubResponse
	GetStatusCode() *int32
	SetBody(v *ListSubResponseBody) *ListSubResponse
	GetBody() *ListSubResponseBody
}

type ListSubResponse struct {
	Headers    map[string]*string   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubResponse) GoString() string {
	return s.String()
}

func (s *ListSubResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubResponse) GetBody() *ListSubResponseBody {
	return s.Body
}

func (s *ListSubResponse) SetHeaders(v map[string]*string) *ListSubResponse {
	s.Headers = v
	return s
}

func (s *ListSubResponse) SetStatusCode(v int32) *ListSubResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubResponse) SetBody(v *ListSubResponseBody) *ListSubResponse {
	s.Body = v
	return s
}

func (s *ListSubResponse) Validate() error {
	return dara.Validate(s)
}
