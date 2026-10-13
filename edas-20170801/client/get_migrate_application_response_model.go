// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrateApplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMigrateApplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMigrateApplicationResponse
	GetStatusCode() *int32
	SetBody(v *GetMigrateApplicationResponseBody) *GetMigrateApplicationResponse
	GetBody() *GetMigrateApplicationResponseBody
}

type GetMigrateApplicationResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMigrateApplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMigrateApplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMigrateApplicationResponse) GoString() string {
	return s.String()
}

func (s *GetMigrateApplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMigrateApplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMigrateApplicationResponse) GetBody() *GetMigrateApplicationResponseBody {
	return s.Body
}

func (s *GetMigrateApplicationResponse) SetHeaders(v map[string]*string) *GetMigrateApplicationResponse {
	s.Headers = v
	return s
}

func (s *GetMigrateApplicationResponse) SetStatusCode(v int32) *GetMigrateApplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMigrateApplicationResponse) SetBody(v *GetMigrateApplicationResponseBody) *GetMigrateApplicationResponse {
	s.Body = v
	return s
}

func (s *GetMigrateApplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
