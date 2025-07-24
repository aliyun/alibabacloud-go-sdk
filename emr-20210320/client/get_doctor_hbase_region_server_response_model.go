// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseRegionServerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorHBaseRegionServerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorHBaseRegionServerResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorHBaseRegionServerResponseBody) *GetDoctorHBaseRegionServerResponse
	GetBody() *GetDoctorHBaseRegionServerResponseBody
}

type GetDoctorHBaseRegionServerResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorHBaseRegionServerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorHBaseRegionServerResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseRegionServerResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseRegionServerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorHBaseRegionServerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorHBaseRegionServerResponse) GetBody() *GetDoctorHBaseRegionServerResponseBody {
	return s.Body
}

func (s *GetDoctorHBaseRegionServerResponse) SetHeaders(v map[string]*string) *GetDoctorHBaseRegionServerResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponse) SetStatusCode(v int32) *GetDoctorHBaseRegionServerResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorHBaseRegionServerResponse) SetBody(v *GetDoctorHBaseRegionServerResponseBody) *GetDoctorHBaseRegionServerResponse {
	s.Body = v
	return s
}

func (s *GetDoctorHBaseRegionServerResponse) Validate() error {
	return dara.Validate(s)
}
