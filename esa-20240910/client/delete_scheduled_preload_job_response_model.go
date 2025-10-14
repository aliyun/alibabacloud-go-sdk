// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScheduledPreloadJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteScheduledPreloadJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteScheduledPreloadJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteScheduledPreloadJobResponseBody) *DeleteScheduledPreloadJobResponse
	GetBody() *DeleteScheduledPreloadJobResponseBody
}

type DeleteScheduledPreloadJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteScheduledPreloadJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteScheduledPreloadJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteScheduledPreloadJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteScheduledPreloadJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteScheduledPreloadJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteScheduledPreloadJobResponse) GetBody() *DeleteScheduledPreloadJobResponseBody {
	return s.Body
}

func (s *DeleteScheduledPreloadJobResponse) SetHeaders(v map[string]*string) *DeleteScheduledPreloadJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteScheduledPreloadJobResponse) SetStatusCode(v int32) *DeleteScheduledPreloadJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteScheduledPreloadJobResponse) SetBody(v *DeleteScheduledPreloadJobResponseBody) *DeleteScheduledPreloadJobResponse {
	s.Body = v
	return s
}

func (s *DeleteScheduledPreloadJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
