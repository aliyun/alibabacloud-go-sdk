// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoRenewalAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAutoRenewalAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAutoRenewalAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAutoRenewalAttributeResponseBody) *ModifyAutoRenewalAttributeResponse
	GetBody() *ModifyAutoRenewalAttributeResponseBody
}

type ModifyAutoRenewalAttributeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAutoRenewalAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAutoRenewalAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoRenewalAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyAutoRenewalAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAutoRenewalAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAutoRenewalAttributeResponse) GetBody() *ModifyAutoRenewalAttributeResponseBody {
	return s.Body
}

func (s *ModifyAutoRenewalAttributeResponse) SetHeaders(v map[string]*string) *ModifyAutoRenewalAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyAutoRenewalAttributeResponse) SetStatusCode(v int32) *ModifyAutoRenewalAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAutoRenewalAttributeResponse) SetBody(v *ModifyAutoRenewalAttributeResponseBody) *ModifyAutoRenewalAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyAutoRenewalAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
