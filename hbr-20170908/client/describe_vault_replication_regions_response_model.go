// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVaultReplicationRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVaultReplicationRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVaultReplicationRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVaultReplicationRegionsResponseBody) *DescribeVaultReplicationRegionsResponse
	GetBody() *DescribeVaultReplicationRegionsResponseBody
}

type DescribeVaultReplicationRegionsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVaultReplicationRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVaultReplicationRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVaultReplicationRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVaultReplicationRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVaultReplicationRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVaultReplicationRegionsResponse) GetBody() *DescribeVaultReplicationRegionsResponseBody {
	return s.Body
}

func (s *DescribeVaultReplicationRegionsResponse) SetHeaders(v map[string]*string) *DescribeVaultReplicationRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVaultReplicationRegionsResponse) SetStatusCode(v int32) *DescribeVaultReplicationRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVaultReplicationRegionsResponse) SetBody(v *DescribeVaultReplicationRegionsResponseBody) *DescribeVaultReplicationRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeVaultReplicationRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
