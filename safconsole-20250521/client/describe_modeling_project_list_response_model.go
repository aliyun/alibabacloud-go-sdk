// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelingProjectListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelingProjectListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelingProjectListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelingProjectListResponseBody) *DescribeModelingProjectListResponse
	GetBody() *DescribeModelingProjectListResponseBody
}

type DescribeModelingProjectListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelingProjectListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelingProjectListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelingProjectListResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelingProjectListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelingProjectListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelingProjectListResponse) GetBody() *DescribeModelingProjectListResponseBody {
	return s.Body
}

func (s *DescribeModelingProjectListResponse) SetHeaders(v map[string]*string) *DescribeModelingProjectListResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelingProjectListResponse) SetStatusCode(v int32) *DescribeModelingProjectListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelingProjectListResponse) SetBody(v *DescribeModelingProjectListResponseBody) *DescribeModelingProjectListResponse {
	s.Body = v
	return s
}

func (s *DescribeModelingProjectListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
