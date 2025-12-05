// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDebuggingJMeterSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDebuggingJMeterSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDebuggingJMeterSceneResponse
	GetStatusCode() *int32
	SetBody(v *StartDebuggingJMeterSceneResponseBody) *StartDebuggingJMeterSceneResponse
	GetBody() *StartDebuggingJMeterSceneResponseBody
}

type StartDebuggingJMeterSceneResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDebuggingJMeterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDebuggingJMeterSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDebuggingJMeterSceneResponse) GoString() string {
	return s.String()
}

func (s *StartDebuggingJMeterSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDebuggingJMeterSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDebuggingJMeterSceneResponse) GetBody() *StartDebuggingJMeterSceneResponseBody {
	return s.Body
}

func (s *StartDebuggingJMeterSceneResponse) SetHeaders(v map[string]*string) *StartDebuggingJMeterSceneResponse {
	s.Headers = v
	return s
}

func (s *StartDebuggingJMeterSceneResponse) SetStatusCode(v int32) *StartDebuggingJMeterSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDebuggingJMeterSceneResponse) SetBody(v *StartDebuggingJMeterSceneResponseBody) *StartDebuggingJMeterSceneResponse {
	s.Body = v
	return s
}

func (s *StartDebuggingJMeterSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
