// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnycastEipAddressAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAnycastEipAddressAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAnycastEipAddressAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAnycastEipAddressAttributeResponseBody) *ModifyAnycastEipAddressAttributeResponse
	GetBody() *ModifyAnycastEipAddressAttributeResponseBody
}

type ModifyAnycastEipAddressAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAnycastEipAddressAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAnycastEipAddressAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnycastEipAddressAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAnycastEipAddressAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAnycastEipAddressAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAnycastEipAddressAttributeResponse) GetBody() *ModifyAnycastEipAddressAttributeResponseBody {
	return s.Body
}

func (s *ModifyAnycastEipAddressAttributeResponse) SetHeaders(v map[string]*string) *ModifyAnycastEipAddressAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAnycastEipAddressAttributeResponse) SetStatusCode(v int32) *ModifyAnycastEipAddressAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAnycastEipAddressAttributeResponse) SetBody(v *ModifyAnycastEipAddressAttributeResponseBody) *ModifyAnycastEipAddressAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyAnycastEipAddressAttributeResponse) Validate() error {
	return dara.Validate(s)
}
