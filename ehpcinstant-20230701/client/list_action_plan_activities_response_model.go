// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListActionPlanActivitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListActionPlanActivitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListActionPlanActivitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListActionPlanActivitiesResponseBody) *ListActionPlanActivitiesResponse
	GetBody() *ListActionPlanActivitiesResponseBody
}

type ListActionPlanActivitiesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListActionPlanActivitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListActionPlanActivitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListActionPlanActivitiesResponse) GoString() string {
	return s.String()
}

func (s *ListActionPlanActivitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListActionPlanActivitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListActionPlanActivitiesResponse) GetBody() *ListActionPlanActivitiesResponseBody {
	return s.Body
}

func (s *ListActionPlanActivitiesResponse) SetHeaders(v map[string]*string) *ListActionPlanActivitiesResponse {
	s.Headers = v
	return s
}

func (s *ListActionPlanActivitiesResponse) SetStatusCode(v int32) *ListActionPlanActivitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListActionPlanActivitiesResponse) SetBody(v *ListActionPlanActivitiesResponseBody) *ListActionPlanActivitiesResponse {
	s.Body = v
	return s
}

func (s *ListActionPlanActivitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
