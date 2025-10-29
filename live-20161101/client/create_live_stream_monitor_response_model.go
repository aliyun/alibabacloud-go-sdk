// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLiveStreamMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLiveStreamMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLiveStreamMonitorResponse
	GetStatusCode() *int32
	SetBody(v *CreateLiveStreamMonitorResponseBody) *CreateLiveStreamMonitorResponse
	GetBody() *CreateLiveStreamMonitorResponseBody
}

type CreateLiveStreamMonitorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLiveStreamMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLiveStreamMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLiveStreamMonitorResponse) GoString() string {
	return s.String()
}

func (s *CreateLiveStreamMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLiveStreamMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLiveStreamMonitorResponse) GetBody() *CreateLiveStreamMonitorResponseBody {
	return s.Body
}

func (s *CreateLiveStreamMonitorResponse) SetHeaders(v map[string]*string) *CreateLiveStreamMonitorResponse {
	s.Headers = v
	return s
}

func (s *CreateLiveStreamMonitorResponse) SetStatusCode(v int32) *CreateLiveStreamMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLiveStreamMonitorResponse) SetBody(v *CreateLiveStreamMonitorResponseBody) *CreateLiveStreamMonitorResponse {
	s.Body = v
	return s
}

func (s *CreateLiveStreamMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
