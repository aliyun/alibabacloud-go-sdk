// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCompletionMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCompletionMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCompletionMessageResponse
	GetStatusCode() *int32
	SetBody(v *RunCompletionMessageResponseBody) *RunCompletionMessageResponse
	GetBody() *RunCompletionMessageResponseBody
}

type RunCompletionMessageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCompletionMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCompletionMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCompletionMessageResponse) GoString() string {
	return s.String()
}

func (s *RunCompletionMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCompletionMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCompletionMessageResponse) GetBody() *RunCompletionMessageResponseBody {
	return s.Body
}

func (s *RunCompletionMessageResponse) SetHeaders(v map[string]*string) *RunCompletionMessageResponse {
	s.Headers = v
	return s
}

func (s *RunCompletionMessageResponse) SetStatusCode(v int32) *RunCompletionMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCompletionMessageResponse) SetBody(v *RunCompletionMessageResponseBody) *RunCompletionMessageResponse {
	s.Body = v
	return s
}

func (s *RunCompletionMessageResponse) Validate() error {
	return dara.Validate(s)
}
