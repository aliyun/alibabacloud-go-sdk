// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLicenseInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLicenseInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLicenseInfoResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLicenseInfoResponseBody) *ModifyLicenseInfoResponse
	GetBody() *ModifyLicenseInfoResponseBody
}

type ModifyLicenseInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLicenseInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLicenseInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLicenseInfoResponse) GoString() string {
	return s.String()
}

func (s *ModifyLicenseInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLicenseInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLicenseInfoResponse) GetBody() *ModifyLicenseInfoResponseBody {
	return s.Body
}

func (s *ModifyLicenseInfoResponse) SetHeaders(v map[string]*string) *ModifyLicenseInfoResponse {
	s.Headers = v
	return s
}

func (s *ModifyLicenseInfoResponse) SetStatusCode(v int32) *ModifyLicenseInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLicenseInfoResponse) SetBody(v *ModifyLicenseInfoResponseBody) *ModifyLicenseInfoResponse {
	s.Body = v
	return s
}

func (s *ModifyLicenseInfoResponse) Validate() error {
	return dara.Validate(s)
}
