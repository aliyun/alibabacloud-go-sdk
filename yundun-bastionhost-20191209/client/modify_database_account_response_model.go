// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDatabaseAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDatabaseAccountResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDatabaseAccountResponseBody) *ModifyDatabaseAccountResponse
	GetBody() *ModifyDatabaseAccountResponseBody
}

type ModifyDatabaseAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabaseAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseAccountResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDatabaseAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDatabaseAccountResponse) GetBody() *ModifyDatabaseAccountResponseBody {
	return s.Body
}

func (s *ModifyDatabaseAccountResponse) SetHeaders(v map[string]*string) *ModifyDatabaseAccountResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseAccountResponse) SetStatusCode(v int32) *ModifyDatabaseAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseAccountResponse) SetBody(v *ModifyDatabaseAccountResponseBody) *ModifyDatabaseAccountResponse {
	s.Body = v
	return s
}

func (s *ModifyDatabaseAccountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
