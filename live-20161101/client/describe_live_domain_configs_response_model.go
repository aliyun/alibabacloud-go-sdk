// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainConfigsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainConfigsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainConfigsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainConfigsResponseBody) *DescribeLiveDomainConfigsResponse
	GetBody() *DescribeLiveDomainConfigsResponseBody
}

type DescribeLiveDomainConfigsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainConfigsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainConfigsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainConfigsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainConfigsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainConfigsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainConfigsResponse) GetBody() *DescribeLiveDomainConfigsResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainConfigsResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainConfigsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainConfigsResponse) SetStatusCode(v int32) *DescribeLiveDomainConfigsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainConfigsResponse) SetBody(v *DescribeLiveDomainConfigsResponseBody) *DescribeLiveDomainConfigsResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainConfigsResponse) Validate() error {
	return dara.Validate(s)
}
