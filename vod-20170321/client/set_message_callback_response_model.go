// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetMessageCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetMessageCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetMessageCallbackResponse
	GetStatusCode() *int32
	SetBody(v *SetMessageCallbackResponseBody) *SetMessageCallbackResponse
	GetBody() *SetMessageCallbackResponseBody
}

type SetMessageCallbackResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetMessageCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetMessageCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s SetMessageCallbackResponse) GoString() string {
	return s.String()
}

func (s *SetMessageCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetMessageCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetMessageCallbackResponse) GetBody() *SetMessageCallbackResponseBody {
	return s.Body
}

func (s *SetMessageCallbackResponse) SetHeaders(v map[string]*string) *SetMessageCallbackResponse {
	s.Headers = v
	return s
}

func (s *SetMessageCallbackResponse) SetStatusCode(v int32) *SetMessageCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *SetMessageCallbackResponse) SetBody(v *SetMessageCallbackResponseBody) *SetMessageCallbackResponse {
	s.Body = v
	return s
}

func (s *SetMessageCallbackResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
