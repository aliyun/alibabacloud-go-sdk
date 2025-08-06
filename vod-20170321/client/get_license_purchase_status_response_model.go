// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicensePurchaseStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLicensePurchaseStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLicensePurchaseStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetLicensePurchaseStatusResponseBody) *GetLicensePurchaseStatusResponse
	GetBody() *GetLicensePurchaseStatusResponseBody
}

type GetLicensePurchaseStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLicensePurchaseStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLicensePurchaseStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLicensePurchaseStatusResponse) GoString() string {
	return s.String()
}

func (s *GetLicensePurchaseStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLicensePurchaseStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLicensePurchaseStatusResponse) GetBody() *GetLicensePurchaseStatusResponseBody {
	return s.Body
}

func (s *GetLicensePurchaseStatusResponse) SetHeaders(v map[string]*string) *GetLicensePurchaseStatusResponse {
	s.Headers = v
	return s
}

func (s *GetLicensePurchaseStatusResponse) SetStatusCode(v int32) *GetLicensePurchaseStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLicensePurchaseStatusResponse) SetBody(v *GetLicensePurchaseStatusResponseBody) *GetLicensePurchaseStatusResponse {
	s.Body = v
	return s
}

func (s *GetLicensePurchaseStatusResponse) Validate() error {
	return dara.Validate(s)
}
