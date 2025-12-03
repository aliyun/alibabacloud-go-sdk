// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupSummaryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupSummaryResponseBody) *DescribeBackupSummaryResponse
	GetBody() *DescribeBackupSummaryResponseBody
}

type DescribeBackupSummaryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSummaryResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupSummaryResponse) GetBody() *DescribeBackupSummaryResponseBody {
	return s.Body
}

func (s *DescribeBackupSummaryResponse) SetHeaders(v map[string]*string) *DescribeBackupSummaryResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSummaryResponse) SetStatusCode(v int32) *DescribeBackupSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupSummaryResponse) SetBody(v *DescribeBackupSummaryResponseBody) *DescribeBackupSummaryResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
