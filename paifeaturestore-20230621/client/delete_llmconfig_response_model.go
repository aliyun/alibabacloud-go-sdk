// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLLMConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLLMConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLLMConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLLMConfigResponseBody) *DeleteLLMConfigResponse
	GetBody() *DeleteLLMConfigResponseBody
}

type DeleteLLMConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLLMConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLLMConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLLMConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLLMConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLLMConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLLMConfigResponse) GetBody() *DeleteLLMConfigResponseBody {
	return s.Body
}

func (s *DeleteLLMConfigResponse) SetHeaders(v map[string]*string) *DeleteLLMConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLLMConfigResponse) SetStatusCode(v int32) *DeleteLLMConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLLMConfigResponse) SetBody(v *DeleteLLMConfigResponseBody) *DeleteLLMConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLLMConfigResponse) Validate() error {
	return dara.Validate(s)
}
