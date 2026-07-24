// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckSqlContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckSqlContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckSqlContentResponse
	GetStatusCode() *int32
	SetBody(v *CheckSqlContentResponseBody) *CheckSqlContentResponse
	GetBody() *CheckSqlContentResponseBody
}

type CheckSqlContentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckSqlContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckSqlContentResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckSqlContentResponse) GoString() string {
	return s.String()
}

func (s *CheckSqlContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckSqlContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckSqlContentResponse) GetBody() *CheckSqlContentResponseBody {
	return s.Body
}

func (s *CheckSqlContentResponse) SetHeaders(v map[string]*string) *CheckSqlContentResponse {
	s.Headers = v
	return s
}

func (s *CheckSqlContentResponse) SetStatusCode(v int32) *CheckSqlContentResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckSqlContentResponse) SetBody(v *CheckSqlContentResponseBody) *CheckSqlContentResponse {
	s.Body = v
	return s
}

func (s *CheckSqlContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
