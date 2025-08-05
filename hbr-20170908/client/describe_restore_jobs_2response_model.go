// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreJobs2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRestoreJobs2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRestoreJobs2Response
	GetStatusCode() *int32
	SetBody(v *DescribeRestoreJobs2ResponseBody) *DescribeRestoreJobs2Response
	GetBody() *DescribeRestoreJobs2ResponseBody
}

type DescribeRestoreJobs2Response struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRestoreJobs2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRestoreJobs2Response) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreJobs2Response) GoString() string {
	return s.String()
}

func (s *DescribeRestoreJobs2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRestoreJobs2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRestoreJobs2Response) GetBody() *DescribeRestoreJobs2ResponseBody {
	return s.Body
}

func (s *DescribeRestoreJobs2Response) SetHeaders(v map[string]*string) *DescribeRestoreJobs2Response {
	s.Headers = v
	return s
}

func (s *DescribeRestoreJobs2Response) SetStatusCode(v int32) *DescribeRestoreJobs2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeRestoreJobs2Response) SetBody(v *DescribeRestoreJobs2ResponseBody) *DescribeRestoreJobs2Response {
	s.Body = v
	return s
}

func (s *DescribeRestoreJobs2Response) Validate() error {
	return dara.Validate(s)
}
