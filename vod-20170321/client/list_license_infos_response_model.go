// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLicenseInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLicenseInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLicenseInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListLicenseInfosResponseBody) *ListLicenseInfosResponse
	GetBody() *ListLicenseInfosResponseBody
}

type ListLicenseInfosResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLicenseInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLicenseInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLicenseInfosResponse) GoString() string {
	return s.String()
}

func (s *ListLicenseInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLicenseInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLicenseInfosResponse) GetBody() *ListLicenseInfosResponseBody {
	return s.Body
}

func (s *ListLicenseInfosResponse) SetHeaders(v map[string]*string) *ListLicenseInfosResponse {
	s.Headers = v
	return s
}

func (s *ListLicenseInfosResponse) SetStatusCode(v int32) *ListLicenseInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLicenseInfosResponse) SetBody(v *ListLicenseInfosResponseBody) *ListLicenseInfosResponse {
	s.Body = v
	return s
}

func (s *ListLicenseInfosResponse) Validate() error {
	return dara.Validate(s)
}
