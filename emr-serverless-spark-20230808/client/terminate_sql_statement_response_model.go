// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTerminateSqlStatementResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TerminateSqlStatementResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TerminateSqlStatementResponse
	GetStatusCode() *int32
	SetBody(v *TerminateSqlStatementResponseBody) *TerminateSqlStatementResponse
	GetBody() *TerminateSqlStatementResponseBody
}

type TerminateSqlStatementResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TerminateSqlStatementResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TerminateSqlStatementResponse) String() string {
	return dara.Prettify(s)
}

func (s TerminateSqlStatementResponse) GoString() string {
	return s.String()
}

func (s *TerminateSqlStatementResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TerminateSqlStatementResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TerminateSqlStatementResponse) GetBody() *TerminateSqlStatementResponseBody {
	return s.Body
}

func (s *TerminateSqlStatementResponse) SetHeaders(v map[string]*string) *TerminateSqlStatementResponse {
	s.Headers = v
	return s
}

func (s *TerminateSqlStatementResponse) SetStatusCode(v int32) *TerminateSqlStatementResponse {
	s.StatusCode = &v
	return s
}

func (s *TerminateSqlStatementResponse) SetBody(v *TerminateSqlStatementResponseBody) *TerminateSqlStatementResponse {
	s.Body = v
	return s
}

func (s *TerminateSqlStatementResponse) Validate() error {
	return dara.Validate(s)
}
