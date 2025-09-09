// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDbInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsDbInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsDbInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsDbInstancesResponseBody) *DescribeDrdsDbInstancesResponse
	GetBody() *DescribeDrdsDbInstancesResponseBody
}

type DescribeDrdsDbInstancesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDbInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDbInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDbInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDbInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsDbInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsDbInstancesResponse) GetBody() *DescribeDrdsDbInstancesResponseBody {
	return s.Body
}

func (s *DescribeDrdsDbInstancesResponse) SetHeaders(v map[string]*string) *DescribeDrdsDbInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDbInstancesResponse) SetStatusCode(v int32) *DescribeDrdsDbInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDbInstancesResponse) SetBody(v *DescribeDrdsDbInstancesResponseBody) *DescribeDrdsDbInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsDbInstancesResponse) Validate() error {
	return dara.Validate(s)
}
