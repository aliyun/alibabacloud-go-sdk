// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOssObjectDetailV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOssObjectDetailV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOssObjectDetailV2Response
	GetStatusCode() *int32
	SetBody(v *DescribeOssObjectDetailV2ResponseBody) *DescribeOssObjectDetailV2Response
	GetBody() *DescribeOssObjectDetailV2ResponseBody
}

type DescribeOssObjectDetailV2Response struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOssObjectDetailV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOssObjectDetailV2Response) String() string {
	return dara.Prettify(s)
}

func (s DescribeOssObjectDetailV2Response) GoString() string {
	return s.String()
}

func (s *DescribeOssObjectDetailV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOssObjectDetailV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOssObjectDetailV2Response) GetBody() *DescribeOssObjectDetailV2ResponseBody {
	return s.Body
}

func (s *DescribeOssObjectDetailV2Response) SetHeaders(v map[string]*string) *DescribeOssObjectDetailV2Response {
	s.Headers = v
	return s
}

func (s *DescribeOssObjectDetailV2Response) SetStatusCode(v int32) *DescribeOssObjectDetailV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeOssObjectDetailV2Response) SetBody(v *DescribeOssObjectDetailV2ResponseBody) *DescribeOssObjectDetailV2Response {
	s.Body = v
	return s
}

func (s *DescribeOssObjectDetailV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
