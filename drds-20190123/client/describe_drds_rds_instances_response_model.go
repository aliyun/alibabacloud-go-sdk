// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsRdsInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsRdsInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsRdsInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsRdsInstancesResponseBody) *DescribeDrdsRdsInstancesResponse
	GetBody() *DescribeDrdsRdsInstancesResponseBody
}

type DescribeDrdsRdsInstancesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsRdsInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsRdsInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsRdsInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsRdsInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsRdsInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsRdsInstancesResponse) GetBody() *DescribeDrdsRdsInstancesResponseBody {
	return s.Body
}

func (s *DescribeDrdsRdsInstancesResponse) SetHeaders(v map[string]*string) *DescribeDrdsRdsInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsRdsInstancesResponse) SetStatusCode(v int32) *DescribeDrdsRdsInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsRdsInstancesResponse) SetBody(v *DescribeDrdsRdsInstancesResponseBody) *DescribeDrdsRdsInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsRdsInstancesResponse) Validate() error {
	return dara.Validate(s)
}
