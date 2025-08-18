// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecordsResponseBody) *ListRecordsResponse
	GetBody() *ListRecordsResponseBody
}

type ListRecordsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecordsResponse) GetBody() *ListRecordsResponseBody {
	return s.Body
}

func (s *ListRecordsResponse) SetHeaders(v map[string]*string) *ListRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListRecordsResponse) SetStatusCode(v int32) *ListRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecordsResponse) SetBody(v *ListRecordsResponseBody) *ListRecordsResponse {
	s.Body = v
	return s
}

func (s *ListRecordsResponse) Validate() error {
	return dara.Validate(s)
}
