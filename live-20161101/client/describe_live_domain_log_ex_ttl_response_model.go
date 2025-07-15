// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainLogExTtlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainLogExTtlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainLogExTtlResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainLogExTtlResponseBody) *DescribeLiveDomainLogExTtlResponse
	GetBody() *DescribeLiveDomainLogExTtlResponseBody
}

type DescribeLiveDomainLogExTtlResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainLogExTtlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainLogExTtlResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainLogExTtlResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainLogExTtlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainLogExTtlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainLogExTtlResponse) GetBody() *DescribeLiveDomainLogExTtlResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainLogExTtlResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainLogExTtlResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponse) SetStatusCode(v int32) *DescribeLiveDomainLogExTtlResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponse) SetBody(v *DescribeLiveDomainLogExTtlResponseBody) *DescribeLiveDomainLogExTtlResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainLogExTtlResponse) Validate() error {
	return dara.Validate(s)
}
