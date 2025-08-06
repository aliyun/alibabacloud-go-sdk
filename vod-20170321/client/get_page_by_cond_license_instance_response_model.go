// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageByCondLicenseInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPageByCondLicenseInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPageByCondLicenseInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetPageByCondLicenseInstanceResponseBody) *GetPageByCondLicenseInstanceResponse
	GetBody() *GetPageByCondLicenseInstanceResponseBody
}

type GetPageByCondLicenseInstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPageByCondLicenseInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPageByCondLicenseInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPageByCondLicenseInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetPageByCondLicenseInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPageByCondLicenseInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPageByCondLicenseInstanceResponse) GetBody() *GetPageByCondLicenseInstanceResponseBody {
	return s.Body
}

func (s *GetPageByCondLicenseInstanceResponse) SetHeaders(v map[string]*string) *GetPageByCondLicenseInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetPageByCondLicenseInstanceResponse) SetStatusCode(v int32) *GetPageByCondLicenseInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPageByCondLicenseInstanceResponse) SetBody(v *GetPageByCondLicenseInstanceResponseBody) *GetPageByCondLicenseInstanceResponse {
	s.Body = v
	return s
}

func (s *GetPageByCondLicenseInstanceResponse) Validate() error {
	return dara.Validate(s)
}
