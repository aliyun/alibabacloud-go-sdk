// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMessageCallbackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMessageCallbackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMessageCallbackResponse
	GetStatusCode() *int32
	SetBody(v *GetMessageCallbackResponseBody) *GetMessageCallbackResponse
	GetBody() *GetMessageCallbackResponseBody
}

type GetMessageCallbackResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMessageCallbackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMessageCallbackResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMessageCallbackResponse) GoString() string {
	return s.String()
}

func (s *GetMessageCallbackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMessageCallbackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMessageCallbackResponse) GetBody() *GetMessageCallbackResponseBody {
	return s.Body
}

func (s *GetMessageCallbackResponse) SetHeaders(v map[string]*string) *GetMessageCallbackResponse {
	s.Headers = v
	return s
}

func (s *GetMessageCallbackResponse) SetStatusCode(v int32) *GetMessageCallbackResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMessageCallbackResponse) SetBody(v *GetMessageCallbackResponseBody) *GetMessageCallbackResponse {
	s.Body = v
	return s
}

func (s *GetMessageCallbackResponse) Validate() error {
	return dara.Validate(s)
}
