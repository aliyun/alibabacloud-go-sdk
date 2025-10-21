// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppliedScheduledPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppliedScheduledPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppliedScheduledPlanResponse
	GetStatusCode() *int32
	SetBody(v *GetAppliedScheduledPlanResponseBody) *GetAppliedScheduledPlanResponse
	GetBody() *GetAppliedScheduledPlanResponseBody
}

type GetAppliedScheduledPlanResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppliedScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppliedScheduledPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppliedScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *GetAppliedScheduledPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppliedScheduledPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppliedScheduledPlanResponse) GetBody() *GetAppliedScheduledPlanResponseBody {
	return s.Body
}

func (s *GetAppliedScheduledPlanResponse) SetHeaders(v map[string]*string) *GetAppliedScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *GetAppliedScheduledPlanResponse) SetStatusCode(v int32) *GetAppliedScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppliedScheduledPlanResponse) SetBody(v *GetAppliedScheduledPlanResponseBody) *GetAppliedScheduledPlanResponse {
	s.Body = v
	return s
}

func (s *GetAppliedScheduledPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
