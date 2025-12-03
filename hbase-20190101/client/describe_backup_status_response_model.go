// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupStatusResponseBody) *DescribeBackupStatusResponse
	GetBody() *DescribeBackupStatusResponseBody
}

type DescribeBackupStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupStatusResponse) GetBody() *DescribeBackupStatusResponseBody {
	return s.Body
}

func (s *DescribeBackupStatusResponse) SetHeaders(v map[string]*string) *DescribeBackupStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupStatusResponse) SetStatusCode(v int32) *DescribeBackupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupStatusResponse) SetBody(v *DescribeBackupStatusResponseBody) *DescribeBackupStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
