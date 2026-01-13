// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerModuleOutputInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomerModuleOutputInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomerModuleOutputInfoResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomerModuleOutputInfoResponseBody) *CreateCustomerModuleOutputInfoResponse
	GetBody() *CreateCustomerModuleOutputInfoResponseBody
}

type CreateCustomerModuleOutputInfoResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomerModuleOutputInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomerModuleOutputInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerModuleOutputInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerModuleOutputInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomerModuleOutputInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomerModuleOutputInfoResponse) GetBody() *CreateCustomerModuleOutputInfoResponseBody {
	return s.Body
}

func (s *CreateCustomerModuleOutputInfoResponse) SetHeaders(v map[string]*string) *CreateCustomerModuleOutputInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerModuleOutputInfoResponse) SetStatusCode(v int32) *CreateCustomerModuleOutputInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomerModuleOutputInfoResponse) SetBody(v *CreateCustomerModuleOutputInfoResponseBody) *CreateCustomerModuleOutputInfoResponse {
	s.Body = v
	return s
}

func (s *CreateCustomerModuleOutputInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
