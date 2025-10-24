// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagGroupResponseBody) *DescribeTagGroupResponse
	GetBody() *DescribeTagGroupResponseBody
}

type DescribeTagGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagGroupResponse) GetBody() *DescribeTagGroupResponseBody {
	return s.Body
}

func (s *DescribeTagGroupResponse) SetHeaders(v map[string]*string) *DescribeTagGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagGroupResponse) SetStatusCode(v int32) *DescribeTagGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagGroupResponse) SetBody(v *DescribeTagGroupResponseBody) *DescribeTagGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeTagGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
