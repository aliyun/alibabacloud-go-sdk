// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupSetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupSetsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupSetsResponseBody) *DescribeBackupSetsResponse
	GetBody() *DescribeBackupSetsResponseBody
}

type DescribeBackupSetsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupSetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupSetsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetsResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupSetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupSetsResponse) GetBody() *DescribeBackupSetsResponseBody {
	return s.Body
}

func (s *DescribeBackupSetsResponse) SetHeaders(v map[string]*string) *DescribeBackupSetsResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSetsResponse) SetStatusCode(v int32) *DescribeBackupSetsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupSetsResponse) SetBody(v *DescribeBackupSetsResponseBody) *DescribeBackupSetsResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupSetsResponse) Validate() error {
	return dara.Validate(s)
}
