// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorHBaseClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorHBaseClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorHBaseClusterResponseBody) *GetDoctorHBaseClusterResponse
	GetBody() *GetDoctorHBaseClusterResponseBody
}

type GetDoctorHBaseClusterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorHBaseClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorHBaseClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseClusterResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorHBaseClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorHBaseClusterResponse) GetBody() *GetDoctorHBaseClusterResponseBody {
	return s.Body
}

func (s *GetDoctorHBaseClusterResponse) SetHeaders(v map[string]*string) *GetDoctorHBaseClusterResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorHBaseClusterResponse) SetStatusCode(v int32) *GetDoctorHBaseClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorHBaseClusterResponse) SetBody(v *GetDoctorHBaseClusterResponseBody) *GetDoctorHBaseClusterResponse {
	s.Body = v
	return s
}

func (s *GetDoctorHBaseClusterResponse) Validate() error {
	return dara.Validate(s)
}
