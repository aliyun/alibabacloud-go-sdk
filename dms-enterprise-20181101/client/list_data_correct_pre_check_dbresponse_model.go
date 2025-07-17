// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCorrectPreCheckDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDataCorrectPreCheckDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDataCorrectPreCheckDBResponse
	GetStatusCode() *int32
	SetBody(v *ListDataCorrectPreCheckDBResponseBody) *ListDataCorrectPreCheckDBResponse
	GetBody() *ListDataCorrectPreCheckDBResponseBody
}

type ListDataCorrectPreCheckDBResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDataCorrectPreCheckDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDataCorrectPreCheckDBResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDataCorrectPreCheckDBResponse) GoString() string {
	return s.String()
}

func (s *ListDataCorrectPreCheckDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDataCorrectPreCheckDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDataCorrectPreCheckDBResponse) GetBody() *ListDataCorrectPreCheckDBResponseBody {
	return s.Body
}

func (s *ListDataCorrectPreCheckDBResponse) SetHeaders(v map[string]*string) *ListDataCorrectPreCheckDBResponse {
	s.Headers = v
	return s
}

func (s *ListDataCorrectPreCheckDBResponse) SetStatusCode(v int32) *ListDataCorrectPreCheckDBResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDataCorrectPreCheckDBResponse) SetBody(v *ListDataCorrectPreCheckDBResponseBody) *ListDataCorrectPreCheckDBResponse {
	s.Body = v
	return s
}

func (s *ListDataCorrectPreCheckDBResponse) Validate() error {
	return dara.Validate(s)
}
