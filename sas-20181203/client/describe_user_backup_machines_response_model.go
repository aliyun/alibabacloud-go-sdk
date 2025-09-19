// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBackupMachinesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserBackupMachinesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserBackupMachinesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserBackupMachinesResponseBody) *DescribeUserBackupMachinesResponse
	GetBody() *DescribeUserBackupMachinesResponseBody
}

type DescribeUserBackupMachinesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserBackupMachinesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserBackupMachinesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBackupMachinesResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBackupMachinesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserBackupMachinesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserBackupMachinesResponse) GetBody() *DescribeUserBackupMachinesResponseBody {
	return s.Body
}

func (s *DescribeUserBackupMachinesResponse) SetHeaders(v map[string]*string) *DescribeUserBackupMachinesResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBackupMachinesResponse) SetStatusCode(v int32) *DescribeUserBackupMachinesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBackupMachinesResponse) SetBody(v *DescribeUserBackupMachinesResponseBody) *DescribeUserBackupMachinesResponse {
	s.Body = v
	return s
}

func (s *DescribeUserBackupMachinesResponse) Validate() error {
	return dara.Validate(s)
}
