// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDatabaseResponseBody) *ModifyDatabaseResponse
	GetBody() *ModifyDatabaseResponseBody
}

type ModifyDatabaseResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDatabaseResponse) GoString() string {
	return s.String()
}

func (s *ModifyDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDatabaseResponse) GetBody() *ModifyDatabaseResponseBody {
	return s.Body
}

func (s *ModifyDatabaseResponse) SetHeaders(v map[string]*string) *ModifyDatabaseResponse {
	s.Headers = v
	return s
}

func (s *ModifyDatabaseResponse) SetStatusCode(v int32) *ModifyDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDatabaseResponse) SetBody(v *ModifyDatabaseResponseBody) *ModifyDatabaseResponse {
	s.Body = v
	return s
}

func (s *ModifyDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
