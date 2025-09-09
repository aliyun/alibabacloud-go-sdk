// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCategoryTemplateRuleListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCategoryTemplateRuleListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCategoryTemplateRuleListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCategoryTemplateRuleListResponseBody) *DescribeCategoryTemplateRuleListResponse
	GetBody() *DescribeCategoryTemplateRuleListResponseBody
}

type DescribeCategoryTemplateRuleListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCategoryTemplateRuleListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCategoryTemplateRuleListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryTemplateRuleListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateRuleListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCategoryTemplateRuleListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCategoryTemplateRuleListResponse) GetBody() *DescribeCategoryTemplateRuleListResponseBody {
	return s.Body
}

func (s *DescribeCategoryTemplateRuleListResponse) SetHeaders(v map[string]*string) *DescribeCategoryTemplateRuleListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponse) SetStatusCode(v int32) *DescribeCategoryTemplateRuleListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponse) SetBody(v *DescribeCategoryTemplateRuleListResponseBody) *DescribeCategoryTemplateRuleListResponse {
	s.Body = v
	return s
}

func (s *DescribeCategoryTemplateRuleListResponse) Validate() error {
	return dara.Validate(s)
}
