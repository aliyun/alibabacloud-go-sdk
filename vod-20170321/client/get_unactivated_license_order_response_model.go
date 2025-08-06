// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUnactivatedLicenseOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUnactivatedLicenseOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUnactivatedLicenseOrderResponse
	GetStatusCode() *int32
	SetBody(v *GetUnactivatedLicenseOrderResponseBody) *GetUnactivatedLicenseOrderResponse
	GetBody() *GetUnactivatedLicenseOrderResponseBody
}

type GetUnactivatedLicenseOrderResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUnactivatedLicenseOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUnactivatedLicenseOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUnactivatedLicenseOrderResponse) GoString() string {
	return s.String()
}

func (s *GetUnactivatedLicenseOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUnactivatedLicenseOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUnactivatedLicenseOrderResponse) GetBody() *GetUnactivatedLicenseOrderResponseBody {
	return s.Body
}

func (s *GetUnactivatedLicenseOrderResponse) SetHeaders(v map[string]*string) *GetUnactivatedLicenseOrderResponse {
	s.Headers = v
	return s
}

func (s *GetUnactivatedLicenseOrderResponse) SetStatusCode(v int32) *GetUnactivatedLicenseOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUnactivatedLicenseOrderResponse) SetBody(v *GetUnactivatedLicenseOrderResponseBody) *GetUnactivatedLicenseOrderResponse {
	s.Body = v
	return s
}

func (s *GetUnactivatedLicenseOrderResponse) Validate() error {
	return dara.Validate(s)
}
