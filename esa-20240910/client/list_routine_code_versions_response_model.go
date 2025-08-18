// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineCodeVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoutineCodeVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoutineCodeVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListRoutineCodeVersionsResponseBody) *ListRoutineCodeVersionsResponse
	GetBody() *ListRoutineCodeVersionsResponseBody
}

type ListRoutineCodeVersionsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoutineCodeVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoutineCodeVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineCodeVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListRoutineCodeVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoutineCodeVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoutineCodeVersionsResponse) GetBody() *ListRoutineCodeVersionsResponseBody {
	return s.Body
}

func (s *ListRoutineCodeVersionsResponse) SetHeaders(v map[string]*string) *ListRoutineCodeVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListRoutineCodeVersionsResponse) SetStatusCode(v int32) *ListRoutineCodeVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoutineCodeVersionsResponse) SetBody(v *ListRoutineCodeVersionsResponseBody) *ListRoutineCodeVersionsResponse {
	s.Body = v
	return s
}

func (s *ListRoutineCodeVersionsResponse) Validate() error {
	return dara.Validate(s)
}
