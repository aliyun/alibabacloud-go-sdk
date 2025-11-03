// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIPv6TranslatorEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIPv6TranslatorEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIPv6TranslatorEntryResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIPv6TranslatorEntryResponseBody) *DeleteIPv6TranslatorEntryResponse
	GetBody() *DeleteIPv6TranslatorEntryResponseBody
}

type DeleteIPv6TranslatorEntryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIPv6TranslatorEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIPv6TranslatorEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIPv6TranslatorEntryResponse) GoString() string {
	return s.String()
}

func (s *DeleteIPv6TranslatorEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIPv6TranslatorEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIPv6TranslatorEntryResponse) GetBody() *DeleteIPv6TranslatorEntryResponseBody {
	return s.Body
}

func (s *DeleteIPv6TranslatorEntryResponse) SetHeaders(v map[string]*string) *DeleteIPv6TranslatorEntryResponse {
	s.Headers = v
	return s
}

func (s *DeleteIPv6TranslatorEntryResponse) SetStatusCode(v int32) *DeleteIPv6TranslatorEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIPv6TranslatorEntryResponse) SetBody(v *DeleteIPv6TranslatorEntryResponseBody) *DeleteIPv6TranslatorEntryResponse {
	s.Body = v
	return s
}

func (s *DeleteIPv6TranslatorEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
