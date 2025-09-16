// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIndexRecoverRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIndexRecoverRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIndexRecoverRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListIndexRecoverRecordsResponseBody) *ListIndexRecoverRecordsResponse
	GetBody() *ListIndexRecoverRecordsResponseBody
}

type ListIndexRecoverRecordsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIndexRecoverRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIndexRecoverRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIndexRecoverRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListIndexRecoverRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIndexRecoverRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIndexRecoverRecordsResponse) GetBody() *ListIndexRecoverRecordsResponseBody {
	return s.Body
}

func (s *ListIndexRecoverRecordsResponse) SetHeaders(v map[string]*string) *ListIndexRecoverRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListIndexRecoverRecordsResponse) SetStatusCode(v int32) *ListIndexRecoverRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIndexRecoverRecordsResponse) SetBody(v *ListIndexRecoverRecordsResponseBody) *ListIndexRecoverRecordsResponse {
	s.Body = v
	return s
}

func (s *ListIndexRecoverRecordsResponse) Validate() error {
	return dara.Validate(s)
}
