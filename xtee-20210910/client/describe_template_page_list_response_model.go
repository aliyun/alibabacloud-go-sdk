// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatePageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTemplatePageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTemplatePageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTemplatePageListResponseBody) *DescribeTemplatePageListResponse
	GetBody() *DescribeTemplatePageListResponseBody
}

type DescribeTemplatePageListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTemplatePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTemplatePageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatePageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTemplatePageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTemplatePageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTemplatePageListResponse) GetBody() *DescribeTemplatePageListResponseBody {
	return s.Body
}

func (s *DescribeTemplatePageListResponse) SetHeaders(v map[string]*string) *DescribeTemplatePageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTemplatePageListResponse) SetStatusCode(v int32) *DescribeTemplatePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTemplatePageListResponse) SetBody(v *DescribeTemplatePageListResponseBody) *DescribeTemplatePageListResponse {
	s.Body = v
	return s
}

func (s *DescribeTemplatePageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
