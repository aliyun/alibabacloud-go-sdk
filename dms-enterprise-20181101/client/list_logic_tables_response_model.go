// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogicTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogicTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListLogicTablesResponseBody) *ListLogicTablesResponse
	GetBody() *ListLogicTablesResponseBody
}

type ListLogicTablesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogicTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogicTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogicTablesResponse) GoString() string {
	return s.String()
}

func (s *ListLogicTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogicTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogicTablesResponse) GetBody() *ListLogicTablesResponseBody {
	return s.Body
}

func (s *ListLogicTablesResponse) SetHeaders(v map[string]*string) *ListLogicTablesResponse {
	s.Headers = v
	return s
}

func (s *ListLogicTablesResponse) SetStatusCode(v int32) *ListLogicTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogicTablesResponse) SetBody(v *ListLogicTablesResponseBody) *ListLogicTablesResponse {
	s.Body = v
	return s
}

func (s *ListLogicTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
