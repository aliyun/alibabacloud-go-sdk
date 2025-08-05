// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTargetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTargetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTargetsResponse
	GetStatusCode() *int32
	SetBody(v *ListTargetsResponseBody) *ListTargetsResponse
	GetBody() *ListTargetsResponseBody
}

type ListTargetsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTargetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTargetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTargetsResponse) GoString() string {
	return s.String()
}

func (s *ListTargetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTargetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTargetsResponse) GetBody() *ListTargetsResponseBody {
	return s.Body
}

func (s *ListTargetsResponse) SetHeaders(v map[string]*string) *ListTargetsResponse {
	s.Headers = v
	return s
}

func (s *ListTargetsResponse) SetStatusCode(v int32) *ListTargetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTargetsResponse) SetBody(v *ListTargetsResponseBody) *ListTargetsResponse {
	s.Body = v
	return s
}

func (s *ListTargetsResponse) Validate() error {
	return dara.Validate(s)
}
