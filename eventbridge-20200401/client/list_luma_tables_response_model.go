// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLumaTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLumaTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLumaTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListLumaTablesResponseBody) *ListLumaTablesResponse
	GetBody() *ListLumaTablesResponseBody
}

type ListLumaTablesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLumaTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLumaTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLumaTablesResponse) GoString() string {
	return s.String()
}

func (s *ListLumaTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLumaTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLumaTablesResponse) GetBody() *ListLumaTablesResponseBody {
	return s.Body
}

func (s *ListLumaTablesResponse) SetHeaders(v map[string]*string) *ListLumaTablesResponse {
	s.Headers = v
	return s
}

func (s *ListLumaTablesResponse) SetStatusCode(v int32) *ListLumaTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLumaTablesResponse) SetBody(v *ListLumaTablesResponseBody) *ListLumaTablesResponse {
	s.Body = v
	return s
}

func (s *ListLumaTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
