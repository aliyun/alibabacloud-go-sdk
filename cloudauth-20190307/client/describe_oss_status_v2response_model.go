// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssStatusV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssStatusV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssStatusV2Response
	GetStatusCode() *int32
	SetBody(v *DescribeOssStatusV2ResponseBody) *DescribeOssStatusV2Response
	GetBody() *DescribeOssStatusV2ResponseBody
}

type DescribeOssStatusV2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssStatusV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssStatusV2Response) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssStatusV2Response) GoString() string {
	return s.String()
}

func (s *DescribeOssStatusV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssStatusV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssStatusV2Response) GetBody() *DescribeOssStatusV2ResponseBody {
	return s.Body
}

func (s *DescribeOssStatusV2Response) SetHeaders(v map[string]*string) *DescribeOssStatusV2Response {
	s.Headers = v
	return s
}

func (s *DescribeOssStatusV2Response) SetStatusCode(v int32) *DescribeOssStatusV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssStatusV2Response) SetBody(v *DescribeOssStatusV2ResponseBody) *DescribeOssStatusV2Response {
	s.Body = v
	return s
}

func (s *DescribeOssStatusV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
