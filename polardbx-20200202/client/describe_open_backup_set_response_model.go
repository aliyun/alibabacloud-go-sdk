// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenBackupSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenBackupSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenBackupSetResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenBackupSetResponseBody) *DescribeOpenBackupSetResponse
	GetBody() *DescribeOpenBackupSetResponseBody
}

type DescribeOpenBackupSetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenBackupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenBackupSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenBackupSetResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenBackupSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenBackupSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenBackupSetResponse) GetBody() *DescribeOpenBackupSetResponseBody {
	return s.Body
}

func (s *DescribeOpenBackupSetResponse) SetHeaders(v map[string]*string) *DescribeOpenBackupSetResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenBackupSetResponse) SetStatusCode(v int32) *DescribeOpenBackupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenBackupSetResponse) SetBody(v *DescribeOpenBackupSetResponseBody) *DescribeOpenBackupSetResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenBackupSetResponse) Validate() error {
	return dara.Validate(s)
}
