// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodDomainStagingConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodDomainStagingConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodDomainStagingConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodDomainStagingConfigResponseBody) *DescribeVodDomainStagingConfigResponse
	GetBody() *DescribeVodDomainStagingConfigResponseBody
}

type DescribeVodDomainStagingConfigResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodDomainStagingConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodDomainStagingConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodDomainStagingConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodDomainStagingConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodDomainStagingConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodDomainStagingConfigResponse) GetBody() *DescribeVodDomainStagingConfigResponseBody {
	return s.Body
}

func (s *DescribeVodDomainStagingConfigResponse) SetHeaders(v map[string]*string) *DescribeVodDomainStagingConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodDomainStagingConfigResponse) SetStatusCode(v int32) *DescribeVodDomainStagingConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodDomainStagingConfigResponse) SetBody(v *DescribeVodDomainStagingConfigResponseBody) *DescribeVodDomainStagingConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeVodDomainStagingConfigResponse) Validate() error {
	return dara.Validate(s)
}
