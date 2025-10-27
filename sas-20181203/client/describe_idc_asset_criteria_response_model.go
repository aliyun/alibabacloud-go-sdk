// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIdcAssetCriteriaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIdcAssetCriteriaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIdcAssetCriteriaResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIdcAssetCriteriaResponseBody) *DescribeIdcAssetCriteriaResponse
	GetBody() *DescribeIdcAssetCriteriaResponseBody
}

type DescribeIdcAssetCriteriaResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIdcAssetCriteriaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIdcAssetCriteriaResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIdcAssetCriteriaResponse) GoString() string {
	return s.String()
}

func (s *DescribeIdcAssetCriteriaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIdcAssetCriteriaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIdcAssetCriteriaResponse) GetBody() *DescribeIdcAssetCriteriaResponseBody {
	return s.Body
}

func (s *DescribeIdcAssetCriteriaResponse) SetHeaders(v map[string]*string) *DescribeIdcAssetCriteriaResponse {
	s.Headers = v
	return s
}

func (s *DescribeIdcAssetCriteriaResponse) SetStatusCode(v int32) *DescribeIdcAssetCriteriaResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIdcAssetCriteriaResponse) SetBody(v *DescribeIdcAssetCriteriaResponseBody) *DescribeIdcAssetCriteriaResponse {
	s.Body = v
	return s
}

func (s *DescribeIdcAssetCriteriaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
