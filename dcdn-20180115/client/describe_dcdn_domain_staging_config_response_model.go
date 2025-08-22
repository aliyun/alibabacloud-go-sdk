// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnDomainStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDcdnDomainStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDcdnDomainStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDcdnDomainStagingConfigResponseBody) *DescribeDcdnDomainStagingConfigResponse
	GetBody() *DescribeDcdnDomainStagingConfigResponseBody
}

type DescribeDcdnDomainStagingConfigResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDcdnDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDcdnDomainStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeDcdnDomainStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDcdnDomainStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDcdnDomainStagingConfigResponse) GetBody() *DescribeDcdnDomainStagingConfigResponseBody {
	return s.Body
}

func (s *DescribeDcdnDomainStagingConfigResponse) SetHeaders(v map[string]*string) *DescribeDcdnDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponse) SetStatusCode(v int32) *DescribeDcdnDomainStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponse) SetBody(v *DescribeDcdnDomainStagingConfigResponseBody) *DescribeDcdnDomainStagingConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeDcdnDomainStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
