// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIPv6TranslatorResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIPv6TranslatorResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIPv6TranslatorResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIPv6TranslatorResponseBody) *DeleteIPv6TranslatorResponse
	GetBody() *DeleteIPv6TranslatorResponseBody
}

type DeleteIPv6TranslatorResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIPv6TranslatorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIPv6TranslatorResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIPv6TranslatorResponse) GoString() string {
	return s.String()
}

func (s *DeleteIPv6TranslatorResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIPv6TranslatorResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIPv6TranslatorResponse) GetBody() *DeleteIPv6TranslatorResponseBody {
	return s.Body
}

func (s *DeleteIPv6TranslatorResponse) SetHeaders(v map[string]*string) *DeleteIPv6TranslatorResponse {
	s.Headers = v
	return s
}

func (s *DeleteIPv6TranslatorResponse) SetStatusCode(v int32) *DeleteIPv6TranslatorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIPv6TranslatorResponse) SetBody(v *DeleteIPv6TranslatorResponseBody) *DeleteIPv6TranslatorResponse {
	s.Body = v
	return s
}

func (s *DeleteIPv6TranslatorResponse) Validate() error {
	return dara.Validate(s)
}
