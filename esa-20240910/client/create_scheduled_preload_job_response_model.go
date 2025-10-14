// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScheduledPreloadJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateScheduledPreloadJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateScheduledPreloadJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateScheduledPreloadJobResponseBody) *CreateScheduledPreloadJobResponse
	GetBody() *CreateScheduledPreloadJobResponseBody
}

type CreateScheduledPreloadJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateScheduledPreloadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateScheduledPreloadJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateScheduledPreloadJobResponse) GoString() string {
	return s.String()
}

func (s *CreateScheduledPreloadJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateScheduledPreloadJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateScheduledPreloadJobResponse) GetBody() *CreateScheduledPreloadJobResponseBody {
	return s.Body
}

func (s *CreateScheduledPreloadJobResponse) SetHeaders(v map[string]*string) *CreateScheduledPreloadJobResponse {
	s.Headers = v
	return s
}

func (s *CreateScheduledPreloadJobResponse) SetStatusCode(v int32) *CreateScheduledPreloadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateScheduledPreloadJobResponse) SetBody(v *CreateScheduledPreloadJobResponseBody) *CreateScheduledPreloadJobResponse {
	s.Body = v
	return s
}

func (s *CreateScheduledPreloadJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
