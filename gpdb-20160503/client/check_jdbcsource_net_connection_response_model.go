// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckJDBCSourceNetConnectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckJDBCSourceNetConnectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckJDBCSourceNetConnectionResponse
	GetStatusCode() *int32
	SetBody(v *CheckJDBCSourceNetConnectionResponseBody) *CheckJDBCSourceNetConnectionResponse
	GetBody() *CheckJDBCSourceNetConnectionResponseBody
}

type CheckJDBCSourceNetConnectionResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckJDBCSourceNetConnectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckJDBCSourceNetConnectionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckJDBCSourceNetConnectionResponse) GoString() string {
	return s.String()
}

func (s *CheckJDBCSourceNetConnectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckJDBCSourceNetConnectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckJDBCSourceNetConnectionResponse) GetBody() *CheckJDBCSourceNetConnectionResponseBody {
	return s.Body
}

func (s *CheckJDBCSourceNetConnectionResponse) SetHeaders(v map[string]*string) *CheckJDBCSourceNetConnectionResponse {
	s.Headers = v
	return s
}

func (s *CheckJDBCSourceNetConnectionResponse) SetStatusCode(v int32) *CheckJDBCSourceNetConnectionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckJDBCSourceNetConnectionResponse) SetBody(v *CheckJDBCSourceNetConnectionResponseBody) *CheckJDBCSourceNetConnectionResponse {
	s.Body = v
	return s
}

func (s *CheckJDBCSourceNetConnectionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
