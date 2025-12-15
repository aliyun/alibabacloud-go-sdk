// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParameterGroupTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParameterGroupTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParameterGroupTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParameterGroupTemplateListResponseBody) *DescribeParameterGroupTemplateListResponse
	GetBody() *DescribeParameterGroupTemplateListResponseBody
}

type DescribeParameterGroupTemplateListResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParameterGroupTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParameterGroupTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParameterGroupTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeParameterGroupTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParameterGroupTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParameterGroupTemplateListResponse) GetBody() *DescribeParameterGroupTemplateListResponseBody {
	return s.Body
}

func (s *DescribeParameterGroupTemplateListResponse) SetHeaders(v map[string]*string) *DescribeParameterGroupTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeParameterGroupTemplateListResponse) SetStatusCode(v int32) *DescribeParameterGroupTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParameterGroupTemplateListResponse) SetBody(v *DescribeParameterGroupTemplateListResponseBody) *DescribeParameterGroupTemplateListResponse {
	s.Body = v
	return s
}

func (s *DescribeParameterGroupTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
