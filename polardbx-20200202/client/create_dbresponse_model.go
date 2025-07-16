// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDBResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDBResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDBResponse
	GetStatusCode() *int32
	SetBody(v *CreateDBResponseBody) *CreateDBResponse
	GetBody() *CreateDBResponseBody
}

type CreateDBResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDBResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDBResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDBResponse) GoString() string {
	return s.String()
}

func (s *CreateDBResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDBResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDBResponse) GetBody() *CreateDBResponseBody {
	return s.Body
}

func (s *CreateDBResponse) SetHeaders(v map[string]*string) *CreateDBResponse {
	s.Headers = v
	return s
}

func (s *CreateDBResponse) SetStatusCode(v int32) *CreateDBResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDBResponse) SetBody(v *CreateDBResponseBody) *CreateDBResponse {
	s.Body = v
	return s
}

func (s *CreateDBResponse) Validate() error {
	return dara.Validate(s)
}
