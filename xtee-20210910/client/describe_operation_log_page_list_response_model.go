// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperationLogPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOperationLogPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOperationLogPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOperationLogPageListResponseBody) *DescribeOperationLogPageListResponse
	GetBody() *DescribeOperationLogPageListResponseBody
}

type DescribeOperationLogPageListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOperationLogPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOperationLogPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperationLogPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperationLogPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOperationLogPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOperationLogPageListResponse) GetBody() *DescribeOperationLogPageListResponseBody {
	return s.Body
}

func (s *DescribeOperationLogPageListResponse) SetHeaders(v map[string]*string) *DescribeOperationLogPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperationLogPageListResponse) SetStatusCode(v int32) *DescribeOperationLogPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperationLogPageListResponse) SetBody(v *DescribeOperationLogPageListResponseBody) *DescribeOperationLogPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeOperationLogPageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
