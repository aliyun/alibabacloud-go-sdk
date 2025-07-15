// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainLimitResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainLimitResponseBody) *DescribeLiveDomainLimitResponse
	GetBody() *DescribeLiveDomainLimitResponseBody
}

type DescribeLiveDomainLimitResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLimitResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainLimitResponse) GetBody() *DescribeLiveDomainLimitResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainLimitResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainLimitResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainLimitResponse) SetStatusCode(v int32) *DescribeLiveDomainLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainLimitResponse) SetBody(v *DescribeLiveDomainLimitResponseBody) *DescribeLiveDomainLimitResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainLimitResponse) Validate() error {
	return dara.Validate(s)
}
