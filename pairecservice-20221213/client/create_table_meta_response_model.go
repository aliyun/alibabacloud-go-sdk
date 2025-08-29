// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTableMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTableMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTableMetaResponse
	GetStatusCode() *int32
	SetBody(v *CreateTableMetaResponseBody) *CreateTableMetaResponse
	GetBody() *CreateTableMetaResponseBody
}

type CreateTableMetaResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTableMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTableMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTableMetaResponse) GoString() string {
	return s.String()
}

func (s *CreateTableMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTableMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTableMetaResponse) GetBody() *CreateTableMetaResponseBody {
	return s.Body
}

func (s *CreateTableMetaResponse) SetHeaders(v map[string]*string) *CreateTableMetaResponse {
	s.Headers = v
	return s
}

func (s *CreateTableMetaResponse) SetStatusCode(v int32) *CreateTableMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTableMetaResponse) SetBody(v *CreateTableMetaResponseBody) *CreateTableMetaResponse {
	s.Body = v
	return s
}

func (s *CreateTableMetaResponse) Validate() error {
	return dara.Validate(s)
}
