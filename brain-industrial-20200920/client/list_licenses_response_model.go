// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLicensesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLicensesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLicensesResponse
	GetStatusCode() *int32
	SetBody(v *ListLicensesResponseBody) *ListLicensesResponse
	GetBody() *ListLicensesResponseBody
}

type ListLicensesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLicensesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLicensesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLicensesResponse) GoString() string {
	return s.String()
}

func (s *ListLicensesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLicensesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLicensesResponse) GetBody() *ListLicensesResponseBody {
	return s.Body
}

func (s *ListLicensesResponse) SetHeaders(v map[string]*string) *ListLicensesResponse {
	s.Headers = v
	return s
}

func (s *ListLicensesResponse) SetStatusCode(v int32) *ListLicensesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLicensesResponse) SetBody(v *ListLicensesResponseBody) *ListLicensesResponse {
	s.Body = v
	return s
}

func (s *ListLicensesResponse) Validate() error {
	return dara.Validate(s)
}
