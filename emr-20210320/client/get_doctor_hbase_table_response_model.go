// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDoctorHBaseTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDoctorHBaseTableResponse
	GetStatusCode() *int32
	SetBody(v *GetDoctorHBaseTableResponseBody) *GetDoctorHBaseTableResponse
	GetBody() *GetDoctorHBaseTableResponseBody
}

type GetDoctorHBaseTableResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDoctorHBaseTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDoctorHBaseTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableResponse) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDoctorHBaseTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDoctorHBaseTableResponse) GetBody() *GetDoctorHBaseTableResponseBody {
	return s.Body
}

func (s *GetDoctorHBaseTableResponse) SetHeaders(v map[string]*string) *GetDoctorHBaseTableResponse {
	s.Headers = v
	return s
}

func (s *GetDoctorHBaseTableResponse) SetStatusCode(v int32) *GetDoctorHBaseTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDoctorHBaseTableResponse) SetBody(v *GetDoctorHBaseTableResponseBody) *GetDoctorHBaseTableResponse {
	s.Body = v
	return s
}

func (s *GetDoctorHBaseTableResponse) Validate() error {
	return dara.Validate(s)
}
