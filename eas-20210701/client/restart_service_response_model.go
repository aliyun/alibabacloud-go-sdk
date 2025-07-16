// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartServiceResponse
	GetStatusCode() *int32
	SetBody(v *RestartServiceResponseBody) *RestartServiceResponse
	GetBody() *RestartServiceResponseBody
}

type RestartServiceResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartServiceResponse) GoString() string {
	return s.String()
}

func (s *RestartServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartServiceResponse) GetBody() *RestartServiceResponseBody {
	return s.Body
}

func (s *RestartServiceResponse) SetHeaders(v map[string]*string) *RestartServiceResponse {
	s.Headers = v
	return s
}

func (s *RestartServiceResponse) SetStatusCode(v int32) *RestartServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartServiceResponse) SetBody(v *RestartServiceResponseBody) *RestartServiceResponse {
	s.Body = v
	return s
}

func (s *RestartServiceResponse) Validate() error {
	return dara.Validate(s)
}
