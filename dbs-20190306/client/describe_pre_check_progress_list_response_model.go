// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePreCheckProgressListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePreCheckProgressListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePreCheckProgressListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePreCheckProgressListResponseBody) *DescribePreCheckProgressListResponse
	GetBody() *DescribePreCheckProgressListResponseBody
}

type DescribePreCheckProgressListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePreCheckProgressListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePreCheckProgressListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePreCheckProgressListResponse) GoString() string {
	return s.String()
}

func (s *DescribePreCheckProgressListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePreCheckProgressListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePreCheckProgressListResponse) GetBody() *DescribePreCheckProgressListResponseBody {
	return s.Body
}

func (s *DescribePreCheckProgressListResponse) SetHeaders(v map[string]*string) *DescribePreCheckProgressListResponse {
	s.Headers = v
	return s
}

func (s *DescribePreCheckProgressListResponse) SetStatusCode(v int32) *DescribePreCheckProgressListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePreCheckProgressListResponse) SetBody(v *DescribePreCheckProgressListResponseBody) *DescribePreCheckProgressListResponse {
	s.Body = v
	return s
}

func (s *DescribePreCheckProgressListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
