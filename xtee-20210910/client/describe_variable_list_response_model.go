// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVariableListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVariableListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVariableListResponseBody) *DescribeVariableListResponse
	GetBody() *DescribeVariableListResponseBody
}

type DescribeVariableListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVariableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVariableListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVariableListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVariableListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVariableListResponse) GetBody() *DescribeVariableListResponseBody {
	return s.Body
}

func (s *DescribeVariableListResponse) SetHeaders(v map[string]*string) *DescribeVariableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVariableListResponse) SetStatusCode(v int32) *DescribeVariableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVariableListResponse) SetBody(v *DescribeVariableListResponseBody) *DescribeVariableListResponse {
	s.Body = v
	return s
}

func (s *DescribeVariableListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
