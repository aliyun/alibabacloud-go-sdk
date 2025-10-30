// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeScoreListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeScoreListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeScoreListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeScoreListResponseBody) *DescribeScoreListResponse
	GetBody() *DescribeScoreListResponseBody
}

type DescribeScoreListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeScoreListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeScoreListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeScoreListResponse) GoString() string {
	return s.String()
}

func (s *DescribeScoreListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeScoreListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeScoreListResponse) GetBody() *DescribeScoreListResponseBody {
	return s.Body
}

func (s *DescribeScoreListResponse) SetHeaders(v map[string]*string) *DescribeScoreListResponse {
	s.Headers = v
	return s
}

func (s *DescribeScoreListResponse) SetStatusCode(v int32) *DescribeScoreListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeScoreListResponse) SetBody(v *DescribeScoreListResponseBody) *DescribeScoreListResponse {
	s.Body = v
	return s
}

func (s *DescribeScoreListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
