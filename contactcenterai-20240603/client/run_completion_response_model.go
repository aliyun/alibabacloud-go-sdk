// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCompletionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCompletionResponse
	GetStatusCode() *int32
	SetBody(v *RunCompletionResponseBody) *RunCompletionResponse
	GetBody() *RunCompletionResponseBody
}

type RunCompletionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCompletionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCompletionResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionResponse) GoString() string {
	return s.String()
}

func (s *RunCompletionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCompletionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCompletionResponse) GetBody() *RunCompletionResponseBody {
	return s.Body
}

func (s *RunCompletionResponse) SetHeaders(v map[string]*string) *RunCompletionResponse {
	s.Headers = v
	return s
}

func (s *RunCompletionResponse) SetStatusCode(v int32) *RunCompletionResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCompletionResponse) SetBody(v *RunCompletionResponseBody) *RunCompletionResponse {
	s.Body = v
	return s
}

func (s *RunCompletionResponse) Validate() error {
	return dara.Validate(s)
}
