// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebSocketSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GenerateWebSocketSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GenerateWebSocketSignResponse
	GetStatusCode() *int32
	SetBody(v *GenerateWebSocketSignResponseBody) *GenerateWebSocketSignResponse
	GetBody() *GenerateWebSocketSignResponseBody
}

type GenerateWebSocketSignResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GenerateWebSocketSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GenerateWebSocketSignResponse) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebSocketSignResponse) GoString() string {
	return s.String()
}

func (s *GenerateWebSocketSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GenerateWebSocketSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GenerateWebSocketSignResponse) GetBody() *GenerateWebSocketSignResponseBody {
	return s.Body
}

func (s *GenerateWebSocketSignResponse) SetHeaders(v map[string]*string) *GenerateWebSocketSignResponse {
	s.Headers = v
	return s
}

func (s *GenerateWebSocketSignResponse) SetStatusCode(v int32) *GenerateWebSocketSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GenerateWebSocketSignResponse) SetBody(v *GenerateWebSocketSignResponseBody) *GenerateWebSocketSignResponse {
	s.Body = v
	return s
}

func (s *GenerateWebSocketSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
