// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLaunchServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LaunchServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LaunchServiceResponse
	GetStatusCode() *int32
	SetBody(v *LaunchServiceResponseBody) *LaunchServiceResponse
	GetBody() *LaunchServiceResponseBody
}

type LaunchServiceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LaunchServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LaunchServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s LaunchServiceResponse) GoString() string {
	return s.String()
}

func (s *LaunchServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LaunchServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LaunchServiceResponse) GetBody() *LaunchServiceResponseBody {
	return s.Body
}

func (s *LaunchServiceResponse) SetHeaders(v map[string]*string) *LaunchServiceResponse {
	s.Headers = v
	return s
}

func (s *LaunchServiceResponse) SetStatusCode(v int32) *LaunchServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *LaunchServiceResponse) SetBody(v *LaunchServiceResponseBody) *LaunchServiceResponse {
	s.Body = v
	return s
}

func (s *LaunchServiceResponse) Validate() error {
	return dara.Validate(s)
}
