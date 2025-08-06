// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListAppInfoResponseBody) *ListAppInfoResponse
	GetBody() *ListAppInfoResponseBody
}

type ListAppInfoResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppInfoResponse) GoString() string {
	return s.String()
}

func (s *ListAppInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppInfoResponse) GetBody() *ListAppInfoResponseBody {
	return s.Body
}

func (s *ListAppInfoResponse) SetHeaders(v map[string]*string) *ListAppInfoResponse {
	s.Headers = v
	return s
}

func (s *ListAppInfoResponse) SetStatusCode(v int32) *ListAppInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppInfoResponse) SetBody(v *ListAppInfoResponseBody) *ListAppInfoResponse {
	s.Body = v
	return s
}

func (s *ListAppInfoResponse) Validate() error {
	return dara.Validate(s)
}
