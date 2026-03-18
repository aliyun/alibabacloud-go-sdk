// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdMemberListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdMemberListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdMemberListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdMemberListResponseBody) *DescribeRdMemberListResponse
	GetBody() *DescribeRdMemberListResponseBody
}

type DescribeRdMemberListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdMemberListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdMemberListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdMemberListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdMemberListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdMemberListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdMemberListResponse) GetBody() *DescribeRdMemberListResponseBody {
	return s.Body
}

func (s *DescribeRdMemberListResponse) SetHeaders(v map[string]*string) *DescribeRdMemberListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdMemberListResponse) SetStatusCode(v int32) *DescribeRdMemberListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdMemberListResponse) SetBody(v *DescribeRdMemberListResponseBody) *DescribeRdMemberListResponse {
	s.Body = v
	return s
}

func (s *DescribeRdMemberListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
