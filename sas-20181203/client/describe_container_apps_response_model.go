// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContainerAppsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContainerAppsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContainerAppsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContainerAppsResponseBody) *DescribeContainerAppsResponse
	GetBody() *DescribeContainerAppsResponseBody
}

type DescribeContainerAppsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContainerAppsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContainerAppsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContainerAppsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContainerAppsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContainerAppsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContainerAppsResponse) GetBody() *DescribeContainerAppsResponseBody {
	return s.Body
}

func (s *DescribeContainerAppsResponse) SetHeaders(v map[string]*string) *DescribeContainerAppsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContainerAppsResponse) SetStatusCode(v int32) *DescribeContainerAppsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContainerAppsResponse) SetBody(v *DescribeContainerAppsResponseBody) *DescribeContainerAppsResponse {
	s.Body = v
	return s
}

func (s *DescribeContainerAppsResponse) Validate() error {
	return dara.Validate(s)
}
