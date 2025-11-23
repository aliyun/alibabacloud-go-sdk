// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPhysicalDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPhysicalDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPhysicalDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *GetPhysicalDatabaseResponseBody) *GetPhysicalDatabaseResponse
	GetBody() *GetPhysicalDatabaseResponseBody
}

type GetPhysicalDatabaseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPhysicalDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPhysicalDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPhysicalDatabaseResponse) GoString() string {
	return s.String()
}

func (s *GetPhysicalDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPhysicalDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPhysicalDatabaseResponse) GetBody() *GetPhysicalDatabaseResponseBody {
	return s.Body
}

func (s *GetPhysicalDatabaseResponse) SetHeaders(v map[string]*string) *GetPhysicalDatabaseResponse {
	s.Headers = v
	return s
}

func (s *GetPhysicalDatabaseResponse) SetStatusCode(v int32) *GetPhysicalDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPhysicalDatabaseResponse) SetBody(v *GetPhysicalDatabaseResponseBody) *GetPhysicalDatabaseResponse {
	s.Body = v
	return s
}

func (s *GetPhysicalDatabaseResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
