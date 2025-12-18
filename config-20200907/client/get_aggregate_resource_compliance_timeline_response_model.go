// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceComplianceTimelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateResourceComplianceTimelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateResourceComplianceTimelineResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateResourceComplianceTimelineResponseBody) *GetAggregateResourceComplianceTimelineResponse
	GetBody() *GetAggregateResourceComplianceTimelineResponseBody
}

type GetAggregateResourceComplianceTimelineResponse struct {
	Headers    map[string]*string                                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateResourceComplianceTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateResourceComplianceTimelineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceComplianceTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceComplianceTimelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateResourceComplianceTimelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateResourceComplianceTimelineResponse) GetBody() *GetAggregateResourceComplianceTimelineResponseBody {
	return s.Body
}

func (s *GetAggregateResourceComplianceTimelineResponse) SetHeaders(v map[string]*string) *GetAggregateResourceComplianceTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponse) SetStatusCode(v int32) *GetAggregateResourceComplianceTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponse) SetBody(v *GetAggregateResourceComplianceTimelineResponseBody) *GetAggregateResourceComplianceTimelineResponse {
	s.Body = v
	return s
}

func (s *GetAggregateResourceComplianceTimelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
