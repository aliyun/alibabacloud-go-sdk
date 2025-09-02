// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCallbackExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CallbackExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CallbackExtensionResponse
	GetStatusCode() *int32
	SetBody(v *CallbackExtensionResponseBody) *CallbackExtensionResponse
	GetBody() *CallbackExtensionResponseBody
}

type CallbackExtensionResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CallbackExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CallbackExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s CallbackExtensionResponse) GoString() string {
	return s.String()
}

func (s *CallbackExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CallbackExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CallbackExtensionResponse) GetBody() *CallbackExtensionResponseBody {
	return s.Body
}

func (s *CallbackExtensionResponse) SetHeaders(v map[string]*string) *CallbackExtensionResponse {
	s.Headers = v
	return s
}

func (s *CallbackExtensionResponse) SetStatusCode(v int32) *CallbackExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *CallbackExtensionResponse) SetBody(v *CallbackExtensionResponseBody) *CallbackExtensionResponse {
	s.Body = v
	return s
}

func (s *CallbackExtensionResponse) Validate() error {
	return dara.Validate(s)
}
