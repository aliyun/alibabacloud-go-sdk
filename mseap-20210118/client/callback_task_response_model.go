// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallbackTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CallbackTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CallbackTaskResponse
	GetStatusCode() *int32
	SetBody(v *CallbackTaskResponseBody) *CallbackTaskResponse
	GetBody() *CallbackTaskResponseBody
}

type CallbackTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CallbackTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CallbackTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CallbackTaskResponse) GoString() string {
	return s.String()
}

func (s *CallbackTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CallbackTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CallbackTaskResponse) GetBody() *CallbackTaskResponseBody {
	return s.Body
}

func (s *CallbackTaskResponse) SetHeaders(v map[string]*string) *CallbackTaskResponse {
	s.Headers = v
	return s
}

func (s *CallbackTaskResponse) SetStatusCode(v int32) *CallbackTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CallbackTaskResponse) SetBody(v *CallbackTaskResponseBody) *CallbackTaskResponse {
	s.Body = v
	return s
}

func (s *CallbackTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
