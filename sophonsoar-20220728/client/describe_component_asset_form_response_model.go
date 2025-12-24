// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentAssetFormResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComponentAssetFormResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComponentAssetFormResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComponentAssetFormResponseBody) *DescribeComponentAssetFormResponse
	GetBody() *DescribeComponentAssetFormResponseBody
}

type DescribeComponentAssetFormResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentAssetFormResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentAssetFormResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentAssetFormResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentAssetFormResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComponentAssetFormResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComponentAssetFormResponse) GetBody() *DescribeComponentAssetFormResponseBody {
	return s.Body
}

func (s *DescribeComponentAssetFormResponse) SetHeaders(v map[string]*string) *DescribeComponentAssetFormResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentAssetFormResponse) SetStatusCode(v int32) *DescribeComponentAssetFormResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentAssetFormResponse) SetBody(v *DescribeComponentAssetFormResponseBody) *DescribeComponentAssetFormResponse {
	s.Body = v
	return s
}

func (s *DescribeComponentAssetFormResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
