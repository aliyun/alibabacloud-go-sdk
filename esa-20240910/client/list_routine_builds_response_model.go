// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineBuildsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoutineBuildsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoutineBuildsResponse
	GetStatusCode() *int32
	SetBody(v *ListRoutineBuildsResponseBody) *ListRoutineBuildsResponse
	GetBody() *ListRoutineBuildsResponseBody
}

type ListRoutineBuildsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoutineBuildsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoutineBuildsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineBuildsResponse) GoString() string {
	return s.String()
}

func (s *ListRoutineBuildsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoutineBuildsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoutineBuildsResponse) GetBody() *ListRoutineBuildsResponseBody {
	return s.Body
}

func (s *ListRoutineBuildsResponse) SetHeaders(v map[string]*string) *ListRoutineBuildsResponse {
	s.Headers = v
	return s
}

func (s *ListRoutineBuildsResponse) SetStatusCode(v int32) *ListRoutineBuildsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoutineBuildsResponse) SetBody(v *ListRoutineBuildsResponseBody) *ListRoutineBuildsResponse {
	s.Body = v
	return s
}

func (s *ListRoutineBuildsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
