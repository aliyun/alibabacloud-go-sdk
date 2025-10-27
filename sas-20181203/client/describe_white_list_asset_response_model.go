// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListAssetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhiteListAssetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhiteListAssetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhiteListAssetResponseBody) *DescribeWhiteListAssetResponse
	GetBody() *DescribeWhiteListAssetResponseBody
}

type DescribeWhiteListAssetResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhiteListAssetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhiteListAssetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListAssetResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListAssetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhiteListAssetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhiteListAssetResponse) GetBody() *DescribeWhiteListAssetResponseBody {
	return s.Body
}

func (s *DescribeWhiteListAssetResponse) SetHeaders(v map[string]*string) *DescribeWhiteListAssetResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteListAssetResponse) SetStatusCode(v int32) *DescribeWhiteListAssetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhiteListAssetResponse) SetBody(v *DescribeWhiteListAssetResponseBody) *DescribeWhiteListAssetResponse {
	s.Body = v
	return s
}

func (s *DescribeWhiteListAssetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
