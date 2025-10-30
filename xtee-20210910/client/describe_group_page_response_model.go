// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupPageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupPageResponseBody) *DescribeGroupPageResponse
	GetBody() *DescribeGroupPageResponseBody
}

type DescribeGroupPageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupPageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupPageResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupPageResponse) GetBody() *DescribeGroupPageResponseBody {
	return s.Body
}

func (s *DescribeGroupPageResponse) SetHeaders(v map[string]*string) *DescribeGroupPageResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupPageResponse) SetStatusCode(v int32) *DescribeGroupPageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupPageResponse) SetBody(v *DescribeGroupPageResponseBody) *DescribeGroupPageResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupPageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
