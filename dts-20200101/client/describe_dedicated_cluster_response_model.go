// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedClusterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedClusterResponseBody) *DescribeDedicatedClusterResponse
	GetBody() *DescribeDedicatedClusterResponseBody
}

type DescribeDedicatedClusterResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedClusterResponse) GetBody() *DescribeDedicatedClusterResponseBody {
	return s.Body
}

func (s *DescribeDedicatedClusterResponse) SetHeaders(v map[string]*string) *DescribeDedicatedClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedClusterResponse) SetStatusCode(v int32) *DescribeDedicatedClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedClusterResponse) SetBody(v *DescribeDedicatedClusterResponseBody) *DescribeDedicatedClusterResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedClusterResponse) Validate() error {
	return dara.Validate(s)
}
