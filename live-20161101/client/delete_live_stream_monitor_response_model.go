// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamMonitorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveStreamMonitorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveStreamMonitorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveStreamMonitorResponseBody) *DeleteLiveStreamMonitorResponse
	GetBody() *DeleteLiveStreamMonitorResponseBody
}

type DeleteLiveStreamMonitorResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveStreamMonitorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveStreamMonitorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamMonitorResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamMonitorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveStreamMonitorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveStreamMonitorResponse) GetBody() *DeleteLiveStreamMonitorResponseBody {
	return s.Body
}

func (s *DeleteLiveStreamMonitorResponse) SetHeaders(v map[string]*string) *DeleteLiveStreamMonitorResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveStreamMonitorResponse) SetStatusCode(v int32) *DeleteLiveStreamMonitorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveStreamMonitorResponse) SetBody(v *DeleteLiveStreamMonitorResponseBody) *DeleteLiveStreamMonitorResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveStreamMonitorResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
