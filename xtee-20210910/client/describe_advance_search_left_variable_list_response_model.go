// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvanceSearchLeftVariableListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvanceSearchLeftVariableListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvanceSearchLeftVariableListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvanceSearchLeftVariableListResponseBody) *DescribeAdvanceSearchLeftVariableListResponse
	GetBody() *DescribeAdvanceSearchLeftVariableListResponseBody
}

type DescribeAdvanceSearchLeftVariableListResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvanceSearchLeftVariableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvanceSearchLeftVariableListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvanceSearchLeftVariableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvanceSearchLeftVariableListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvanceSearchLeftVariableListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvanceSearchLeftVariableListResponse) GetBody() *DescribeAdvanceSearchLeftVariableListResponseBody {
	return s.Body
}

func (s *DescribeAdvanceSearchLeftVariableListResponse) SetHeaders(v map[string]*string) *DescribeAdvanceSearchLeftVariableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponse) SetStatusCode(v int32) *DescribeAdvanceSearchLeftVariableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponse) SetBody(v *DescribeAdvanceSearchLeftVariableListResponseBody) *DescribeAdvanceSearchLeftVariableListResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvanceSearchLeftVariableListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
