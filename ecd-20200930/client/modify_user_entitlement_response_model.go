// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserEntitlementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyUserEntitlementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyUserEntitlementResponse
	GetStatusCode() *int32
	SetBody(v *ModifyUserEntitlementResponseBody) *ModifyUserEntitlementResponse
	GetBody() *ModifyUserEntitlementResponseBody
}

type ModifyUserEntitlementResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyUserEntitlementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyUserEntitlementResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserEntitlementResponse) GoString() string {
	return s.String()
}

func (s *ModifyUserEntitlementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyUserEntitlementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyUserEntitlementResponse) GetBody() *ModifyUserEntitlementResponseBody {
	return s.Body
}

func (s *ModifyUserEntitlementResponse) SetHeaders(v map[string]*string) *ModifyUserEntitlementResponse {
	s.Headers = v
	return s
}

func (s *ModifyUserEntitlementResponse) SetStatusCode(v int32) *ModifyUserEntitlementResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyUserEntitlementResponse) SetBody(v *ModifyUserEntitlementResponseBody) *ModifyUserEntitlementResponse {
	s.Body = v
	return s
}

func (s *ModifyUserEntitlementResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
