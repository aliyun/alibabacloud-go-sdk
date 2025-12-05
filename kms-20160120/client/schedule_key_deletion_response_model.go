// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleKeyDeletionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScheduleKeyDeletionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScheduleKeyDeletionResponse
	GetStatusCode() *int32
	SetBody(v *ScheduleKeyDeletionResponseBody) *ScheduleKeyDeletionResponse
	GetBody() *ScheduleKeyDeletionResponseBody
}

type ScheduleKeyDeletionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScheduleKeyDeletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScheduleKeyDeletionResponse) String() string {
	return dara.Prettify(s)
}

func (s ScheduleKeyDeletionResponse) GoString() string {
	return s.String()
}

func (s *ScheduleKeyDeletionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScheduleKeyDeletionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScheduleKeyDeletionResponse) GetBody() *ScheduleKeyDeletionResponseBody {
	return s.Body
}

func (s *ScheduleKeyDeletionResponse) SetHeaders(v map[string]*string) *ScheduleKeyDeletionResponse {
	s.Headers = v
	return s
}

func (s *ScheduleKeyDeletionResponse) SetStatusCode(v int32) *ScheduleKeyDeletionResponse {
	s.StatusCode = &v
	return s
}

func (s *ScheduleKeyDeletionResponse) SetBody(v *ScheduleKeyDeletionResponseBody) *ScheduleKeyDeletionResponse {
	s.Body = v
	return s
}

func (s *ScheduleKeyDeletionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
