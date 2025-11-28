// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateResourceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateResourceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateResourceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *MigrateResourceInstanceResponseBody) *MigrateResourceInstanceResponse
	GetBody() *MigrateResourceInstanceResponseBody
}

type MigrateResourceInstanceResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateResourceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateResourceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateResourceInstanceResponse) GoString() string {
	return s.String()
}

func (s *MigrateResourceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateResourceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateResourceInstanceResponse) GetBody() *MigrateResourceInstanceResponseBody {
	return s.Body
}

func (s *MigrateResourceInstanceResponse) SetHeaders(v map[string]*string) *MigrateResourceInstanceResponse {
	s.Headers = v
	return s
}

func (s *MigrateResourceInstanceResponse) SetStatusCode(v int32) *MigrateResourceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateResourceInstanceResponse) SetBody(v *MigrateResourceInstanceResponseBody) *MigrateResourceInstanceResponse {
	s.Body = v
	return s
}

func (s *MigrateResourceInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
