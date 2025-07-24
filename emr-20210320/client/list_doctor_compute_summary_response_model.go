// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorComputeSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorComputeSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorComputeSummaryResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorComputeSummaryResponseBody) *ListDoctorComputeSummaryResponse
	GetBody() *ListDoctorComputeSummaryResponseBody
}

type ListDoctorComputeSummaryResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorComputeSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorComputeSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorComputeSummaryResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorComputeSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorComputeSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorComputeSummaryResponse) GetBody() *ListDoctorComputeSummaryResponseBody {
	return s.Body
}

func (s *ListDoctorComputeSummaryResponse) SetHeaders(v map[string]*string) *ListDoctorComputeSummaryResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorComputeSummaryResponse) SetStatusCode(v int32) *ListDoctorComputeSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorComputeSummaryResponse) SetBody(v *ListDoctorComputeSummaryResponseBody) *ListDoctorComputeSummaryResponse {
	s.Body = v
	return s
}

func (s *ListDoctorComputeSummaryResponse) Validate() error {
	return dara.Validate(s)
}
