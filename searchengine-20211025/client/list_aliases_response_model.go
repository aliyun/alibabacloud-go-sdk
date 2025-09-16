// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAliasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAliasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAliasesResponse
	GetStatusCode() *int32
	SetBody(v *ListAliasesResponseBody) *ListAliasesResponse
	GetBody() *ListAliasesResponseBody
}

type ListAliasesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAliasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAliasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAliasesResponse) GoString() string {
	return s.String()
}

func (s *ListAliasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAliasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAliasesResponse) GetBody() *ListAliasesResponseBody {
	return s.Body
}

func (s *ListAliasesResponse) SetHeaders(v map[string]*string) *ListAliasesResponse {
	s.Headers = v
	return s
}

func (s *ListAliasesResponse) SetStatusCode(v int32) *ListAliasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAliasesResponse) SetBody(v *ListAliasesResponseBody) *ListAliasesResponse {
	s.Body = v
	return s
}

func (s *ListAliasesResponse) Validate() error {
	return dara.Validate(s)
}
