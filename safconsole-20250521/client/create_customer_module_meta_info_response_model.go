// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomerModuleMetaInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomerModuleMetaInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomerModuleMetaInfoResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomerModuleMetaInfoResponseBody) *CreateCustomerModuleMetaInfoResponse
	GetBody() *CreateCustomerModuleMetaInfoResponseBody
}

type CreateCustomerModuleMetaInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomerModuleMetaInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomerModuleMetaInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomerModuleMetaInfoResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomerModuleMetaInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomerModuleMetaInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomerModuleMetaInfoResponse) GetBody() *CreateCustomerModuleMetaInfoResponseBody {
	return s.Body
}

func (s *CreateCustomerModuleMetaInfoResponse) SetHeaders(v map[string]*string) *CreateCustomerModuleMetaInfoResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomerModuleMetaInfoResponse) SetStatusCode(v int32) *CreateCustomerModuleMetaInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomerModuleMetaInfoResponse) SetBody(v *CreateCustomerModuleMetaInfoResponseBody) *CreateCustomerModuleMetaInfoResponse {
	s.Body = v
	return s
}

func (s *CreateCustomerModuleMetaInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
