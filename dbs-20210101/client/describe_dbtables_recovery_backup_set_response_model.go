// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBTablesRecoveryBackupSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBTablesRecoveryBackupSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBTablesRecoveryBackupSetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBTablesRecoveryBackupSetResponseBody) *DescribeDBTablesRecoveryBackupSetResponse
	GetBody() *DescribeDBTablesRecoveryBackupSetResponseBody
}

type DescribeDBTablesRecoveryBackupSetResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBTablesRecoveryBackupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBTablesRecoveryBackupSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBTablesRecoveryBackupSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) GetBody() *DescribeDBTablesRecoveryBackupSetResponseBody {
	return s.Body
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) SetHeaders(v map[string]*string) *DescribeDBTablesRecoveryBackupSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) SetStatusCode(v int32) *DescribeDBTablesRecoveryBackupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) SetBody(v *DescribeDBTablesRecoveryBackupSetResponseBody) *DescribeDBTablesRecoveryBackupSetResponse {
	s.Body = v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
