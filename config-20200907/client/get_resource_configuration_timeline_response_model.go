// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceConfigurationTimelineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceConfigurationTimelineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceConfigurationTimelineResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceConfigurationTimelineResponseBody) *GetResourceConfigurationTimelineResponse
	GetBody() *GetResourceConfigurationTimelineResponseBody
}

type GetResourceConfigurationTimelineResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceConfigurationTimelineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceConfigurationTimelineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceConfigurationTimelineResponse) GoString() string {
	return s.String()
}

func (s *GetResourceConfigurationTimelineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceConfigurationTimelineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceConfigurationTimelineResponse) GetBody() *GetResourceConfigurationTimelineResponseBody {
	return s.Body
}

func (s *GetResourceConfigurationTimelineResponse) SetHeaders(v map[string]*string) *GetResourceConfigurationTimelineResponse {
	s.Headers = v
	return s
}

func (s *GetResourceConfigurationTimelineResponse) SetStatusCode(v int32) *GetResourceConfigurationTimelineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceConfigurationTimelineResponse) SetBody(v *GetResourceConfigurationTimelineResponseBody) *GetResourceConfigurationTimelineResponse {
	s.Body = v
	return s
}

func (s *GetResourceConfigurationTimelineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
