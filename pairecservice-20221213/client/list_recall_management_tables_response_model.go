// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecallManagementTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecallManagementTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecallManagementTablesResponse
	GetStatusCode() *int32
	SetBody(v *ListRecallManagementTablesResponseBody) *ListRecallManagementTablesResponse
	GetBody() *ListRecallManagementTablesResponseBody
}

type ListRecallManagementTablesResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecallManagementTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecallManagementTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecallManagementTablesResponse) GoString() string {
	return s.String()
}

func (s *ListRecallManagementTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecallManagementTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecallManagementTablesResponse) GetBody() *ListRecallManagementTablesResponseBody {
	return s.Body
}

func (s *ListRecallManagementTablesResponse) SetHeaders(v map[string]*string) *ListRecallManagementTablesResponse {
	s.Headers = v
	return s
}

func (s *ListRecallManagementTablesResponse) SetStatusCode(v int32) *ListRecallManagementTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecallManagementTablesResponse) SetBody(v *ListRecallManagementTablesResponseBody) *ListRecallManagementTablesResponse {
	s.Body = v
	return s
}

func (s *ListRecallManagementTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
