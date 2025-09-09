// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDrdsDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDrdsDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDrdsDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDrdsDBClusterResponseBody) *DescribeDrdsDBClusterResponse
	GetBody() *DescribeDrdsDBClusterResponseBody
}

type DescribeDrdsDBClusterResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDrdsDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDrdsDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDrdsDBClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeDrdsDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDrdsDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDrdsDBClusterResponse) GetBody() *DescribeDrdsDBClusterResponseBody {
	return s.Body
}

func (s *DescribeDrdsDBClusterResponse) SetHeaders(v map[string]*string) *DescribeDrdsDBClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeDrdsDBClusterResponse) SetStatusCode(v int32) *DescribeDrdsDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDrdsDBClusterResponse) SetBody(v *DescribeDrdsDBClusterResponseBody) *DescribeDrdsDBClusterResponse {
	s.Body = v
	return s
}

func (s *DescribeDrdsDBClusterResponse) Validate() error {
	return dara.Validate(s)
}
