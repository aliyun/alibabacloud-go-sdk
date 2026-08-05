// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTablesResponse
	GetStatusCode() *int32
	SetBody(v *GetTablesResponseBody) *GetTablesResponse
	GetBody() *GetTablesResponseBody
}

type GetTablesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTablesResponse) GoString() string {
	return s.String()
}

func (s *GetTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTablesResponse) GetBody() *GetTablesResponseBody {
	return s.Body
}

func (s *GetTablesResponse) SetHeaders(v map[string]*string) *GetTablesResponse {
	s.Headers = v
	return s
}

func (s *GetTablesResponse) SetStatusCode(v int32) *GetTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTablesResponse) SetBody(v *GetTablesResponseBody) *GetTablesResponse {
	s.Body = v
	return s
}

func (s *GetTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
