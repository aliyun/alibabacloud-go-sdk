// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIPv6TranslatorAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIPv6TranslatorAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIPv6TranslatorAttributeResponseBody) *ModifyIPv6TranslatorAttributeResponse
	GetBody() *ModifyIPv6TranslatorAttributeResponseBody
}

type ModifyIPv6TranslatorAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIPv6TranslatorAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIPv6TranslatorAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIPv6TranslatorAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIPv6TranslatorAttributeResponse) GetBody() *ModifyIPv6TranslatorAttributeResponseBody {
	return s.Body
}

func (s *ModifyIPv6TranslatorAttributeResponse) SetHeaders(v map[string]*string) *ModifyIPv6TranslatorAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyIPv6TranslatorAttributeResponse) SetStatusCode(v int32) *ModifyIPv6TranslatorAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIPv6TranslatorAttributeResponse) SetBody(v *ModifyIPv6TranslatorAttributeResponseBody) *ModifyIPv6TranslatorAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyIPv6TranslatorAttributeResponse) Validate() error {
	return dara.Validate(s)
}
