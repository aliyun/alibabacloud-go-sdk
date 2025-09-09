// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartServiceInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartServiceInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartServiceInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartServiceInstanceResponseBody) *StartServiceInstanceResponse
	GetBody() *StartServiceInstanceResponseBody
}

type StartServiceInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartServiceInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartServiceInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartServiceInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartServiceInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartServiceInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartServiceInstanceResponse) GetBody() *StartServiceInstanceResponseBody {
	return s.Body
}

func (s *StartServiceInstanceResponse) SetHeaders(v map[string]*string) *StartServiceInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartServiceInstanceResponse) SetStatusCode(v int32) *StartServiceInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartServiceInstanceResponse) SetBody(v *StartServiceInstanceResponseBody) *StartServiceInstanceResponse {
	s.Body = v
	return s
}

func (s *StartServiceInstanceResponse) Validate() error {
	return dara.Validate(s)
}
