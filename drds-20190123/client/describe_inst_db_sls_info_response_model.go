// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstDbSlsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstDbSlsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstDbSlsInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstDbSlsInfoResponseBody) *DescribeInstDbSlsInfoResponse
	GetBody() *DescribeInstDbSlsInfoResponseBody
}

type DescribeInstDbSlsInfoResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstDbSlsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstDbSlsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstDbSlsInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstDbSlsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstDbSlsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstDbSlsInfoResponse) GetBody() *DescribeInstDbSlsInfoResponseBody {
	return s.Body
}

func (s *DescribeInstDbSlsInfoResponse) SetHeaders(v map[string]*string) *DescribeInstDbSlsInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstDbSlsInfoResponse) SetStatusCode(v int32) *DescribeInstDbSlsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstDbSlsInfoResponse) SetBody(v *DescribeInstDbSlsInfoResponseBody) *DescribeInstDbSlsInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeInstDbSlsInfoResponse) Validate() error {
	return dara.Validate(s)
}
