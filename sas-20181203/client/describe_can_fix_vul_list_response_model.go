// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCanFixVulListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCanFixVulListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCanFixVulListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCanFixVulListResponseBody) *DescribeCanFixVulListResponse
	GetBody() *DescribeCanFixVulListResponseBody
}

type DescribeCanFixVulListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCanFixVulListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCanFixVulListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCanFixVulListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCanFixVulListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCanFixVulListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCanFixVulListResponse) GetBody() *DescribeCanFixVulListResponseBody {
	return s.Body
}

func (s *DescribeCanFixVulListResponse) SetHeaders(v map[string]*string) *DescribeCanFixVulListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCanFixVulListResponse) SetStatusCode(v int32) *DescribeCanFixVulListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCanFixVulListResponse) SetBody(v *DescribeCanFixVulListResponseBody) *DescribeCanFixVulListResponse {
	s.Body = v
	return s
}

func (s *DescribeCanFixVulListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
