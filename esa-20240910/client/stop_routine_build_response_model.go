// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopRoutineBuildResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopRoutineBuildResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopRoutineBuildResponse
	GetStatusCode() *int32
	SetBody(v *StopRoutineBuildResponseBody) *StopRoutineBuildResponse
	GetBody() *StopRoutineBuildResponseBody
}

type StopRoutineBuildResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopRoutineBuildResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopRoutineBuildResponse) String() string {
	return dara.Prettify(s)
}

func (s StopRoutineBuildResponse) GoString() string {
	return s.String()
}

func (s *StopRoutineBuildResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopRoutineBuildResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopRoutineBuildResponse) GetBody() *StopRoutineBuildResponseBody {
	return s.Body
}

func (s *StopRoutineBuildResponse) SetHeaders(v map[string]*string) *StopRoutineBuildResponse {
	s.Headers = v
	return s
}

func (s *StopRoutineBuildResponse) SetStatusCode(v int32) *StopRoutineBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *StopRoutineBuildResponse) SetBody(v *StopRoutineBuildResponseBody) *StopRoutineBuildResponse {
	s.Body = v
	return s
}

func (s *StopRoutineBuildResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
