// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckBatchTableAccessPermissionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckBatchTableAccessPermissionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckBatchTableAccessPermissionResponse
	GetStatusCode() *int32
	SetBody(v *CheckBatchTableAccessPermissionResponseBody) *CheckBatchTableAccessPermissionResponse
	GetBody() *CheckBatchTableAccessPermissionResponseBody
}

type CheckBatchTableAccessPermissionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckBatchTableAccessPermissionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckBatchTableAccessPermissionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckBatchTableAccessPermissionResponse) GoString() string {
	return s.String()
}

func (s *CheckBatchTableAccessPermissionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckBatchTableAccessPermissionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckBatchTableAccessPermissionResponse) GetBody() *CheckBatchTableAccessPermissionResponseBody {
	return s.Body
}

func (s *CheckBatchTableAccessPermissionResponse) SetHeaders(v map[string]*string) *CheckBatchTableAccessPermissionResponse {
	s.Headers = v
	return s
}

func (s *CheckBatchTableAccessPermissionResponse) SetStatusCode(v int32) *CheckBatchTableAccessPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckBatchTableAccessPermissionResponse) SetBody(v *CheckBatchTableAccessPermissionResponseBody) *CheckBatchTableAccessPermissionResponse {
	s.Body = v
	return s
}

func (s *CheckBatchTableAccessPermissionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
