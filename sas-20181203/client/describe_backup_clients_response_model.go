// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupClientsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupClientsResponseBody) *DescribeBackupClientsResponse
	GetBody() *DescribeBackupClientsResponseBody
}

type DescribeBackupClientsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupClientsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupClientsResponse) GetBody() *DescribeBackupClientsResponseBody {
	return s.Body
}

func (s *DescribeBackupClientsResponse) SetHeaders(v map[string]*string) *DescribeBackupClientsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupClientsResponse) SetStatusCode(v int32) *DescribeBackupClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupClientsResponse) SetBody(v *DescribeBackupClientsResponseBody) *DescribeBackupClientsResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
