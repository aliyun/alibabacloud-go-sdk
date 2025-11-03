// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIPv6TranslatorEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIPv6TranslatorEntryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIPv6TranslatorEntryResponseBody) *ModifyIPv6TranslatorEntryResponse
	GetBody() *ModifyIPv6TranslatorEntryResponseBody
}

type ModifyIPv6TranslatorEntryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIPv6TranslatorEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIPv6TranslatorEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorEntryResponse) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIPv6TranslatorEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIPv6TranslatorEntryResponse) GetBody() *ModifyIPv6TranslatorEntryResponseBody {
	return s.Body
}

func (s *ModifyIPv6TranslatorEntryResponse) SetHeaders(v map[string]*string) *ModifyIPv6TranslatorEntryResponse {
	s.Headers = v
	return s
}

func (s *ModifyIPv6TranslatorEntryResponse) SetStatusCode(v int32) *ModifyIPv6TranslatorEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIPv6TranslatorEntryResponse) SetBody(v *ModifyIPv6TranslatorEntryResponseBody) *ModifyIPv6TranslatorEntryResponse {
	s.Body = v
	return s
}

func (s *ModifyIPv6TranslatorEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
