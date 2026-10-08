// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduleTemplatesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScheduleTemplatesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScheduleTemplatesResponse
	GetStatusCode() *int32
	SetBody(v *ListScheduleTemplatesResponseBody) *ListScheduleTemplatesResponse
	GetBody() *ListScheduleTemplatesResponseBody
}

type ListScheduleTemplatesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduleTemplatesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduleTemplatesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduleTemplatesResponse) GoString() string {
	return s.String()
}

func (s *ListScheduleTemplatesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScheduleTemplatesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScheduleTemplatesResponse) GetBody() *ListScheduleTemplatesResponseBody {
	return s.Body
}

func (s *ListScheduleTemplatesResponse) SetHeaders(v map[string]*string) *ListScheduleTemplatesResponse {
	s.Headers = v
	return s
}

func (s *ListScheduleTemplatesResponse) SetStatusCode(v int32) *ListScheduleTemplatesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduleTemplatesResponse) SetBody(v *ListScheduleTemplatesResponseBody) *ListScheduleTemplatesResponse {
	s.Body = v
	return s
}

func (s *ListScheduleTemplatesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
