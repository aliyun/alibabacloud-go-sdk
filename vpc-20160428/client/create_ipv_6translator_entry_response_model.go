// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIPv6TranslatorEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateIPv6TranslatorEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateIPv6TranslatorEntryResponse
	GetStatusCode() *int32
	SetBody(v *CreateIPv6TranslatorEntryResponseBody) *CreateIPv6TranslatorEntryResponse
	GetBody() *CreateIPv6TranslatorEntryResponseBody
}

type CreateIPv6TranslatorEntryResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateIPv6TranslatorEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateIPv6TranslatorEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateIPv6TranslatorEntryResponse) GoString() string {
	return s.String()
}

func (s *CreateIPv6TranslatorEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateIPv6TranslatorEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateIPv6TranslatorEntryResponse) GetBody() *CreateIPv6TranslatorEntryResponseBody {
	return s.Body
}

func (s *CreateIPv6TranslatorEntryResponse) SetHeaders(v map[string]*string) *CreateIPv6TranslatorEntryResponse {
	s.Headers = v
	return s
}

func (s *CreateIPv6TranslatorEntryResponse) SetStatusCode(v int32) *CreateIPv6TranslatorEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateIPv6TranslatorEntryResponse) SetBody(v *CreateIPv6TranslatorEntryResponseBody) *CreateIPv6TranslatorEntryResponse {
	s.Body = v
	return s
}

func (s *CreateIPv6TranslatorEntryResponse) Validate() error {
	return dara.Validate(s)
}
