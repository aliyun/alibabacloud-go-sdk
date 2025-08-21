// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListActionRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListActionRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListActionRecordsResponseBody) *ListActionRecordsResponse
	GetBody() *ListActionRecordsResponseBody
}

type ListActionRecordsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListActionRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListActionRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListActionRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListActionRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListActionRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListActionRecordsResponse) GetBody() *ListActionRecordsResponseBody {
	return s.Body
}

func (s *ListActionRecordsResponse) SetHeaders(v map[string]*string) *ListActionRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListActionRecordsResponse) SetStatusCode(v int32) *ListActionRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActionRecordsResponse) SetBody(v *ListActionRecordsResponseBody) *ListActionRecordsResponse {
	s.Body = v
	return s
}

func (s *ListActionRecordsResponse) Validate() error {
	return dara.Validate(s)
}
