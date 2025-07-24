// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseRegionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorHBaseRegionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorHBaseRegionResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorHBaseRegionResponseBody) *GetDoctorHBaseRegionResponse
	GetBody() *GetDoctorHBaseRegionResponseBody
}

type GetDoctorHBaseRegionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorHBaseRegionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorHBaseRegionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorHBaseRegionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorHBaseRegionResponse) GetBody() *GetDoctorHBaseRegionResponseBody {
	return s.Body
}

func (s *GetDoctorHBaseRegionResponse) SetHeaders(v map[string]*string) *GetDoctorHBaseRegionResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorHBaseRegionResponse) SetStatusCode(v int32) *GetDoctorHBaseRegionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorHBaseRegionResponse) SetBody(v *GetDoctorHBaseRegionResponseBody) *GetDoctorHBaseRegionResponse {
	s.Body = v
	return s
}

func (s *GetDoctorHBaseRegionResponse) Validate() error {
	return dara.Validate(s)
}
