// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDepartmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncDepartmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncDepartmentResponse
	GetStatusCode() *int32
	SetBody(v *SyncDepartmentResponseBody) *SyncDepartmentResponse
	GetBody() *SyncDepartmentResponseBody
}

type SyncDepartmentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDepartmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDepartmentResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentResponse) GoString() string {
	return s.String()
}

func (s *SyncDepartmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncDepartmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncDepartmentResponse) GetBody() *SyncDepartmentResponseBody {
	return s.Body
}

func (s *SyncDepartmentResponse) SetHeaders(v map[string]*string) *SyncDepartmentResponse {
	s.Headers = v
	return s
}

func (s *SyncDepartmentResponse) SetStatusCode(v int32) *SyncDepartmentResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDepartmentResponse) SetBody(v *SyncDepartmentResponseBody) *SyncDepartmentResponse {
	s.Body = v
	return s
}

func (s *SyncDepartmentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
