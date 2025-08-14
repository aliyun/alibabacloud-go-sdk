// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListVariablePageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNameListVariablePageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNameListVariablePageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNameListVariablePageListResponseBody) *DescribeNameListVariablePageListResponse
	GetBody() *DescribeNameListVariablePageListResponseBody
}

type DescribeNameListVariablePageListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNameListVariablePageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNameListVariablePageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListVariablePageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNameListVariablePageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNameListVariablePageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNameListVariablePageListResponse) GetBody() *DescribeNameListVariablePageListResponseBody {
	return s.Body
}

func (s *DescribeNameListVariablePageListResponse) SetHeaders(v map[string]*string) *DescribeNameListVariablePageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNameListVariablePageListResponse) SetStatusCode(v int32) *DescribeNameListVariablePageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNameListVariablePageListResponse) SetBody(v *DescribeNameListVariablePageListResponseBody) *DescribeNameListVariablePageListResponse {
	s.Body = v
	return s
}

func (s *DescribeNameListVariablePageListResponse) Validate() error {
	return dara.Validate(s)
}
