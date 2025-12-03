// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreIncrDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreIncrDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreIncrDetailResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreIncrDetailResponseBody) *DescribeRestoreIncrDetailResponse
	GetBody() *DescribeRestoreIncrDetailResponseBody
}

type DescribeRestoreIncrDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreIncrDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreIncrDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreIncrDetailResponse) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreIncrDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreIncrDetailResponse) GetBody() *DescribeRestoreIncrDetailResponseBody {
	return s.Body
}

func (s *DescribeRestoreIncrDetailResponse) SetHeaders(v map[string]*string) *DescribeRestoreIncrDetailResponse {
	s.Headers = v
	return s
}

func (s *DescribeRestoreIncrDetailResponse) SetStatusCode(v int32) *DescribeRestoreIncrDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponse) SetBody(v *DescribeRestoreIncrDetailResponseBody) *DescribeRestoreIncrDetailResponse {
	s.Body = v
	return s
}

func (s *DescribeRestoreIncrDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
