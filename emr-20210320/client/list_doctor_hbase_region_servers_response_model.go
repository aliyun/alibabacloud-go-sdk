// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDoctorHBaseRegionServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDoctorHBaseRegionServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDoctorHBaseRegionServersResponse
	GetStatusCode() *int32
	SetBody(v *ListDoctorHBaseRegionServersResponseBody) *ListDoctorHBaseRegionServersResponse
	GetBody() *ListDoctorHBaseRegionServersResponseBody
}

type ListDoctorHBaseRegionServersResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDoctorHBaseRegionServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDoctorHBaseRegionServersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDoctorHBaseRegionServersResponse) GoString() string {
	return s.String()
}

func (s *ListDoctorHBaseRegionServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDoctorHBaseRegionServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDoctorHBaseRegionServersResponse) GetBody() *ListDoctorHBaseRegionServersResponseBody {
	return s.Body
}

func (s *ListDoctorHBaseRegionServersResponse) SetHeaders(v map[string]*string) *ListDoctorHBaseRegionServersResponse {
	s.Headers = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponse) SetStatusCode(v int32) *ListDoctorHBaseRegionServersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDoctorHBaseRegionServersResponse) SetBody(v *ListDoctorHBaseRegionServersResponseBody) *ListDoctorHBaseRegionServersResponse {
	s.Body = v
	return s
}

func (s *ListDoctorHBaseRegionServersResponse) Validate() error {
	return dara.Validate(s)
}
