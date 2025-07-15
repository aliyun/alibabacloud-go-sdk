// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAclEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAclEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAclEntriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAclEntriesResponseBody) *DescribeAclEntriesResponse
	GetBody() *DescribeAclEntriesResponseBody
}

type DescribeAclEntriesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAclEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAclEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAclEntriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeAclEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAclEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAclEntriesResponse) GetBody() *DescribeAclEntriesResponseBody {
	return s.Body
}

func (s *DescribeAclEntriesResponse) SetHeaders(v map[string]*string) *DescribeAclEntriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeAclEntriesResponse) SetStatusCode(v int32) *DescribeAclEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAclEntriesResponse) SetBody(v *DescribeAclEntriesResponseBody) *DescribeAclEntriesResponse {
	s.Body = v
	return s
}

func (s *DescribeAclEntriesResponse) Validate() error {
	return dara.Validate(s)
}
