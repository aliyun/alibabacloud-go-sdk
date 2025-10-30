// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindAxnExtensionFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BindAxnExtensionFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BindAxnExtensionFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *BindAxnExtensionFixedLineResponseBody) *BindAxnExtensionFixedLineResponse
	GetBody() *BindAxnExtensionFixedLineResponseBody
}

type BindAxnExtensionFixedLineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BindAxnExtensionFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BindAxnExtensionFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s BindAxnExtensionFixedLineResponse) GoString() string {
	return s.String()
}

func (s *BindAxnExtensionFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BindAxnExtensionFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BindAxnExtensionFixedLineResponse) GetBody() *BindAxnExtensionFixedLineResponseBody {
	return s.Body
}

func (s *BindAxnExtensionFixedLineResponse) SetHeaders(v map[string]*string) *BindAxnExtensionFixedLineResponse {
	s.Headers = v
	return s
}

func (s *BindAxnExtensionFixedLineResponse) SetStatusCode(v int32) *BindAxnExtensionFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *BindAxnExtensionFixedLineResponse) SetBody(v *BindAxnExtensionFixedLineResponseBody) *BindAxnExtensionFixedLineResponse {
	s.Body = v
	return s
}

func (s *BindAxnExtensionFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
