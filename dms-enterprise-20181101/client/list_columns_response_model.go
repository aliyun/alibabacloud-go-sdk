// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListColumnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListColumnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListColumnsResponse
	GetStatusCode() *int32
	SetBody(v *ListColumnsResponseBody) *ListColumnsResponse
	GetBody() *ListColumnsResponseBody
}

type ListColumnsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListColumnsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListColumnsResponse) GoString() string {
	return s.String()
}

func (s *ListColumnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListColumnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListColumnsResponse) GetBody() *ListColumnsResponseBody {
	return s.Body
}

func (s *ListColumnsResponse) SetHeaders(v map[string]*string) *ListColumnsResponse {
	s.Headers = v
	return s
}

func (s *ListColumnsResponse) SetStatusCode(v int32) *ListColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListColumnsResponse) SetBody(v *ListColumnsResponseBody) *ListColumnsResponse {
	s.Body = v
	return s
}

func (s *ListColumnsResponse) Validate() error {
	return dara.Validate(s)
}
