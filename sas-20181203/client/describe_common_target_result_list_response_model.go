// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCommonTargetResultListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCommonTargetResultListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCommonTargetResultListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCommonTargetResultListResponseBody) *DescribeCommonTargetResultListResponse
	GetBody() *DescribeCommonTargetResultListResponseBody
}

type DescribeCommonTargetResultListResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCommonTargetResultListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCommonTargetResultListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCommonTargetResultListResponse) GoString() string {
	return s.String()
}

func (s *DescribeCommonTargetResultListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCommonTargetResultListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCommonTargetResultListResponse) GetBody() *DescribeCommonTargetResultListResponseBody {
	return s.Body
}

func (s *DescribeCommonTargetResultListResponse) SetHeaders(v map[string]*string) *DescribeCommonTargetResultListResponse {
	s.Headers = v
	return s
}

func (s *DescribeCommonTargetResultListResponse) SetStatusCode(v int32) *DescribeCommonTargetResultListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCommonTargetResultListResponse) SetBody(v *DescribeCommonTargetResultListResponseBody) *DescribeCommonTargetResultListResponse {
	s.Body = v
	return s
}

func (s *DescribeCommonTargetResultListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
