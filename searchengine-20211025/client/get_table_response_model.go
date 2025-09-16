// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTableResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTableResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTableResponse
	GetStatusCode() *int32
	SetBody(v *GetTableResponseBody) *GetTableResponse
	GetBody() *GetTableResponseBody
}

type GetTableResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTableResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTableResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTableResponse) GoString() string {
	return s.String()
}

func (s *GetTableResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTableResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTableResponse) GetBody() *GetTableResponseBody {
	return s.Body
}

func (s *GetTableResponse) SetHeaders(v map[string]*string) *GetTableResponse {
	s.Headers = v
	return s
}

func (s *GetTableResponse) SetStatusCode(v int32) *GetTableResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTableResponse) SetBody(v *GetTableResponseBody) *GetTableResponse {
	s.Body = v
	return s
}

func (s *GetTableResponse) Validate() error {
	return dara.Validate(s)
}
