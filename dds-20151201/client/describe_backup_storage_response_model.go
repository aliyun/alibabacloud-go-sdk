// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupStorageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupStorageResponseBody) *DescribeBackupStorageResponse
	GetBody() *DescribeBackupStorageResponseBody
}

type DescribeBackupStorageResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupStorageResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupStorageResponse) GetBody() *DescribeBackupStorageResponseBody {
	return s.Body
}

func (s *DescribeBackupStorageResponse) SetHeaders(v map[string]*string) *DescribeBackupStorageResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupStorageResponse) SetStatusCode(v int32) *DescribeBackupStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupStorageResponse) SetBody(v *DescribeBackupStorageResponseBody) *DescribeBackupStorageResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
