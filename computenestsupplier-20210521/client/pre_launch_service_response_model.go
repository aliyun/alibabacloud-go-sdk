// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreLaunchServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PreLaunchServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PreLaunchServiceResponse
	GetStatusCode() *int32
	SetBody(v *PreLaunchServiceResponseBody) *PreLaunchServiceResponse
	GetBody() *PreLaunchServiceResponseBody
}

type PreLaunchServiceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PreLaunchServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PreLaunchServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s PreLaunchServiceResponse) GoString() string {
	return s.String()
}

func (s *PreLaunchServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PreLaunchServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PreLaunchServiceResponse) GetBody() *PreLaunchServiceResponseBody {
	return s.Body
}

func (s *PreLaunchServiceResponse) SetHeaders(v map[string]*string) *PreLaunchServiceResponse {
	s.Headers = v
	return s
}

func (s *PreLaunchServiceResponse) SetStatusCode(v int32) *PreLaunchServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *PreLaunchServiceResponse) SetBody(v *PreLaunchServiceResponseBody) *PreLaunchServiceResponse {
	s.Body = v
	return s
}

func (s *PreLaunchServiceResponse) Validate() error {
	return dara.Validate(s)
}
