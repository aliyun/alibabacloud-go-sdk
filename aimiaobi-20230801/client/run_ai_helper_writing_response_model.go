// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunAiHelperWritingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunAiHelperWritingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunAiHelperWritingResponse
	GetStatusCode() *int32
	SetBody(v *RunAiHelperWritingResponseBody) *RunAiHelperWritingResponse
	GetBody() *RunAiHelperWritingResponseBody
}

type RunAiHelperWritingResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunAiHelperWritingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunAiHelperWritingResponse) String() string {
	return dara.Prettify(s)
}

func (s RunAiHelperWritingResponse) GoString() string {
	return s.String()
}

func (s *RunAiHelperWritingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunAiHelperWritingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunAiHelperWritingResponse) GetBody() *RunAiHelperWritingResponseBody {
	return s.Body
}

func (s *RunAiHelperWritingResponse) SetHeaders(v map[string]*string) *RunAiHelperWritingResponse {
	s.Headers = v
	return s
}

func (s *RunAiHelperWritingResponse) SetStatusCode(v int32) *RunAiHelperWritingResponse {
	s.StatusCode = &v
	return s
}

func (s *RunAiHelperWritingResponse) SetBody(v *RunAiHelperWritingResponseBody) *RunAiHelperWritingResponse {
	s.Body = v
	return s
}

func (s *RunAiHelperWritingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
