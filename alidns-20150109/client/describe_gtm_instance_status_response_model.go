// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGtmInstanceStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGtmInstanceStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGtmInstanceStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGtmInstanceStatusResponseBody) *DescribeGtmInstanceStatusResponse
	GetBody() *DescribeGtmInstanceStatusResponseBody
}

type DescribeGtmInstanceStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGtmInstanceStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGtmInstanceStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGtmInstanceStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeGtmInstanceStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGtmInstanceStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGtmInstanceStatusResponse) GetBody() *DescribeGtmInstanceStatusResponseBody {
	return s.Body
}

func (s *DescribeGtmInstanceStatusResponse) SetHeaders(v map[string]*string) *DescribeGtmInstanceStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeGtmInstanceStatusResponse) SetStatusCode(v int32) *DescribeGtmInstanceStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGtmInstanceStatusResponse) SetBody(v *DescribeGtmInstanceStatusResponseBody) *DescribeGtmInstanceStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeGtmInstanceStatusResponse) Validate() error {
	return dara.Validate(s)
}
