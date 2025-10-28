// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertRoleResponse
	GetStatusCode() *int32
	SetBody(v *InsertRoleResponseBody) *InsertRoleResponse
	GetBody() *InsertRoleResponseBody
}

type InsertRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertRoleResponse) GoString() string {
	return s.String()
}

func (s *InsertRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertRoleResponse) GetBody() *InsertRoleResponseBody {
	return s.Body
}

func (s *InsertRoleResponse) SetHeaders(v map[string]*string) *InsertRoleResponse {
	s.Headers = v
	return s
}

func (s *InsertRoleResponse) SetStatusCode(v int32) *InsertRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertRoleResponse) SetBody(v *InsertRoleResponseBody) *InsertRoleResponse {
	s.Body = v
	return s
}

func (s *InsertRoleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
