// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDownloadBackupSetStorageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDownloadBackupSetStorageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDownloadBackupSetStorageInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDownloadBackupSetStorageInfoResponseBody) *DescribeDownloadBackupSetStorageInfoResponse
	GetBody() *DescribeDownloadBackupSetStorageInfoResponseBody
}

type DescribeDownloadBackupSetStorageInfoResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDownloadBackupSetStorageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDownloadBackupSetStorageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDownloadBackupSetStorageInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) GetBody() *DescribeDownloadBackupSetStorageInfoResponseBody {
	return s.Body
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) SetHeaders(v map[string]*string) *DescribeDownloadBackupSetStorageInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) SetStatusCode(v int32) *DescribeDownloadBackupSetStorageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) SetBody(v *DescribeDownloadBackupSetStorageInfoResponseBody) *DescribeDownloadBackupSetStorageInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeDownloadBackupSetStorageInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
