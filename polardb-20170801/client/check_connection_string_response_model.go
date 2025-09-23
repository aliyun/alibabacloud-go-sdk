// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckConnectionStringResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckConnectionStringResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckConnectionStringResponse
	GetStatusCode() *int32
	SetBody(v *CheckConnectionStringResponseBody) *CheckConnectionStringResponse
	GetBody() *CheckConnectionStringResponseBody
}

type CheckConnectionStringResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckConnectionStringResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckConnectionStringResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckConnectionStringResponse) GoString() string {
	return s.String()
}

func (s *CheckConnectionStringResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckConnectionStringResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckConnectionStringResponse) GetBody() *CheckConnectionStringResponseBody {
	return s.Body
}

func (s *CheckConnectionStringResponse) SetHeaders(v map[string]*string) *CheckConnectionStringResponse {
	s.Headers = v
	return s
}

func (s *CheckConnectionStringResponse) SetStatusCode(v int32) *CheckConnectionStringResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckConnectionStringResponse) SetBody(v *CheckConnectionStringResponseBody) *CheckConnectionStringResponse {
	s.Body = v
	return s
}

func (s *CheckConnectionStringResponse) Validate() error {
	return dara.Validate(s)
}
