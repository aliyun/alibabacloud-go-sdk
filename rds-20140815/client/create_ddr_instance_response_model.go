// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDdrInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDdrInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDdrInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateDdrInstanceResponseBody) *CreateDdrInstanceResponse
	GetBody() *CreateDdrInstanceResponseBody
}

type CreateDdrInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDdrInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDdrInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDdrInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateDdrInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDdrInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDdrInstanceResponse) GetBody() *CreateDdrInstanceResponseBody {
	return s.Body
}

func (s *CreateDdrInstanceResponse) SetHeaders(v map[string]*string) *CreateDdrInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateDdrInstanceResponse) SetStatusCode(v int32) *CreateDdrInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDdrInstanceResponse) SetBody(v *CreateDdrInstanceResponseBody) *CreateDdrInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateDdrInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
