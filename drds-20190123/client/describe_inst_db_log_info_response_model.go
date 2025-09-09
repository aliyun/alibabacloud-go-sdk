// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstDbLogInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstDbLogInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstDbLogInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstDbLogInfoResponseBody) *DescribeInstDbLogInfoResponse
	GetBody() *DescribeInstDbLogInfoResponseBody
}

type DescribeInstDbLogInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstDbLogInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstDbLogInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstDbLogInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstDbLogInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstDbLogInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstDbLogInfoResponse) GetBody() *DescribeInstDbLogInfoResponseBody {
	return s.Body
}

func (s *DescribeInstDbLogInfoResponse) SetHeaders(v map[string]*string) *DescribeInstDbLogInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstDbLogInfoResponse) SetStatusCode(v int32) *DescribeInstDbLogInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstDbLogInfoResponse) SetBody(v *DescribeInstDbLogInfoResponseBody) *DescribeInstDbLogInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeInstDbLogInfoResponse) Validate() error {
	return dara.Validate(s)
}
