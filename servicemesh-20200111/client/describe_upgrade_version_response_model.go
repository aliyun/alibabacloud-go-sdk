// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpgradeVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUpgradeVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUpgradeVersionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUpgradeVersionResponseBody) *DescribeUpgradeVersionResponse
	GetBody() *DescribeUpgradeVersionResponseBody
}

type DescribeUpgradeVersionResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpgradeVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpgradeVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpgradeVersionResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpgradeVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUpgradeVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUpgradeVersionResponse) GetBody() *DescribeUpgradeVersionResponseBody {
	return s.Body
}

func (s *DescribeUpgradeVersionResponse) SetHeaders(v map[string]*string) *DescribeUpgradeVersionResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpgradeVersionResponse) SetStatusCode(v int32) *DescribeUpgradeVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpgradeVersionResponse) SetBody(v *DescribeUpgradeVersionResponseBody) *DescribeUpgradeVersionResponse {
	s.Body = v
	return s
}

func (s *DescribeUpgradeVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
