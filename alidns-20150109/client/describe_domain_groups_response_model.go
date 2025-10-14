// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainGroupsResponseBody) *DescribeDomainGroupsResponse
	GetBody() *DescribeDomainGroupsResponseBody
}

type DescribeDomainGroupsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainGroupsResponse) GetBody() *DescribeDomainGroupsResponseBody {
	return s.Body
}

func (s *DescribeDomainGroupsResponse) SetHeaders(v map[string]*string) *DescribeDomainGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainGroupsResponse) SetStatusCode(v int32) *DescribeDomainGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainGroupsResponse) SetBody(v *DescribeDomainGroupsResponseBody) *DescribeDomainGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
