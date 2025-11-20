// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupJobs2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupJobs2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupJobs2Response
	GetStatusCode() *int32
	SetBody(v *DescribeBackupJobs2ResponseBody) *DescribeBackupJobs2Response
	GetBody() *DescribeBackupJobs2ResponseBody
}

type DescribeBackupJobs2Response struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupJobs2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupJobs2Response) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupJobs2Response) GoString() string {
	return s.String()
}

func (s *DescribeBackupJobs2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupJobs2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupJobs2Response) GetBody() *DescribeBackupJobs2ResponseBody {
	return s.Body
}

func (s *DescribeBackupJobs2Response) SetHeaders(v map[string]*string) *DescribeBackupJobs2Response {
	s.Headers = v
	return s
}

func (s *DescribeBackupJobs2Response) SetStatusCode(v int32) *DescribeBackupJobs2Response {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupJobs2Response) SetBody(v *DescribeBackupJobs2ResponseBody) *DescribeBackupJobs2Response {
	s.Body = v
	return s
}

func (s *DescribeBackupJobs2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
