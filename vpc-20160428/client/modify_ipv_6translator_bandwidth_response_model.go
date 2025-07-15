// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIPv6TranslatorBandwidthResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyIPv6TranslatorBandwidthResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyIPv6TranslatorBandwidthResponse
	GetStatusCode() *int32
	SetBody(v *ModifyIPv6TranslatorBandwidthResponseBody) *ModifyIPv6TranslatorBandwidthResponse
	GetBody() *ModifyIPv6TranslatorBandwidthResponseBody
}

type ModifyIPv6TranslatorBandwidthResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyIPv6TranslatorBandwidthResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyIPv6TranslatorBandwidthResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyIPv6TranslatorBandwidthResponse) GoString() string {
	return s.String()
}

func (s *ModifyIPv6TranslatorBandwidthResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyIPv6TranslatorBandwidthResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyIPv6TranslatorBandwidthResponse) GetBody() *ModifyIPv6TranslatorBandwidthResponseBody {
	return s.Body
}

func (s *ModifyIPv6TranslatorBandwidthResponse) SetHeaders(v map[string]*string) *ModifyIPv6TranslatorBandwidthResponse {
	s.Headers = v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthResponse) SetStatusCode(v int32) *ModifyIPv6TranslatorBandwidthResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthResponse) SetBody(v *ModifyIPv6TranslatorBandwidthResponseBody) *ModifyIPv6TranslatorBandwidthResponse {
	s.Body = v
	return s
}

func (s *ModifyIPv6TranslatorBandwidthResponse) Validate() error {
	return dara.Validate(s)
}
