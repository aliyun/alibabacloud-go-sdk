// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceMeshUpgradeStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeServiceMeshUpgradeStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeServiceMeshUpgradeStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeServiceMeshUpgradeStatusResponseBody) *DescribeServiceMeshUpgradeStatusResponse
	GetBody() *DescribeServiceMeshUpgradeStatusResponseBody
}

type DescribeServiceMeshUpgradeStatusResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeServiceMeshUpgradeStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeServiceMeshUpgradeStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceMeshUpgradeStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeServiceMeshUpgradeStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeServiceMeshUpgradeStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeServiceMeshUpgradeStatusResponse) GetBody() *DescribeServiceMeshUpgradeStatusResponseBody {
	return s.Body
}

func (s *DescribeServiceMeshUpgradeStatusResponse) SetHeaders(v map[string]*string) *DescribeServiceMeshUpgradeStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponse) SetStatusCode(v int32) *DescribeServiceMeshUpgradeStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponse) SetBody(v *DescribeServiceMeshUpgradeStatusResponseBody) *DescribeServiceMeshUpgradeStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeServiceMeshUpgradeStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
