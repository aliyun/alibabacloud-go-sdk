// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIPv6TranslatorAclListEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddIPv6TranslatorAclListEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddIPv6TranslatorAclListEntryResponse
	GetStatusCode() *int32
	SetBody(v *AddIPv6TranslatorAclListEntryResponseBody) *AddIPv6TranslatorAclListEntryResponse
	GetBody() *AddIPv6TranslatorAclListEntryResponseBody
}

type AddIPv6TranslatorAclListEntryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddIPv6TranslatorAclListEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddIPv6TranslatorAclListEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddIPv6TranslatorAclListEntryResponse) GoString() string {
	return s.String()
}

func (s *AddIPv6TranslatorAclListEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddIPv6TranslatorAclListEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddIPv6TranslatorAclListEntryResponse) GetBody() *AddIPv6TranslatorAclListEntryResponseBody {
	return s.Body
}

func (s *AddIPv6TranslatorAclListEntryResponse) SetHeaders(v map[string]*string) *AddIPv6TranslatorAclListEntryResponse {
	s.Headers = v
	return s
}

func (s *AddIPv6TranslatorAclListEntryResponse) SetStatusCode(v int32) *AddIPv6TranslatorAclListEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddIPv6TranslatorAclListEntryResponse) SetBody(v *AddIPv6TranslatorAclListEntryResponseBody) *AddIPv6TranslatorAclListEntryResponse {
	s.Body = v
	return s
}

func (s *AddIPv6TranslatorAclListEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
