// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComponentAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComponentAssetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComponentAssetsResponseBody) *DescribeComponentAssetsResponse
	GetBody() *DescribeComponentAssetsResponseBody
}

type DescribeComponentAssetsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComponentAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComponentAssetsResponse) GetBody() *DescribeComponentAssetsResponseBody {
	return s.Body
}

func (s *DescribeComponentAssetsResponse) SetHeaders(v map[string]*string) *DescribeComponentAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentAssetsResponse) SetStatusCode(v int32) *DescribeComponentAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentAssetsResponse) SetBody(v *DescribeComponentAssetsResponseBody) *DescribeComponentAssetsResponse {
	s.Body = v
	return s
}

func (s *DescribeComponentAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
