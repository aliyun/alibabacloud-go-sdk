// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupDBsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupDBsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupDBsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupDBsResponseBody) *DescribeBackupDBsResponse
	GetBody() *DescribeBackupDBsResponseBody
}

type DescribeBackupDBsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupDBsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupDBsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupDBsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupDBsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupDBsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupDBsResponse) GetBody() *DescribeBackupDBsResponseBody {
	return s.Body
}

func (s *DescribeBackupDBsResponse) SetHeaders(v map[string]*string) *DescribeBackupDBsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupDBsResponse) SetStatusCode(v int32) *DescribeBackupDBsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupDBsResponse) SetBody(v *DescribeBackupDBsResponseBody) *DescribeBackupDBsResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupDBsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
