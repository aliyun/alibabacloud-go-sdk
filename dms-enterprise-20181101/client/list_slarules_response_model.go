// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSLARulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSLARulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSLARulesResponse
	GetStatusCode() *int32
	SetBody(v *ListSLARulesResponseBody) *ListSLARulesResponse
	GetBody() *ListSLARulesResponseBody
}

type ListSLARulesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSLARulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSLARulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSLARulesResponse) GoString() string {
	return s.String()
}

func (s *ListSLARulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSLARulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSLARulesResponse) GetBody() *ListSLARulesResponseBody {
	return s.Body
}

func (s *ListSLARulesResponse) SetHeaders(v map[string]*string) *ListSLARulesResponse {
	s.Headers = v
	return s
}

func (s *ListSLARulesResponse) SetStatusCode(v int32) *ListSLARulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSLARulesResponse) SetBody(v *ListSLARulesResponseBody) *ListSLARulesResponse {
	s.Body = v
	return s
}

func (s *ListSLARulesResponse) Validate() error {
	return dara.Validate(s)
}
