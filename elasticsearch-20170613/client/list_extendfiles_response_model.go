// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExtendfilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExtendfilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExtendfilesResponse
	GetStatusCode() *int32
	SetBody(v *ListExtendfilesResponseBody) *ListExtendfilesResponse
	GetBody() *ListExtendfilesResponseBody
}

type ListExtendfilesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExtendfilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExtendfilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExtendfilesResponse) GoString() string {
	return s.String()
}

func (s *ListExtendfilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExtendfilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExtendfilesResponse) GetBody() *ListExtendfilesResponseBody {
	return s.Body
}

func (s *ListExtendfilesResponse) SetHeaders(v map[string]*string) *ListExtendfilesResponse {
	s.Headers = v
	return s
}

func (s *ListExtendfilesResponse) SetStatusCode(v int32) *ListExtendfilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExtendfilesResponse) SetBody(v *ListExtendfilesResponseBody) *ListExtendfilesResponse {
	s.Body = v
	return s
}

func (s *ListExtendfilesResponse) Validate() error {
	return dara.Validate(s)
}
