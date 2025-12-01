// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupSetDownloadRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupSetDownloadRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupSetDownloadRulesResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupSetDownloadRulesResponseBody) *ModifyBackupSetDownloadRulesResponse
	GetBody() *ModifyBackupSetDownloadRulesResponseBody
}

type ModifyBackupSetDownloadRulesResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupSetDownloadRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupSetDownloadRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupSetDownloadRulesResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupSetDownloadRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupSetDownloadRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupSetDownloadRulesResponse) GetBody() *ModifyBackupSetDownloadRulesResponseBody {
	return s.Body
}

func (s *ModifyBackupSetDownloadRulesResponse) SetHeaders(v map[string]*string) *ModifyBackupSetDownloadRulesResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponse) SetStatusCode(v int32) *ModifyBackupSetDownloadRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponse) SetBody(v *ModifyBackupSetDownloadRulesResponseBody) *ModifyBackupSetDownloadRulesResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupSetDownloadRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
