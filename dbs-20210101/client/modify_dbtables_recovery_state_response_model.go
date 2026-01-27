// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBTablesRecoveryStateResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBTablesRecoveryStateResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBTablesRecoveryStateResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBTablesRecoveryStateResponseBody) *ModifyDBTablesRecoveryStateResponse
	GetBody() *ModifyDBTablesRecoveryStateResponseBody
}

type ModifyDBTablesRecoveryStateResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBTablesRecoveryStateResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBTablesRecoveryStateResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBTablesRecoveryStateResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBTablesRecoveryStateResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBTablesRecoveryStateResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBTablesRecoveryStateResponse) GetBody() *ModifyDBTablesRecoveryStateResponseBody {
	return s.Body
}

func (s *ModifyDBTablesRecoveryStateResponse) SetHeaders(v map[string]*string) *ModifyDBTablesRecoveryStateResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponse) SetStatusCode(v int32) *ModifyDBTablesRecoveryStateResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponse) SetBody(v *ModifyDBTablesRecoveryStateResponseBody) *ModifyDBTablesRecoveryStateResponse {
	s.Body = v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
