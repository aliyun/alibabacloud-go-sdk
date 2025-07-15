// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveIPv6TranslatorAclListEntryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveIPv6TranslatorAclListEntryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveIPv6TranslatorAclListEntryResponse
	GetStatusCode() *int32
	SetBody(v *RemoveIPv6TranslatorAclListEntryResponseBody) *RemoveIPv6TranslatorAclListEntryResponse
	GetBody() *RemoveIPv6TranslatorAclListEntryResponseBody
}

type RemoveIPv6TranslatorAclListEntryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveIPv6TranslatorAclListEntryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveIPv6TranslatorAclListEntryResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveIPv6TranslatorAclListEntryResponse) GoString() string {
	return s.String()
}

func (s *RemoveIPv6TranslatorAclListEntryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveIPv6TranslatorAclListEntryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveIPv6TranslatorAclListEntryResponse) GetBody() *RemoveIPv6TranslatorAclListEntryResponseBody {
	return s.Body
}

func (s *RemoveIPv6TranslatorAclListEntryResponse) SetHeaders(v map[string]*string) *RemoveIPv6TranslatorAclListEntryResponse {
	s.Headers = v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryResponse) SetStatusCode(v int32) *RemoveIPv6TranslatorAclListEntryResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryResponse) SetBody(v *RemoveIPv6TranslatorAclListEntryResponseBody) *RemoveIPv6TranslatorAclListEntryResponse {
	s.Body = v
	return s
}

func (s *RemoveIPv6TranslatorAclListEntryResponse) Validate() error {
	return dara.Validate(s)
}
