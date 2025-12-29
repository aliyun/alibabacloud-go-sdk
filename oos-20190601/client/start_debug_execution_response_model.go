// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDebugExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDebugExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDebugExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StartDebugExecutionResponseBody) *StartDebugExecutionResponse
	GetBody() *StartDebugExecutionResponseBody
}

type StartDebugExecutionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDebugExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDebugExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDebugExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartDebugExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDebugExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDebugExecutionResponse) GetBody() *StartDebugExecutionResponseBody {
	return s.Body
}

func (s *StartDebugExecutionResponse) SetHeaders(v map[string]*string) *StartDebugExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartDebugExecutionResponse) SetStatusCode(v int32) *StartDebugExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDebugExecutionResponse) SetBody(v *StartDebugExecutionResponseBody) *StartDebugExecutionResponse {
	s.Body = v
	return s
}

func (s *StartDebugExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
