// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonMetadataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterAddonMetadataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterAddonMetadataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterAddonMetadataResponseBody) *DescribeClusterAddonMetadataResponse
	GetBody() *DescribeClusterAddonMetadataResponseBody
}

type DescribeClusterAddonMetadataResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterAddonMetadataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterAddonMetadataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonMetadataResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonMetadataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterAddonMetadataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterAddonMetadataResponse) GetBody() *DescribeClusterAddonMetadataResponseBody {
	return s.Body
}

func (s *DescribeClusterAddonMetadataResponse) SetHeaders(v map[string]*string) *DescribeClusterAddonMetadataResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterAddonMetadataResponse) SetStatusCode(v int32) *DescribeClusterAddonMetadataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterAddonMetadataResponse) SetBody(v *DescribeClusterAddonMetadataResponseBody) *DescribeClusterAddonMetadataResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterAddonMetadataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
