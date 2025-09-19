// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBackupRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListBackupRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListBackupRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListBackupRecordResponseBody) *ListBackupRecordResponse
	GetBody() *ListBackupRecordResponseBody
}

type ListBackupRecordResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListBackupRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListBackupRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListBackupRecordResponse) GoString() string {
	return s.String()
}

func (s *ListBackupRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListBackupRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListBackupRecordResponse) GetBody() *ListBackupRecordResponseBody {
	return s.Body
}

func (s *ListBackupRecordResponse) SetHeaders(v map[string]*string) *ListBackupRecordResponse {
	s.Headers = v
	return s
}

func (s *ListBackupRecordResponse) SetStatusCode(v int32) *ListBackupRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListBackupRecordResponse) SetBody(v *ListBackupRecordResponseBody) *ListBackupRecordResponse {
	s.Body = v
	return s
}

func (s *ListBackupRecordResponse) Validate() error {
	return dara.Validate(s)
}
