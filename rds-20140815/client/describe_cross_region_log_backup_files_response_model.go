// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionLogBackupFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCrossRegionLogBackupFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCrossRegionLogBackupFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCrossRegionLogBackupFilesResponseBody) *DescribeCrossRegionLogBackupFilesResponse
	GetBody() *DescribeCrossRegionLogBackupFilesResponseBody
}

type DescribeCrossRegionLogBackupFilesResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrossRegionLogBackupFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrossRegionLogBackupFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionLogBackupFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionLogBackupFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCrossRegionLogBackupFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCrossRegionLogBackupFilesResponse) GetBody() *DescribeCrossRegionLogBackupFilesResponseBody {
	return s.Body
}

func (s *DescribeCrossRegionLogBackupFilesResponse) SetHeaders(v map[string]*string) *DescribeCrossRegionLogBackupFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponse) SetStatusCode(v int32) *DescribeCrossRegionLogBackupFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponse) SetBody(v *DescribeCrossRegionLogBackupFilesResponseBody) *DescribeCrossRegionLogBackupFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeCrossRegionLogBackupFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
