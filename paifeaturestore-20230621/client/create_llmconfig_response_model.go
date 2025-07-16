// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLLMConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLLMConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLLMConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateLLMConfigResponseBody) *CreateLLMConfigResponse
	GetBody() *CreateLLMConfigResponseBody
}

type CreateLLMConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLLMConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLLMConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLLMConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateLLMConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLLMConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLLMConfigResponse) GetBody() *CreateLLMConfigResponseBody {
	return s.Body
}

func (s *CreateLLMConfigResponse) SetHeaders(v map[string]*string) *CreateLLMConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateLLMConfigResponse) SetStatusCode(v int32) *CreateLLMConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLLMConfigResponse) SetBody(v *CreateLLMConfigResponseBody) *CreateLLMConfigResponse {
	s.Body = v
	return s
}

func (s *CreateLLMConfigResponse) Validate() error {
	return dara.Validate(s)
}
