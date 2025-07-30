// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedClusterInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedClusterInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedClusterInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedClusterInstanceListResponseBody) *DescribeDedicatedClusterInstanceListResponse
	GetBody() *DescribeDedicatedClusterInstanceListResponseBody
}

type DescribeDedicatedClusterInstanceListResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedClusterInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedClusterInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedClusterInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedClusterInstanceListResponse) GetBody() *DescribeDedicatedClusterInstanceListResponseBody {
	return s.Body
}

func (s *DescribeDedicatedClusterInstanceListResponse) SetHeaders(v map[string]*string) *DescribeDedicatedClusterInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponse) SetStatusCode(v int32) *DescribeDedicatedClusterInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponse) SetBody(v *DescribeDedicatedClusterInstanceListResponseBody) *DescribeDedicatedClusterInstanceListResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedClusterInstanceListResponse) Validate() error {
	return dara.Validate(s)
}
