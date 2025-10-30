// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnExtensionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAxnExtensionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAxnExtensionResponse
	GetStatusCode() *int32
	SetBody(v *BindAxnExtensionResponseBody) *BindAxnExtensionResponse
	GetBody() *BindAxnExtensionResponseBody
}

type BindAxnExtensionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxnExtensionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxnExtensionResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionResponse) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAxnExtensionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAxnExtensionResponse) GetBody() *BindAxnExtensionResponseBody {
	return s.Body
}

func (s *BindAxnExtensionResponse) SetHeaders(v map[string]*string) *BindAxnExtensionResponse {
	s.Headers = v
	return s
}

func (s *BindAxnExtensionResponse) SetStatusCode(v int32) *BindAxnExtensionResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxnExtensionResponse) SetBody(v *BindAxnExtensionResponseBody) *BindAxnExtensionResponse {
	s.Body = v
	return s
}

func (s *BindAxnExtensionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
