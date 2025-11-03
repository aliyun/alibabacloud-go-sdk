// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorAclListEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIPv6TranslatorAclListEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIPv6TranslatorAclListEntryResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIPv6TranslatorAclListEntryResponseBody) *ModifyIPv6TranslatorAclListEntryResponse
	GetBody() *ModifyIPv6TranslatorAclListEntryResponseBody
}

type ModifyIPv6TranslatorAclListEntryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIPv6TranslatorAclListEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIPv6TranslatorAclListEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorAclListEntryResponse) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorAclListEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIPv6TranslatorAclListEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIPv6TranslatorAclListEntryResponse) GetBody() *ModifyIPv6TranslatorAclListEntryResponseBody {
	return s.Body
}

func (s *ModifyIPv6TranslatorAclListEntryResponse) SetHeaders(v map[string]*string) *ModifyIPv6TranslatorAclListEntryResponse {
	s.Headers = v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryResponse) SetStatusCode(v int32) *ModifyIPv6TranslatorAclListEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryResponse) SetBody(v *ModifyIPv6TranslatorAclListEntryResponseBody) *ModifyIPv6TranslatorAclListEntryResponse {
	s.Body = v
	return s
}

func (s *ModifyIPv6TranslatorAclListEntryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
