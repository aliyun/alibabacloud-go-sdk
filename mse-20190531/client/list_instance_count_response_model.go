// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceCountResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceCountResponseBody) *ListInstanceCountResponse
	GetBody() *ListInstanceCountResponseBody
}

type ListInstanceCountResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceCountResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceCountResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceCountResponse) GetBody() *ListInstanceCountResponseBody {
	return s.Body
}

func (s *ListInstanceCountResponse) SetHeaders(v map[string]*string) *ListInstanceCountResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceCountResponse) SetStatusCode(v int32) *ListInstanceCountResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceCountResponse) SetBody(v *ListInstanceCountResponseBody) *ListInstanceCountResponse {
	s.Body = v
	return s
}

func (s *ListInstanceCountResponse) Validate() error {
	return dara.Validate(s)
}
