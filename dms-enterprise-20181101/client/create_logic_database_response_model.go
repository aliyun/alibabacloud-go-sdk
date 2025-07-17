// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogicDatabaseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLogicDatabaseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLogicDatabaseResponse
	GetStatusCode() *int32
	SetBody(v *CreateLogicDatabaseResponseBody) *CreateLogicDatabaseResponse
	GetBody() *CreateLogicDatabaseResponseBody
}

type CreateLogicDatabaseResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogicDatabaseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogicDatabaseResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLogicDatabaseResponse) GoString() string {
	return s.String()
}

func (s *CreateLogicDatabaseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLogicDatabaseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLogicDatabaseResponse) GetBody() *CreateLogicDatabaseResponseBody {
	return s.Body
}

func (s *CreateLogicDatabaseResponse) SetHeaders(v map[string]*string) *CreateLogicDatabaseResponse {
	s.Headers = v
	return s
}

func (s *CreateLogicDatabaseResponse) SetStatusCode(v int32) *CreateLogicDatabaseResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogicDatabaseResponse) SetBody(v *CreateLogicDatabaseResponseBody) *CreateLogicDatabaseResponse {
	s.Body = v
	return s
}

func (s *CreateLogicDatabaseResponse) Validate() error {
	return dara.Validate(s)
}
