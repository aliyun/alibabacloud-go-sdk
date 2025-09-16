// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTableResponse
	GetStatusCode() *int32
	SetBody(v *CreateTableResponseBody) *CreateTableResponse
	GetBody() *CreateTableResponseBody
}

type CreateTableResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTableResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTableResponse) GoString() string {
	return s.String()
}

func (s *CreateTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTableResponse) GetBody() *CreateTableResponseBody {
	return s.Body
}

func (s *CreateTableResponse) SetHeaders(v map[string]*string) *CreateTableResponse {
	s.Headers = v
	return s
}

func (s *CreateTableResponse) SetStatusCode(v int32) *CreateTableResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTableResponse) SetBody(v *CreateTableResponseBody) *CreateTableResponse {
	s.Body = v
	return s
}

func (s *CreateTableResponse) Validate() error {
	return dara.Validate(s)
}
