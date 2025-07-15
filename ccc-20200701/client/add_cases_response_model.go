// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCasesResponse
	GetStatusCode() *int32
	SetBody(v *AddCasesResponseBody) *AddCasesResponse
	GetBody() *AddCasesResponseBody
}

type AddCasesResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCasesResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCasesResponse) GoString() string {
	return s.String()
}

func (s *AddCasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCasesResponse) GetBody() *AddCasesResponseBody {
	return s.Body
}

func (s *AddCasesResponse) SetHeaders(v map[string]*string) *AddCasesResponse {
	s.Headers = v
	return s
}

func (s *AddCasesResponse) SetStatusCode(v int32) *AddCasesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCasesResponse) SetBody(v *AddCasesResponseBody) *AddCasesResponse {
	s.Body = v
	return s
}

func (s *AddCasesResponse) Validate() error {
	return dara.Validate(s)
}
