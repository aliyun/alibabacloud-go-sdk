// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEngineMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyEngineMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyEngineMigrationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyEngineMigrationResponseBody) *ModifyEngineMigrationResponse
	GetBody() *ModifyEngineMigrationResponseBody
}

type ModifyEngineMigrationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyEngineMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyEngineMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyEngineMigrationResponse) GoString() string {
	return s.String()
}

func (s *ModifyEngineMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyEngineMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyEngineMigrationResponse) GetBody() *ModifyEngineMigrationResponseBody {
	return s.Body
}

func (s *ModifyEngineMigrationResponse) SetHeaders(v map[string]*string) *ModifyEngineMigrationResponse {
	s.Headers = v
	return s
}

func (s *ModifyEngineMigrationResponse) SetStatusCode(v int32) *ModifyEngineMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyEngineMigrationResponse) SetBody(v *ModifyEngineMigrationResponseBody) *ModifyEngineMigrationResponse {
	s.Body = v
	return s
}

func (s *ModifyEngineMigrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
