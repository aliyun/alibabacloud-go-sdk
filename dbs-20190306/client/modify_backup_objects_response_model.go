// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupObjectsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupObjectsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupObjectsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupObjectsResponseBody) *ModifyBackupObjectsResponse
	GetBody() *ModifyBackupObjectsResponseBody
}

type ModifyBackupObjectsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupObjectsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupObjectsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupObjectsResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupObjectsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupObjectsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupObjectsResponse) GetBody() *ModifyBackupObjectsResponseBody {
	return s.Body
}

func (s *ModifyBackupObjectsResponse) SetHeaders(v map[string]*string) *ModifyBackupObjectsResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupObjectsResponse) SetStatusCode(v int32) *ModifyBackupObjectsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupObjectsResponse) SetBody(v *ModifyBackupObjectsResponseBody) *ModifyBackupObjectsResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupObjectsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
