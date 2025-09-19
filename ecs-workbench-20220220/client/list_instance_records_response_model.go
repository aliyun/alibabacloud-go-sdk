// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceRecordsResponseBody) *ListInstanceRecordsResponse
	GetBody() *ListInstanceRecordsResponseBody
}

type ListInstanceRecordsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceRecordsResponse) GetBody() *ListInstanceRecordsResponseBody {
	return s.Body
}

func (s *ListInstanceRecordsResponse) SetHeaders(v map[string]*string) *ListInstanceRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceRecordsResponse) SetStatusCode(v int32) *ListInstanceRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceRecordsResponse) SetBody(v *ListInstanceRecordsResponseBody) *ListInstanceRecordsResponse {
	s.Body = v
	return s
}

func (s *ListInstanceRecordsResponse) Validate() error {
	return dara.Validate(s)
}
