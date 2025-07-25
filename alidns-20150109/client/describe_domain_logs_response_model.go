// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainLogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainLogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainLogsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainLogsResponseBody) *DescribeDomainLogsResponse
	GetBody() *DescribeDomainLogsResponseBody
}

type DescribeDomainLogsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainLogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainLogsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainLogsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainLogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainLogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainLogsResponse) GetBody() *DescribeDomainLogsResponseBody {
	return s.Body
}

func (s *DescribeDomainLogsResponse) SetHeaders(v map[string]*string) *DescribeDomainLogsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainLogsResponse) SetStatusCode(v int32) *DescribeDomainLogsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainLogsResponse) SetBody(v *DescribeDomainLogsResponseBody) *DescribeDomainLogsResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainLogsResponse) Validate() error {
	return dara.Validate(s)
}
