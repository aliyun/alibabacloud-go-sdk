// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetGroupResponseBody) *DescribeAssetGroupResponse
	GetBody() *DescribeAssetGroupResponseBody
}

type DescribeAssetGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetGroupResponse) GetBody() *DescribeAssetGroupResponseBody {
	return s.Body
}

func (s *DescribeAssetGroupResponse) SetHeaders(v map[string]*string) *DescribeAssetGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetGroupResponse) SetStatusCode(v int32) *DescribeAssetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetGroupResponse) SetBody(v *DescribeAssetGroupResponseBody) *DescribeAssetGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
