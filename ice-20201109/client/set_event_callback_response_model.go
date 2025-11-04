// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEventCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetEventCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetEventCallbackResponse
	GetStatusCode() *int32
	SetBody(v *SetEventCallbackResponseBody) *SetEventCallbackResponse
	GetBody() *SetEventCallbackResponseBody
}

type SetEventCallbackResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetEventCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetEventCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s SetEventCallbackResponse) GoString() string {
	return s.String()
}

func (s *SetEventCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetEventCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetEventCallbackResponse) GetBody() *SetEventCallbackResponseBody {
	return s.Body
}

func (s *SetEventCallbackResponse) SetHeaders(v map[string]*string) *SetEventCallbackResponse {
	s.Headers = v
	return s
}

func (s *SetEventCallbackResponse) SetStatusCode(v int32) *SetEventCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *SetEventCallbackResponse) SetBody(v *SetEventCallbackResponseBody) *SetEventCallbackResponse {
	s.Body = v
	return s
}

func (s *SetEventCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
