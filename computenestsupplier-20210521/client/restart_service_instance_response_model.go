// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *RestartServiceInstanceResponseBody) *RestartServiceInstanceResponse
	GetBody() *RestartServiceInstanceResponseBody
}

type RestartServiceInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *RestartServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartServiceInstanceResponse) GetBody() *RestartServiceInstanceResponseBody {
	return s.Body
}

func (s *RestartServiceInstanceResponse) SetHeaders(v map[string]*string) *RestartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *RestartServiceInstanceResponse) SetStatusCode(v int32) *RestartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartServiceInstanceResponse) SetBody(v *RestartServiceInstanceResponseBody) *RestartServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *RestartServiceInstanceResponse) Validate() error {
	return dara.Validate(s)
}
