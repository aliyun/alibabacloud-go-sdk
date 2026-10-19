// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvancedBackupStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvancedBackupStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvancedBackupStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvancedBackupStatusResponseBody) *DescribeAdvancedBackupStatusResponse
	GetBody() *DescribeAdvancedBackupStatusResponseBody
}

type DescribeAdvancedBackupStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvancedBackupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvancedBackupStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvancedBackupStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvancedBackupStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvancedBackupStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvancedBackupStatusResponse) GetBody() *DescribeAdvancedBackupStatusResponseBody {
	return s.Body
}

func (s *DescribeAdvancedBackupStatusResponse) SetHeaders(v map[string]*string) *DescribeAdvancedBackupStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvancedBackupStatusResponse) SetStatusCode(v int32) *DescribeAdvancedBackupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvancedBackupStatusResponse) SetBody(v *DescribeAdvancedBackupStatusResponseBody) *DescribeAdvancedBackupStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvancedBackupStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
