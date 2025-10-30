// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAxnExtensionBindFixedLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAxnExtensionBindFixedLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAxnExtensionBindFixedLineResponse
	GetStatusCode() *int32
	SetBody(v *QueryAxnExtensionBindFixedLineResponseBody) *QueryAxnExtensionBindFixedLineResponse
	GetBody() *QueryAxnExtensionBindFixedLineResponseBody
}

type QueryAxnExtensionBindFixedLineResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAxnExtensionBindFixedLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAxnExtensionBindFixedLineResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAxnExtensionBindFixedLineResponse) GoString() string {
	return s.String()
}

func (s *QueryAxnExtensionBindFixedLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAxnExtensionBindFixedLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAxnExtensionBindFixedLineResponse) GetBody() *QueryAxnExtensionBindFixedLineResponseBody {
	return s.Body
}

func (s *QueryAxnExtensionBindFixedLineResponse) SetHeaders(v map[string]*string) *QueryAxnExtensionBindFixedLineResponse {
	s.Headers = v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponse) SetStatusCode(v int32) *QueryAxnExtensionBindFixedLineResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponse) SetBody(v *QueryAxnExtensionBindFixedLineResponseBody) *QueryAxnExtensionBindFixedLineResponse {
	s.Body = v
	return s
}

func (s *QueryAxnExtensionBindFixedLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
