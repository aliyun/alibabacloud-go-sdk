// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMetaTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMetaTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMetaTablesResponse
	GetStatusCode() *int32
	SetBody(v *SearchMetaTablesResponseBody) *SearchMetaTablesResponse
	GetBody() *SearchMetaTablesResponseBody
}

type SearchMetaTablesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMetaTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMetaTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMetaTablesResponse) GoString() string {
	return s.String()
}

func (s *SearchMetaTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMetaTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMetaTablesResponse) GetBody() *SearchMetaTablesResponseBody {
	return s.Body
}

func (s *SearchMetaTablesResponse) SetHeaders(v map[string]*string) *SearchMetaTablesResponse {
	s.Headers = v
	return s
}

func (s *SearchMetaTablesResponse) SetStatusCode(v int32) *SearchMetaTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMetaTablesResponse) SetBody(v *SearchMetaTablesResponseBody) *SearchMetaTablesResponse {
	s.Body = v
	return s
}

func (s *SearchMetaTablesResponse) Validate() error {
	return dara.Validate(s)
}
