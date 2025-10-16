// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDesktopsInGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDesktopsInGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDesktopsInGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDesktopsInGroupResponseBody) *DescribeDesktopsInGroupResponse
	GetBody() *DescribeDesktopsInGroupResponseBody
}

type DescribeDesktopsInGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDesktopsInGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDesktopsInGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDesktopsInGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeDesktopsInGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDesktopsInGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDesktopsInGroupResponse) GetBody() *DescribeDesktopsInGroupResponseBody {
	return s.Body
}

func (s *DescribeDesktopsInGroupResponse) SetHeaders(v map[string]*string) *DescribeDesktopsInGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetStatusCode(v int32) *DescribeDesktopsInGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDesktopsInGroupResponse) SetBody(v *DescribeDesktopsInGroupResponseBody) *DescribeDesktopsInGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeDesktopsInGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
