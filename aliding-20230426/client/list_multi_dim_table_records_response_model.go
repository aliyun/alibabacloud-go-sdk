// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMultiDimTableRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMultiDimTableRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMultiDimTableRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListMultiDimTableRecordsResponseBody) *ListMultiDimTableRecordsResponse
	GetBody() *ListMultiDimTableRecordsResponseBody
}

type ListMultiDimTableRecordsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMultiDimTableRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMultiDimTableRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMultiDimTableRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListMultiDimTableRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMultiDimTableRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMultiDimTableRecordsResponse) GetBody() *ListMultiDimTableRecordsResponseBody {
	return s.Body
}

func (s *ListMultiDimTableRecordsResponse) SetHeaders(v map[string]*string) *ListMultiDimTableRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListMultiDimTableRecordsResponse) SetStatusCode(v int32) *ListMultiDimTableRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMultiDimTableRecordsResponse) SetBody(v *ListMultiDimTableRecordsResponseBody) *ListMultiDimTableRecordsResponse {
	s.Body = v
	return s
}

func (s *ListMultiDimTableRecordsResponse) Validate() error {
	return dara.Validate(s)
}
