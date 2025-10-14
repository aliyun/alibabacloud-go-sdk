// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeRoutineRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListEdgeRoutineRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListEdgeRoutineRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListEdgeRoutineRecordsResponseBody) *ListEdgeRoutineRecordsResponse
	GetBody() *ListEdgeRoutineRecordsResponseBody
}

type ListEdgeRoutineRecordsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListEdgeRoutineRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListEdgeRoutineRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeRoutineRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListEdgeRoutineRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListEdgeRoutineRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListEdgeRoutineRecordsResponse) GetBody() *ListEdgeRoutineRecordsResponseBody {
	return s.Body
}

func (s *ListEdgeRoutineRecordsResponse) SetHeaders(v map[string]*string) *ListEdgeRoutineRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListEdgeRoutineRecordsResponse) SetStatusCode(v int32) *ListEdgeRoutineRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListEdgeRoutineRecordsResponse) SetBody(v *ListEdgeRoutineRecordsResponseBody) *ListEdgeRoutineRecordsResponse {
	s.Body = v
	return s
}

func (s *ListEdgeRoutineRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
