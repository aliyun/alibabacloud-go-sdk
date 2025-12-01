// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeColumnsV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeColumnsV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeColumnsV2Response
	GetStatusCode() *int32
	SetBody(v *DescribeColumnsV2ResponseBody) *DescribeColumnsV2Response
	GetBody() *DescribeColumnsV2ResponseBody
}

type DescribeColumnsV2Response struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeColumnsV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeColumnsV2Response) String() string {
	return dara.Prettify(s)
}

func (s DescribeColumnsV2Response) GoString() string {
	return s.String()
}

func (s *DescribeColumnsV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeColumnsV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeColumnsV2Response) GetBody() *DescribeColumnsV2ResponseBody {
	return s.Body
}

func (s *DescribeColumnsV2Response) SetHeaders(v map[string]*string) *DescribeColumnsV2Response {
	s.Headers = v
	return s
}

func (s *DescribeColumnsV2Response) SetStatusCode(v int32) *DescribeColumnsV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeColumnsV2Response) SetBody(v *DescribeColumnsV2ResponseBody) *DescribeColumnsV2Response {
	s.Body = v
	return s
}

func (s *DescribeColumnsV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
