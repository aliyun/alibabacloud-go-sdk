// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAffectedAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAffectedAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAffectedAssetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAffectedAssetsResponseBody) *DescribeAffectedAssetsResponse
	GetBody() *DescribeAffectedAssetsResponseBody
}

type DescribeAffectedAssetsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAffectedAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAffectedAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAffectedAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeAffectedAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAffectedAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAffectedAssetsResponse) GetBody() *DescribeAffectedAssetsResponseBody {
	return s.Body
}

func (s *DescribeAffectedAssetsResponse) SetHeaders(v map[string]*string) *DescribeAffectedAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeAffectedAssetsResponse) SetStatusCode(v int32) *DescribeAffectedAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAffectedAssetsResponse) SetBody(v *DescribeAffectedAssetsResponseBody) *DescribeAffectedAssetsResponse {
	s.Body = v
	return s
}

func (s *DescribeAffectedAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
