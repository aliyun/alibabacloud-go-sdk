// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDepartmentUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SyncDepartmentUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SyncDepartmentUserResponse
	GetStatusCode() *int32
	SetBody(v *SyncDepartmentUserResponseBody) *SyncDepartmentUserResponse
	GetBody() *SyncDepartmentUserResponseBody
}

type SyncDepartmentUserResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SyncDepartmentUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SyncDepartmentUserResponse) String() string {
	return dara.Prettify(s)
}

func (s SyncDepartmentUserResponse) GoString() string {
	return s.String()
}

func (s *SyncDepartmentUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SyncDepartmentUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SyncDepartmentUserResponse) GetBody() *SyncDepartmentUserResponseBody {
	return s.Body
}

func (s *SyncDepartmentUserResponse) SetHeaders(v map[string]*string) *SyncDepartmentUserResponse {
	s.Headers = v
	return s
}

func (s *SyncDepartmentUserResponse) SetStatusCode(v int32) *SyncDepartmentUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SyncDepartmentUserResponse) SetBody(v *SyncDepartmentUserResponseBody) *SyncDepartmentUserResponse {
	s.Body = v
	return s
}

func (s *SyncDepartmentUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
