// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccesskeyLeakListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAccesskeyLeakListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAccesskeyLeakListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAccesskeyLeakListResponseBody) *DescribeAccesskeyLeakListResponse
	GetBody() *DescribeAccesskeyLeakListResponseBody
}

type DescribeAccesskeyLeakListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAccesskeyLeakListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAccesskeyLeakListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccesskeyLeakListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAccesskeyLeakListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAccesskeyLeakListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAccesskeyLeakListResponse) GetBody() *DescribeAccesskeyLeakListResponseBody {
	return s.Body
}

func (s *DescribeAccesskeyLeakListResponse) SetHeaders(v map[string]*string) *DescribeAccesskeyLeakListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAccesskeyLeakListResponse) SetStatusCode(v int32) *DescribeAccesskeyLeakListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAccesskeyLeakListResponse) SetBody(v *DescribeAccesskeyLeakListResponseBody) *DescribeAccesskeyLeakListResponse {
	s.Body = v
	return s
}

func (s *DescribeAccesskeyLeakListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
