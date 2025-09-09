// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCategoryTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCategoryTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCategoryTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCategoryTemplateListResponseBody) *DescribeCategoryTemplateListResponse
	GetBody() *DescribeCategoryTemplateListResponseBody
}

type DescribeCategoryTemplateListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCategoryTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCategoryTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCategoryTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCategoryTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCategoryTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCategoryTemplateListResponse) GetBody() *DescribeCategoryTemplateListResponseBody {
	return s.Body
}

func (s *DescribeCategoryTemplateListResponse) SetHeaders(v map[string]*string) *DescribeCategoryTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCategoryTemplateListResponse) SetStatusCode(v int32) *DescribeCategoryTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCategoryTemplateListResponse) SetBody(v *DescribeCategoryTemplateListResponseBody) *DescribeCategoryTemplateListResponse {
	s.Body = v
	return s
}

func (s *DescribeCategoryTemplateListResponse) Validate() error {
	return dara.Validate(s)
}
