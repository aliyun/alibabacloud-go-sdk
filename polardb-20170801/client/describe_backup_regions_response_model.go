// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupRegionsResponseBody) *DescribeBackupRegionsResponse
	GetBody() *DescribeBackupRegionsResponseBody
}

type DescribeBackupRegionsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupRegionsResponse) GetBody() *DescribeBackupRegionsResponseBody {
	return s.Body
}

func (s *DescribeBackupRegionsResponse) SetHeaders(v map[string]*string) *DescribeBackupRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupRegionsResponse) SetStatusCode(v int32) *DescribeBackupRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupRegionsResponse) SetBody(v *DescribeBackupRegionsResponseBody) *DescribeBackupRegionsResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupRegionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
