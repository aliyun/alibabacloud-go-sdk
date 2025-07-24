// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorComputeSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorComputeSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorComputeSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorComputeSummaryResponseBody) *GetDoctorComputeSummaryResponse
	GetBody() *GetDoctorComputeSummaryResponseBody
}

type GetDoctorComputeSummaryResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorComputeSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorComputeSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorComputeSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorComputeSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorComputeSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorComputeSummaryResponse) GetBody() *GetDoctorComputeSummaryResponseBody {
	return s.Body
}

func (s *GetDoctorComputeSummaryResponse) SetHeaders(v map[string]*string) *GetDoctorComputeSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorComputeSummaryResponse) SetStatusCode(v int32) *GetDoctorComputeSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorComputeSummaryResponse) SetBody(v *GetDoctorComputeSummaryResponseBody) *GetDoctorComputeSummaryResponse {
	s.Body = v
	return s
}

func (s *GetDoctorComputeSummaryResponse) Validate() error {
	return dara.Validate(s)
}
