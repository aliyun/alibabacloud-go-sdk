// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAggregateResourceConfigurationTimelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAggregateResourceConfigurationTimelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAggregateResourceConfigurationTimelineResponse
	GetStatusCode() *int32
	SetBody(v *GetAggregateResourceConfigurationTimelineResponseBody) *GetAggregateResourceConfigurationTimelineResponse
	GetBody() *GetAggregateResourceConfigurationTimelineResponseBody
}

type GetAggregateResourceConfigurationTimelineResponse struct {
	Headers    map[string]*string                                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAggregateResourceConfigurationTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAggregateResourceConfigurationTimelineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAggregateResourceConfigurationTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetAggregateResourceConfigurationTimelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAggregateResourceConfigurationTimelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAggregateResourceConfigurationTimelineResponse) GetBody() *GetAggregateResourceConfigurationTimelineResponseBody {
	return s.Body
}

func (s *GetAggregateResourceConfigurationTimelineResponse) SetHeaders(v map[string]*string) *GetAggregateResourceConfigurationTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponse) SetStatusCode(v int32) *GetAggregateResourceConfigurationTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponse) SetBody(v *GetAggregateResourceConfigurationTimelineResponseBody) *GetAggregateResourceConfigurationTimelineResponse {
	s.Body = v
	return s
}

func (s *GetAggregateResourceConfigurationTimelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
