// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerModuleBasicInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomerModuleBasicInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomerModuleBasicInfoResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomerModuleBasicInfoResponseBody) *CreateCustomerModuleBasicInfoResponse
	GetBody() *CreateCustomerModuleBasicInfoResponseBody
}

type CreateCustomerModuleBasicInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomerModuleBasicInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomerModuleBasicInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerModuleBasicInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerModuleBasicInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomerModuleBasicInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomerModuleBasicInfoResponse) GetBody() *CreateCustomerModuleBasicInfoResponseBody {
	return s.Body
}

func (s *CreateCustomerModuleBasicInfoResponse) SetHeaders(v map[string]*string) *CreateCustomerModuleBasicInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerModuleBasicInfoResponse) SetStatusCode(v int32) *CreateCustomerModuleBasicInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomerModuleBasicInfoResponse) SetBody(v *CreateCustomerModuleBasicInfoResponseBody) *CreateCustomerModuleBasicInfoResponse {
	s.Body = v
	return s
}

func (s *CreateCustomerModuleBasicInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
