// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContactListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContactListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContactListResponseBody) *DescribeContactListResponse
	GetBody() *DescribeContactListResponseBody
}

type DescribeContactListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContactListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContactListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactListResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContactListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContactListResponse) GetBody() *DescribeContactListResponseBody {
	return s.Body
}

func (s *DescribeContactListResponse) SetHeaders(v map[string]*string) *DescribeContactListResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactListResponse) SetStatusCode(v int32) *DescribeContactListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContactListResponse) SetBody(v *DescribeContactListResponseBody) *DescribeContactListResponse {
	s.Body = v
	return s
}

func (s *DescribeContactListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
