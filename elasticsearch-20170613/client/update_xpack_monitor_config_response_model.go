// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateXpackMonitorConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateXpackMonitorConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateXpackMonitorConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateXpackMonitorConfigResponseBody) *UpdateXpackMonitorConfigResponse
	GetBody() *UpdateXpackMonitorConfigResponseBody
}

type UpdateXpackMonitorConfigResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateXpackMonitorConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateXpackMonitorConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateXpackMonitorConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateXpackMonitorConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateXpackMonitorConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateXpackMonitorConfigResponse) GetBody() *UpdateXpackMonitorConfigResponseBody {
	return s.Body
}

func (s *UpdateXpackMonitorConfigResponse) SetHeaders(v map[string]*string) *UpdateXpackMonitorConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateXpackMonitorConfigResponse) SetStatusCode(v int32) *UpdateXpackMonitorConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateXpackMonitorConfigResponse) SetBody(v *UpdateXpackMonitorConfigResponseBody) *UpdateXpackMonitorConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateXpackMonitorConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
