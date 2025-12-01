// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupStrategyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupStrategyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupStrategyResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupStrategyResponseBody) *ModifyBackupStrategyResponse
	GetBody() *ModifyBackupStrategyResponseBody
}

type ModifyBackupStrategyResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupStrategyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupStrategyResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupStrategyResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupStrategyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupStrategyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupStrategyResponse) GetBody() *ModifyBackupStrategyResponseBody {
	return s.Body
}

func (s *ModifyBackupStrategyResponse) SetHeaders(v map[string]*string) *ModifyBackupStrategyResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupStrategyResponse) SetStatusCode(v int32) *ModifyBackupStrategyResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupStrategyResponse) SetBody(v *ModifyBackupStrategyResponseBody) *ModifyBackupStrategyResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupStrategyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
