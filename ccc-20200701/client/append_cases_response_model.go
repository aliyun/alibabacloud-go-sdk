// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAppendCasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AppendCasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AppendCasesResponse
	GetStatusCode() *int32
	SetBody(v *AppendCasesResponseBody) *AppendCasesResponse
	GetBody() *AppendCasesResponseBody
}

type AppendCasesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AppendCasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AppendCasesResponse) String() string {
	return dara.Prettify(s)
}

func (s AppendCasesResponse) GoString() string {
	return s.String()
}

func (s *AppendCasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AppendCasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AppendCasesResponse) GetBody() *AppendCasesResponseBody {
	return s.Body
}

func (s *AppendCasesResponse) SetHeaders(v map[string]*string) *AppendCasesResponse {
	s.Headers = v
	return s
}

func (s *AppendCasesResponse) SetStatusCode(v int32) *AppendCasesResponse {
	s.StatusCode = &v
	return s
}

func (s *AppendCasesResponse) SetBody(v *AppendCasesResponseBody) *AppendCasesResponse {
	s.Body = v
	return s
}

func (s *AppendCasesResponse) Validate() error {
	return dara.Validate(s)
}
