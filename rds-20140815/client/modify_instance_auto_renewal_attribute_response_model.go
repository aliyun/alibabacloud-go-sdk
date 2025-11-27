// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoRenewalAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyInstanceAutoRenewalAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyInstanceAutoRenewalAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyInstanceAutoRenewalAttributeResponseBody) *ModifyInstanceAutoRenewalAttributeResponse
	GetBody() *ModifyInstanceAutoRenewalAttributeResponseBody
}

type ModifyInstanceAutoRenewalAttributeResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyInstanceAutoRenewalAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyInstanceAutoRenewalAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoRenewalAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoRenewalAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyInstanceAutoRenewalAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyInstanceAutoRenewalAttributeResponse) GetBody() *ModifyInstanceAutoRenewalAttributeResponseBody {
	return s.Body
}

func (s *ModifyInstanceAutoRenewalAttributeResponse) SetHeaders(v map[string]*string) *ModifyInstanceAutoRenewalAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeResponse) SetStatusCode(v int32) *ModifyInstanceAutoRenewalAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeResponse) SetBody(v *ModifyInstanceAutoRenewalAttributeResponseBody) *ModifyInstanceAutoRenewalAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyInstanceAutoRenewalAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
