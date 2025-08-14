// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventsVariableListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventsVariableListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventsVariableListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventsVariableListResponseBody) *DescribeEventsVariableListResponse
	GetBody() *DescribeEventsVariableListResponseBody
}

type DescribeEventsVariableListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventsVariableListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventsVariableListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventsVariableListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventsVariableListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventsVariableListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventsVariableListResponse) GetBody() *DescribeEventsVariableListResponseBody {
	return s.Body
}

func (s *DescribeEventsVariableListResponse) SetHeaders(v map[string]*string) *DescribeEventsVariableListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventsVariableListResponse) SetStatusCode(v int32) *DescribeEventsVariableListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventsVariableListResponse) SetBody(v *DescribeEventsVariableListResponseBody) *DescribeEventsVariableListResponse {
	s.Body = v
	return s
}

func (s *DescribeEventsVariableListResponse) Validate() error {
	return dara.Validate(s)
}
