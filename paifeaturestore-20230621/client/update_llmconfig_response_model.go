// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLLMConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLLMConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLLMConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLLMConfigResponseBody) *UpdateLLMConfigResponse
	GetBody() *UpdateLLMConfigResponseBody
}

type UpdateLLMConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLLMConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLLMConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLLMConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLLMConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLLMConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLLMConfigResponse) GetBody() *UpdateLLMConfigResponseBody {
	return s.Body
}

func (s *UpdateLLMConfigResponse) SetHeaders(v map[string]*string) *UpdateLLMConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLLMConfigResponse) SetStatusCode(v int32) *UpdateLLMConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLLMConfigResponse) SetBody(v *UpdateLLMConfigResponseBody) *UpdateLLMConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLLMConfigResponse) Validate() error {
	return dara.Validate(s)
}
