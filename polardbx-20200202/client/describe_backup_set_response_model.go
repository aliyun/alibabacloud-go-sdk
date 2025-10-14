// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackupSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackupSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackupSetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackupSetResponseBody) *DescribeBackupSetResponse
	GetBody() *DescribeBackupSetResponseBody
}

type DescribeBackupSetResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackupSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackupSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackupSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackupSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackupSetResponse) GetBody() *DescribeBackupSetResponseBody {
	return s.Body
}

func (s *DescribeBackupSetResponse) SetHeaders(v map[string]*string) *DescribeBackupSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackupSetResponse) SetStatusCode(v int32) *DescribeBackupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackupSetResponse) SetBody(v *DescribeBackupSetResponseBody) *DescribeBackupSetResponse {
	s.Body = v
	return s
}

func (s *DescribeBackupSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
