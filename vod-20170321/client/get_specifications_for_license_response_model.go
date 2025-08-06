// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSpecificationsForLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSpecificationsForLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSpecificationsForLicenseResponse
	GetStatusCode() *int32
	SetBody(v *GetSpecificationsForLicenseResponseBody) *GetSpecificationsForLicenseResponse
	GetBody() *GetSpecificationsForLicenseResponseBody
}

type GetSpecificationsForLicenseResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSpecificationsForLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSpecificationsForLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSpecificationsForLicenseResponse) GoString() string {
	return s.String()
}

func (s *GetSpecificationsForLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSpecificationsForLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSpecificationsForLicenseResponse) GetBody() *GetSpecificationsForLicenseResponseBody {
	return s.Body
}

func (s *GetSpecificationsForLicenseResponse) SetHeaders(v map[string]*string) *GetSpecificationsForLicenseResponse {
	s.Headers = v
	return s
}

func (s *GetSpecificationsForLicenseResponse) SetStatusCode(v int32) *GetSpecificationsForLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSpecificationsForLicenseResponse) SetBody(v *GetSpecificationsForLicenseResponseBody) *GetSpecificationsForLicenseResponse {
	s.Body = v
	return s
}

func (s *GetSpecificationsForLicenseResponse) Validate() error {
	return dara.Validate(s)
}
