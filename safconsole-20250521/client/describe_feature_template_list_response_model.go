// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFeatureTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeFeatureTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeFeatureTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeFeatureTemplateListResponseBody) *DescribeFeatureTemplateListResponse
	GetBody() *DescribeFeatureTemplateListResponseBody
}

type DescribeFeatureTemplateListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeFeatureTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeFeatureTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeFeatureTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeFeatureTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeFeatureTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeFeatureTemplateListResponse) GetBody() *DescribeFeatureTemplateListResponseBody {
	return s.Body
}

func (s *DescribeFeatureTemplateListResponse) SetHeaders(v map[string]*string) *DescribeFeatureTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeFeatureTemplateListResponse) SetStatusCode(v int32) *DescribeFeatureTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeFeatureTemplateListResponse) SetBody(v *DescribeFeatureTemplateListResponseBody) *DescribeFeatureTemplateListResponse {
	s.Body = v
	return s
}

func (s *DescribeFeatureTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
