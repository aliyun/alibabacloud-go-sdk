// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrGetVirtualLicenseOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrGetVirtualLicenseOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrGetVirtualLicenseOrderResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrGetVirtualLicenseOrderResponseBody) *CreateOrGetVirtualLicenseOrderResponse
	GetBody() *CreateOrGetVirtualLicenseOrderResponseBody
}

type CreateOrGetVirtualLicenseOrderResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrGetVirtualLicenseOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrGetVirtualLicenseOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrGetVirtualLicenseOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrGetVirtualLicenseOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrGetVirtualLicenseOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrGetVirtualLicenseOrderResponse) GetBody() *CreateOrGetVirtualLicenseOrderResponseBody {
	return s.Body
}

func (s *CreateOrGetVirtualLicenseOrderResponse) SetHeaders(v map[string]*string) *CreateOrGetVirtualLicenseOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponse) SetStatusCode(v int32) *CreateOrGetVirtualLicenseOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponse) SetBody(v *CreateOrGetVirtualLicenseOrderResponseBody) *CreateOrGetVirtualLicenseOrderResponse {
	s.Body = v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponse) Validate() error {
	return dara.Validate(s)
}
