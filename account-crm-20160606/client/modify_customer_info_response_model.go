// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCustomerInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyCustomerInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyCustomerInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyCustomerInfoResponseBody) *ModifyCustomerInfoResponse
	GetBody() *ModifyCustomerInfoResponseBody
}

type ModifyCustomerInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyCustomerInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyCustomerInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyCustomerInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyCustomerInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyCustomerInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyCustomerInfoResponse) GetBody() *ModifyCustomerInfoResponseBody {
	return s.Body
}

func (s *ModifyCustomerInfoResponse) SetHeaders(v map[string]*string) *ModifyCustomerInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyCustomerInfoResponse) SetStatusCode(v int32) *ModifyCustomerInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyCustomerInfoResponse) SetBody(v *ModifyCustomerInfoResponseBody) *ModifyCustomerInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyCustomerInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
