// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StartExecutionResponseBody) *StartExecutionResponse
	GetBody() *StartExecutionResponseBody
}

type StartExecutionResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartExecutionResponse) GetBody() *StartExecutionResponseBody {
	return s.Body
}

func (s *StartExecutionResponse) SetHeaders(v map[string]*string) *StartExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartExecutionResponse) SetStatusCode(v int32) *StartExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartExecutionResponse) SetBody(v *StartExecutionResponseBody) *StartExecutionResponse {
	s.Body = v
	return s
}

func (s *StartExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
