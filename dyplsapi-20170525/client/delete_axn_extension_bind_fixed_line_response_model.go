// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxnExtensionBindFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAxnExtensionBindFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAxnExtensionBindFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAxnExtensionBindFixedLineResponseBody) *DeleteAxnExtensionBindFixedLineResponse
	GetBody() *DeleteAxnExtensionBindFixedLineResponseBody
}

type DeleteAxnExtensionBindFixedLineResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAxnExtensionBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAxnExtensionBindFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxnExtensionBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *DeleteAxnExtensionBindFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAxnExtensionBindFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAxnExtensionBindFixedLineResponse) GetBody() *DeleteAxnExtensionBindFixedLineResponseBody {
	return s.Body
}

func (s *DeleteAxnExtensionBindFixedLineResponse) SetHeaders(v map[string]*string) *DeleteAxnExtensionBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponse) SetStatusCode(v int32) *DeleteAxnExtensionBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponse) SetBody(v *DeleteAxnExtensionBindFixedLineResponseBody) *DeleteAxnExtensionBindFixedLineResponse {
	s.Body = v
	return s
}

func (s *DeleteAxnExtensionBindFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
