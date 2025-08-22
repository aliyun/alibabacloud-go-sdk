// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppNamesResponse
	GetStatusCode() *int32
	SetBody(v *ListAppNamesResponseBody) *ListAppNamesResponse
	GetBody() *ListAppNamesResponseBody
}

type ListAppNamesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppNamesResponse) GoString() string {
	return s.String()
}

func (s *ListAppNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppNamesResponse) GetBody() *ListAppNamesResponseBody {
	return s.Body
}

func (s *ListAppNamesResponse) SetHeaders(v map[string]*string) *ListAppNamesResponse {
	s.Headers = v
	return s
}

func (s *ListAppNamesResponse) SetStatusCode(v int32) *ListAppNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppNamesResponse) SetBody(v *ListAppNamesResponseBody) *ListAppNamesResponse {
	s.Body = v
	return s
}

func (s *ListAppNamesResponse) Validate() error {
	return dara.Validate(s)
}
