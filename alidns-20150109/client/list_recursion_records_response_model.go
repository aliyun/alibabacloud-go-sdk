// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecursionRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecursionRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecursionRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListRecursionRecordsResponseBody) *ListRecursionRecordsResponse
	GetBody() *ListRecursionRecordsResponseBody
}

type ListRecursionRecordsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecursionRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecursionRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecursionRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListRecursionRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecursionRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecursionRecordsResponse) GetBody() *ListRecursionRecordsResponseBody {
	return s.Body
}

func (s *ListRecursionRecordsResponse) SetHeaders(v map[string]*string) *ListRecursionRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListRecursionRecordsResponse) SetStatusCode(v int32) *ListRecursionRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecursionRecordsResponse) SetBody(v *ListRecursionRecordsResponseBody) *ListRecursionRecordsResponse {
	s.Body = v
	return s
}

func (s *ListRecursionRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
