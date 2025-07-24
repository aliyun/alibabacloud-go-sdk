// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHiveClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorHiveClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorHiveClusterResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorHiveClusterResponseBody) *GetDoctorHiveClusterResponse
	GetBody() *GetDoctorHiveClusterResponseBody
}

type GetDoctorHiveClusterResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorHiveClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorHiveClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveClusterResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorHiveClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorHiveClusterResponse) GetBody() *GetDoctorHiveClusterResponseBody {
	return s.Body
}

func (s *GetDoctorHiveClusterResponse) SetHeaders(v map[string]*string) *GetDoctorHiveClusterResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorHiveClusterResponse) SetStatusCode(v int32) *GetDoctorHiveClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorHiveClusterResponse) SetBody(v *GetDoctorHiveClusterResponseBody) *GetDoctorHiveClusterResponse {
	s.Body = v
	return s
}

func (s *GetDoctorHiveClusterResponse) Validate() error {
	return dara.Validate(s)
}
