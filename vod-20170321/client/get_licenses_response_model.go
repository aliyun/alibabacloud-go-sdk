// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLicensesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLicensesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLicensesResponse
	GetStatusCode() *int32
	SetBody(v *GetLicensesResponseBody) *GetLicensesResponse
	GetBody() *GetLicensesResponseBody
}

type GetLicensesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLicensesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLicensesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLicensesResponse) GoString() string {
	return s.String()
}

func (s *GetLicensesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLicensesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLicensesResponse) GetBody() *GetLicensesResponseBody {
	return s.Body
}

func (s *GetLicensesResponse) SetHeaders(v map[string]*string) *GetLicensesResponse {
	s.Headers = v
	return s
}

func (s *GetLicensesResponse) SetStatusCode(v int32) *GetLicensesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLicensesResponse) SetBody(v *GetLicensesResponseBody) *GetLicensesResponse {
	s.Body = v
	return s
}

func (s *GetLicensesResponse) Validate() error {
	return dara.Validate(s)
}
