// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessCustomIMCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProcessCustomIMCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProcessCustomIMCallbackResponse
	GetStatusCode() *int32
	SetBody(v *ProcessCustomIMCallbackResponseBody) *ProcessCustomIMCallbackResponse
	GetBody() *ProcessCustomIMCallbackResponseBody
}

type ProcessCustomIMCallbackResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProcessCustomIMCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProcessCustomIMCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s ProcessCustomIMCallbackResponse) GoString() string {
	return s.String()
}

func (s *ProcessCustomIMCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProcessCustomIMCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProcessCustomIMCallbackResponse) GetBody() *ProcessCustomIMCallbackResponseBody {
	return s.Body
}

func (s *ProcessCustomIMCallbackResponse) SetHeaders(v map[string]*string) *ProcessCustomIMCallbackResponse {
	s.Headers = v
	return s
}

func (s *ProcessCustomIMCallbackResponse) SetStatusCode(v int32) *ProcessCustomIMCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ProcessCustomIMCallbackResponse) SetBody(v *ProcessCustomIMCallbackResponseBody) *ProcessCustomIMCallbackResponse {
	s.Body = v
	return s
}

func (s *ProcessCustomIMCallbackResponse) Validate() error {
	return dara.Validate(s)
}
