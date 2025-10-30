// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAxnExtensionBindFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAxnExtensionBindFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAxnExtensionBindFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAxnExtensionBindFixedLineResponseBody) *UpdateAxnExtensionBindFixedLineResponse
	GetBody() *UpdateAxnExtensionBindFixedLineResponseBody
}

type UpdateAxnExtensionBindFixedLineResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAxnExtensionBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAxnExtensionBindFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAxnExtensionBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *UpdateAxnExtensionBindFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAxnExtensionBindFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAxnExtensionBindFixedLineResponse) GetBody() *UpdateAxnExtensionBindFixedLineResponseBody {
	return s.Body
}

func (s *UpdateAxnExtensionBindFixedLineResponse) SetHeaders(v map[string]*string) *UpdateAxnExtensionBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponse) SetStatusCode(v int32) *UpdateAxnExtensionBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponse) SetBody(v *UpdateAxnExtensionBindFixedLineResponseBody) *UpdateAxnExtensionBindFixedLineResponse {
	s.Body = v
	return s
}

func (s *UpdateAxnExtensionBindFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
