// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupUserResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupUserResponseBody) *DescribeGroupUserResponse
	GetBody() *DescribeGroupUserResponseBody
}

type DescribeGroupUserResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupUserResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupUserResponse) GetBody() *DescribeGroupUserResponseBody {
	return s.Body
}

func (s *DescribeGroupUserResponse) SetHeaders(v map[string]*string) *DescribeGroupUserResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupUserResponse) SetStatusCode(v int32) *DescribeGroupUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupUserResponse) SetBody(v *DescribeGroupUserResponseBody) *DescribeGroupUserResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
