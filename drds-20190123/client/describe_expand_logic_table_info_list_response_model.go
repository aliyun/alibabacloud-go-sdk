// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpandLogicTableInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeExpandLogicTableInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeExpandLogicTableInfoListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeExpandLogicTableInfoListResponseBody) *DescribeExpandLogicTableInfoListResponse
	GetBody() *DescribeExpandLogicTableInfoListResponseBody
}

type DescribeExpandLogicTableInfoListResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeExpandLogicTableInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeExpandLogicTableInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponse) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeExpandLogicTableInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeExpandLogicTableInfoListResponse) GetBody() *DescribeExpandLogicTableInfoListResponseBody {
	return s.Body
}

func (s *DescribeExpandLogicTableInfoListResponse) SetHeaders(v map[string]*string) *DescribeExpandLogicTableInfoListResponse {
	s.Headers = v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponse) SetStatusCode(v int32) *DescribeExpandLogicTableInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponse) SetBody(v *DescribeExpandLogicTableInfoListResponseBody) *DescribeExpandLogicTableInfoListResponse {
	s.Body = v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponse) Validate() error {
	return dara.Validate(s)
}
