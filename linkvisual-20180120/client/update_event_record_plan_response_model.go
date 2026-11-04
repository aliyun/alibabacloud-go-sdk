// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEventRecordPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEventRecordPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEventRecordPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEventRecordPlanResponseBody) *UpdateEventRecordPlanResponse
	GetBody() *UpdateEventRecordPlanResponseBody
}

type UpdateEventRecordPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEventRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEventRecordPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEventRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateEventRecordPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEventRecordPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEventRecordPlanResponse) GetBody() *UpdateEventRecordPlanResponseBody {
	return s.Body
}

func (s *UpdateEventRecordPlanResponse) SetHeaders(v map[string]*string) *UpdateEventRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateEventRecordPlanResponse) SetStatusCode(v int32) *UpdateEventRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEventRecordPlanResponse) SetBody(v *UpdateEventRecordPlanResponseBody) *UpdateEventRecordPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateEventRecordPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
