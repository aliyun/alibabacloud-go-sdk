// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossRegionBackupDBInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCrossRegionBackupDBInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCrossRegionBackupDBInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCrossRegionBackupDBInstanceResponseBody) *DescribeCrossRegionBackupDBInstanceResponse
	GetBody() *DescribeCrossRegionBackupDBInstanceResponseBody
}

type DescribeCrossRegionBackupDBInstanceResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrossRegionBackupDBInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrossRegionBackupDBInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossRegionBackupDBInstanceResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrossRegionBackupDBInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCrossRegionBackupDBInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCrossRegionBackupDBInstanceResponse) GetBody() *DescribeCrossRegionBackupDBInstanceResponseBody {
	return s.Body
}

func (s *DescribeCrossRegionBackupDBInstanceResponse) SetHeaders(v map[string]*string) *DescribeCrossRegionBackupDBInstanceResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponse) SetStatusCode(v int32) *DescribeCrossRegionBackupDBInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponse) SetBody(v *DescribeCrossRegionBackupDBInstanceResponseBody) *DescribeCrossRegionBackupDBInstanceResponse {
	s.Body = v
	return s
}

func (s *DescribeCrossRegionBackupDBInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
