// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainCustomLogConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDomainCustomLogConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDomainCustomLogConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDomainCustomLogConfigResponseBody) *DescribeDomainCustomLogConfigResponse
	GetBody() *DescribeDomainCustomLogConfigResponseBody
}

type DescribeDomainCustomLogConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDomainCustomLogConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDomainCustomLogConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainCustomLogConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDomainCustomLogConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDomainCustomLogConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDomainCustomLogConfigResponse) GetBody() *DescribeDomainCustomLogConfigResponseBody {
	return s.Body
}

func (s *DescribeDomainCustomLogConfigResponse) SetHeaders(v map[string]*string) *DescribeDomainCustomLogConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDomainCustomLogConfigResponse) SetStatusCode(v int32) *DescribeDomainCustomLogConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDomainCustomLogConfigResponse) SetBody(v *DescribeDomainCustomLogConfigResponseBody) *DescribeDomainCustomLogConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDomainCustomLogConfigResponse) Validate() error {
	return dara.Validate(s)
}
