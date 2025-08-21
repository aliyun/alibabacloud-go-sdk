// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainOverviewResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainOverviewResponseBody) *DescribeDomainOverviewResponse
	GetBody() *DescribeDomainOverviewResponseBody
}

type DescribeDomainOverviewResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainOverviewResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainOverviewResponse) GetBody() *DescribeDomainOverviewResponseBody {
	return s.Body
}

func (s *DescribeDomainOverviewResponse) SetHeaders(v map[string]*string) *DescribeDomainOverviewResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainOverviewResponse) SetStatusCode(v int32) *DescribeDomainOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainOverviewResponse) SetBody(v *DescribeDomainOverviewResponseBody) *DescribeDomainOverviewResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainOverviewResponse) Validate() error {
	return dara.Validate(s)
}
