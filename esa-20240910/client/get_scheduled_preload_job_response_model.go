// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledPreloadJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetScheduledPreloadJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetScheduledPreloadJobResponse
	GetStatusCode() *int32
	SetBody(v *GetScheduledPreloadJobResponseBody) *GetScheduledPreloadJobResponse
	GetBody() *GetScheduledPreloadJobResponseBody
}

type GetScheduledPreloadJobResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetScheduledPreloadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetScheduledPreloadJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledPreloadJobResponse) GoString() string {
	return s.String()
}

func (s *GetScheduledPreloadJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetScheduledPreloadJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetScheduledPreloadJobResponse) GetBody() *GetScheduledPreloadJobResponseBody {
	return s.Body
}

func (s *GetScheduledPreloadJobResponse) SetHeaders(v map[string]*string) *GetScheduledPreloadJobResponse {
	s.Headers = v
	return s
}

func (s *GetScheduledPreloadJobResponse) SetStatusCode(v int32) *GetScheduledPreloadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetScheduledPreloadJobResponse) SetBody(v *GetScheduledPreloadJobResponseBody) *GetScheduledPreloadJobResponse {
	s.Body = v
	return s
}

func (s *GetScheduledPreloadJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
