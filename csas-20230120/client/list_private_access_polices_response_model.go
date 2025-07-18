// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrivateAccessPolicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPrivateAccessPolicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPrivateAccessPolicesResponse
	GetStatusCode() *int32
	SetBody(v *ListPrivateAccessPolicesResponseBody) *ListPrivateAccessPolicesResponse
	GetBody() *ListPrivateAccessPolicesResponseBody
}

type ListPrivateAccessPolicesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPrivateAccessPolicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPrivateAccessPolicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPrivateAccessPolicesResponse) GoString() string {
	return s.String()
}

func (s *ListPrivateAccessPolicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPrivateAccessPolicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPrivateAccessPolicesResponse) GetBody() *ListPrivateAccessPolicesResponseBody {
	return s.Body
}

func (s *ListPrivateAccessPolicesResponse) SetHeaders(v map[string]*string) *ListPrivateAccessPolicesResponse {
	s.Headers = v
	return s
}

func (s *ListPrivateAccessPolicesResponse) SetStatusCode(v int32) *ListPrivateAccessPolicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPrivateAccessPolicesResponse) SetBody(v *ListPrivateAccessPolicesResponseBody) *ListPrivateAccessPolicesResponse {
	s.Body = v
	return s
}

func (s *ListPrivateAccessPolicesResponse) Validate() error {
	return dara.Validate(s)
}
