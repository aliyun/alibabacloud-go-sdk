// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenAttachedChildInstancesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCenAttachedChildInstancesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCenAttachedChildInstancesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCenAttachedChildInstancesResponseBody) *DescribeCenAttachedChildInstancesResponse
	GetBody() *DescribeCenAttachedChildInstancesResponseBody
}

type DescribeCenAttachedChildInstancesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCenAttachedChildInstancesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCenAttachedChildInstancesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenAttachedChildInstancesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCenAttachedChildInstancesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCenAttachedChildInstancesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCenAttachedChildInstancesResponse) GetBody() *DescribeCenAttachedChildInstancesResponseBody {
	return s.Body
}

func (s *DescribeCenAttachedChildInstancesResponse) SetHeaders(v map[string]*string) *DescribeCenAttachedChildInstancesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponse) SetStatusCode(v int32) *DescribeCenAttachedChildInstancesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponse) SetBody(v *DescribeCenAttachedChildInstancesResponseBody) *DescribeCenAttachedChildInstancesResponse {
	s.Body = v
	return s
}

func (s *DescribeCenAttachedChildInstancesResponse) Validate() error {
	return dara.Validate(s)
}
