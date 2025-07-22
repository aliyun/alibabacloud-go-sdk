// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceShutdownTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateInstanceShutdownTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateInstanceShutdownTimerResponse
	GetStatusCode() *int32
	SetBody(v *CreateInstanceShutdownTimerResponseBody) *CreateInstanceShutdownTimerResponse
	GetBody() *CreateInstanceShutdownTimerResponseBody
}

type CreateInstanceShutdownTimerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateInstanceShutdownTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceShutdownTimerResponse) GoString() string {
	return s.String()
}

func (s *CreateInstanceShutdownTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateInstanceShutdownTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateInstanceShutdownTimerResponse) GetBody() *CreateInstanceShutdownTimerResponseBody {
	return s.Body
}

func (s *CreateInstanceShutdownTimerResponse) SetHeaders(v map[string]*string) *CreateInstanceShutdownTimerResponse {
	s.Headers = v
	return s
}

func (s *CreateInstanceShutdownTimerResponse) SetStatusCode(v int32) *CreateInstanceShutdownTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateInstanceShutdownTimerResponse) SetBody(v *CreateInstanceShutdownTimerResponseBody) *CreateInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

func (s *CreateInstanceShutdownTimerResponse) Validate() error {
	return dara.Validate(s)
}
