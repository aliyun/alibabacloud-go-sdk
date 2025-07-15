// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEntitlementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEntitlementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEntitlementResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEntitlementResponseBody) *ModifyEntitlementResponse
	GetBody() *ModifyEntitlementResponseBody
}

type ModifyEntitlementResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEntitlementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEntitlementResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEntitlementResponse) GoString() string {
	return s.String()
}

func (s *ModifyEntitlementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEntitlementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEntitlementResponse) GetBody() *ModifyEntitlementResponseBody {
	return s.Body
}

func (s *ModifyEntitlementResponse) SetHeaders(v map[string]*string) *ModifyEntitlementResponse {
	s.Headers = v
	return s
}

func (s *ModifyEntitlementResponse) SetStatusCode(v int32) *ModifyEntitlementResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEntitlementResponse) SetBody(v *ModifyEntitlementResponseBody) *ModifyEntitlementResponse {
	s.Body = v
	return s
}

func (s *ModifyEntitlementResponse) Validate() error {
	return dara.Validate(s)
}
