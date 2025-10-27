// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUniBackupStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUniBackupStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUniBackupStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUniBackupStatisticsResponseBody) *DescribeUniBackupStatisticsResponse
	GetBody() *DescribeUniBackupStatisticsResponseBody
}

type DescribeUniBackupStatisticsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUniBackupStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUniBackupStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUniBackupStatisticsResponse) GoString() string {
	return s.String()
}

func (s *DescribeUniBackupStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUniBackupStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUniBackupStatisticsResponse) GetBody() *DescribeUniBackupStatisticsResponseBody {
	return s.Body
}

func (s *DescribeUniBackupStatisticsResponse) SetHeaders(v map[string]*string) *DescribeUniBackupStatisticsResponse {
	s.Headers = v
	return s
}

func (s *DescribeUniBackupStatisticsResponse) SetStatusCode(v int32) *DescribeUniBackupStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUniBackupStatisticsResponse) SetBody(v *DescribeUniBackupStatisticsResponseBody) *DescribeUniBackupStatisticsResponse {
	s.Body = v
	return s
}

func (s *DescribeUniBackupStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
