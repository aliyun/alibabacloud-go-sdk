// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpnAttachmentsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVpnAttachmentsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVpnAttachmentsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVpnAttachmentsResponseBody) *DescribeVpnAttachmentsResponse
	GetBody() *DescribeVpnAttachmentsResponseBody
}

type DescribeVpnAttachmentsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVpnAttachmentsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVpnAttachmentsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpnAttachmentsResponse) GoString() string {
	return s.String()
}

func (s *DescribeVpnAttachmentsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVpnAttachmentsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVpnAttachmentsResponse) GetBody() *DescribeVpnAttachmentsResponseBody {
	return s.Body
}

func (s *DescribeVpnAttachmentsResponse) SetHeaders(v map[string]*string) *DescribeVpnAttachmentsResponse {
	s.Headers = v
	return s
}

func (s *DescribeVpnAttachmentsResponse) SetStatusCode(v int32) *DescribeVpnAttachmentsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVpnAttachmentsResponse) SetBody(v *DescribeVpnAttachmentsResponseBody) *DescribeVpnAttachmentsResponse {
	s.Body = v
	return s
}

func (s *DescribeVpnAttachmentsResponse) Validate() error {
	return dara.Validate(s)
}
