// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceShutdownTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceShutdownTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceShutdownTimerResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceShutdownTimerResponseBody) *GetInstanceShutdownTimerResponse
	GetBody() *GetInstanceShutdownTimerResponseBody
}

type GetInstanceShutdownTimerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceShutdownTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceShutdownTimerResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceShutdownTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceShutdownTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceShutdownTimerResponse) GetBody() *GetInstanceShutdownTimerResponseBody {
	return s.Body
}

func (s *GetInstanceShutdownTimerResponse) SetHeaders(v map[string]*string) *GetInstanceShutdownTimerResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceShutdownTimerResponse) SetStatusCode(v int32) *GetInstanceShutdownTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceShutdownTimerResponse) SetBody(v *GetInstanceShutdownTimerResponseBody) *GetInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

func (s *GetInstanceShutdownTimerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
