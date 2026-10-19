// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleWorkItemResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScheduleWorkItemResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScheduleWorkItemResponse
	GetStatusCode() *int32
	SetBody(v *ScheduleWorkItemResponseBody) *ScheduleWorkItemResponse
	GetBody() *ScheduleWorkItemResponseBody
}

type ScheduleWorkItemResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScheduleWorkItemResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScheduleWorkItemResponse) String() string {
	return dara.Prettify(s)
}

func (s ScheduleWorkItemResponse) GoString() string {
	return s.String()
}

func (s *ScheduleWorkItemResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScheduleWorkItemResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScheduleWorkItemResponse) GetBody() *ScheduleWorkItemResponseBody {
	return s.Body
}

func (s *ScheduleWorkItemResponse) SetHeaders(v map[string]*string) *ScheduleWorkItemResponse {
	s.Headers = v
	return s
}

func (s *ScheduleWorkItemResponse) SetStatusCode(v int32) *ScheduleWorkItemResponse {
	s.StatusCode = &v
	return s
}

func (s *ScheduleWorkItemResponse) SetBody(v *ScheduleWorkItemResponseBody) *ScheduleWorkItemResponse {
	s.Body = v
	return s
}

func (s *ScheduleWorkItemResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
