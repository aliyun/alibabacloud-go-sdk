// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentCallDetailRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecentCallDetailRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecentCallDetailRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecentCallDetailRecordsResponseBody) *ListRecentCallDetailRecordsResponse
	GetBody() *ListRecentCallDetailRecordsResponseBody
}

type ListRecentCallDetailRecordsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecentCallDetailRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecentCallDetailRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecentCallDetailRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListRecentCallDetailRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecentCallDetailRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecentCallDetailRecordsResponse) GetBody() *ListRecentCallDetailRecordsResponseBody {
	return s.Body
}

func (s *ListRecentCallDetailRecordsResponse) SetHeaders(v map[string]*string) *ListRecentCallDetailRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListRecentCallDetailRecordsResponse) SetStatusCode(v int32) *ListRecentCallDetailRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentCallDetailRecordsResponse) SetBody(v *ListRecentCallDetailRecordsResponseBody) *ListRecentCallDetailRecordsResponse {
	s.Body = v
	return s
}

func (s *ListRecentCallDetailRecordsResponse) Validate() error {
	return dara.Validate(s)
}
