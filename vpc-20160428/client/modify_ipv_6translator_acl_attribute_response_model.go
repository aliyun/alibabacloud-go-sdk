// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorAclAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIPv6TranslatorAclAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIPv6TranslatorAclAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIPv6TranslatorAclAttributeResponseBody) *ModifyIPv6TranslatorAclAttributeResponse
	GetBody() *ModifyIPv6TranslatorAclAttributeResponseBody
}

type ModifyIPv6TranslatorAclAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIPv6TranslatorAclAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIPv6TranslatorAclAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorAclAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorAclAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIPv6TranslatorAclAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIPv6TranslatorAclAttributeResponse) GetBody() *ModifyIPv6TranslatorAclAttributeResponseBody {
	return s.Body
}

func (s *ModifyIPv6TranslatorAclAttributeResponse) SetHeaders(v map[string]*string) *ModifyIPv6TranslatorAclAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeResponse) SetStatusCode(v int32) *ModifyIPv6TranslatorAclAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeResponse) SetBody(v *ModifyIPv6TranslatorAclAttributeResponseBody) *ModifyIPv6TranslatorAclAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyIPv6TranslatorAclAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
