// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainTrafficDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainTrafficDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainTrafficDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainTrafficDataResponseBody) *DescribeLiveDomainTrafficDataResponse
	GetBody() *DescribeLiveDomainTrafficDataResponseBody
}

type DescribeLiveDomainTrafficDataResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainTrafficDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainTrafficDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainTrafficDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainTrafficDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainTrafficDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainTrafficDataResponse) GetBody() *DescribeLiveDomainTrafficDataResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainTrafficDataResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainTrafficDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponse) SetStatusCode(v int32) *DescribeLiveDomainTrafficDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponse) SetBody(v *DescribeLiveDomainTrafficDataResponseBody) *DescribeLiveDomainTrafficDataResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainTrafficDataResponse) Validate() error {
	return dara.Validate(s)
}
