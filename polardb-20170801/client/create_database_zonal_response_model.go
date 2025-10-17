// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseZonalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDatabaseZonalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDatabaseZonalResponse
	GetStatusCode() *int32
	SetBody(v *CreateDatabaseZonalResponseBody) *CreateDatabaseZonalResponse
	GetBody() *CreateDatabaseZonalResponseBody
}

type CreateDatabaseZonalResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDatabaseZonalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDatabaseZonalResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseZonalResponse) GoString() string {
	return s.String()
}

func (s *CreateDatabaseZonalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDatabaseZonalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDatabaseZonalResponse) GetBody() *CreateDatabaseZonalResponseBody {
	return s.Body
}

func (s *CreateDatabaseZonalResponse) SetHeaders(v map[string]*string) *CreateDatabaseZonalResponse {
	s.Headers = v
	return s
}

func (s *CreateDatabaseZonalResponse) SetStatusCode(v int32) *CreateDatabaseZonalResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDatabaseZonalResponse) SetBody(v *CreateDatabaseZonalResponseBody) *CreateDatabaseZonalResponse {
	s.Body = v
	return s
}

func (s *CreateDatabaseZonalResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
