// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCdnDomainStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCdnDomainStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCdnDomainStagingConfigResponseBody) *DescribeCdnDomainStagingConfigResponse
	GetBody() *DescribeCdnDomainStagingConfigResponseBody
}

type DescribeCdnDomainStagingConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCdnDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCdnDomainStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCdnDomainStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCdnDomainStagingConfigResponse) GetBody() *DescribeCdnDomainStagingConfigResponseBody {
	return s.Body
}

func (s *DescribeCdnDomainStagingConfigResponse) SetHeaders(v map[string]*string) *DescribeCdnDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponse) SetStatusCode(v int32) *DescribeCdnDomainStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponse) SetBody(v *DescribeCdnDomainStagingConfigResponseBody) *DescribeCdnDomainStagingConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
