// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveStreamMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveStreamMonitorResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveStreamMonitorResponseBody) *UpdateLiveStreamMonitorResponse
	GetBody() *UpdateLiveStreamMonitorResponseBody
}

type UpdateLiveStreamMonitorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveStreamMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveStreamMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamMonitorResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveStreamMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveStreamMonitorResponse) GetBody() *UpdateLiveStreamMonitorResponseBody {
	return s.Body
}

func (s *UpdateLiveStreamMonitorResponse) SetHeaders(v map[string]*string) *UpdateLiveStreamMonitorResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveStreamMonitorResponse) SetStatusCode(v int32) *UpdateLiveStreamMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveStreamMonitorResponse) SetBody(v *UpdateLiveStreamMonitorResponseBody) *UpdateLiveStreamMonitorResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveStreamMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
