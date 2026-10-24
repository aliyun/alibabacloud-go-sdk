// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetScopeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetScopeResponseBody) *DescribeAssetScopeResponse
	GetBody() *DescribeAssetScopeResponseBody
}

type DescribeAssetScopeResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetScopeResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetScopeResponse) GetBody() *DescribeAssetScopeResponseBody {
	return s.Body
}

func (s *DescribeAssetScopeResponse) SetHeaders(v map[string]*string) *DescribeAssetScopeResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetScopeResponse) SetStatusCode(v int32) *DescribeAssetScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetScopeResponse) SetBody(v *DescribeAssetScopeResponseBody) *DescribeAssetScopeResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
