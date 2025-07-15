// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoNotCallNumbersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoNotCallNumbersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoNotCallNumbersResponse
	GetStatusCode() *int32
	SetBody(v *ListDoNotCallNumbersResponseBody) *ListDoNotCallNumbersResponse
	GetBody() *ListDoNotCallNumbersResponseBody
}

type ListDoNotCallNumbersResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoNotCallNumbersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoNotCallNumbersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoNotCallNumbersResponse) GoString() string {
	return s.String()
}

func (s *ListDoNotCallNumbersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoNotCallNumbersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoNotCallNumbersResponse) GetBody() *ListDoNotCallNumbersResponseBody {
	return s.Body
}

func (s *ListDoNotCallNumbersResponse) SetHeaders(v map[string]*string) *ListDoNotCallNumbersResponse {
	s.Headers = v
	return s
}

func (s *ListDoNotCallNumbersResponse) SetStatusCode(v int32) *ListDoNotCallNumbersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoNotCallNumbersResponse) SetBody(v *ListDoNotCallNumbersResponseBody) *ListDoNotCallNumbersResponse {
	s.Body = v
	return s
}

func (s *ListDoNotCallNumbersResponse) Validate() error {
	return dara.Validate(s)
}
