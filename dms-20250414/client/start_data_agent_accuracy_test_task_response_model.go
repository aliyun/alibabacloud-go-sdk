// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataAgentAccuracyTestTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDataAgentAccuracyTestTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDataAgentAccuracyTestTaskResponse
	GetStatusCode() *int32
	SetBody(v *StartDataAgentAccuracyTestTaskResponseBody) *StartDataAgentAccuracyTestTaskResponse
	GetBody() *StartDataAgentAccuracyTestTaskResponseBody
}

type StartDataAgentAccuracyTestTaskResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDataAgentAccuracyTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDataAgentAccuracyTestTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDataAgentAccuracyTestTaskResponse) GoString() string {
	return s.String()
}

func (s *StartDataAgentAccuracyTestTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDataAgentAccuracyTestTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDataAgentAccuracyTestTaskResponse) GetBody() *StartDataAgentAccuracyTestTaskResponseBody {
	return s.Body
}

func (s *StartDataAgentAccuracyTestTaskResponse) SetHeaders(v map[string]*string) *StartDataAgentAccuracyTestTaskResponse {
	s.Headers = v
	return s
}

func (s *StartDataAgentAccuracyTestTaskResponse) SetStatusCode(v int32) *StartDataAgentAccuracyTestTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDataAgentAccuracyTestTaskResponse) SetBody(v *StartDataAgentAccuracyTestTaskResponseBody) *StartDataAgentAccuracyTestTaskResponse {
	s.Body = v
	return s
}

func (s *StartDataAgentAccuracyTestTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
