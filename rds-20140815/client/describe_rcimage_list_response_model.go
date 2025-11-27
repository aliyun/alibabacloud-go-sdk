// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCImageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCImageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCImageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCImageListResponseBody) *DescribeRCImageListResponse
	GetBody() *DescribeRCImageListResponseBody
}

type DescribeRCImageListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCImageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCImageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCImageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCImageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCImageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCImageListResponse) GetBody() *DescribeRCImageListResponseBody {
	return s.Body
}

func (s *DescribeRCImageListResponse) SetHeaders(v map[string]*string) *DescribeRCImageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCImageListResponse) SetStatusCode(v int32) *DescribeRCImageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCImageListResponse) SetBody(v *DescribeRCImageListResponseBody) *DescribeRCImageListResponse {
	s.Body = v
	return s
}

func (s *DescribeRCImageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
