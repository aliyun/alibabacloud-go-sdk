// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCallDetailRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCallDetailRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCallDetailRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListCallDetailRecordsResponseBody) *ListCallDetailRecordsResponse
	GetBody() *ListCallDetailRecordsResponseBody
}

type ListCallDetailRecordsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCallDetailRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCallDetailRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCallDetailRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListCallDetailRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCallDetailRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCallDetailRecordsResponse) GetBody() *ListCallDetailRecordsResponseBody {
	return s.Body
}

func (s *ListCallDetailRecordsResponse) SetHeaders(v map[string]*string) *ListCallDetailRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListCallDetailRecordsResponse) SetStatusCode(v int32) *ListCallDetailRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCallDetailRecordsResponse) SetBody(v *ListCallDetailRecordsResponseBody) *ListCallDetailRecordsResponse {
	s.Body = v
	return s
}

func (s *ListCallDetailRecordsResponse) Validate() error {
	return dara.Validate(s)
}
