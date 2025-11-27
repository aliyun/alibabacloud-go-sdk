// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDedicatedHostGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDedicatedHostGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDedicatedHostGroupsResponseBody) *DescribeDedicatedHostGroupsResponse
	GetBody() *DescribeDedicatedHostGroupsResponseBody
}

type DescribeDedicatedHostGroupsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDedicatedHostGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDedicatedHostGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDedicatedHostGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDedicatedHostGroupsResponse) GetBody() *DescribeDedicatedHostGroupsResponseBody {
	return s.Body
}

func (s *DescribeDedicatedHostGroupsResponse) SetHeaders(v map[string]*string) *DescribeDedicatedHostGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponse) SetStatusCode(v int32) *DescribeDedicatedHostGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDedicatedHostGroupsResponse) SetBody(v *DescribeDedicatedHostGroupsResponseBody) *DescribeDedicatedHostGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDedicatedHostGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
