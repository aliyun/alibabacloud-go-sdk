// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResultCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResultCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResultCallbackResponse
	GetStatusCode() *int32
	SetBody(v *ResultCallbackResponseBody) *ResultCallbackResponse
	GetBody() *ResultCallbackResponseBody
}

type ResultCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResultCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResultCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s ResultCallbackResponse) GoString() string {
	return s.String()
}

func (s *ResultCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResultCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResultCallbackResponse) GetBody() *ResultCallbackResponseBody {
	return s.Body
}

func (s *ResultCallbackResponse) SetHeaders(v map[string]*string) *ResultCallbackResponse {
	s.Headers = v
	return s
}

func (s *ResultCallbackResponse) SetStatusCode(v int32) *ResultCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ResultCallbackResponse) SetBody(v *ResultCallbackResponseBody) *ResultCallbackResponse {
	s.Body = v
	return s
}

func (s *ResultCallbackResponse) Validate() error {
	return dara.Validate(s)
}
