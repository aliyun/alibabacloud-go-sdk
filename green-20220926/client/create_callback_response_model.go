// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCallbackResponse
	GetStatusCode() *int32
	SetBody(v *CreateCallbackResponseBody) *CreateCallbackResponse
	GetBody() *CreateCallbackResponseBody
}

type CreateCallbackResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCallbackResponse) GoString() string {
	return s.String()
}

func (s *CreateCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCallbackResponse) GetBody() *CreateCallbackResponseBody {
	return s.Body
}

func (s *CreateCallbackResponse) SetHeaders(v map[string]*string) *CreateCallbackResponse {
	s.Headers = v
	return s
}

func (s *CreateCallbackResponse) SetStatusCode(v int32) *CreateCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCallbackResponse) SetBody(v *CreateCallbackResponseBody) *CreateCallbackResponse {
	s.Body = v
	return s
}

func (s *CreateCallbackResponse) Validate() error {
	return dara.Validate(s)
}
