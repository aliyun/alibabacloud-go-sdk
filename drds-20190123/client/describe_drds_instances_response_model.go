// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsInstancesResponseBody) *DescribeDrdsInstancesResponse
	GetBody() *DescribeDrdsInstancesResponseBody
}

type DescribeDrdsInstancesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsInstancesResponse) GetBody() *DescribeDrdsInstancesResponseBody {
	return s.Body
}

func (s *DescribeDrdsInstancesResponse) SetHeaders(v map[string]*string) *DescribeDrdsInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsInstancesResponse) SetStatusCode(v int32) *DescribeDrdsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsInstancesResponse) SetBody(v *DescribeDrdsInstancesResponseBody) *DescribeDrdsInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsInstancesResponse) Validate() error {
	return dara.Validate(s)
}
