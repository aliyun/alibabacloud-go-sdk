// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScheduledPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListScheduledPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListScheduledPlanResponse
	GetStatusCode() *int32
	SetBody(v *ListScheduledPlanResponseBody) *ListScheduledPlanResponse
	GetBody() *ListScheduledPlanResponseBody
}

type ListScheduledPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListScheduledPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ListScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *ListScheduledPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListScheduledPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListScheduledPlanResponse) GetBody() *ListScheduledPlanResponseBody {
	return s.Body
}

func (s *ListScheduledPlanResponse) SetHeaders(v map[string]*string) *ListScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *ListScheduledPlanResponse) SetStatusCode(v int32) *ListScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ListScheduledPlanResponse) SetBody(v *ListScheduledPlanResponseBody) *ListScheduledPlanResponse {
	s.Body = v
	return s
}

func (s *ListScheduledPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
