// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEventRecordPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEventRecordPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEventRecordPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateEventRecordPlanResponseBody) *CreateEventRecordPlanResponse
	GetBody() *CreateEventRecordPlanResponseBody
}

type CreateEventRecordPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEventRecordPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEventRecordPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEventRecordPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateEventRecordPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEventRecordPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEventRecordPlanResponse) GetBody() *CreateEventRecordPlanResponseBody {
	return s.Body
}

func (s *CreateEventRecordPlanResponse) SetHeaders(v map[string]*string) *CreateEventRecordPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateEventRecordPlanResponse) SetStatusCode(v int32) *CreateEventRecordPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEventRecordPlanResponse) SetBody(v *CreateEventRecordPlanResponseBody) *CreateEventRecordPlanResponse {
	s.Body = v
	return s
}

func (s *CreateEventRecordPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
