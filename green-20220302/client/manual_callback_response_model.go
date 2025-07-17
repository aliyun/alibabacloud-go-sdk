// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManualCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManualCallbackResponse
	GetStatusCode() *int32
	SetBody(v *ManualCallbackResponseBody) *ManualCallbackResponse
	GetBody() *ManualCallbackResponseBody
}

type ManualCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManualCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManualCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s ManualCallbackResponse) GoString() string {
	return s.String()
}

func (s *ManualCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManualCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManualCallbackResponse) GetBody() *ManualCallbackResponseBody {
	return s.Body
}

func (s *ManualCallbackResponse) SetHeaders(v map[string]*string) *ManualCallbackResponse {
	s.Headers = v
	return s
}

func (s *ManualCallbackResponse) SetStatusCode(v int32) *ManualCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *ManualCallbackResponse) SetBody(v *ManualCallbackResponseBody) *ManualCallbackResponse {
	s.Body = v
	return s
}

func (s *ManualCallbackResponse) Validate() error {
	return dara.Validate(s)
}
