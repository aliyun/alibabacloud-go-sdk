// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvancedBackupStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAdvancedBackupStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAdvancedBackupStatusResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAdvancedBackupStatusResponseBody) *ModifyAdvancedBackupStatusResponse
	GetBody() *ModifyAdvancedBackupStatusResponseBody
}

type ModifyAdvancedBackupStatusResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAdvancedBackupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAdvancedBackupStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvancedBackupStatusResponse) GoString() string {
	return s.String()
}

func (s *ModifyAdvancedBackupStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAdvancedBackupStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAdvancedBackupStatusResponse) GetBody() *ModifyAdvancedBackupStatusResponseBody {
	return s.Body
}

func (s *ModifyAdvancedBackupStatusResponse) SetHeaders(v map[string]*string) *ModifyAdvancedBackupStatusResponse {
	s.Headers = v
	return s
}

func (s *ModifyAdvancedBackupStatusResponse) SetStatusCode(v int32) *ModifyAdvancedBackupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAdvancedBackupStatusResponse) SetBody(v *ModifyAdvancedBackupStatusResponseBody) *ModifyAdvancedBackupStatusResponse {
	s.Body = v
	return s
}

func (s *ModifyAdvancedBackupStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
