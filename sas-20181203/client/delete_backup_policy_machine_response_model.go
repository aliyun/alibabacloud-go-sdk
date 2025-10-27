// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPolicyMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupPolicyMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupPolicyMachineResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupPolicyMachineResponseBody) *DeleteBackupPolicyMachineResponse
	GetBody() *DeleteBackupPolicyMachineResponseBody
}

type DeleteBackupPolicyMachineResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupPolicyMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupPolicyMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPolicyMachineResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupPolicyMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupPolicyMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupPolicyMachineResponse) GetBody() *DeleteBackupPolicyMachineResponseBody {
	return s.Body
}

func (s *DeleteBackupPolicyMachineResponse) SetHeaders(v map[string]*string) *DeleteBackupPolicyMachineResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupPolicyMachineResponse) SetStatusCode(v int32) *DeleteBackupPolicyMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupPolicyMachineResponse) SetBody(v *DeleteBackupPolicyMachineResponseBody) *DeleteBackupPolicyMachineResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupPolicyMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
