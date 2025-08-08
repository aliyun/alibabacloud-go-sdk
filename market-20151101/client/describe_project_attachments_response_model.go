// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProjectAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProjectAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProjectAttachmentsResponseBody) *DescribeProjectAttachmentsResponse
	GetBody() *DescribeProjectAttachmentsResponseBody
}

type DescribeProjectAttachmentsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProjectAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProjectAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProjectAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProjectAttachmentsResponse) GetBody() *DescribeProjectAttachmentsResponseBody {
	return s.Body
}

func (s *DescribeProjectAttachmentsResponse) SetHeaders(v map[string]*string) *DescribeProjectAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectAttachmentsResponse) SetStatusCode(v int32) *DescribeProjectAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectAttachmentsResponse) SetBody(v *DescribeProjectAttachmentsResponseBody) *DescribeProjectAttachmentsResponse {
	s.Body = v
	return s
}

func (s *DescribeProjectAttachmentsResponse) Validate() error {
	return dara.Validate(s)
}
