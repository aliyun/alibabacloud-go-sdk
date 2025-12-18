// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceComplianceTimelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceComplianceTimelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceComplianceTimelineResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceComplianceTimelineResponseBody) *GetResourceComplianceTimelineResponse
	GetBody() *GetResourceComplianceTimelineResponseBody
}

type GetResourceComplianceTimelineResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceComplianceTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceComplianceTimelineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceComplianceTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetResourceComplianceTimelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceComplianceTimelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceComplianceTimelineResponse) GetBody() *GetResourceComplianceTimelineResponseBody {
	return s.Body
}

func (s *GetResourceComplianceTimelineResponse) SetHeaders(v map[string]*string) *GetResourceComplianceTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetResourceComplianceTimelineResponse) SetStatusCode(v int32) *GetResourceComplianceTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceComplianceTimelineResponse) SetBody(v *GetResourceComplianceTimelineResponseBody) *GetResourceComplianceTimelineResponse {
	s.Body = v
	return s
}

func (s *GetResourceComplianceTimelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
