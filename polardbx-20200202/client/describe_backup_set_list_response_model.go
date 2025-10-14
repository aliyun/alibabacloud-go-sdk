// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupSetListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupSetListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupSetListResponseBody) *DescribeBackupSetListResponse
	GetBody() *DescribeBackupSetListResponseBody
}

type DescribeBackupSetListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupSetListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupSetListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupSetListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupSetListResponse) GetBody() *DescribeBackupSetListResponseBody {
	return s.Body
}

func (s *DescribeBackupSetListResponse) SetHeaders(v map[string]*string) *DescribeBackupSetListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSetListResponse) SetStatusCode(v int32) *DescribeBackupSetListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupSetListResponse) SetBody(v *DescribeBackupSetListResponseBody) *DescribeBackupSetListResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupSetListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
