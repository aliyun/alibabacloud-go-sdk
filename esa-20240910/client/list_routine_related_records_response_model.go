// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineRelatedRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoutineRelatedRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoutineRelatedRecordsResponse
	GetStatusCode() *int32
	SetBody(v *ListRoutineRelatedRecordsResponseBody) *ListRoutineRelatedRecordsResponse
	GetBody() *ListRoutineRelatedRecordsResponseBody
}

type ListRoutineRelatedRecordsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoutineRelatedRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoutineRelatedRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineRelatedRecordsResponse) GoString() string {
	return s.String()
}

func (s *ListRoutineRelatedRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoutineRelatedRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoutineRelatedRecordsResponse) GetBody() *ListRoutineRelatedRecordsResponseBody {
	return s.Body
}

func (s *ListRoutineRelatedRecordsResponse) SetHeaders(v map[string]*string) *ListRoutineRelatedRecordsResponse {
	s.Headers = v
	return s
}

func (s *ListRoutineRelatedRecordsResponse) SetStatusCode(v int32) *ListRoutineRelatedRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoutineRelatedRecordsResponse) SetBody(v *ListRoutineRelatedRecordsResponseBody) *ListRoutineRelatedRecordsResponse {
	s.Body = v
	return s
}

func (s *ListRoutineRelatedRecordsResponse) Validate() error {
	return dara.Validate(s)
}
