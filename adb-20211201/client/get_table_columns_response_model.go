// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableColumnsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableColumnsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableColumnsResponse
	GetStatusCode() *int32
	SetBody(v *GetTableColumnsResponseBody) *GetTableColumnsResponse
	GetBody() *GetTableColumnsResponseBody
}

type GetTableColumnsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableColumnsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableColumnsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableColumnsResponse) GoString() string {
	return s.String()
}

func (s *GetTableColumnsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableColumnsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableColumnsResponse) GetBody() *GetTableColumnsResponseBody {
	return s.Body
}

func (s *GetTableColumnsResponse) SetHeaders(v map[string]*string) *GetTableColumnsResponse {
	s.Headers = v
	return s
}

func (s *GetTableColumnsResponse) SetStatusCode(v int32) *GetTableColumnsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableColumnsResponse) SetBody(v *GetTableColumnsResponseBody) *GetTableColumnsResponse {
	s.Body = v
	return s
}

func (s *GetTableColumnsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
