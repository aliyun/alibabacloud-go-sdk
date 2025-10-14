// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRoutineCanaryAreasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRoutineCanaryAreasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRoutineCanaryAreasResponse
	GetStatusCode() *int32
	SetBody(v *ListRoutineCanaryAreasResponseBody) *ListRoutineCanaryAreasResponse
	GetBody() *ListRoutineCanaryAreasResponseBody
}

type ListRoutineCanaryAreasResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRoutineCanaryAreasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRoutineCanaryAreasResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRoutineCanaryAreasResponse) GoString() string {
	return s.String()
}

func (s *ListRoutineCanaryAreasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRoutineCanaryAreasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRoutineCanaryAreasResponse) GetBody() *ListRoutineCanaryAreasResponseBody {
	return s.Body
}

func (s *ListRoutineCanaryAreasResponse) SetHeaders(v map[string]*string) *ListRoutineCanaryAreasResponse {
	s.Headers = v
	return s
}

func (s *ListRoutineCanaryAreasResponse) SetStatusCode(v int32) *ListRoutineCanaryAreasResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRoutineCanaryAreasResponse) SetBody(v *ListRoutineCanaryAreasResponseBody) *ListRoutineCanaryAreasResponse {
	s.Body = v
	return s
}

func (s *ListRoutineCanaryAreasResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
