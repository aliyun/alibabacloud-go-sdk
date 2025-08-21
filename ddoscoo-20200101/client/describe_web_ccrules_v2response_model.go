// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCCRulesV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeWebCCRulesV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeWebCCRulesV2Response
	GetStatusCode() *int32
	SetBody(v *DescribeWebCCRulesV2ResponseBody) *DescribeWebCCRulesV2Response
	GetBody() *DescribeWebCCRulesV2ResponseBody
}

type DescribeWebCCRulesV2Response struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeWebCCRulesV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeWebCCRulesV2Response) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCCRulesV2Response) GoString() string {
	return s.String()
}

func (s *DescribeWebCCRulesV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeWebCCRulesV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeWebCCRulesV2Response) GetBody() *DescribeWebCCRulesV2ResponseBody {
	return s.Body
}

func (s *DescribeWebCCRulesV2Response) SetHeaders(v map[string]*string) *DescribeWebCCRulesV2Response {
	s.Headers = v
	return s
}

func (s *DescribeWebCCRulesV2Response) SetStatusCode(v int32) *DescribeWebCCRulesV2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeWebCCRulesV2Response) SetBody(v *DescribeWebCCRulesV2ResponseBody) *DescribeWebCCRulesV2Response {
	s.Body = v
	return s
}

func (s *DescribeWebCCRulesV2Response) Validate() error {
	return dara.Validate(s)
}
