// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleTimesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScheduleTimesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScheduleTimesResponse
	GetStatusCode() *int32
	SetBody(v *ListScheduleTimesResponseBody) *ListScheduleTimesResponse
	GetBody() *ListScheduleTimesResponseBody
}

type ListScheduleTimesResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduleTimesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduleTimesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTimesResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleTimesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScheduleTimesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScheduleTimesResponse) GetBody() *ListScheduleTimesResponseBody {
	return s.Body
}

func (s *ListScheduleTimesResponse) SetHeaders(v map[string]*string) *ListScheduleTimesResponse {
	s.Headers = v
	return s
}

func (s *ListScheduleTimesResponse) SetStatusCode(v int32) *ListScheduleTimesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduleTimesResponse) SetBody(v *ListScheduleTimesResponseBody) *ListScheduleTimesResponse {
	s.Body = v
	return s
}

func (s *ListScheduleTimesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
