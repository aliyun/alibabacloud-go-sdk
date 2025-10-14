// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceShutdownTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceShutdownTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceShutdownTimerResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceShutdownTimerResponseBody) *DeleteInstanceShutdownTimerResponse
	GetBody() *DeleteInstanceShutdownTimerResponseBody
}

type DeleteInstanceShutdownTimerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceShutdownTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceShutdownTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceShutdownTimerResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceShutdownTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceShutdownTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceShutdownTimerResponse) GetBody() *DeleteInstanceShutdownTimerResponseBody {
	return s.Body
}

func (s *DeleteInstanceShutdownTimerResponse) SetHeaders(v map[string]*string) *DeleteInstanceShutdownTimerResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceShutdownTimerResponse) SetStatusCode(v int32) *DeleteInstanceShutdownTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceShutdownTimerResponse) SetBody(v *DeleteInstanceShutdownTimerResponseBody) *DeleteInstanceShutdownTimerResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceShutdownTimerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
