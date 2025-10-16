// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCtrlInstanceMemberAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCtrlInstanceMemberAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCtrlInstanceMemberAccountsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCtrlInstanceMemberAccountsResponseBody) *DescribeCtrlInstanceMemberAccountsResponse
	GetBody() *DescribeCtrlInstanceMemberAccountsResponseBody
}

type DescribeCtrlInstanceMemberAccountsResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCtrlInstanceMemberAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCtrlInstanceMemberAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCtrlInstanceMemberAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCtrlInstanceMemberAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCtrlInstanceMemberAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCtrlInstanceMemberAccountsResponse) GetBody() *DescribeCtrlInstanceMemberAccountsResponseBody {
	return s.Body
}

func (s *DescribeCtrlInstanceMemberAccountsResponse) SetHeaders(v map[string]*string) *DescribeCtrlInstanceMemberAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCtrlInstanceMemberAccountsResponse) SetStatusCode(v int32) *DescribeCtrlInstanceMemberAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCtrlInstanceMemberAccountsResponse) SetBody(v *DescribeCtrlInstanceMemberAccountsResponseBody) *DescribeCtrlInstanceMemberAccountsResponse {
	s.Body = v
	return s
}

func (s *DescribeCtrlInstanceMemberAccountsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
