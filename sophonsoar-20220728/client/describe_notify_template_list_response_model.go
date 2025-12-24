// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNotifyTemplateListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNotifyTemplateListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNotifyTemplateListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNotifyTemplateListResponseBody) *DescribeNotifyTemplateListResponse
	GetBody() *DescribeNotifyTemplateListResponseBody
}

type DescribeNotifyTemplateListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNotifyTemplateListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNotifyTemplateListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNotifyTemplateListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNotifyTemplateListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNotifyTemplateListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNotifyTemplateListResponse) GetBody() *DescribeNotifyTemplateListResponseBody {
	return s.Body
}

func (s *DescribeNotifyTemplateListResponse) SetHeaders(v map[string]*string) *DescribeNotifyTemplateListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNotifyTemplateListResponse) SetStatusCode(v int32) *DescribeNotifyTemplateListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNotifyTemplateListResponse) SetBody(v *DescribeNotifyTemplateListResponseBody) *DescribeNotifyTemplateListResponse {
	s.Body = v
	return s
}

func (s *DescribeNotifyTemplateListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
