// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceResourceTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceResourceTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceResourceTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceResourceTablesResponseBody) *ListInstanceResourceTablesResponse
	GetBody() *ListInstanceResourceTablesResponseBody
}

type ListInstanceResourceTablesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceResourceTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceResourceTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceResourceTablesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceResourceTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceResourceTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceResourceTablesResponse) GetBody() *ListInstanceResourceTablesResponseBody {
	return s.Body
}

func (s *ListInstanceResourceTablesResponse) SetHeaders(v map[string]*string) *ListInstanceResourceTablesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceResourceTablesResponse) SetStatusCode(v int32) *ListInstanceResourceTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceResourceTablesResponse) SetBody(v *ListInstanceResourceTablesResponseBody) *ListInstanceResourceTablesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceResourceTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
