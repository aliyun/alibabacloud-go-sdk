// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhiteListEffectiveAssetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWhiteListEffectiveAssetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWhiteListEffectiveAssetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeWhiteListEffectiveAssetsResponseBody) *DescribeWhiteListEffectiveAssetsResponse
	GetBody() *DescribeWhiteListEffectiveAssetsResponseBody
}

type DescribeWhiteListEffectiveAssetsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWhiteListEffectiveAssetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWhiteListEffectiveAssetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhiteListEffectiveAssetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeWhiteListEffectiveAssetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWhiteListEffectiveAssetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWhiteListEffectiveAssetsResponse) GetBody() *DescribeWhiteListEffectiveAssetsResponseBody {
	return s.Body
}

func (s *DescribeWhiteListEffectiveAssetsResponse) SetHeaders(v map[string]*string) *DescribeWhiteListEffectiveAssetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponse) SetStatusCode(v int32) *DescribeWhiteListEffectiveAssetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponse) SetBody(v *DescribeWhiteListEffectiveAssetsResponseBody) *DescribeWhiteListEffectiveAssetsResponse {
	s.Body = v
	return s
}

func (s *DescribeWhiteListEffectiveAssetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
