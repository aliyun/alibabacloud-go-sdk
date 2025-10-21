// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScheduledPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScheduledPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScheduledPlanResponseBody) *DeleteScheduledPlanResponse
	GetBody() *DeleteScheduledPlanResponseBody
}

type DeleteScheduledPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScheduledPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScheduledPlanResponse) GetBody() *DeleteScheduledPlanResponseBody {
	return s.Body
}

func (s *DeleteScheduledPlanResponse) SetHeaders(v map[string]*string) *DeleteScheduledPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledPlanResponse) SetStatusCode(v int32) *DeleteScheduledPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledPlanResponse) SetBody(v *DeleteScheduledPlanResponseBody) *DeleteScheduledPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteScheduledPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
