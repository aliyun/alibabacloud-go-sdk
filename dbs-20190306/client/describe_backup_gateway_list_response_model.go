// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupGatewayListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupGatewayListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupGatewayListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupGatewayListResponseBody) *DescribeBackupGatewayListResponse
	GetBody() *DescribeBackupGatewayListResponseBody
}

type DescribeBackupGatewayListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupGatewayListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupGatewayListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupGatewayListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupGatewayListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupGatewayListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupGatewayListResponse) GetBody() *DescribeBackupGatewayListResponseBody {
	return s.Body
}

func (s *DescribeBackupGatewayListResponse) SetHeaders(v map[string]*string) *DescribeBackupGatewayListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupGatewayListResponse) SetStatusCode(v int32) *DescribeBackupGatewayListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupGatewayListResponse) SetBody(v *DescribeBackupGatewayListResponseBody) *DescribeBackupGatewayListResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupGatewayListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
