// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeParentInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeParentInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeParentInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeParentInstanceResponseBody) *DescribeParentInstanceResponse
	GetBody() *DescribeParentInstanceResponseBody
}

type DescribeParentInstanceResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeParentInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeParentInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeParentInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeParentInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeParentInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeParentInstanceResponse) GetBody() *DescribeParentInstanceResponseBody {
	return s.Body
}

func (s *DescribeParentInstanceResponse) SetHeaders(v map[string]*string) *DescribeParentInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeParentInstanceResponse) SetStatusCode(v int32) *DescribeParentInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeParentInstanceResponse) SetBody(v *DescribeParentInstanceResponseBody) *DescribeParentInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeParentInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
