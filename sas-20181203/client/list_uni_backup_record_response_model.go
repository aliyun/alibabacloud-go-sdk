// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUniBackupRecordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUniBackupRecordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUniBackupRecordResponse
	GetStatusCode() *int32
	SetBody(v *ListUniBackupRecordResponseBody) *ListUniBackupRecordResponse
	GetBody() *ListUniBackupRecordResponseBody
}

type ListUniBackupRecordResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUniBackupRecordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUniBackupRecordResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUniBackupRecordResponse) GoString() string {
	return s.String()
}

func (s *ListUniBackupRecordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUniBackupRecordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUniBackupRecordResponse) GetBody() *ListUniBackupRecordResponseBody {
	return s.Body
}

func (s *ListUniBackupRecordResponse) SetHeaders(v map[string]*string) *ListUniBackupRecordResponse {
	s.Headers = v
	return s
}

func (s *ListUniBackupRecordResponse) SetStatusCode(v int32) *ListUniBackupRecordResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUniBackupRecordResponse) SetBody(v *ListUniBackupRecordResponseBody) *ListUniBackupRecordResponse {
	s.Body = v
	return s
}

func (s *ListUniBackupRecordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
