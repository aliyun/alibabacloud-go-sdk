// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchMCTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchMCTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchMCTablesResponse
	GetStatusCode() *int32
	SetBody(v *SearchMCTablesResponseBody) *SearchMCTablesResponse
	GetBody() *SearchMCTablesResponseBody
}

type SearchMCTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchMCTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchMCTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchMCTablesResponse) GoString() string {
	return s.String()
}

func (s *SearchMCTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchMCTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchMCTablesResponse) GetBody() *SearchMCTablesResponseBody {
	return s.Body
}

func (s *SearchMCTablesResponse) SetHeaders(v map[string]*string) *SearchMCTablesResponse {
	s.Headers = v
	return s
}

func (s *SearchMCTablesResponse) SetStatusCode(v int32) *SearchMCTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchMCTablesResponse) SetBody(v *SearchMCTablesResponseBody) *SearchMCTablesResponse {
	s.Body = v
	return s
}

func (s *SearchMCTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
