// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupMachineStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupMachineStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupMachineStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupMachineStatusResponseBody) *DescribeBackupMachineStatusResponse
	GetBody() *DescribeBackupMachineStatusResponseBody
}

type DescribeBackupMachineStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupMachineStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupMachineStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupMachineStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupMachineStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupMachineStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupMachineStatusResponse) GetBody() *DescribeBackupMachineStatusResponseBody {
	return s.Body
}

func (s *DescribeBackupMachineStatusResponse) SetHeaders(v map[string]*string) *DescribeBackupMachineStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupMachineStatusResponse) SetStatusCode(v int32) *DescribeBackupMachineStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupMachineStatusResponse) SetBody(v *DescribeBackupMachineStatusResponseBody) *DescribeBackupMachineStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupMachineStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
