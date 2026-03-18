// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAssetGroupToInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAssetGroupToInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAssetGroupToInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAssetGroupToInstanceResponseBody) *DescribeAssetGroupToInstanceResponse
	GetBody() *DescribeAssetGroupToInstanceResponseBody
}

type DescribeAssetGroupToInstanceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAssetGroupToInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAssetGroupToInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAssetGroupToInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAssetGroupToInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAssetGroupToInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAssetGroupToInstanceResponse) GetBody() *DescribeAssetGroupToInstanceResponseBody {
	return s.Body
}

func (s *DescribeAssetGroupToInstanceResponse) SetHeaders(v map[string]*string) *DescribeAssetGroupToInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAssetGroupToInstanceResponse) SetStatusCode(v int32) *DescribeAssetGroupToInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAssetGroupToInstanceResponse) SetBody(v *DescribeAssetGroupToInstanceResponseBody) *DescribeAssetGroupToInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeAssetGroupToInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
