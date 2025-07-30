// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTCInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTCInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTCInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateTCInstanceResponseBody) *CreateTCInstanceResponse
	GetBody() *CreateTCInstanceResponseBody
}

type CreateTCInstanceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTCInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTCInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTCInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateTCInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTCInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTCInstanceResponse) GetBody() *CreateTCInstanceResponseBody {
	return s.Body
}

func (s *CreateTCInstanceResponse) SetHeaders(v map[string]*string) *CreateTCInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateTCInstanceResponse) SetStatusCode(v int32) *CreateTCInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTCInstanceResponse) SetBody(v *CreateTCInstanceResponseBody) *CreateTCInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateTCInstanceResponse) Validate() error {
	return dara.Validate(s)
}
