// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIncrementBackupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIncrementBackupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIncrementBackupListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIncrementBackupListResponseBody) *DescribeIncrementBackupListResponse
	GetBody() *DescribeIncrementBackupListResponseBody
}

type DescribeIncrementBackupListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIncrementBackupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIncrementBackupListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIncrementBackupListResponse) GoString() string {
	return s.String()
}

func (s *DescribeIncrementBackupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIncrementBackupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIncrementBackupListResponse) GetBody() *DescribeIncrementBackupListResponseBody {
	return s.Body
}

func (s *DescribeIncrementBackupListResponse) SetHeaders(v map[string]*string) *DescribeIncrementBackupListResponse {
	s.Headers = v
	return s
}

func (s *DescribeIncrementBackupListResponse) SetStatusCode(v int32) *DescribeIncrementBackupListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIncrementBackupListResponse) SetBody(v *DescribeIncrementBackupListResponseBody) *DescribeIncrementBackupListResponse {
	s.Body = v
	return s
}

func (s *DescribeIncrementBackupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
