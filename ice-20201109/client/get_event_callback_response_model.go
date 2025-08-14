// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEventCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEventCallbackResponse
	GetStatusCode() *int32
	SetBody(v *GetEventCallbackResponseBody) *GetEventCallbackResponse
	GetBody() *GetEventCallbackResponseBody
}

type GetEventCallbackResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEventCallbackResponse) GoString() string {
	return s.String()
}

func (s *GetEventCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEventCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEventCallbackResponse) GetBody() *GetEventCallbackResponseBody {
	return s.Body
}

func (s *GetEventCallbackResponse) SetHeaders(v map[string]*string) *GetEventCallbackResponse {
	s.Headers = v
	return s
}

func (s *GetEventCallbackResponse) SetStatusCode(v int32) *GetEventCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventCallbackResponse) SetBody(v *GetEventCallbackResponseBody) *GetEventCallbackResponse {
	s.Body = v
	return s
}

func (s *GetEventCallbackResponse) Validate() error {
	return dara.Validate(s)
}
