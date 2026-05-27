// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsTablesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateMmsTablesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateMmsTablesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateMmsTablesResponseBody) *UpdateMmsTablesResponse
	GetBody() *UpdateMmsTablesResponseBody
}

type UpdateMmsTablesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateMmsTablesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateMmsTablesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsTablesResponse) GoString() string {
	return s.String()
}

func (s *UpdateMmsTablesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateMmsTablesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateMmsTablesResponse) GetBody() *UpdateMmsTablesResponseBody {
	return s.Body
}

func (s *UpdateMmsTablesResponse) SetHeaders(v map[string]*string) *UpdateMmsTablesResponse {
	s.Headers = v
	return s
}

func (s *UpdateMmsTablesResponse) SetStatusCode(v int32) *UpdateMmsTablesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateMmsTablesResponse) SetBody(v *UpdateMmsTablesResponseBody) *UpdateMmsTablesResponse {
	s.Body = v
	return s
}

func (s *UpdateMmsTablesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
