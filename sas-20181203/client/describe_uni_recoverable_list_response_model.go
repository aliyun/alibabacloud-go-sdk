// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniRecoverableListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUniRecoverableListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUniRecoverableListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUniRecoverableListResponseBody) *DescribeUniRecoverableListResponse
	GetBody() *DescribeUniRecoverableListResponseBody
}

type DescribeUniRecoverableListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUniRecoverableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUniRecoverableListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniRecoverableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeUniRecoverableListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUniRecoverableListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUniRecoverableListResponse) GetBody() *DescribeUniRecoverableListResponseBody {
	return s.Body
}

func (s *DescribeUniRecoverableListResponse) SetHeaders(v map[string]*string) *DescribeUniRecoverableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeUniRecoverableListResponse) SetStatusCode(v int32) *DescribeUniRecoverableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUniRecoverableListResponse) SetBody(v *DescribeUniRecoverableListResponseBody) *DescribeUniRecoverableListResponse {
	s.Body = v
	return s
}

func (s *DescribeUniRecoverableListResponse) Validate() error {
	return dara.Validate(s)
}
