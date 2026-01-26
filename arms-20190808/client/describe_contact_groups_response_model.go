// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContactGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContactGroupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContactGroupsResponseBody) *DescribeContactGroupsResponse
	GetBody() *DescribeContactGroupsResponseBody
}

type DescribeContactGroupsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContactGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContactGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactGroupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContactGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContactGroupsResponse) GetBody() *DescribeContactGroupsResponseBody {
	return s.Body
}

func (s *DescribeContactGroupsResponse) SetHeaders(v map[string]*string) *DescribeContactGroupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactGroupsResponse) SetStatusCode(v int32) *DescribeContactGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContactGroupsResponse) SetBody(v *DescribeContactGroupsResponseBody) *DescribeContactGroupsResponse {
	s.Body = v
	return s
}

func (s *DescribeContactGroupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
