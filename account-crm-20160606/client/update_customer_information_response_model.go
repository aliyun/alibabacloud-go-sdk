// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomerInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomerInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomerInformationResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomerInformationResponseBody) *UpdateCustomerInformationResponse
	GetBody() *UpdateCustomerInformationResponseBody
}

type UpdateCustomerInformationResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomerInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomerInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomerInformationResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomerInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomerInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomerInformationResponse) GetBody() *UpdateCustomerInformationResponseBody {
	return s.Body
}

func (s *UpdateCustomerInformationResponse) SetHeaders(v map[string]*string) *UpdateCustomerInformationResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomerInformationResponse) SetStatusCode(v int32) *UpdateCustomerInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomerInformationResponse) SetBody(v *UpdateCustomerInformationResponseBody) *UpdateCustomerInformationResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomerInformationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
