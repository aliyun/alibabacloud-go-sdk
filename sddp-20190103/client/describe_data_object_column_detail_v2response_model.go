// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataObjectColumnDetailV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataObjectColumnDetailV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataObjectColumnDetailV2Response
	GetStatusCode() *int32
	SetBody(v *DescribeDataObjectColumnDetailV2ResponseBody) *DescribeDataObjectColumnDetailV2Response
	GetBody() *DescribeDataObjectColumnDetailV2ResponseBody
}

type DescribeDataObjectColumnDetailV2Response struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataObjectColumnDetailV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataObjectColumnDetailV2Response) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataObjectColumnDetailV2Response) GoString() string {
	return s.String()
}

func (s *DescribeDataObjectColumnDetailV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataObjectColumnDetailV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataObjectColumnDetailV2Response) GetBody() *DescribeDataObjectColumnDetailV2ResponseBody {
	return s.Body
}

func (s *DescribeDataObjectColumnDetailV2Response) SetHeaders(v map[string]*string) *DescribeDataObjectColumnDetailV2Response {
	s.Headers = v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Response) SetStatusCode(v int32) *DescribeDataObjectColumnDetailV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Response) SetBody(v *DescribeDataObjectColumnDetailV2ResponseBody) *DescribeDataObjectColumnDetailV2Response {
	s.Body = v
	return s
}

func (s *DescribeDataObjectColumnDetailV2Response) Validate() error {
	return dara.Validate(s)
}
